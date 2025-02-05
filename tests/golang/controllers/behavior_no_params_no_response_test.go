package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/stretchr/testify/assert"
)

func TestBehaviorNoParamsNoResponse(t *testing.T) {
	setupRouter := func() (*behaviorControllerNoParamsNoResponseTestActions, http.Handler) {
		testActions := &behaviorControllerNoParamsNoResponseTestActions{}
		controller := behaviorControllerNoParamsNoResponse{testActions}
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterBehaviorNoParamsNoResponseIsolatedRoutesV2(
			&controller,
			handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())),
		)
		return testActions, router.mux
	}

	type testCase = routeTestCase[*behaviorControllerNoParamsNoResponseTestActions]

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
