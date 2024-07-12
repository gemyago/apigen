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
		controller := newErrorHandlingController()
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.MountErrorHandlingRoutes(
			controller,
			handlers.NewHttpApp(router, handlers.WithLogger(newLogger())),
		)
		return router
	}

	type testCase struct {
		name       string
		path       string
		query      url.Values
		wantErrors []handlers.BindingError
	}

	runTestCase := func(t *testing.T, tc testCase) {
		router := setupRouter()
		testReq := httptest.NewRequest(
			"GET",
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
		runTestCase(t, testCase{
			name: "respond with 400 if parsing fails",
			path: fmt.Sprintf("/error-handling/parsing-errors/%[1]s/%[1]s?requiredQuery1=%[1]s&requiredQuery2=%[1]s", fake.Lorem().Word()),
			wantErrors: []handlers.BindingError{
				{Field: "pathParam1", Location: "path", Code: "BAD_FORMAT"},
				{Field: "pathParam2", Location: "path", Code: "BAD_FORMAT"},
				{Field: "requiredQuery1", Location: "query", Code: "BAD_FORMAT"},
				{Field: "requiredQuery2", Location: "query", Code: "BAD_FORMAT"},
			},
		})
	})
}
