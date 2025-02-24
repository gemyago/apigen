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

func TestBehaviorTransform(t *testing.T) {
	fake := faker.New()

	type controllerTestActions = behaviorControllerTransformTestActions
	type testCase = routeTestCase[*controllerTestActions]

	setupRouter := func(tc testCase) (*controllerTestActions, http.Handler) {
		testActions := tc.setupActions(&behaviorControllerTransformTestActions{})
		controller := &behaviorControllerTransform{
			testActions:   testActions,
			stdController: behaviorController{testActions: &behaviorControllerTestActions{}},
		}
		rootHandler := handlers.
			NewRootHandler(
				&routerAdapter{mux: http.NewServeMux()},
				tc.appendRootHandlerOpts(handlers.WithLogger(newLogger()))...,
			).
			RegisterBehaviorRoutes(controller)
		return testActions, rootHandler
	}

	t.Run("noParamsWithResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			wantResponse := &models.BehaviorNoParamsWithResponse202Response{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-with-response",
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.noParamsWithResponse.nextResult = wantResponse
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.noParamsWithResponse.calls, 1)
					assert.Equal(t, wantResponse, testActions.noParamsWithResponse.unmarshalResult(t, recorder.Body))
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-with-response",
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.noParamsWithResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
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
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsNoResponse.calls, 1)
					assert.Equal(t, wantParams, testActions.withParamsNoResponse.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should fail with error", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{fake.Lorem().Word()}},
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.withParamsNoResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsNoResponse.calls, 1)
				},
			}
		})

		runRouteTestCase(t, "should fail with error if http action", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{fake.Lorem().Word()}},
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.withParamsNoResponse.isHTTPAction = true
					testActions.withParamsNoResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
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
			wantResponse := &models.BehaviorWithParamsAndResponseResponseBody{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam1": []string{wantParams.QueryParam1}},
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.withParamsAndResponse.nextResult =
						(*transformedBehaviorWithParamsAndResponseResponseBody)(wantResponse)
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsAndResponse.calls, 1)
					assert.Equal(t,
						wantParams,
						(*handlers.BehaviorBehaviorWithParamsAndResponseRequest)(
							testActions.withParamsAndResponse.calls[0].params,
						),
					)
					assert.Equal(t,
						(*transformedBehaviorWithParamsAndResponseResponseBody)(wantResponse),
						testActions.withParamsAndResponse.unmarshalResult(t, recorder.Body),
					)

					assert.True(t, testActions.behaviorWithParamsAndResponseTransformer.lastReqProvided)
					assert.True(t, testActions.behaviorWithParamsAndResponseTransformer.lastCtxProvided)
				},
			}
		})

		runRouteTestCase(t, "should fail if req transformation fails", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam1": []string{fake.Lorem().Word()}},
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.behaviorWithParamsAndResponseTransformer.nextTransformRequestErr = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Empty(t, testActions.withParamsAndResponse.calls)
				},
			}
		})

		runRouteTestCase(t, "should fail if res transformation fails", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam1": []string{fake.Lorem().Word()}},
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.behaviorWithParamsAndResponseTransformer.nextTransformResponseErr = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Len(t, testActions.withParamsAndResponse.calls, 1)
				},
			}
		})
	})
}
