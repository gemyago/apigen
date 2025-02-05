package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandling(t *testing.T) {
	fake := faker.New()
	setupRouter := func() *routerAdapter {
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		rootLogger := newLogger()
		handlers.RegisterErrorHandlingRoutesV2(
			&errorHandlingController{},
			handlers.NewHTTPApp(router,
				handlers.WithLogger(rootLogger),
			),
		)
		return router
	}

	type testCase struct {
		name       string
		path       string
		query      url.Values
		wantErrors []expectedBindingError
	}

	runBadRequestTestCase := func(t *testing.T, tc testCase) {
		router := setupRouter()
		testReq := httptest.NewRequest(
			http.MethodGet,
			tc.path,
			http.NoBody,
		)
		if len(tc.query) != 0 {
			testReq.URL.RawQuery = tc.query.Encode()
		}
		recorder := httptest.NewRecorder()
		router.mux.ServeHTTP(recorder, testReq)

		assert.Equal(t, 400, recorder.Code)
		assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("content-type"))
		gotErrors := unmarshalBindingErrors(t, recorder.Body)
		if !assert.Len(t, gotErrors.Errors, len(tc.wantErrors)) {
			return
		}
		for _, fe := range tc.wantErrors {
			assertFieldError(t, gotErrors, fe.Location, fe.Field, fe.Code)
		}
	}

	t.Run("parsing-errors", func(t *testing.T) {
		runBadRequestTestCase(t, testCase{
			name: "respond with 400 if parsing fails",
			path: fmt.Sprintf(
				"/error-handling/parsing-errors/%[1]s/%[1]s?requiredQuery1=%[1]s&requiredQuery2=%[1]s",
				fake.Lorem().Word(),
			),
			wantErrors: []expectedBindingError{
				{Field: "pathParam1", Location: "path", Code: "BAD_FORMAT"},
				{Field: "pathParam2", Location: "path", Code: "BAD_FORMAT"},
				{Field: "requiredQuery1", Location: "query", Code: "BAD_FORMAT"},
				{Field: "requiredQuery2", Location: "query", Code: "BAD_FORMAT"},
			},
		})
	})

	t.Run("validation-errors", func(t *testing.T) {
		runBadRequestTestCase(t, testCase{
			name: "respond with 400 if validation fails",
			path: "/error-handling/validation-errors?requiredQuery1=1&requiredQuery2=2",
			wantErrors: []expectedBindingError{
				{Field: "requiredQuery1", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
				{Field: "requiredQuery2", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
			},
		})
	})

	t.Run("action error with default action error handler", func(t *testing.T) {
		router := setupRouter()
		testReq := httptest.NewRequest(
			http.MethodGet,
			"/error-handling/action-errors",
			http.NoBody,
		)
		recorder := httptest.NewRecorder()
		router.mux.ServeHTTP(recorder, testReq)

		assert.Equal(t, 500, recorder.Code)
		assert.Equal(t, "text/plain; charset=utf-8", recorder.Header().Get("content-type"))
		assert.Equal(t, "Internal Server Error\n", recorder.Body.String())
	})
}
