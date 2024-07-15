package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/stretchr/testify/assert"
)

type routerAdapter struct {
	mux           *http.ServeMux
	handledErrors []error
}

func (r *routerAdapter) PathValue(req *http.Request, paramName string) string {
	return req.PathValue(paramName)
}

func (r *routerAdapter) HandleRoute(method, pathPattern string, h http.Handler) {
	r.mux.Handle(method+" "+pathPattern, h)
}

func (r *routerAdapter) HandleError(req *http.Request, w http.ResponseWriter, err error) {
	r.handledErrors = append(r.handledErrors, err)
	w.WriteHeader(http.StatusInternalServerError)
}

// openTestLogFile will open a log file in a project root directory
func openTestLogFile() *os.File {
	_, filename, _, _ := runtime.Caller(0) // Will be current file
	testFilePath := filepath.Join(filename, "..", "..", "..", "..", "golang-test.log")
	f, err := os.OpenFile(testFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666)
	if err != nil {
		err = fmt.Errorf("fail to open log file %q for test logging: %w", testFilePath, err)
		panic(err)
	}
	return f
}

var testOutput = openTestLogFile()

func newLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(testOutput, nil))
}

type routeTestCase[TActions any] struct {
	path   string
	query  url.Values
	expect routeTestCaseExpectFn[TActions]
}

type routeTestCaseSetupFn[TActions any] func() (TActions, http.Handler)

func runRouteTestCase[TActions any](
	t *testing.T,
	name string,
	setupFn routeTestCaseSetupFn[TActions],
	tc func() routeTestCase[TActions],
) {
	t.Run(name, func(t *testing.T) {
		tc := tc()
		testActions, router := setupFn()
		testReq := httptest.NewRequest(
			"GET",
			tc.path,
			http.NoBody,
		)
		if len(tc.query) != 0 {
			testReq.URL.RawQuery = tc.query.Encode()
		}
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, testReq)
		tc.expect(t, testActions, recorder)
	})
}

type routeTestCaseExpectFn[TActions any] func(t *testing.T, testActions TActions, recorder *httptest.ResponseRecorder)

func expectBindingErrors[TActions any](wantErrors []handlers.BindingError) routeTestCaseExpectFn[TActions] {
	return func(
		t *testing.T,
		testActions TActions,
		recorder *httptest.ResponseRecorder,
	) {
		assert.Equal(t, 400, recorder.Code)
		assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("content-type"))
		gotErrors := unmarshalBindingErrors(t, recorder.Body)
		if !assert.Len(t, gotErrors.Errors, len(wantErrors)) {
			return
		}
		for _, fe := range wantErrors {
			assertFieldError(t, gotErrors, fe.Location, fe.Field, fe.Code)
		}
	}
}

func unmarshalBindingErrors(
	t *testing.T,
	body *bytes.Buffer,
) *handlers.AggregatedBindingError {
	var gotErrors handlers.AggregatedBindingError
	if err := json.Unmarshal(body.Bytes(), &gotErrors); !assert.NoError(t, err) {
		t.FailNow()
		return nil
	}
	return &gotErrors
}

func assertFieldError(
	t *testing.T,
	err *handlers.AggregatedBindingError,
	location string,
	field string,
	code handlers.BindingErrorCode,
) bool {
	for _, fieldErr := range err.Errors {
		if fieldErr.Location == location && fieldErr.Field == field {
			return assert.Equal(t, code, fieldErr.Code, "field %s: unexpected error code for", field)
		}
	}
	assert.Fail(t, fmt.Sprintf("no error found for field %s, code %s", field, code))
	return false
}

type mockActionCall[TParams any] struct {
	params TParams
}

type mockAction[TParams any] struct {
	calls []mockActionCall[TParams]
}

func (c *mockAction[TParams]) action(
	_ context.Context, params TParams,
) error {
	c.calls = append(c.calls, mockActionCall[TParams]{
		params: params,
	})
	return nil
}
