package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/jaswdr/faker"
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
		handlers.MountStringTypesRoutes(controller, handlers.NewHttpApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	type testCase = routeTestCase[*stringTypesControllerTestActions]

	t.Run("parsing", func(t *testing.T) {
		randomReq := func() *handlers.StringTypesStringTypesParsingRequest {
			return &handlers.StringTypesStringTypesParsingRequest{
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
			}
		}
		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := url.Values{}
			query.Add("unformattedStrInQuery", originalReq.UnformattedStrInQuery)
			query.Add("customFormatStrInQuery", originalReq.CustomFormatStrInQuery)
			query.Add("dateStrInQuery", originalReq.DateStrInQuery.Format(time.RFC3339Nano))
			query.Add("dateTimeStrInQuery", originalReq.DateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("byteStrInQuery", originalReq.ByteStrInQuery)

			return testCase{
				path:  fmt.Sprintf("/string-types/parsing/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.RFC3339Nano), originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					assert.Equal(t, 204, recorder.Code)

					wantReq := *originalReq
					wantReq.DateStr = originalReq.DateStr.Truncate(24 * time.Hour)
					wantReq.DateStrInQuery = originalReq.DateStrInQuery.Truncate(24 * time.Hour)
					assert.Equal(t, &wantReq, testActions.stringTypesParsing.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should fail if bad values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := url.Values{}
			query.Add("unformattedStrInQuery", originalReq.UnformattedStrInQuery)
			query.Add("customFormatStrInQuery", originalReq.CustomFormatStrInQuery)
			query.Add("dateStrInQuery", fake.Lorem().Word())
			query.Add("dateTimeStrInQuery", fake.Lorem().Word())
			query.Add("byteStrInQuery", originalReq.ByteStrInQuery)

			return testCase{
				path:  fmt.Sprintf("/string-types/parsing/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr, fake.Lorem().Word(), fake.Lorem().Word(), originalReq.ByteStr),
				query: query,
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]handlers.BindingError{
						// path
						{Field: "dateStr", Location: "path", Code: handlers.ErrBadValueFormat},
						{Field: "dateTimeStr", Location: "path", Code: handlers.ErrBadValueFormat},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: handlers.ErrBadValueFormat},
						{Field: "dateTimeStrInQuery", Location: "query", Code: handlers.ErrBadValueFormat},
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
			query.Add("dateStrInQuery", wantReq.DateStrInQuery.Format(time.RFC3339Nano))
			query.Add("dateTimeStrInQuery", wantReq.DateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("byteStrInQuery", wantReq.ByteStrInQuery)
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			return testCase{
				path:  fmt.Sprintf("/string-types/range-validation/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.RFC3339Nano), originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = originalReq.DateStr.Truncate(24 * time.Hour)
					wantReq.DateStrInQuery = originalReq.DateStrInQuery.Truncate(24 * time.Hour)
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
			})
			query := buildQuery(originalReq)

			return testCase{
				path:  fmt.Sprintf("/string-types/range-validation/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.RFC3339Nano), originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]handlers.BindingError{
						// path
						{Field: "unformattedStr", Location: "path", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "customFormatStr", Location: "path", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "byteStr", Location: "path", Code: handlers.ErrInvalidValueOutOfRange},

						// query
						{Field: "unformattedStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "customFormatStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "byteStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
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
			})
			query := buildQuery(originalReq)

			return testCase{
				path:  fmt.Sprintf("/string-types/range-validation/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.RFC3339Nano), originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = originalReq.DateStr.Truncate(24 * time.Hour)
					wantReq.DateStrInQuery = originalReq.DateStrInQuery.Truncate(24 * time.Hour)
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
			})
			query := buildQuery(originalReq)

			return testCase{
				path:  fmt.Sprintf("/string-types/range-validation/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.RFC3339Nano), originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]handlers.BindingError{
						// path
						{Field: "unformattedStr", Location: "path", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "customFormatStr", Location: "path", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "byteStr", Location: "path", Code: handlers.ErrInvalidValueOutOfRange},

						// query
						{Field: "unformattedStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "customFormatStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "byteStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
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
			})
			query := buildQuery(originalReq)

			return testCase{
				path:  fmt.Sprintf("/string-types/range-validation/%v/%v/%v/%v/%v", originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.RFC3339Nano), originalReq.DateTimeStr.Format(time.RFC3339Nano), originalReq.ByteStr),
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = originalReq.DateStr.Truncate(24 * time.Hour)
					wantReq.DateStrInQuery = originalReq.DateStrInQuery.Truncate(24 * time.Hour)
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
			query.Add("dateStrInQuery", wantReq.DateStrInQuery.Format(time.RFC3339Nano))
			query.Add("dateTimeStrInQuery", wantReq.DateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("byteStrInQuery", wantReq.ByteStrInQuery)

			query.Add("optionalUnformattedStrInQuery", wantReq.OptionalUnformattedStrInQuery)
			query.Add("optionalCustomFormatStrInQuery", wantReq.OptionalCustomFormatStrInQuery)
			query.Add("optionalDateStrInQuery", wantReq.OptionalDateStrInQuery.Format(time.RFC3339Nano))
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
					wantReq.DateStrInQuery = originalReq.DateStrInQuery.Truncate(24 * time.Hour)
					wantReq.OptionalDateStrInQuery = originalReq.OptionalDateStrInQuery.Truncate(24 * time.Hour)
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
					wantReq.DateStrInQuery = originalReq.DateStrInQuery.Truncate(24 * time.Hour)
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
					[]handlers.BindingError{
						{Field: "unformattedStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "customFormatStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "dateStrInQuery", Location: "query", Code: handlers.ErrBadValueFormat},
						{Field: "dateTimeStrInQuery", Location: "query", Code: handlers.ErrBadValueFormat},
						{Field: "byteStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},

						{Field: "optionalUnformattedStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "optionalCustomFormatStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
						{Field: "optionalDateStrInQuery", Location: "query", Code: handlers.ErrBadValueFormat},
						{Field: "optionalDateTimeStrInQuery", Location: "query", Code: handlers.ErrBadValueFormat},
						{Field: "optionalByteStrInQuery", Location: "query", Code: handlers.ErrInvalidValueOutOfRange},
					},
				),
			}
		})
	})
}
