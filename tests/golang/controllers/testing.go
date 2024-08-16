package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/samber/lo"
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

func (r *routerAdapter) HandleError(_ *http.Request, w http.ResponseWriter, err error) {
	r.handledErrors = append(r.handledErrors, err)
	w.WriteHeader(http.StatusInternalServerError)
}

// openTestLogFile will open a log file in a project root directory.
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

var testOutput = openTestLogFile() //nolint:gochecknoglobals // we want to open it once for all tests

func newLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(testOutput, nil))
}

type routeTestCase[TActions any] struct {
	method string
	path   string
	query  url.Values
	body   io.Reader
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
		method := tc.method
		if method == "" {
			method = http.MethodGet
		}
		body := tc.body
		if body == nil {
			body = http.NoBody
		}
		testReq := httptest.NewRequest(
			method,
			tc.path,
			body,
		)
		if len(tc.query) != 0 {
			testReq.URL.RawQuery = tc.query.Encode()
		}
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, testReq)
		tc.expect(t, testActions, recorder)
	})
}

func marshalJSONDataAsReader(t *testing.T, data any) io.Reader {
	body, err := json.Marshal(data)
	if err != nil {
		t.Errorf("failed to marshal body: %v", err)
		t.FailNow()
	}
	return bytes.NewBuffer(body)
}

type routeTestCaseExpectFn[TActions any] func(t *testing.T, testActions TActions, recorder *httptest.ResponseRecorder)

type expectedBindingError struct {
	Field    string
	Location string
	Code     string
}

type fieldBindingError struct {
	Location string `json:"location"`
	Code     string `json:"code"`
}

type aggregatedBindingError struct {
	Errors []fieldBindingError `json:"errors"`
}

func expectBindingErrors[TActions any](wantErrors []expectedBindingError) routeTestCaseExpectFn[TActions] {
	return func(
		t *testing.T,
		_ TActions,
		recorder *httptest.ResponseRecorder,
	) {
		if !assert.Equal(t, http.StatusBadRequest, recorder.Code, "Unexpected response: %v", recorder.Body) {
			return
		}
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
) *aggregatedBindingError {
	var gotErrors aggregatedBindingError
	if err := json.Unmarshal(body.Bytes(), &gotErrors); !assert.NoError(t, err) {
		t.FailNow()
		return nil
	}
	return &gotErrors
}

func assertFieldError(
	t *testing.T,
	err *aggregatedBindingError,
	location string,
	field string,
	code string,
) {
	fieldPath := lo.If(field == "", location).Else(location + "." + field)
	for _, fieldErr := range err.Errors {
		if fieldErr.Location == fieldPath {
			assert.Equal(t, code, fieldErr.Code, "field %s: unexpected error code for", fieldPath)
			return
		}
	}
	assert.FailNow(t,
		fmt.Sprintf("no error found for field %s, code %s. Got Errors: %v",
			fieldPath, code, err.Errors,
		))
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

func injectValueRandomly[T any](fake faker.Faker, values []T, value T) (int, []T) {
	index := fake.IntBetween(0, len(values))
	newValues := make([]T, len(values)+1)
	copy(newValues[:index], values[:index])
	newValues[index] = value
	copy(newValues[index+1:], values[index:])
	return index, newValues
}
