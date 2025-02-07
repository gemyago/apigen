package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestBehavior(t *testing.T) {
	fake := faker.New()

	type testCase = routeTestCase[*behaviorControllerTestActions]

	setupRouter := func(tc testCase) (*behaviorControllerTestActions, http.Handler) {
		testActions := tc.setupActions(&behaviorControllerTestActions{})
		controller := &behaviorController{testActions}
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterBehaviorRoutesV2(
			controller,
			handlers.NewHTTPApp(router, tc.appendHTTPAppOpts(handlers.WithLogger(newLogger()))...),
		)
		return testActions, router.mux
	}

	t.Run("app", func(t *testing.T) {
		runRouteTestCase(t, "should handle parsing errors with default handler", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-and-response",
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
				},
			}
		})
	})

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

		runRouteTestCase(t, "should process the request with http action", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-no-response",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.noParamsNoResponse.isHTTPAction = true
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsNoResponse.calls, 1)
					lastCall := testActions.noParamsNoResponse.calls[0]
					assert.NotNil(t, lastCall.req)
					assert.NotNil(t, lastCall.res)
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-no-response",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.noParamsNoResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
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
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.noParamsWithResponse.nextResult = wantResponse
					return testActions
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

		runRouteTestCase(t, "should process the request with http action", setupRouter, func() testCase {
			wantResponse := &models.BehaviorNoParamsWithResponse202Response{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-with-response",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.noParamsWithResponse.isHTTPAction = true
					testActions.noParamsWithResponse.nextResult = wantResponse
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsWithResponse.calls, 1)
					lastCall := testActions.noParamsWithResponse.calls[0]
					assert.NotNil(t, lastCall.req)
					assert.NotNil(t, lastCall.res)
					assert.Equal(t, wantResponse, testActions.noParamsWithResponse.unmarshalResult(t, recorder.Body))
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-with-response",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.noParamsWithResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsWithResponse.calls, 1)
				},
			}
		})
	})

	t.Run("withParamsNoResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			wantParams := &handlers.BehaviorBehaviorWithParamsNoResponseRequest{
				QueryParam1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{wantParams.QueryParam1}},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsNoResponse.calls, 1)
					assert.Equal(t, wantParams, testActions.withParamsNoResponse.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should process the request with http action", setupRouter, func() testCase {
			wantParams := &handlers.BehaviorBehaviorWithParamsNoResponseRequest{
				QueryParam1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{wantParams.QueryParam1}},
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsNoResponse.isHTTPAction = true
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsNoResponse.calls, 1)
					lastCall := testActions.withParamsNoResponse.calls[0]
					assert.NotNil(t, lastCall.req)
					assert.NotNil(t, lastCall.res)
					assert.Equal(t, wantParams, testActions.withParamsNoResponse.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{fake.Lorem().Word()}},
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsNoResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsNoResponse.calls, 1)
				},
			}
		})
	})

	t.Run("withParamsAndResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			wantParams := &handlers.BehaviorBehaviorWithParamsAndResponseRequest{
				QueryParam1: fake.Lorem().Word(),
			}
			wantResponse := &models.BehaviorNoParamsWithResponse202Response{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam1": []string{wantParams.QueryParam1}},
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsAndResponse.nextResult = wantResponse
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsAndResponse.calls, 1)
					assert.Equal(t, wantParams, testActions.withParamsAndResponse.calls[0].params)
					assert.Equal(t, wantResponse, testActions.withParamsAndResponse.unmarshalResult(t, recorder.Body))
				},
			}
		})

		runRouteTestCase(t, "should process the request with http action", setupRouter, func() testCase {
			wantParams := &handlers.BehaviorBehaviorWithParamsAndResponseRequest{
				QueryParam1: fake.Lorem().Word(),
			}
			wantResponse := &models.BehaviorNoParamsWithResponse202Response{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam1": []string{wantParams.QueryParam1}},
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsAndResponse.isHTTPAction = true
					testActions.withParamsAndResponse.nextResult = wantResponse
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsAndResponse.calls, 1)
					lastCall := testActions.withParamsAndResponse.calls[0]
					assert.NotNil(t, lastCall.req)
					assert.NotNil(t, lastCall.res)
					assert.Equal(t, wantParams, testActions.withParamsAndResponse.calls[0].params)
					assert.Equal(t, wantResponse, testActions.withParamsAndResponse.unmarshalResult(t, recorder.Body))
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam1": []string{fake.Lorem().Word()}},
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsAndResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsAndResponse.calls, 1)
				},
			}
		})
	})

	t.Run("noStatusDefined", func(t *testing.T) {
		runRouteTestCase(t, "should use default status", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-status-defined",
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 200, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noStatusDefined.calls, 1)
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-status-defined",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.noStatusDefined.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noStatusDefined.calls, 1)
				},
			}
		})
	})

	t.Run("withStatusDefined", func(t *testing.T) {
		runRouteTestCase(t, "should use defined status", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-status-defined",
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withStatusDefined.calls, 1)
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-status-defined",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withStatusDefined.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withStatusDefined.calls, 1)
				},
			}
		})
	})
}
