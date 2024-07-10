package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

	t.Run("parsing-errors", func(t *testing.T) {
		t.Run("respond with 400 if parsing fails", func(t *testing.T) {
			router := setupRouter()
			testReq := httptest.NewRequest(
				"GET",
				fmt.Sprintf("/error-handling/parsing-errors?requiredQuery1=%s&requiredQuery2=%[1]s", fake.Lorem().Word()),
				http.NoBody,
			)
			recorder := httptest.NewRecorder()
			router.mux.ServeHTTP(recorder, testReq)

			assert.Equal(t, 400, recorder.Code)
			assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("content-type"))
			assert.JSONEq(t, `{
				"errors": [
					{
						"field": "requiredQuery1",
						"location": "query",
						"code": "BAD_FORMAT"
					},
					{
						"field": "requiredQuery2",
						"location": "query",
						"code": "BAD_FORMAT"
					}
				]
			}`, recorder.Body.String())
		})
		t.Run("respond with 400 if validation fails", func(t *testing.T) {
			router := setupRouter()
			testReq := httptest.NewRequest(
				"GET",
				"/error-handling/parsing-errors",
				http.NoBody,
			)
			recorder := httptest.NewRecorder()
			router.mux.ServeHTTP(recorder, testReq)

			assert.Equal(t, 400, recorder.Code)
			assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("content-type"))
			assert.JSONEq(t, `{
				"errors": [
					{
						"field": "requiredQuery1",
						"location": "query",
						"code": "INVALID"
					},
					{
						"field": "requiredQuery2",
						"location": "query",
						"code": "INVALID"
					}
				]
			}`, recorder.Body.String())
		})
	})
}
