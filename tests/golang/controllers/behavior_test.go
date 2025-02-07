package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	t.Run("errors", func(t *testing.T) {
		runRouteTestCase(t, "should handle parsing errors with default handler", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam2": []string{fake.Lorem().Word()}},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 400, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Empty(t, testActions.withParamsAndResponse.calls)
				},
			}
		})
		runRouteTestCase(t, "should handle request path params validation errors", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam2": []string{strconv.Itoa(fake.IntBetween(5001, 1000000))}},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 400, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Empty(t, testActions.withParamsAndResponse.calls)
				},
			}
		})
		runRouteTestCase(t, "should handle request body validation errors", setupRouter, func() testCase {
			body := fmt.Sprintf(`{"field2":%v}`, fake.IntBetween(5001, 1000000))
			return testCase{
				method: http.MethodPost,
				body:   bytes.NewReader([]byte(body)),
				path:   "/behavior/with-params-and-response",
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 400, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Empty(t, testActions.withParamsAndResponse.calls)
				},
			}
		})
		runRouteTestCase(t, "should handle response body serialization errors", setupRouter, func() testCase {
			resBody := models.BehaviorWithParamsAndResponseResponseBody{
				Field1: fake.Lorem().Word(),
			}
			resBody.Field2 = &resBody // Circular reference will cause serialization error
			bodyData, err := json.Marshal(resBody)
			require.Errorf(t, err, "body was marshalled successfully: %v", bodyData)
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsAndResponse.nextResult = &resBody
					return testActions
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 500, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Empty(t, testActions.withParamsAndResponse.calls)
				},
			}
		})
		runRouteTestCase(t, "should handle parsing errors with custom handler", setupRouter, func() testCase {
			wantStatusCode := fake.IntBetween(400, 500)
			wantErrorBody := fake.Lorem().Word()
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				query:  url.Values{"queryParam2": []string{fake.Lorem().Word()}},
				appendHTTPAppOpts: func(opts ...handlers.HTTPAppOpt) []handlers.HTTPAppOpt {
					return append(opts, handlers.WithParsingErrorHandler(
						func(_ *http.Request, w http.ResponseWriter, err error) {
							w.WriteHeader(wantStatusCode)
							_, _ = w.Write([]byte(wantErrorBody + ":" + err.Error()))
						},
					))
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, wantStatusCode, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Contains(t, recorder.Body.String(), wantErrorBody)
					assert.Empty(t, testActions.withParamsAndResponse.calls)
				},
			}
		})
		runRouteTestCase(t, "should handle action errors with default handler", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
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
		runRouteTestCase(t, "should handle action errors with custom handler", setupRouter, func() testCase {
			wantStatusCode := fake.IntBetween(400, 500)
			wantErrorBody := fake.Lorem().Word()
			return testCase{
				method: http.MethodPost,
				path:   "/behavior/with-params-and-response",
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsAndResponse.nextError = errors.New(fake.Lorem().Word())
					return testActions
				},
				appendHTTPAppOpts: func(opts ...handlers.HTTPAppOpt) []handlers.HTTPAppOpt {
					return append(opts, handlers.WithActionErrorHandler(
						func(_ *http.Request, w http.ResponseWriter, err error) {
							w.WriteHeader(wantStatusCode)
							_, _ = w.Write([]byte(wantErrorBody + ":" + err.Error()))
						},
					))
				},
				expect: func(t *testing.T, testActions *behaviorControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, wantStatusCode, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Contains(t, recorder.Body.String(), wantErrorBody)
					assert.Len(t, testActions.withParamsAndResponse.calls, 1)
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

		runRouteTestCase(t, "should fail with error if http action", setupRouter, func() testCase {
			return testCase{
				method: http.MethodGet,
				path:   "/behavior/with-params-no-response",
				query:  url.Values{"queryParam1": []string{fake.Lorem().Word()}},
				setupActions: func(testActions *behaviorControllerTestActions) *behaviorControllerTestActions {
					testActions.withParamsNoResponse.isHTTPAction = true
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
			wantResponse := &models.BehaviorWithParamsAndResponseResponseBody{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodPost,
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
			wantResponse := &models.BehaviorWithParamsAndResponseResponseBody{
				Field1: fake.Lorem().Word(),
			}
			return testCase{
				method: http.MethodPost,
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
				method: http.MethodPost,
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
