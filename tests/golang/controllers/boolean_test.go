package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestBoolean(t *testing.T) {
	fake := faker.New()

	type testCase = routeTestCase[*booleanControllerTestActions]
	setupRouter := func(_ testCase) (*booleanControllerTestActions, http.Handler) {
		testActions := &booleanControllerTestActions{}
		controller := &booleanController{testActions}
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterBooleanRoutes(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	randomBooleans := func(n int) []bool {
		vals := make([]bool, n)
		for i := range n {
			vals[i] = fake.Bool()
		}
		return vals
	}

	booleansToString := func(vals []bool) []string {
		res := make([]string, len(vals))
		for i, v := range vals {
			res[i] = strconv.FormatBool(v)
		}
		return res
	}

	t.Run("parsing", func(t *testing.T) {
		randomReq := func() *handlers.BooleanBooleanParsingRequest {
			return &handlers.BooleanBooleanParsingRequest{
				// path
				BoolParam1: fake.Bool(),
				BoolParam2: fake.Bool(),

				// query
				BoolParam1InQuery: fake.Bool(),
				BoolParam2InQuery: fake.Bool(),

				// body
				Payload: &models.BooleanParsingRequest{
					BoolParam1: fake.Bool(),
					BoolParam2: fake.Bool(),
				},
			}
		}

		buildQuery := func(wantReq *handlers.BooleanBooleanParsingRequest) url.Values {
			query := url.Values{}
			query.Add("boolParam1InQuery", strconv.FormatBool(wantReq.BoolParam1InQuery))
			query.Add("boolParam2InQuery", strconv.FormatBool(wantReq.BoolParam2InQuery))
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := randomReq()
				query := buildQuery(wantReq)

				return routeTestCase[*booleanControllerTestActions]{
					method: http.MethodPost,
					path: fmt.Sprintf("/boolean/parsing/%v/%v",
						wantReq.BoolParam1, wantReq.BoolParam2),
					query: query,
					body:  marshalJSONDataAsReader(t, wantReq.Payload),
					expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, wantReq, testActions.booleanParsing.calls[0].params)
					},
				}
			})

		nonBooleanValues := []string{
			"True",
			// "False",
			// "0",
			// "1",
			// fake.Lorem().Word(),
		}

		for _, val := range nonBooleanValues {
			runRouteTestCase(t, "should fail if non boolean value "+val, setupRouter, func() testCase {
				query := url.Values{}
				query.Add("boolParam1InQuery", val)
				query.Add("boolParam2InQuery", val)
				query.Add("optionalBoolParam1InQuery", val)
				query.Add("optionalBoolParam2InQuery", val)

				return testCase{
					method: http.MethodPost,
					path:   fmt.Sprintf("/boolean/parsing/%v/%v", val, val),
					query:  query,
					body: bytes.NewBufferString(`{
						"boolParam1": %v,
						"boolParam2": %v
					}`),
					expect: expectBindingErrors[*booleanControllerTestActions](
						[]expectedBindingError{
							{Field: "boolParam1", Location: "path", Code: "BAD_FORMAT"},
							{Field: "boolParam2", Location: "path", Code: "BAD_FORMAT"},
							{Field: "boolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "boolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},
							{Location: "body", Code: "BAD_FORMAT"},
						},
					),
				}
			})
		}
	})

	t.Run("required-validation", func(t *testing.T) {
		buildQuery := func(wantReq *handlers.BooleanBooleanRequiredValidationRequest) url.Values {
			query := url.Values{}
			query.Add("boolParam1InQuery", strconv.FormatBool(wantReq.BoolParam1InQuery))
			query.Add("boolParam2InQuery", strconv.FormatBool(wantReq.BoolParam2InQuery))
			query.Add("optionalBoolParam1InQuery", strconv.FormatBool(wantReq.OptionalBoolParam1InQuery))
			query.Add("optionalBoolParam2InQuery", strconv.FormatBool(wantReq.OptionalBoolParam2InQuery))
			return query
		}

		runRouteTestCase(t, "should ignore missing optional params", setupRouter, func() testCase {
			wantReq := &handlers.BooleanBooleanRequiredValidationRequest{
				// query
				BoolParam1InQuery: fake.Bool(),
				BoolParam2InQuery: fake.Bool(),

				Payload: &models.BooleanRequiredValidationRequest{},
			}

			query := buildQuery(wantReq)
			query.Del("optionalBoolParam1InQuery")
			query.Del("optionalBoolParam2InQuery")

			return testCase{
				method: http.MethodPost,
				path:   "/boolean/required-validation",
				query:  query,
				body:   marshalJSONDataAsReader(t, wantReq.Payload),
				expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, wantReq, testActions.booleanRequiredValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should parse optional params", setupRouter, func() testCase {
			wantReq := &handlers.BooleanBooleanRequiredValidationRequest{
				// query
				BoolParam1InQuery:         fake.Bool(),
				BoolParam2InQuery:         fake.Bool(),
				OptionalBoolParam1InQuery: fake.Bool(),
				OptionalBoolParam2InQuery: fake.Bool(),

				Payload: &models.BooleanRequiredValidationRequest{
					OptionalBoolParam1: fake.Bool(),
					OptionalBoolParam2: fake.Bool(),
				},
			}

			return testCase{
				method: http.MethodPost,
				path:   "/boolean/required-validation",
				query:  buildQuery(wantReq),
				body:   marshalJSONDataAsReader(t, wantReq.Payload),
				expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, wantReq, testActions.booleanRequiredValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should validate empty params", setupRouter, func() testCase {
			query := url.Values{}
			query.Add("boolParam1InQuery", "")
			query.Add("boolParam2InQuery", "")
			query.Add("optionalBoolParam1InQuery", "")
			query.Add("optionalBoolParam2InQuery", "")

			return testCase{
				method: http.MethodPost,
				path:   "/boolean/required-validation",
				query:  query,
				body:   bytes.NewBufferString("{}"),
				expect: expectBindingErrors[*booleanControllerTestActions](
					[]expectedBindingError{
						{Field: "boolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "boolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalBoolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalBoolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},

						// "required" validation of optional bool params in body is not yet supported
					},
				),
			}
		})
		runRouteTestCase(t, "should ensure required params", setupRouter, func() testCase {
			query := url.Values{}
			query.Del("boolParam1InQuery")
			query.Del("boolParam2InQuery")

			return testCase{
				method: http.MethodPost,
				path:   "/boolean/required-validation",
				query:  query,
				body:   bytes.NewBufferString("{}"),
				expect: expectBindingErrors[*booleanControllerTestActions](
					[]expectedBindingError{
						{Field: "boolParam1InQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "boolParam2InQuery", Location: "query", Code: "INVALID_REQUIRED"},

						// "required" validation of optional bool params in body is not yet supported
					},
				),
			}
		})
	})

	t.Run("nullable", func(t *testing.T) {
		randomReq := func() *handlers.BooleanBooleanNullableRequest {
			return &handlers.BooleanBooleanNullableRequest{
				// path
				BoolParam1: lo.ToPtr(fake.Bool()),
				BoolParam2: lo.ToPtr(fake.Bool()),

				// query
				BoolParam1InQuery:         lo.ToPtr(fake.Bool()),
				BoolParam2InQuery:         lo.ToPtr(fake.Bool()),
				OptionalBoolParam1InQuery: lo.ToPtr(fake.Bool()),

				// body
				Payload: &models.BooleanNullableRequest{
					BoolParam1:         lo.ToPtr(fake.Bool()),
					BoolParam2:         lo.ToPtr(fake.Bool()),
					OptionalBoolParam1: lo.ToPtr(fake.Bool()),
				},
			}
		}

		buildQuery := func(wantReq *handlers.BooleanBooleanNullableRequest) url.Values {
			query := url.Values{}
			query.Add("boolParam1InQuery", strconv.FormatBool(lo.FromPtr(wantReq.BoolParam1InQuery)))
			query.Add("boolParam2InQuery", strconv.FormatBool(lo.FromPtr(wantReq.BoolParam2InQuery)))
			query.Add("optionalBoolParam1InQuery", strconv.FormatBool(lo.FromPtr(wantReq.OptionalBoolParam1InQuery)))
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := randomReq()
				query := buildQuery(wantReq)

				return routeTestCase[*booleanControllerTestActions]{
					method: http.MethodPost,
					path: fmt.Sprintf("/boolean/nullable/%v/%v",
						*wantReq.BoolParam1, *wantReq.BoolParam2),
					query: query,
					body:  marshalJSONDataAsReader(t, wantReq.Payload),
					expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, wantReq, testActions.booleanNullable.calls[0].params)
					},
				}
			})

		runRouteTestCase(t, "should accept null values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := &handlers.BooleanBooleanNullableRequest{
					Payload: &models.BooleanNullableRequest{},
				}
				query := url.Values{}
				query.Add("boolParam1InQuery", "null")
				query.Add("boolParam2InQuery", "null")
				query.Add("optionalBoolParam1InQuery", "null")
				return routeTestCase[*booleanControllerTestActions]{
					method: http.MethodPost,
					path:   "/boolean/nullable/null/null",
					query:  query,
					body: bytes.NewBufferString(`{
						"boolParam1": null,
						"boolParam2": null,
						"optionalBoolParam1": null
					}`),
					expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, wantReq, testActions.booleanNullable.calls[0].params)
					},
				}
			})
	})

	t.Run("array-items", func(t *testing.T) {
		randomReq := func() *handlers.BooleanBooleanArrayItemsRequest {
			return &handlers.BooleanBooleanArrayItemsRequest{
				// path
				BoolParam1: randomBooleans(5),
				BoolParam2: randomBooleans(5),

				// query
				BoolParam1InQuery: randomBooleans(5),
				BoolParam2InQuery: randomBooleans(5),

				// body
				Payload: &models.BooleanArrayItemsRequest{
					BoolParam1: randomBooleans(5),
					BoolParam2: randomBooleans(5),
				},
			}
		}

		buildQuery := func(wantReq *handlers.BooleanBooleanArrayItemsRequest) url.Values {
			query := url.Values{}
			query["boolParam1InQuery"] = booleansToString(wantReq.BoolParam1InQuery)
			query["boolParam2InQuery"] = booleansToString(wantReq.BoolParam2InQuery)
			return query
		}

		buildPath := func(wantReq *handlers.BooleanBooleanArrayItemsRequest) string {
			return fmt.Sprintf("/boolean/array-items/%v/%v",
				strings.Join(booleansToString(wantReq.BoolParam1), ","),
				strings.Join(booleansToString(wantReq.BoolParam2), ","),
			)
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := randomReq()

				return routeTestCase[*booleanControllerTestActions]{
					method: http.MethodPost,
					path:   buildPath(wantReq),
					query:  buildQuery(wantReq),
					body:   marshalJSONDataAsReader(t, wantReq.Payload),
					expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, wantReq, testActions.booleanArrayItems.calls[0].params)
					},
				}
			})

		runRouteTestCase(t, "should fail if bad values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := randomReq()
				query := buildQuery(wantReq)

				_, query["boolParam1InQuery"] = injectValueRandomly(fake, query["boolParam1InQuery"], fake.Lorem().Word())
				_, query["boolParam2InQuery"] = injectValueRandomly(fake, query["boolParam2InQuery"], fake.Lorem().Word())

				return routeTestCase[*booleanControllerTestActions]{
					method: http.MethodPost,
					path: fmt.Sprintf("/boolean/array-items/%v/%v",
						strings.Join(booleansToString(wantReq.BoolParam1), ",")+","+fake.Lorem().Word(),
						strings.Join(booleansToString(wantReq.BoolParam2), ",")+","+fake.Lorem().Word(),
					),
					query: query,
					body: bytes.NewBufferString(fmt.Sprintf(`{
						"boolParam1": [%v],
						"boolParam2": [%v]
					}`, fake.Lorem().Word(), fake.Lorem().Word())),
					expect: expectBindingErrors[*booleanControllerTestActions](
						[]expectedBindingError{
							// path
							{Field: "boolParam1", Location: "path", Code: "BAD_FORMAT"},
							{Field: "boolParam2", Location: "path", Code: "BAD_FORMAT"},
							// query
							{Field: "boolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "boolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},
							// body
							{Location: "body", Code: "BAD_FORMAT"},
						},
					),
				}
			})
	})

	t.Run("nullable-array-items", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.BooleanBooleanNullableArrayItemsRequest),
		) *handlers.BooleanBooleanNullableArrayItemsRequest {
			res := &handlers.BooleanBooleanNullableArrayItemsRequest{
				// path
				BoolParam1: lo.ToSlicePtr(randomBooleans(5)),
				BoolParam2: lo.ToSlicePtr(randomBooleans(5)),

				// query
				BoolParam1InQuery: lo.ToSlicePtr(randomBooleans(5)),
				BoolParam2InQuery: lo.ToSlicePtr(randomBooleans(5)),

				// body
				Payload: &models.BooleanNullableArrayItemsRequest{
					BoolParam1: lo.ToSlicePtr(randomBooleans(5)),
					BoolParam2: lo.ToSlicePtr(randomBooleans(5)),
				},
			}
			for _, opt := range opts {
				opt(res)
			}
			return res
		}

		buildQuery := func(wantReq *handlers.BooleanBooleanNullableArrayItemsRequest) url.Values {
			query := url.Values{}
			query["boolParam1InQuery"] = fromNullableItems(wantReq.BoolParam1InQuery, strconv.FormatBool)
			query["boolParam2InQuery"] = fromNullableItems(wantReq.BoolParam2InQuery, strconv.FormatBool)
			return query
		}

		buildPath := func(wantReq *handlers.BooleanBooleanNullableArrayItemsRequest) string {
			return fmt.Sprintf("/boolean/nullable-array-items/%v/%v",
				strings.Join(fromNullableItems(wantReq.BoolParam1, strconv.FormatBool), ","),
				strings.Join(fromNullableItems(wantReq.BoolParam2, strconv.FormatBool), ","),
			)
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := randomReq()

				return routeTestCase[*booleanControllerTestActions]{
					method: http.MethodPost,
					path:   buildPath(wantReq),
					query:  buildQuery(wantReq),
					body:   marshalJSONDataAsReader(t, wantReq.Payload),
					expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, wantReq, testActions.nullableBooleanArrayItems.calls[0].params)
					},
				}
			})

		runRouteTestCase(t, "should parse and bind valid null values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := randomReq(func(req *handlers.BooleanBooleanNullableArrayItemsRequest) {
					_, req.BoolParam1 = injectValueRandomly(fake, req.BoolParam1, nil)
					_, req.BoolParam2 = injectValueRandomly(fake, req.BoolParam2, nil)
					_, req.BoolParam1InQuery = injectValueRandomly(fake, req.BoolParam1InQuery, nil)
					_, req.BoolParam2InQuery = injectValueRandomly(fake, req.BoolParam2InQuery, nil)
					_, req.Payload.BoolParam1 = injectValueRandomly(fake, req.Payload.BoolParam1, nil)
					_, req.Payload.BoolParam2 = injectValueRandomly(fake, req.Payload.BoolParam2, nil)
				})

				return routeTestCase[*booleanControllerTestActions]{
					method: http.MethodPost,
					path:   buildPath(wantReq),
					query:  buildQuery(wantReq),
					body:   marshalJSONDataAsReader(t, wantReq.Payload),
					expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, wantReq, testActions.nullableBooleanArrayItems.calls[0].params)
					},
				}
			})
	})
}
