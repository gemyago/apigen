package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandling(t *testing.T) {
	setupRouter := func() *routerAdapter {
		controller := newErrorHandlingController()
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.MountErrorHandlingRoutes(controller, handlers.NewHttpApp(router))
		return router
	}

	t.Run("parsing-errors", func(t *testing.T) {
		t.Run("respond with 400 and error details by default", func(t *testing.T) {
			router := setupRouter()
			testReq := httptest.NewRequest(
				"GET",
				"/error-handling/parsing-errors",
				http.NoBody,
			)
			recorder := httptest.NewRecorder()
			router.mux.ServeHTTP(recorder, testReq)

			assert.Equal(t, 400, recorder.Code)
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
