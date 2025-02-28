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
	"github.com/stretchr/testify/require"
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
					testActions.noParamsWithResponse.nextResult = (*transformedBehaviorNoParamsWithResponse202Response)(wantResponse)
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					require.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body)
					assert.Len(t, testActions.noParamsWithResponse.calls, 1)
					assert.Equal(t,
						wantResponse,
						(*models.BehaviorNoParamsWithResponse202Response)(
							testActions.noParamsWithResponse.unmarshalResult(t, recorder.Body),
						),
					)
				},
			}
		})

		runRouteTestCase(t, "should fail if res transformation fails", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/no-params-with-response",
				setupActions: func(testActions *controllerTestActions) *controllerTestActions {
					testActions.behaviorNoParamsWithResponseTransformer.nextTransformResponseErr = errors.New(fake.Lorem().Word())
					return testActions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					require.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body)

					assert.Len(t, testActions.noParamsWithResponse.calls, 1)
				},
			}
		})
	})

	t.Run("withParamsNoResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			wantParams := &models.BehaviorWithParamsNoResponseParams{
				QueryParam1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{wantParams.QueryParam1}},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					require.Equal(t, 202, recorder.Code, "Unexpected response: %v", recorder.Body)

					assert.Len(t, testActions.withParamsNoResponse.calls, 1)
					assert.Equal(t,
						wantParams,
						(*models.BehaviorWithParamsNoResponseParams)(
							testActions.withParamsNoResponse.calls[0].params,
						),
					)
				},
			}
		})

		runRouteTestCase(t, "should fail if req transformation fails", setupRouter, func() testCase {
			wantParams := &models.BehaviorWithParamsNoResponseParams{
				QueryParam1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{wantParams.QueryParam1}},
				setupActions: func(actions *controllerTestActions) *controllerTestActions {
					actions.behaviorWithParamsNoResponseTransformer.nextTransformRequestErr = errors.New(fake.Lorem().Word())
					return actions
				},
				expect: func(t *testing.T, testActions *controllerTestActions, recorder *httptest.ResponseRecorder) {
					require.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body)

					assert.Empty(t, testActions.withParamsNoResponse.calls)
				},
			}
		})
	})

	t.Run("withParamsAndResponse", func(t *testing.T) {
		runRouteTestCase(t, "should process the request", setupRouter, func() testCase {
			wantParams := &models.BehaviorWithParamsAndResponseParams{
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
						(*models.BehaviorWithParamsAndResponseParams)(
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
