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

	randomFloat32 := func(minVal, maxVal float32) float32 {
		return fake.Float32(10, int(minVal), int(maxVal))
	}

	randomFloat64 := func(minVal, maxVal float64) float64 {
		return fake.Float64(10, int(minVal), int(maxVal))
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
						[]expectedBindingError{
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
							{Location: "body", Code: "BAD_FORMAT"},
						},
					),
				}
			})
	})

	t.Run("array-items", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.NumericTypesNumericTypesArrayItemsRequest),
		) *handlers.NumericTypesNumericTypesArrayItemsRequest {
			res := &handlers.NumericTypesNumericTypesArrayItemsRequest{
				// path
				NumberAny:    randomNumbers(5, randomFloat32, 101, 200),
				NumberFloat:  randomNumbers(5, randomFloat32, 201, 300),
				NumberDouble: randomNumbers(5, randomFloat64, 301, 400),
				NumberInt:    randomNumbers(5, fake.Int32Between, 401, 500),
				NumberInt32:  randomNumbers(5, fake.Int32Between, 501, 600),
				NumberInt64:  randomNumbers(5, fake.Int64Between, 601, 700),

				// query
				NumberAnyInQuery:    randomNumbers(5, randomFloat32, 101, 200),
				NumberFloatInQuery:  randomNumbers(5, randomFloat32, 201, 300),
				NumberDoubleInQuery: randomNumbers(5, randomFloat64, 301, 400),
				NumberIntInQuery:    randomNumbers(5, fake.Int32Between, 401, 500),
				NumberInt32InQuery:  randomNumbers(5, fake.Int32Between, 501, 600),
				NumberInt64InQuery:  randomNumbers(5, fake.Int64Between, 601, 700),

				// body
				Payload: &models.NumericTypesArrayItemsRequest{
					NumberAny:    randomNumbers(5, randomFloat32, 101, 200),
					NumberFloat:  randomNumbers(5, randomFloat32, 201, 300),
					NumberDouble: randomNumbers(5, randomFloat64, 301, 400),
					NumberInt:    randomNumbers(5, fake.Int32Between, 401, 500),
					NumberInt32:  randomNumbers(5, fake.Int32Between, 501, 600),
					NumberInt64:  randomNumbers(5, fake.Int64Between, 601, 700),
				},
			}
			for _, opt := range opts {
				opt(res)
			}
			return res
		}

		buildPath := func(wantReq *handlers.NumericTypesNumericTypesArrayItemsRequest) string {
			return fmt.Sprintf(
				"/numeric-types/array-items/%v/%v/%v/%v/%v/%v",
				strings.Join(numbersToString(wantReq.NumberAny), ","),
				strings.Join(numbersToString(wantReq.NumberFloat), ","),
				strings.Join(numbersToString(wantReq.NumberDouble), ","),
				strings.Join(numbersToString(wantReq.NumberInt), ","),
				strings.Join(numbersToString(wantReq.NumberInt32), ","),
				strings.Join(numbersToString(wantReq.NumberInt64), ","),
			)
		}

		buildQuery := func(wantReq *handlers.NumericTypesNumericTypesArrayItemsRequest) url.Values {
			query := url.Values{}
			query["numberAnyInQuery"] = numbersToString(wantReq.NumberAnyInQuery)
			query["numberFloatInQuery"] = numbersToString(wantReq.NumberFloatInQuery)
			query["numberDoubleInQuery"] = numbersToString(wantReq.NumberDoubleInQuery)
			query["numberIntInQuery"] = numbersToString(wantReq.NumberIntInQuery)
			query["numberInt32InQuery"] = numbersToString(wantReq.NumberInt32InQuery)
			query["numberInt64InQuery"] = numbersToString(wantReq.NumberInt64InQuery)
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter,
			func() routeTestCase[*numericTypesControllerTestActions] {
				wantReq := randomReq()
				return routeTestCase[*numericTypesControllerTestActions]{
					method: http.MethodPost,
					path:   buildPath(wantReq),
					query:  buildQuery(wantReq),
					body:   marshalJSONDataAsReader(t, wantReq.Payload),
					expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
						assert.Equal(t, 204, recorder.Code)
						assert.Equal(t, wantReq, testActions.numericTypesArrayItems.calls[0].params)
					},
				}
			})

		runRouteTestCase(t, "should fail if bad values", setupRouter,
			func() routeTestCase[*numericTypesControllerTestActions] {
				wantReq := randomReq()
				query := buildQuery(wantReq)

				_, query["numberAnyInQuery"] = injectValueRandomly(fake, query["numberAnyInQuery"], fake.Lorem().Word())
				_, query["numberFloatInQuery"] = injectValueRandomly(fake, query["numberFloatInQuery"], fake.Lorem().Word())
				_, query["numberDoubleInQuery"] = injectValueRandomly(fake, query["numberDoubleInQuery"], fake.Lorem().Word())
				_, query["numberIntInQuery"] = injectValueRandomly(fake, query["numberIntInQuery"], fake.Lorem().Word())
				_, query["numberInt32InQuery"] = injectValueRandomly(fake, query["numberInt32InQuery"], fake.Lorem().Word())
				_, query["numberInt64InQuery"] = injectValueRandomly(fake, query["numberInt64InQuery"], fake.Lorem().Word())

				return routeTestCase[*numericTypesControllerTestActions]{
					method: http.MethodPost,
					path: fmt.Sprintf(
						"/numeric-types/array-items/%v/%v/%v/%v/%v/%v",
						strings.Join(numbersToString(wantReq.NumberAny), ",")+","+fake.Lorem().Word(),
						strings.Join(numbersToString(wantReq.NumberFloat), ",")+","+fake.Lorem().Word(),
						strings.Join(numbersToString(wantReq.NumberDouble), ",")+","+fake.Lorem().Word(),
						strings.Join(numbersToString(wantReq.NumberInt), ",")+","+fake.Lorem().Word(),
						strings.Join(numbersToString(wantReq.NumberInt32), ",")+","+fake.Lorem().Word(),
						strings.Join(numbersToString(wantReq.NumberInt64), ",")+","+fake.Lorem().Word(),
					),
					query: query,
					body: bytes.NewBuffer(([]byte)(fmt.Sprintf(`{
						"numberAny": %v
					}`, fake.Lorem().Word()))),
					expect: expectBindingErrors[*numericTypesControllerTestActions](
						[]expectedBindingError{
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
							{Location: "body", Code: "BAD_FORMAT"},
						},
					),
				}
			})

		runRouteTestCase(t, "should fail if invalid values", setupRouter,
			func() routeTestCase[*numericTypesControllerTestActions] {
				badIndices := make([]int, 18)
				wantReq := randomReq(func(req *handlers.NumericTypesNumericTypesArrayItemsRequest) {
					badIndices[0], req.NumberAny = injectValueRandomly(fake, req.NumberAny, fake.Float32(10, 1, 100))
					badIndices[1], req.NumberFloat = injectValueRandomly(fake, req.NumberFloat, fake.Float32(10, 301, 500))
					badIndices[2], req.NumberDouble = injectValueRandomly(fake, req.NumberDouble, fake.Float64(10, 1, 300))
					badIndices[3], req.NumberInt = injectValueRandomly(fake, req.NumberInt, fake.Int32Between(1, 400))
					badIndices[4], req.NumberInt32 = injectValueRandomly(fake, req.NumberInt32, fake.Int32Between(1, 500))
					badIndices[5], req.NumberInt64 = injectValueRandomly(fake, req.NumberInt64, fake.Int64Between(1, 600))

					badIndices[6], req.NumberAnyInQuery = injectValueRandomly(fake, req.NumberAnyInQuery, fake.Float32(10, 1, 100))
					badIndices[7], req.NumberFloatInQuery = injectValueRandomly(fake, req.NumberFloatInQuery, fake.Float32(10, 301, 500))
					badIndices[8], req.NumberDoubleInQuery = injectValueRandomly(
						fake, req.NumberDoubleInQuery, fake.Float64(10, 1, 300),
					)
					badIndices[9], req.NumberIntInQuery = injectValueRandomly(fake, req.NumberIntInQuery, fake.Int32Between(1, 400))
					badIndices[10], req.NumberInt32InQuery = injectValueRandomly(
						fake, req.NumberInt32InQuery, fake.Int32Between(1, 500),
					)
					badIndices[11], req.NumberInt64InQuery = injectValueRandomly(
						fake, req.NumberInt64InQuery, fake.Int64Between(1, 600),
					)

					badIndices[12], req.Payload.NumberAny = injectValueRandomly(
						fake, req.Payload.NumberAny, fake.Float32(10, 1, 100),
					)
					badIndices[13], req.Payload.NumberFloat = injectValueRandomly(
						fake, req.Payload.NumberFloat, fake.Float32(10, 301, 500),
					)
					badIndices[14], req.Payload.NumberDouble = injectValueRandomly(
						fake, req.Payload.NumberDouble, fake.Float64(10, 1, 300),
					)
					badIndices[15], req.Payload.NumberInt = injectValueRandomly(fake, req.Payload.NumberInt, fake.Int32Between(1, 400))
					badIndices[16], req.Payload.NumberInt32 = injectValueRandomly(
						fake, req.Payload.NumberInt32, fake.Int32Between(1, 500),
					)
					badIndices[17], req.Payload.NumberInt64 = injectValueRandomly(
						fake, req.Payload.NumberInt64, fake.Int64Between(1, 600),
					)
				})

				return routeTestCase[*numericTypesControllerTestActions]{
					method: http.MethodPost,
					path:   buildPath(wantReq),
					query:  buildQuery(wantReq),
					body:   marshalJSONDataAsReader(t, wantReq.Payload),
					expect: expectBindingErrors[*numericTypesControllerTestActions](
						[]expectedBindingError{
							// path
							{Field: strconv.Itoa(badIndices[0]), Location: "path.numberAny", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[1]), Location: "path.numberFloat", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[2]), Location: "path.numberDouble", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[3]), Location: "path.numberInt", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[4]), Location: "path.numberInt32", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[5]), Location: "path.numberInt64", Code: "INVALID_OUT_OF_RANGE"},
							// query
							{Field: strconv.Itoa(badIndices[6]), Location: "query.numberAnyInQuery", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[7]), Location: "query.numberFloatInQuery", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[8]), Location: "query.numberDoubleInQuery", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[9]), Location: "query.numberIntInQuery", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[10]), Location: "query.numberInt32InQuery", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[11]), Location: "query.numberInt64InQuery", Code: "INVALID_OUT_OF_RANGE"},
							// body
							{Field: strconv.Itoa(badIndices[12]), Location: "body.numberAny", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[13]), Location: "body.numberFloat", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[14]), Location: "body.numberDouble", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[15]), Location: "body.numberInt", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[16]), Location: "body.numberInt32", Code: "INVALID_OUT_OF_RANGE"},
							{Field: strconv.Itoa(badIndices[17]), Location: "body.numberInt64", Code: "INVALID_OUT_OF_RANGE"},
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

				Payload: &models.NumericTypesRangeValidationRequest{
					NumberAny:    fake.Float32(5, 10, 100),
					NumberFloat:  fake.Float32(5, 20, 200),
					NumberDouble: fake.Float64(5, 30, 300),
					NumberInt:    fake.Int32Between(40, 400),
					NumberInt32:  fake.Int32Between(50, 500),
					NumberInt64:  fake.Int64Between(60, 600),
				},
			}
			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]expectedBindingError{
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

						// body
						{Field: "numberAny", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
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

				Payload: &models.NumericTypesRangeValidationRequest{
					NumberAny:    100.01,
					NumberFloat:  200.02,
					NumberDouble: 300.03,
					NumberInt:    400,
					NumberInt32:  500,
					NumberInt64:  600,
				},
			}

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
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

				// body
				Payload: &models.NumericTypesRangeValidationRequest{
					NumberAny:    fake.Float32(5, 201, 1000),
					NumberFloat:  fake.Float32(5, 302, 1000),
					NumberDouble: fake.Float64(5, 403, 1000),
					NumberInt:    fake.Int32Between(500, 1000),
					NumberInt32:  fake.Int32Between(600, 1000),
					NumberInt64:  fake.Int64Between(700, 1000),
				},
			}
			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]expectedBindingError{
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

						// body
						{Field: "numberAny", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
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

				Payload: &models.NumericTypesRangeValidationRequest{
					NumberAny:    200.02,
					NumberFloat:  300.03,
					NumberDouble: 400.04,
					NumberInt:    500,
					NumberInt32:  600,
					NumberInt64:  700,
				},
			}

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/numeric-types/range-validation/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32,
					wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
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

				// body
				Payload: &models.NumericTypesRangeValidationExclusiveRequest{
					NumberAny:    100.01,
					NumberFloat:  200.02,
					NumberDouble: 300.03,
					NumberInt:    400,
					NumberInt32:  500,
					NumberInt64:  600,
				},
			}

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf("/numeric-types/range-validation-exclusive/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt,
					wantReq.NumberInt32, wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]expectedBindingError{
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

						// body
						{Field: "numberAny", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
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

				Payload: &models.NumericTypesRangeValidationExclusiveRequest{
					NumberAny:    200.02,
					NumberFloat:  300.03,
					NumberDouble: 400.04,
					NumberInt:    500,
					NumberInt32:  600,
					NumberInt64:  700,
				},
			}

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf("/numeric-types/range-validation-exclusive/%v/%v/%v/%v/%v/%v",
					wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt,
					wantReq.NumberInt32, wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]expectedBindingError{
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

						// body
						{Field: "numberAny", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
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
					[]expectedBindingError{
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
					[]expectedBindingError{
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
					[]expectedBindingError{
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

	t.Run("nullable", func(t *testing.T) {
		buildQuery := func(wantReq *handlers.NumericTypesNumericTypesNullableRequest) url.Values {
			query := url.Values{}
			query.Add("numberAnyInQuery", fmt.Sprint(lo.FromPtr(wantReq.NumberAnyInQuery)))
			query.Add("optionalNumberAnyInQuery", fmt.Sprint(lo.FromPtr(wantReq.OptionalNumberAnyInQuery)))
			query.Add("numberFloatInQuery", fmt.Sprint(lo.FromPtr(wantReq.NumberFloatInQuery)))
			query.Add("numberDoubleInQuery", fmt.Sprint(lo.FromPtr(wantReq.NumberDoubleInQuery)))
			query.Add("numberIntInQuery", strconv.FormatInt(int64(lo.FromPtr(wantReq.NumberIntInQuery)), 10))
			query.Add("numberInt32InQuery", strconv.FormatInt(int64(lo.FromPtr(wantReq.NumberInt32InQuery)), 10))
			query.Add("numberInt64InQuery", strconv.FormatInt(lo.FromPtr(wantReq.NumberInt64InQuery), 10))
			return query
		}

		runRouteTestCase(t, "should bind valid values", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesNullableRequest{
				// path
				NumberAny:    lo.ToPtr(fake.Float32(5, 100, 200)),
				NumberFloat:  lo.ToPtr(fake.Float32(5, 200, 300)),
				NumberDouble: lo.ToPtr(fake.Float64(5, 300, 400)),
				NumberInt:    lo.ToPtr(fake.Int32Between(400, 500)),
				NumberInt32:  lo.ToPtr(fake.Int32Between(500, 600)),
				NumberInt64:  lo.ToPtr(fake.Int64Between(600, 700)),

				// query
				NumberAnyInQuery:         lo.ToPtr(fake.Float32(5, 100, 200)),
				OptionalNumberAnyInQuery: lo.ToPtr(fake.Float32(5, 100, 200)),
				NumberFloatInQuery:       lo.ToPtr(fake.Float32(5, 200, 300)),
				NumberDoubleInQuery:      lo.ToPtr(fake.Float64(5, 300, 400)),
				NumberIntInQuery:         lo.ToPtr(fake.Int32Between(400, 500)),
				NumberInt32InQuery:       lo.ToPtr(fake.Int32Between(500, 600)),
				NumberInt64InQuery:       lo.ToPtr(fake.Int64Between(600, 700)),

				Payload: &models.NumericTypesNullableRequest{
					NumberAny:         lo.ToPtr(fake.Float32(5, 100, 200)),
					OptionalNumberAny: lo.ToPtr(fake.Float32(5, 100, 200)),
					NumberFloat:       lo.ToPtr(fake.Float32(5, 200, 300)),
					NumberDouble:      lo.ToPtr(fake.Float64(5, 300, 400)),
					NumberInt:         lo.ToPtr(fake.Int32Between(400, 500)),
					NumberInt32:       lo.ToPtr(fake.Int32Between(500, 600)),
					NumberInt64:       lo.ToPtr(fake.Int64Between(600, 700)),
				},
			}
			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/numeric-types/nullable/%v/%v/%v/%v/%v/%v",
					*wantReq.NumberAny, *wantReq.NumberFloat, *wantReq.NumberDouble, *wantReq.NumberInt, *wantReq.NumberInt32,
					*wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
				expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, wantReq, testActions.numericTypesNullable.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate invalid values", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesNullableRequest{
				// path
				NumberAny:    lo.ToPtr(fake.Float32(5, 201, 1000)),
				NumberFloat:  lo.ToPtr(fake.Float32(5, 302, 1000)),
				NumberDouble: lo.ToPtr(fake.Float64(5, 403, 1000)),
				NumberInt:    lo.ToPtr(fake.Int32Between(500, 1000)),
				NumberInt32:  lo.ToPtr(fake.Int32Between(600, 1000)),
				NumberInt64:  lo.ToPtr(fake.Int64Between(700, 1000)),

				// query
				NumberAnyInQuery:         lo.ToPtr(fake.Float32(5, 201, 1000)),
				OptionalNumberAnyInQuery: lo.ToPtr(fake.Float32(5, 201, 1000)),
				NumberFloatInQuery:       lo.ToPtr(fake.Float32(5, 302, 1000)),
				NumberDoubleInQuery:      lo.ToPtr(fake.Float64(5, 403, 1000)),
				NumberIntInQuery:         lo.ToPtr(fake.Int32Between(500, 1000)),
				NumberInt32InQuery:       lo.ToPtr(fake.Int32Between(600, 1000)),
				NumberInt64InQuery:       lo.ToPtr(fake.Int64Between(700, 1000)),

				// body
				Payload: &models.NumericTypesNullableRequest{
					NumberAny:         lo.ToPtr(fake.Float32(5, 201, 1000)),
					OptionalNumberAny: lo.ToPtr(fake.Float32(5, 201, 1000)),
					NumberFloat:       lo.ToPtr(fake.Float32(5, 302, 1000)),
					NumberDouble:      lo.ToPtr(fake.Float64(5, 403, 1000)),
					NumberInt:         lo.ToPtr(fake.Int32Between(500, 1000)),
					NumberInt32:       lo.ToPtr(fake.Int32Between(600, 1000)),
					NumberInt64:       lo.ToPtr(fake.Int64Between(700, 1000)),
				},
			}
			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/numeric-types/nullable/%v/%v/%v/%v/%v/%v",
					*wantReq.NumberAny, *wantReq.NumberFloat, *wantReq.NumberDouble, *wantReq.NumberInt, *wantReq.NumberInt32,
					*wantReq.NumberInt64),
				query: buildQuery(wantReq),
				body:  marshalJSONDataAsReader(t, wantReq.Payload),
				expect: expectBindingErrors[*numericTypesControllerTestActions](
					[]expectedBindingError{
						// path
						{Field: "numberAny", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "path", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "numberAnyInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalNumberAnyInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloatInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDoubleInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberIntInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},

						// body
						{Field: "numberAny", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalNumberAny", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberFloat", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberDouble", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt32", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "numberInt64", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})

		runRouteTestCase(t, "should accept null values", setupRouter, func() testCase {
			wantReq := &handlers.NumericTypesNumericTypesNullableRequest{
				Payload: &models.NumericTypesNullableRequest{},
			}
			query := url.Values{}
			query.Add("numberAnyInQuery", "null")
			query.Add("optionalNumberAnyInQuery", "null")
			query.Add("numberFloatInQuery", "null")
			query.Add("numberDoubleInQuery", "null")
			query.Add("numberIntInQuery", "null")
			query.Add("numberInt32InQuery", "null")
			query.Add("numberInt64InQuery", "null")
			return testCase{
				method: http.MethodPost,
				path:   "/numeric-types/nullable/null/null/null/null/null/null",
				query:  query,
				body: bytes.NewBufferString(`{
					"numberAny": null,
					"optionalNumberAny": null,
					"numberFloat": null,
					"numberDouble": null,
					"numberInt": null,
					"numberInt32": null,
					"numberInt64": null
				}`),
				expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, wantReq, testActions.numericTypesNullable.calls[0].params)
				},
			}
		})
	})
}
