package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestBehavior(t *testing.T) {
	fake := faker.New()

	setupRouter := func() (*behaviorControllerTestActions, http.Handler) {
		testActions := &behaviorControllerTestActions{}
		controller := &behaviorController{testActions}
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterBehaviorRoutesV2(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	type testCase = routeTestCase[*behaviorControllerTestActions]

	t.Run("noParamsNoResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-no-response",
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsNoResponse.calls, 1)
				},
			}
		})
	})

	t.Run("noParamsWithResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			wantResponse := &models.BehaviorNoParamsWithResponse202Response{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-with-response",
				setupActions: func(testActions *behaviorControllerTestActions) {
					testActions.noParamsWithResponse.nextResult = wantResponse
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsWithResponse.calls, 1)
					assert.Equal(t, wantResponse, testActions.noParamsWithResponse.unmarshalResult(t, recorder.Body))
				},
			}
		})
	})

	t.Run("withParamsNoResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsNoResponse.calls, 1)
				},
			}
		})
	})
}
