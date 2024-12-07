package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/stretchr/testify/assert"
)

func TestBehavior(t *testing.T) {
	setupRouter := func() (*behaviorControllerTestActions, http.Handler) {
		testActions := &behaviorControllerTestActions{}
		controller := newBehaviorController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterBehaviorRoutes(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	t.Run("controller builder", func(t *testing.T) {
		t.Run("should panic if actions are not initialized", func(t *testing.T) {
			assert.PanicsWithError(t, "behaviorNoParamsNoResponse action has not been initialized", func() {
				handlers.BuildBehaviorController().Finalize()
			})
		})

		t.Run("should build the controller if all actions are initialized", func(t *testing.T) {
			assert.NotPanics(t, func() {
				newBehaviorController(&behaviorControllerTestActions{})
			})
		})
	})

	type testCase = routeTestCase[*behaviorControllerTestActions]

	t.Run("noParamsNoResponse", func(t *testing.T) {
		t.Run("should process the request", func(t *testing.T) {

		})

		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/no-params-no-response",
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsNoResponse.calls, 1)
				},
			}
		})
	})
}
