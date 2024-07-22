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
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestStringTypes(t *testing.T) {
	fake := faker.New()

	setupRouter := func() (*stringTypesControllerTestActions, http.Handler) {
		testActions := &stringTypesControllerTestActions{}
		controller := newStringTypesController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterStringTypesRoutes(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	type testCase = routeTestCase[*stringTypesControllerTestActions]

	t.Run("parsing", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesParsingRequest),
		) *handlers.StringTypesStringTypesParsingRequest {
			req := &handlers.StringTypesStringTypesParsingRequest{
				// path
				UnformattedStr:  fake.Lorem().Word(),
				CustomFormatStr: fake.Lorem().Word(),
				DateStr:         fake.Time().Time(time.Now()),
				DateTimeStr:     fake.Time().Time(time.Now()),
				ByteStr:         fake.BinaryString().BinaryString(10),

				// query
				UnformattedStrInQuery:  fake.Lorem().Word(),
				CustomFormatStrInQuery: fake.Lorem().Word(),
				DateStrInQuery:         fake.Time().Time(time.Now()),
				DateTimeStrInQuery:     fake.Time().Time(time.Now()),
				ByteStrInQuery:         fake.BinaryString().BinaryString(10),

				// body
				Payload: &models.StringTypesParsingRequest{
					UnformattedStr:  fake.Lorem().Word(),
					CustomFormatStr: fake.Lorem().Word(),
					DateStr:         fake.Time().Time(time.Now()),
					DateTimeStr:     fake.Time().Time(time.Now()),
					ByteStr:         fake.BinaryString().BinaryString(10),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *handlers.StringTypesStringTypesParsingRequest) url.Values {
			query := url.Values{}
			query.Add("unformattedStrInQuery", req.UnformattedStrInQuery)
			query.Add("customFormatStrInQuery", req.CustomFormatStrInQuery)
			query.Add("dateStrInQuery", req.DateStrInQuery.Format(time.DateOnly))
			query.Add("dateTimeStrInQuery", req.DateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("byteStrInQuery", req.ByteStrInQuery)
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/parsing/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr,
				),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = lo.Must(time.Parse(time.DateOnly, wantReq.DateStr.Format(time.DateOnly)))
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, wantReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.stringTypesParsing.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should parse time with locale", setupRouter, func() testCase {
			location := time.FixedZone("", int((time.Duration(fake.IntBetween(2, 30)) * time.Minute).Seconds()))
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesParsingRequest) {
				req.DateTimeStr = req.DateTimeStr.In(location)
				req.DateTimeStrInQuery = req.DateTimeStrInQuery.In(location)
				req.Payload.DateTimeStr = req.Payload.DateTimeStr.In(location)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/parsing/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr,
				),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					assert.Equal(t, wantReq.DateTimeStr, testActions.stringTypesParsing.calls[0].params.DateTimeStr)
					assert.Equal(t, wantReq.DateTimeStrInQuery, testActions.stringTypesParsing.calls[0].params.DateTimeStrInQuery)
				},
			}
		})
		runRouteTestCase(t, "should fail if bad values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)
			query.Set("dateStrInQuery", fake.Lorem().Word())
			query.Set("dateTimeStrInQuery", fake.Lorem().Word())

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/parsing/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, fake.Lorem().Word(),
					fake.Lorem().Word(), originalReq.ByteStr),
				query: query,
				body: bytes.NewBuffer(([]byte)(fmt.Sprintf(`{
					"unformattedStr": %v
				}`, fake.IntBetween(10, 100)))),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "dateStr", Location: "path", Code: "BAD_FORMAT"},
						{Field: "dateTimeStr", Location: "path", Code: "BAD_FORMAT"},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "dateTimeStrInQuery", Location: "query", Code: "BAD_FORMAT"},

						// body
						{Field: "payload", Location: "body", Code: "BAD_FORMAT"},
					},
				),
			}
		})
		runRouteTestCase(t, "should fail if date formatted with time", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)
			query.Set("dateStrInQuery", originalReq.DateStrInQuery.Format(time.RFC3339))

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/parsing/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStrInQuery.Format(time.RFC3339),
					originalReq.DateTimeStrInQuery.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				body: bytes.NewBuffer(([]byte)(fmt.Sprintf(`{
					"unformattedStr": "%v",
					"customFormatStr": "%v",
					"dateStr": "%v",
					"dateTimeStr": "%v",
					"byteStr": "%v"
				}`,
					originalReq.Payload.UnformattedStr,
					originalReq.Payload.CustomFormatStr,
					originalReq.Payload.DateStr.Format(time.RFC3339),
					originalReq.Payload.DateTimeStr.Format(time.RFC3339),
					originalReq.Payload.ByteStr,
				))),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "dateStr", Location: "path", Code: "BAD_FORMAT"},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
					},
				),
			}
		})
	})

	t.Run("range-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesRangeValidationRequest),
		) *handlers.StringTypesStringTypesRangeValidationRequest {
			req := &handlers.StringTypesStringTypesRangeValidationRequest{
				// path
				UnformattedStr:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
				CustomFormatStr: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
				DateStr:         fake.Time().Time(time.Now()),
				DateTimeStr:     fake.Time().Time(time.Now()),
				ByteStr:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),

				// query
				UnformattedStrInQuery:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
				CustomFormatStrInQuery: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
				DateStrInQuery:         fake.Time().Time(time.Now()),
				DateTimeStrInQuery:     fake.Time().Time(time.Now()),
				ByteStrInQuery:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),

				// body
				Payload: &models.StringTypesRangeValidationRequest{
					UnformattedStr:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
					CustomFormatStr: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
					DateStr:         fake.Time().Time(time.Now()),
					DateTimeStr:     fake.Time().Time(time.Now()),
					ByteStr:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(wantReq *handlers.StringTypesStringTypesRangeValidationRequest) url.Values {
			query := url.Values{}
			query.Add("unformattedStrInQuery", wantReq.UnformattedStrInQuery)
			query.Add("customFormatStrInQuery", wantReq.CustomFormatStrInQuery)
			query.Add("dateStrInQuery", wantReq.DateStrInQuery.Format(time.DateOnly))
			query.Add("dateTimeStrInQuery", wantReq.DateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("byteStrInQuery", wantReq.ByteStrInQuery)
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/range-validation/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr,
					originalReq.DateStr.Format(time.DateOnly), originalReq.DateTimeStr.Format(time.RFC3339Nano),
					originalReq.ByteStr),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = lo.Must(time.Parse(time.DateOnly, originalReq.DateStr.Format(time.DateOnly)))
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.StringTypesRangeValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should validate min length", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesRangeValidationRequest) {
				req.UnformattedStr = fake.BinaryString().BinaryString(9)
				req.CustomFormatStr = fake.BinaryString().BinaryString(19)
				req.ByteStr = fake.BinaryString().BinaryString(29)

				req.UnformattedStrInQuery = fake.BinaryString().BinaryString(9)
				req.CustomFormatStrInQuery = fake.BinaryString().BinaryString(19)
				req.ByteStrInQuery = fake.BinaryString().BinaryString(29)

				req.Payload.UnformattedStr = fake.BinaryString().BinaryString(9)
				req.Payload.CustomFormatStr = fake.BinaryString().BinaryString(19)
				req.Payload.ByteStr = fake.BinaryString().BinaryString(29)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/range-validation/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "unformattedStr", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "customFormatStr", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "byteStr", Location: "path", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "customFormatStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "byteStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "unformattedStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "customFormatStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "byteStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
		runRouteTestCase(t, "should allow inclusive min length", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesRangeValidationRequest) {
				req.UnformattedStr = fake.BinaryString().BinaryString(10)
				req.CustomFormatStr = fake.BinaryString().BinaryString(20)
				req.ByteStr = fake.BinaryString().BinaryString(30)

				req.UnformattedStrInQuery = fake.BinaryString().BinaryString(10)
				req.CustomFormatStrInQuery = fake.BinaryString().BinaryString(20)
				req.ByteStrInQuery = fake.BinaryString().BinaryString(30)

				req.Payload.UnformattedStr = fake.BinaryString().BinaryString(10)
				req.Payload.CustomFormatStr = fake.BinaryString().BinaryString(20)
				req.Payload.ByteStr = fake.BinaryString().BinaryString(30)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/range-validation/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = lo.Must(time.Parse(time.DateOnly, originalReq.DateStr.Format(time.DateOnly)))
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.StringTypesRangeValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should validate max length", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesRangeValidationRequest) {
				req.UnformattedStr = fake.Lorem().Text(21)
				req.CustomFormatStr = fake.Lorem().Text(31)
				req.ByteStr = fake.Lorem().Text(41)

				req.UnformattedStrInQuery = fake.Lorem().Text(21)
				req.CustomFormatStrInQuery = fake.Lorem().Text(31)
				req.ByteStrInQuery = fake.Lorem().Text(41)

				req.Payload.UnformattedStr = fake.Lorem().Text(21)
				req.Payload.CustomFormatStr = fake.Lorem().Text(31)
				req.Payload.ByteStr = fake.Lorem().Text(41)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/range-validation/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]fieldBindingError{
						// path
						{Field: "unformattedStr", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "customFormatStr", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "byteStr", Location: "path", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "customFormatStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "byteStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},

						// body
						{Field: "unformattedStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "customFormatStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "byteStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
		runRouteTestCase(t, "should allow inclusive max length", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesRangeValidationRequest) {
				req.UnformattedStr = fake.BinaryString().BinaryString(20)
				req.CustomFormatStr = fake.BinaryString().BinaryString(30)
				req.ByteStr = fake.BinaryString().BinaryString(40)

				req.UnformattedStrInQuery = fake.BinaryString().BinaryString(20)
				req.CustomFormatStrInQuery = fake.BinaryString().BinaryString(30)
				req.ByteStrInQuery = fake.BinaryString().BinaryString(40)

				req.Payload.UnformattedStr = fake.BinaryString().BinaryString(20)
				req.Payload.CustomFormatStr = fake.BinaryString().BinaryString(30)
				req.Payload.ByteStr = fake.BinaryString().BinaryString(40)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/range-validation/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = lo.Must(time.Parse(time.DateOnly, originalReq.DateStr.Format(time.DateOnly)))
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.StringTypesRangeValidation.calls[0].params)
				},
			}
		})
	})

	t.Run("required-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesRequiredValidationRequest),
		) *handlers.StringTypesStringTypesRequiredValidationRequest {
			req := &handlers.StringTypesStringTypesRequiredValidationRequest{
				// query
				UnformattedStrInQuery:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
				CustomFormatStrInQuery: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
				DateStrInQuery:         fake.Time().Time(time.Now()),
				DateTimeStrInQuery:     fake.Time().Time(time.Now()),
				ByteStrInQuery:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),

				OptionalUnformattedStrInQuery:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
				OptionalCustomFormatStrInQuery: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
				OptionalDateStrInQuery:         fake.Time().Time(time.Now()),
				OptionalDateTimeStrInQuery:     fake.Time().Time(time.Now()),
				OptionalByteStrInQuery:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(wantReq *handlers.StringTypesStringTypesRequiredValidationRequest) url.Values {
			query := url.Values{}
			query.Add("unformattedStrInQuery", wantReq.UnformattedStrInQuery)
			query.Add("customFormatStrInQuery", wantReq.CustomFormatStrInQuery)
			query.Add("dateStrInQuery", wantReq.DateStrInQuery.Format(time.DateOnly))
			query.Add("dateTimeStrInQuery", wantReq.DateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("byteStrInQuery", wantReq.ByteStrInQuery)

			query.Add("optionalUnformattedStrInQuery", wantReq.OptionalUnformattedStrInQuery)
			query.Add("optionalCustomFormatStrInQuery", wantReq.OptionalCustomFormatStrInQuery)
			query.Add("optionalDateStrInQuery", wantReq.OptionalDateStrInQuery.Format(time.DateOnly))
			query.Add("optionalDateTimeStrInQuery", wantReq.OptionalDateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("optionalByteStrInQuery", wantReq.OptionalByteStrInQuery)
			return query
		}

		runRouteTestCase(t, "should parse optional params", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)
			return testCase{
				path:  "/string-types/required-validation",
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					wantReq.OptionalDateStrInQuery = lo.Must(
						time.Parse(time.DateOnly, originalReq.OptionalDateStrInQuery.Format(time.DateOnly)),
					)
					assert.Equal(t, &wantReq, testActions.StringTypesRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should ignore missing optional params", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesRequiredValidationRequest) {
				req.OptionalUnformattedStrInQuery = ""
				req.OptionalCustomFormatStrInQuery = ""
				req.OptionalDateStrInQuery = time.Time{}
				req.OptionalDateTimeStrInQuery = time.Time{}
				req.OptionalByteStrInQuery = ""
			})

			query := buildQuery(originalReq)
			query.Del("optionalUnformattedStrInQuery")
			query.Del("optionalCustomFormatStrInQuery")
			query.Del("optionalDateStrInQuery")
			query.Del("optionalDateTimeStrInQuery")
			query.Del("optionalByteStrInQuery")

			return testCase{
				path:  "/string-types/required-validation",
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.StringTypesRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate empty params", setupRouter, func() testCase {
			query := url.Values{}

			query.Set("unformattedStrInQuery", "")
			query.Set("customFormatStrInQuery", "")
			query.Set("dateStrInQuery", "")
			query.Set("dateTimeStrInQuery", "")
			query.Set("byteStrInQuery", "")

			query.Set("optionalUnformattedStrInQuery", "")
			query.Set("optionalCustomFormatStrInQuery", "")
			query.Set("optionalDateStrInQuery", "")
			query.Set("optionalDateTimeStrInQuery", "")
			query.Set("optionalByteStrInQuery", "")

			return testCase{
				path:  "/string-types/required-validation",
				query: query,
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]fieldBindingError{
						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "customFormatStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "dateTimeStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "byteStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},

						{Field: "optionalUnformattedStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalCustomFormatStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalDateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalDateTimeStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalByteStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})
	})

	t.Run("pattern-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesPatternValidationRequest),
		) *handlers.StringTypesStringTypesPatternValidationRequest {
			req := &handlers.StringTypesStringTypesPatternValidationRequest{
				// path
				UnformattedStr:  strings.Repeat(strconv.Itoa(fake.RandomDigit()), 10),
				CustomFormatStr: strings.Repeat(strconv.Itoa(fake.RandomDigit()), 20),
				DateStr:         fake.Time().Time(time.Now()),
				DateTimeStr:     fake.Time().Time(time.Now()),

				// query
				UnformattedStrInQuery:  strings.Repeat(strconv.Itoa(fake.RandomDigit()), 10),
				CustomFormatStrInQuery: strings.Repeat(strconv.Itoa(fake.RandomDigit()), 20),
				DateStrInQuery:         fake.Time().Time(time.Now()),
				DateTimeStrInQuery:     fake.Time().Time(time.Now()),
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(wantReq *handlers.StringTypesStringTypesPatternValidationRequest) url.Values {
			query := url.Values{}
			query.Add("unformattedStrInQuery", wantReq.UnformattedStrInQuery)
			query.Add("customFormatStrInQuery", wantReq.CustomFormatStrInQuery)
			query.Add("dateStrInQuery", wantReq.DateStrInQuery.Format(time.DateOnly))
			query.Add("dateTimeStrInQuery", wantReq.DateTimeStrInQuery.Format(time.RFC3339Nano))
			return query
		}

		runRouteTestCase(t, "should parse valid params", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)
			return testCase{
				path: fmt.Sprintf(
					"/string-types/pattern-validation/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano)),
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = lo.Must(time.Parse(time.DateOnly, originalReq.DateStr.Format(time.DateOnly)))
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.StringTypesPatternValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate pattern", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesPatternValidationRequest) {
				// path
				req.UnformattedStr = strings.Repeat(strconv.Itoa(fake.RandomDigit()), 11)
				req.CustomFormatStr = strings.Repeat(strconv.Itoa(fake.RandomDigit()), 21)

				// query
				req.UnformattedStrInQuery = strings.Repeat(strconv.Itoa(fake.RandomDigit()), 11)
				req.CustomFormatStrInQuery = strings.Repeat(strconv.Itoa(fake.RandomDigit()), 21)
			})
			query := buildQuery(originalReq)
			return testCase{
				path: fmt.Sprintf(
					"/string-types/pattern-validation/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano)),
				query: query,
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]fieldBindingError{
						{Field: "unformattedStr", Location: "path", Code: "INVALID"},
						{Field: "customFormatStr", Location: "path", Code: "INVALID"},

						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID"},
						{Field: "customFormatStrInQuery", Location: "query", Code: "INVALID"},
					},
				),
			}
		})
	})
}
