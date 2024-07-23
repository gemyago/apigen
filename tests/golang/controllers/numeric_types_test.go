package controllers

import (
	"bytes"
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
)

func TestNumericTypes(t *testing.T) {
	fake := faker.New()

	setupRouter := func() (*numericTypesControllerTestActions, http.Handler) {
		testActions := &numericTypesControllerTestActions{}
		controller := newNumericTypesController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterNumericTypesRoutes(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	type testCase = routeTestCase[*numericTypesControllerTestActions]

	t.Run("parsing", func(t *testing.T) {
		randomReq := func() *handlers.NumericTypesNumericTypesParsingRequest {
			return &handlers.NumericTypesNumericTypesParsingRequest{
				// path
				NumberAny:    fake.Float32(10, 1, 100),
				NumberFloat:  fake.Float32(10, 1, 100),
				NumberDouble: fake.Float64(10, 1, 100),
				NumberInt:    fake.Int32(),
				NumberInt32:  fake.Int32(),
				NumberInt64:  fake.Int64(),

				// query
				NumberAnyInQuery:    fake.Float32(10, 1, 100),
				NumberFloatInQuery:  fake.Float32(10, 1, 100),
				NumberDoubleInQuery: fake.Float64(10, 1, 100),
				NumberIntInQuery:    fake.Int32(),
				NumberInt32InQuery:  fake.Int32(),
				NumberInt64InQuery:  fake.Int64(),

				// body
				Payload: &models.NumericTypesParsingRequest{
					NumberAny:    fake.Float32(10, 1, 100),
					NumberFloat:  fake.Float32(10, 1, 100),
					NumberDouble: fake.Float64(10, 1, 100),
					NumberInt:    fake.Int32(),
					NumberInt32:  fake.Int32(),
					NumberInt64:  fake.Int64(),
				},
			}
		}
		runRouteTestCase(t, "should parse and bind valid values", setupRouter,
			func() routeTestCase[*numericTypesControllerTestActions] {
				wantReq := randomReq()
				query := url.Values{}
				query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
				query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
				query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
				query.Add("numberIntInQuery", strconv.FormatInt(int64(wantReq.NumberIntInQuery), 10))
				query.Add("numberInt32InQuery", strconv.FormatInt(int64(wantReq.NumberInt32InQuery), 10))
				query.Add("numberInt64InQuery", strconv.FormatInt(wantReq.NumberInt64InQuery, 10))

				return routeTestCase[*numericTypesControllerTestActions]{
					method: http.MethodPost,
					path: fmt.Sprintf("/numeric-types/parsing/%v/%v/%v/%v/%v/%v",
						wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt,
						wantReq.NumberInt32, wantReq.NumberInt64),
					query: query,
					body:  marshalJSONDataAsReader(t, wantReq.Payload),
					expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
						assert.Equal(t, 204, recorder.Code)
						assert.Equal(t, wantReq, testActions.numericTypesParsing.calls[0].params)
					},
				}
			})
		runRouteTestCase(t, "should fail if invalid values", setupRouter,
			func() routeTestCase[*numericTypesControllerTestActions] {
				query := url.Values{}
				query.Add("numberAnyInQuery", fake.Lorem().Word())
				query.Add("numberFloatInQuery", fake.Lorem().Word())
				query.Add("numberDoubleInQuery", fake.Lorem().Word())
				query.Add("numberIntInQuery", fake.Lorem().Word())
				query.Add("numberInt32InQuery", fake.Lorem().Word())
				query.Add("numberInt64InQuery", fake.Lorem().Word())

				return routeTestCase[*numericTypesControllerTestActions]{
					method: http.MethodPost,
					path: fmt.Sprintf("/numeric-types/parsing/%v/%v/%v/%v/%v/%v",
						fake.Lorem().Word(), fake.Lorem().Word(), fake.Lorem().Word(), fake.Lorem().Word(),
						fake.Lorem().Word(), fake.Lorem().Word()),
					query: query,
					body: bytes.NewBuffer(([]byte)(fmt.Sprintf(`{
						"numberAny": %v
					}`, fake.Lorem().Word()))),
					expect: expectBindingErrors[*numericTypesControllerTestActions](
						[]fieldBindingError{
							// path
							{Field: "numberAny", Location: "path", Code: "BAD_FORMAT"},
							{Field: "numberFloat", Location: "path", Code: "BAD_FORMAT"},
							{Field: "numberDouble", Location: "path", Code: "BAD_FORMAT"},
							{Field: "numberInt", Location: "path", Code: "BAD_FORMAT"},
							{Field: "numberInt32", Location: "path", Code: "BAD_FORMAT"},
							{Field: "numberInt64", Location: "path", Code: "BAD_FORMAT"},

							// query
							{Field: "numberAnyInQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "numberFloatInQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "numberDoubleInQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "numberIntInQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "numberInt32InQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "numberInt64InQuery", Location: "query", Code: "BAD_FORMAT"},

							// body
							{Field: "payload", Location: "body", Code: "BAD_FORMAT"},
						},
					),
				}
			})
	})

	t.Run("range-validation", func(t *testing.T) {
		buildQuery := func(wantReq *handlers.NumericTypesNumericTypesRangeValidationRequest) url.Values {
			query := url.Values{}
			query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
			query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
			query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
			query.Add("numberIntInQuery", strconv.FormatInt(int64(wantReq.NumberIntInQuery), 10))
			query.Add("numberInt32InQuery", strconv.FormatInt(int64(wantReq.NumberInt32InQuery), 10))
			query.Add("numberInt64InQuery", strconv.FormatInt(wantReq.NumberInt64InQuery, 10))
			return query
		}

		runRouteTestCase(t, "should validate min range", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRangeValidationRequest{
				// path
				NumberAny:    fake.Float32(5, 10, 100),
				NumberFloat:  fake.Float32(5, 20, 200),
				NumberDouble: fake.Float64(5, 30, 300),
				NumberInt:    fake.Int32Between(40, 400),
				NumberInt32:  fake.Int32Between(50, 500),
				NumberInt64:  fake.Int64Between(60, 600),

				// query
				NumberAnyInQuery:    fake.Float32(5, 10, 100),
				NumberFloatInQuery:  fake.Float32(5, 20, 200),
				NumberDoubleInQuery: fake.Float64(5, 30, 300),
				NumberIntInQuery:    fake.Int32Between(40, 400),
				NumberInt32InQuery:  fake.Int32Between(50, 500),
				NumberInt64InQuery:  fake.Int64Between(60, 600),
			}
			return testCase{
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "numberAny", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "path", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "numberAnyInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloatInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDoubleInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberIntInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
		runRouteTestCase(t, "should allow min inclusive values by default", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRangeValidationRequest{
				// path
				NumberAny:    100.01,
				NumberFloat:  200.02,
				NumberDouble: 300.03,
				NumberInt:    400,
				NumberInt32:  500,
				NumberInt64:  600,

				// query
				NumberAnyInQuery:    100.01,
				NumberFloatInQuery:  200.02,
				NumberDoubleInQuery: 300.03,
				NumberIntInQuery:    400,
				NumberInt32InQuery:  500,
				NumberInt64InQuery:  600,
			}

			return testCase{
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body)
					assert.Equal(t, wantReq, testActions.numericTypesRangeValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate max range", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRangeValidationRequest{
				// path
				NumberAny:    fake.Float32(5, 201, 1000),
				NumberFloat:  fake.Float32(5, 302, 1000),
				NumberDouble: fake.Float64(5, 403, 1000),
				NumberInt:    fake.Int32Between(500, 1000),
				NumberInt32:  fake.Int32Between(600, 1000),
				NumberInt64:  fake.Int64Between(700, 1000),

				// query
				NumberAnyInQuery:    fake.Float32(5, 201, 1000),
				NumberFloatInQuery:  fake.Float32(5, 302, 1000),
				NumberDoubleInQuery: fake.Float64(5, 403, 1000),
				NumberIntInQuery:    fake.Int32Between(500, 1000),
				NumberInt32InQuery:  fake.Int32Between(600, 1000),
				NumberInt64InQuery:  fake.Int64Between(700, 1000),
			}
			return testCase{
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "numberAny", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "path", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "numberAnyInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloatInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDoubleInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberIntInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
		runRouteTestCase(t, "should allow max inclusive values by default", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRangeValidationRequest{
				// path
				NumberAny:    200.02,
				NumberFloat:  300.03,
				NumberDouble: 400.04,
				NumberInt:    500,
				NumberInt32:  600,
				NumberInt64:  700,

				// query
				NumberAnyInQuery:    200.02,
				NumberFloatInQuery:  300.03,
				NumberDoubleInQuery: 400.04,
				NumberIntInQuery:    500,
				NumberInt32InQuery:  600,
				NumberInt64InQuery:  700,
			}

			return testCase{
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body)
					assert.Equal(t, wantReq, testActions.numericTypesRangeValidation.calls[0].params)
				},
			}
		})
	})

	t.Run("range-validation-exclusive", func(t *testing.T) {
		buildQuery := func(wantReq *handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest) url.Values {
			query := url.Values{}
			query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
			query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
			query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
			query.Add("numberIntInQuery", strconv.FormatInt(int64(wantReq.NumberIntInQuery), 10))
			query.Add("numberInt32InQuery", strconv.FormatInt(int64(wantReq.NumberInt32InQuery), 10))
			query.Add("numberInt64InQuery", strconv.FormatInt(wantReq.NumberInt64InQuery, 10))
			return query
		}

		runRouteTestCase(t, "should not allow min inclusive values", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest{
				// path
				NumberAny:    100.01,
				NumberFloat:  200.02,
				NumberDouble: 300.03,
				NumberInt:    400,
				NumberInt32:  500,
				NumberInt64:  600,

				// query
				NumberAnyInQuery:    100.01,
				NumberFloatInQuery:  200.02,
				NumberDoubleInQuery: 300.03,
				NumberIntInQuery:    400,
				NumberInt32InQuery:  500,
				NumberInt64InQuery:  600,
			}

			return testCase{
				path: fmt.Sprintf("/numeric-types/range-validation-exclusive/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt,
					wantReq.NumberInt32, wantReq.NumberInt64),
				query: buildQuery(wantReq),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "numberAny", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "path", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "numberAnyInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloatInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDoubleInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberIntInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
		runRouteTestCase(t, "should not allow max inclusive values", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest{
				// path
				NumberAny:    200.02,
				NumberFloat:  300.03,
				NumberDouble: 400.04,
				NumberInt:    500,
				NumberInt32:  600,
				NumberInt64:  700,

				// query
				NumberAnyInQuery:    200.02,
				NumberFloatInQuery:  300.03,
				NumberDoubleInQuery: 400.04,
				NumberIntInQuery:    500,
				NumberInt32InQuery:  600,
				NumberInt64InQuery:  700,
			}

			return testCase{
				path: fmt.Sprintf("/numeric-types/range-validation-exclusive/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt,
					wantReq.NumberInt32, wantReq.NumberInt64),
				query: buildQuery(wantReq),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "numberAny", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "path", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "numberAnyInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloatInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDoubleInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberIntInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
	})

	t.Run("required-validation", func(t *testing.T) {
		runRouteTestCase(t, "should ignore missing optional params", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRequiredValidationRequest{
				// query
				NumberAnyInQuery:    fake.Float32(5, 201, 1000),
				NumberFloatInQuery:  fake.Float32(5, 302, 1000),
				NumberDoubleInQuery: fake.Float64(5, 403, 1000),
				NumberIntInQuery:    fake.Int32Between(500, 1000),
				NumberInt32InQuery:  fake.Int32Between(600, 1000),
				NumberInt64InQuery:  fake.Int64Between(700, 1000),
			}

			buildQuery := func(wantReq *handlers.NumericTypesNumericTypesRequiredValidationRequest) url.Values {
				query := url.Values{}
				query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
				query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
				query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
				query.Add("numberIntInQuery", strconv.FormatInt(int64(wantReq.NumberIntInQuery), 10))
				query.Add("numberInt32InQuery", strconv.FormatInt(int64(wantReq.NumberInt32InQuery), 10))
				query.Add("numberInt64InQuery", strconv.FormatInt(wantReq.NumberInt64InQuery, 10))
				return query
			}

			return testCase{
				path:  "/numeric-types/required-validation",
				query: buildQuery(wantReq),
				expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body)
					assert.Equal(t, wantReq, testActions.numericTypesRequiredValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should parse optional params", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRequiredValidationRequest{
				// query
				NumberAnyInQuery:    fake.Float32(5, 201, 1000),
				NumberFloatInQuery:  fake.Float32(5, 302, 1000),
				NumberDoubleInQuery: fake.Float64(5, 403, 1000),
				NumberIntInQuery:    fake.Int32Between(500, 1000),
				NumberInt32InQuery:  fake.Int32Between(600, 1000),
				NumberInt64InQuery:  fake.Int64Between(700, 1000),

				OptionalNumberAnyInQuery:    fake.Float32(5, 101, 1000),
				OptionalNumberFloatInQuery:  fake.Float32(5, 202, 1000),
				OptionalNumberDoubleInQuery: fake.Float64(5, 303, 1000),
				OptionalNumberIntInQuery:    fake.Int32Between(400, 1000),
				OptionalNumberInt32InQuery:  fake.Int32Between(500, 1000),
				OptionalNumberInt64InQuery:  fake.Int64Between(600, 1000),
			}

			buildQuery := func(wantReq *handlers.NumericTypesNumericTypesRequiredValidationRequest) url.Values {
				query := url.Values{}
				query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
				query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
				query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
				query.Add("numberIntInQuery", strconv.FormatInt(int64(wantReq.NumberIntInQuery), 10))
				query.Add("numberInt32InQuery", strconv.FormatInt(int64(wantReq.NumberInt32InQuery), 10))
				query.Add("numberInt64InQuery", strconv.FormatInt(wantReq.NumberInt64InQuery, 10))
				query.Add("optionalNumberAnyInQuery", fmt.Sprint(wantReq.OptionalNumberAnyInQuery))
				query.Add("optionalNumberFloatInQuery", fmt.Sprint(wantReq.OptionalNumberFloatInQuery))
				query.Add("optionalNumberDoubleInQuery", fmt.Sprint(wantReq.OptionalNumberDoubleInQuery))
				query.Add("optionalNumberIntInQuery", strconv.FormatInt(int64(wantReq.OptionalNumberIntInQuery), 10))
				query.Add("optionalNumberInt32InQuery", strconv.FormatInt(int64(wantReq.OptionalNumberInt32InQuery), 10))
				query.Add("optionalNumberInt64InQuery", strconv.FormatInt(wantReq.OptionalNumberInt64InQuery, 10))
				return query
			}

			return testCase{
				path:  "/numeric-types/required-validation",
				query: buildQuery(wantReq),
				expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body)
					assert.Equal(t, wantReq, testActions.numericTypesRequiredValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should validate optional params", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesRequiredValidationRequest{
				// query
				NumberAnyInQuery:    fake.Float32(5, 201, 1000),
				NumberFloatInQuery:  fake.Float32(5, 302, 1000),
				NumberDoubleInQuery: fake.Float64(5, 403, 1000),
				NumberIntInQuery:    fake.Int32Between(500, 1000),
				NumberInt32InQuery:  fake.Int32Between(600, 1000),
				NumberInt64InQuery:  fake.Int64Between(700, 1000),

				OptionalNumberAnyInQuery:    fake.Float32(5, 1, 100),
				OptionalNumberFloatInQuery:  fake.Float32(5, 1, 200),
				OptionalNumberDoubleInQuery: fake.Float64(5, 1, 300),
				OptionalNumberIntInQuery:    fake.Int32Between(1, 399),
				OptionalNumberInt32InQuery:  fake.Int32Between(1, 499),
				OptionalNumberInt64InQuery:  fake.Int64Between(1, 599),
			}

			buildQuery := func(wantReq *handlers.NumericTypesNumericTypesRequiredValidationRequest) url.Values {
				query := url.Values{}
				query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
				query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
				query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
				query.Add("numberIntInQuery", strconv.FormatInt(int64(wantReq.NumberIntInQuery), 10))
				query.Add("numberInt32InQuery", strconv.FormatInt(int64(wantReq.NumberInt32InQuery), 10))
				query.Add("numberInt64InQuery", strconv.FormatInt(wantReq.NumberInt64InQuery, 10))
				query.Add("optionalNumberAnyInQuery", fmt.Sprint(wantReq.OptionalNumberAnyInQuery))
				query.Add("optionalNumberFloatInQuery", fmt.Sprint(wantReq.OptionalNumberFloatInQuery))
				query.Add("optionalNumberDoubleInQuery", fmt.Sprint(wantReq.OptionalNumberDoubleInQuery))
				query.Add("optionalNumberIntInQuery", strconv.FormatInt(int64(wantReq.OptionalNumberIntInQuery), 10))
				query.Add("optionalNumberInt32InQuery", strconv.FormatInt(int64(wantReq.OptionalNumberInt32InQuery), 10))
				query.Add("optionalNumberInt64InQuery", strconv.FormatInt(wantReq.OptionalNumberInt64InQuery, 10))
				return query
			}

			return testCase{
				path:  "/numeric-types/required-validation",
				query: buildQuery(wantReq),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]fieldBindingError{
						{Field: "optionalNumberAnyInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalNumberFloatInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalNumberDoubleInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalNumberIntInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalNumberInt32InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalNumberInt64InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
		runRouteTestCase(t, "should validate empty params", setupRouter, func() testCase {
			buildQuery := func() url.Values {
				query := url.Values{}
				query.Add("numberAnyInQuery", "")
				query.Add("numberFloatInQuery", "")
				query.Add("numberDoubleInQuery", "")
				query.Add("numberIntInQuery", "")
				query.Add("numberInt32InQuery", "")
				query.Add("numberInt64InQuery", "")
				query.Add("optionalNumberAnyInQuery", "")
				query.Add("optionalNumberFloatInQuery", "")
				query.Add("optionalNumberDoubleInQuery", "")
				query.Add("optionalNumberIntInQuery", "")
				query.Add("optionalNumberInt32InQuery", "")
				query.Add("optionalNumberInt64InQuery", "")
				return query
			}

			return testCase{
				path:  "/numeric-types/required-validation",
				query: buildQuery(),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]fieldBindingError{
						{Field: "numberAnyInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "numberFloatInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "numberDoubleInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "numberIntInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "numberInt32InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "numberInt64InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalNumberAnyInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalNumberFloatInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalNumberDoubleInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalNumberIntInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalNumberInt32InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalNumberInt64InQuery", Location: "query", Code: "BAD_FORMAT"},
					},
				),
			}
		})
		runRouteTestCase(t, "should ensure required params", setupRouter, func() testCase {
			return testCase{
				path:  "/numeric-types/required-validation",
				query: url.Values{},
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]fieldBindingError{
						{Field: "numberAnyInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "numberFloatInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "numberDoubleInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "numberIntInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "numberInt32InQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "numberInt64InQuery", Location: "query", Code: "INVALID_REQUIRED"},
					},
				),
			}
		})
	})
}
