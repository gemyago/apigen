package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/stretchr/testify/assert"
)

func TestBehaviorNoParamsNoResponse(t *testing.T) {
	type testCase = routeTestCase[*behaviorControllerNoParamsNoResponseTestActions]
	setupRouter := func(_ testCase) (*behaviorControllerNoParamsNoResponseTestActions, http.Handler) {
		testActions := &behaviorControllerNoParamsNoResponseTestActions{}
		controller := behaviorControllerNoParamsNoResponse{testActions}
		rootHandler := handlers.
			NewRootHandler(&routerAdapter{mux: http.NewServeMux()}, handlers.WithLogger(newLogger())).
			RegisterBehaviorNoParamsNoResponseIsolatedRoutes(&controller)
		return testActions, rootHandler
	}

	t.Run("noParamsNoResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-no-response-isolated",
				expect: func(
					t *testing.T,
					testActions *behaviorControllerNoParamsNoResponseTestActions,
					recorder *httptest.ResponseRecorder,
				) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsNoResponse.calls, 1)
				},
			}
		})
	})
}
