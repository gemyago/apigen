package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
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

	type testCase = routeTestCase[*stringTypesControllerTestActions]
	setupRouter := func(_ testCase) (*stringTypesControllerTestActions, http.Handler) {
		testActions := &stringTypesControllerTestActions{}
		controller := &stringTypesController{testActions: testActions}
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterStringTypesRoutesV2(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	randomStrings := func(n, minLength, maxLength int) []string {
		strs := make([]string, n)
		for i := range n {
			length := fake.IntBetween(minLength, maxLength)
			strs[i] = fake.RandomStringWithLength(length)
		}
		return strs
	}

	randomDates := func(n int) []time.Time {
		dates := make([]time.Time, n)
		for i := range n {
			dates[i] = fake.Time().Time(time.Now())
		}
		return dates
	}

	timesToStr := func(v []time.Time, format string) []string {
		return lo.Map(v, func(v time.Time, _ int) string { return v.Format(format) })
	}

	nullableTimesToStr := func(v []*time.Time, format string) []string {
		return lo.Map(v, func(v *time.Time, _ int) string {
			if v == nil {
				return "null"
			}
			return v.Format(format)
		})
	}

	datePartOnly := func(v []time.Time) []time.Time {
		return lo.Map(v, func(v time.Time, _ int) time.Time {
			return lo.Must(time.Parse(time.DateOnly, v.Format(time.DateOnly)))
		})
	}

	nullableDatePartOnly := func(v []*time.Time) []*time.Time {
		return lo.Map(v, func(v *time.Time, _ int) *time.Time {
			if v == nil {
				return nil
			}
			return lo.ToPtr(lo.Must(time.Parse(time.DateOnly, v.Format(time.DateOnly))))
		})
	}

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
					[]expectedBindingError{
						// path
						{Field: "dateStr", Location: "path", Code: "BAD_FORMAT"},
						{Field: "dateTimeStr", Location: "path", Code: "BAD_FORMAT"},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "dateTimeStrInQuery", Location: "query", Code: "BAD_FORMAT"},

						// body
						{Location: "body", Code: "BAD_FORMAT"},
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
					[]expectedBindingError{
						// path
						{Field: "dateStr", Location: "path", Code: "BAD_FORMAT"},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
					},
				),
			}
		})
	})

	t.Run("nullable-parsing", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesNullableParsingRequest),
		) *handlers.StringTypesStringTypesNullableParsingRequest {
			req := &handlers.StringTypesStringTypesNullableParsingRequest{
				// path
				UnformattedStr:  lo.ToPtr(fake.Lorem().Word()),
				CustomFormatStr: lo.ToPtr(fake.Lorem().Word()),
				DateStr:         lo.ToPtr(fake.Time().Time(time.Now())),
				DateTimeStr:     lo.ToPtr(fake.Time().Time(time.Now())),
				ByteStr:         lo.ToPtr(fake.BinaryString().BinaryString(10)),

				// query
				UnformattedStrInQuery:  lo.ToPtr(fake.Lorem().Word()),
				CustomFormatStrInQuery: lo.ToPtr(fake.Lorem().Word()),
				DateStrInQuery:         lo.ToPtr(fake.Time().Time(time.Now())),
				DateTimeStrInQuery:     lo.ToPtr(fake.Time().Time(time.Now())),
				ByteStrInQuery:         lo.ToPtr(fake.BinaryString().BinaryString(10)),

				// body
				Payload: &models.StringTypesNullableParsingRequest{
					UnformattedStr:  lo.ToPtr(fake.Lorem().Word()),
					CustomFormatStr: lo.ToPtr(fake.Lorem().Word()),
					DateStr:         lo.ToPtr(fake.Time().Time(time.Now())),
					DateTimeStr:     lo.ToPtr(fake.Time().Time(time.Now())),
					ByteStr:         lo.ToPtr(fake.BinaryString().BinaryString(10)),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *handlers.StringTypesStringTypesNullableParsingRequest) url.Values {
			query := url.Values{}
			query.Add("unformattedStrInQuery", lo.FromPtrOr(req.UnformattedStrInQuery, "null"))
			query.Add("customFormatStrInQuery", lo.FromPtrOr(req.CustomFormatStrInQuery, "null"))
			query.Add("dateStrInQuery", req.DateStrInQuery.Format(time.DateOnly))
			query.Add("dateTimeStrInQuery", req.DateTimeStrInQuery.Format(time.RFC3339Nano))
			query.Add("byteStrInQuery", lo.FromPtrOr(req.ByteStrInQuery, "null"))
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/nullable-parsing/%v/%v/%v/%v/%v",
					*originalReq.UnformattedStr, *originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), *originalReq.ByteStr,
				),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = lo.ToPtr(
						lo.Must(time.Parse(time.DateOnly, lo.FromPtr(wantReq.DateStr).Format(time.DateOnly))),
					)
					wantReq.DateStrInQuery = lo.ToPtr(
						lo.Must(time.Parse(time.DateOnly, lo.FromPtr(wantReq.DateStrInQuery).Format(time.DateOnly))),
					)
					assert.Equal(t, &wantReq, testActions.stringTypesNullableParsing.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should parse and bind null values", setupRouter, func() testCase {
			originalReq := &handlers.StringTypesStringTypesNullableParsingRequest{
				Payload: &models.StringTypesNullableParsingRequest{},
			}
			query := url.Values{}
			query.Add("unformattedStrInQuery", "null")
			query.Add("customFormatStrInQuery", "null")
			query.Add("dateStrInQuery", "null")
			query.Add("dateTimeStrInQuery", "null")
			query.Add("byteStrInQuery", "null")

			return testCase{
				method: http.MethodPost,
				path:   "/string-types/nullable-parsing/null/null/null/null/null",
				query:  query,
				body: bytes.NewBufferString(`{
					"unformattedStr": null,
					"customFormatStr": null,
					"dateStr": null,
					"dateTimeStr": null,
					"byteStr": null
				}`),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, originalReq, testActions.stringTypesNullableParsing.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should parse time with locale", setupRouter, func() testCase {
			location := time.FixedZone("", int((time.Duration(fake.IntBetween(2, 30)) * time.Minute).Seconds()))
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesNullableParsingRequest) {
				req.DateTimeStr = lo.ToPtr(lo.FromPtr(req.DateTimeStr).In(location))
				req.DateTimeStrInQuery = lo.ToPtr(lo.FromPtr(req.DateTimeStrInQuery).In(location))
				req.Payload.DateTimeStr = lo.ToPtr(lo.FromPtr(req.Payload.DateTimeStr).In(location))
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/nullable-parsing/%v/%v/%v/%v/%v",
					*originalReq.UnformattedStr, *originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano), *originalReq.ByteStr,
				),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					assert.Equal(t, wantReq.DateTimeStr, testActions.stringTypesNullableParsing.calls[0].params.DateTimeStr)
					assert.Equal(t,
						wantReq.DateTimeStrInQuery,
						testActions.stringTypesNullableParsing.calls[0].params.DateTimeStrInQuery,
					)
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
					"/string-types/nullable-parsing/%v/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, fake.Lorem().Word(),
					fake.Lorem().Word(), originalReq.ByteStr),
				query: query,
				body: bytes.NewBuffer(([]byte)(fmt.Sprintf(`{
					"unformattedStr": %v
				}`, fake.IntBetween(10, 100)))),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						// path
						{Field: "dateStr", Location: "path", Code: "BAD_FORMAT"},
						{Field: "dateTimeStr", Location: "path", Code: "BAD_FORMAT"},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "dateTimeStrInQuery", Location: "query", Code: "BAD_FORMAT"},

						// body
						{Location: "body", Code: "BAD_FORMAT"},
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
					"/string-types/nullable-parsing/%v/%v/%v/%v/%v",
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
					[]expectedBindingError{
						// path
						{Field: "dateStr", Location: "path", Code: "BAD_FORMAT"},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
					},
				),
			}
		})
	})

	t.Run("arrays-parsing", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesArraysParsingRequest),
		) *handlers.StringTypesStringTypesArraysParsingRequest {
			req := &handlers.StringTypesStringTypesArraysParsingRequest{
				// path
				UnformattedStr:  randomStrings(3, 10, 20),
				CustomFormatStr: randomStrings(3, 10, 20),
				DateStr:         randomDates(3),
				DateTimeStr:     randomDates(3),
				ByteStr:         randomStrings(3, 10, 20),

				// query
				UnformattedStrInQuery:  randomStrings(3, 10, 20),
				CustomFormatStrInQuery: randomStrings(3, 10, 20),
				DateStrInQuery:         randomDates(3),
				DateTimeStrInQuery:     randomDates(3),
				ByteStrInQuery: []string{
					fake.BinaryString().BinaryString(10), fake.BinaryString().BinaryString(10), fake.BinaryString().BinaryString(10),
				},

				// body
				Payload: &models.StringTypesArraysParsingRequest{
					UnformattedStr:  randomStrings(3, 10, 20),
					CustomFormatStr: randomStrings(3, 10, 20),
					DateStr:         randomDates(3),
					DateTimeStr:     randomDates(3),
					ByteStr:         randomStrings(3, 10, 20),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *handlers.StringTypesStringTypesArraysParsingRequest) url.Values {
			query := url.Values{}
			query["unformattedStrInQuery"] = req.UnformattedStrInQuery
			query["customFormatStrInQuery"] = req.CustomFormatStrInQuery
			query["dateStrInQuery"] = timesToStr(req.DateStrInQuery, time.DateOnly)
			query["dateTimeStrInQuery"] = timesToStr(req.DateTimeStrInQuery, time.RFC3339Nano)
			query["byteStrInQuery"] = req.ByteStrInQuery
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/arrays-parsing/%v/%v/%v/%v/%v",
					strings.Join(originalReq.UnformattedStr, ","),
					strings.Join(originalReq.CustomFormatStr, ","),
					strings.Join(timesToStr(originalReq.DateStr, time.DateOnly), ","),
					strings.Join(timesToStr(originalReq.DateTimeStr, time.RFC3339Nano), ","),
					strings.Join(originalReq.ByteStr, ","),
				),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = datePartOnly(originalReq.DateStr)
					wantReq.DateStrInQuery = datePartOnly(originalReq.DateStrInQuery)
					assert.Equal(t, &wantReq, testActions.stringTypesArraysParsing.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should fail if bad values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			_, badDateStrValues := injectValueRandomly(fake, timesToStr(originalReq.DateStr, time.DateOnly), fake.Lorem().Word())
			_, badDateTimeStrValues := injectValueRandomly(
				fake, timesToStr(originalReq.DateTimeStr, time.RFC3339Nano), fake.Lorem().Word(),
			)

			query["dateStrInQuery"] = badDateStrValues
			query["dateTimeStrInQuery"] = badDateTimeStrValues

			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/arrays-parsing/%v/%v/%v/%v/%v",
					strings.Join(originalReq.UnformattedStr, ","),
					strings.Join(originalReq.CustomFormatStr, ","),
					strings.Join(badDateStrValues, ","),
					strings.Join(badDateTimeStrValues, ","),
					strings.Join(originalReq.ByteStr, ","),
				),
				query: query,
				body: bytes.NewBuffer(([]byte)(fmt.Sprintf(`{
					"unformattedStr": [%v]
				}`, fake.IntBetween(10, 100)))),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						// path
						{Field: "dateStr", Location: "path", Code: "BAD_FORMAT"},
						{Field: "dateTimeStr", Location: "path", Code: "BAD_FORMAT"},

						// query
						{Field: "dateStrInQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "dateTimeStrInQuery", Location: "query", Code: "BAD_FORMAT"},

						// body
						{Location: "body", Code: "BAD_FORMAT"},
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
					assert.Equal(t, &wantReq, testActions.stringTypesRangeValidation.calls[0].params)
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
					[]expectedBindingError{
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
					assert.Equal(t, &wantReq, testActions.stringTypesRangeValidation.calls[0].params)
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
					[]expectedBindingError{
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
					assert.Equal(t, &wantReq, testActions.stringTypesRangeValidation.calls[0].params)
				},
			}
		})
	})

	t.Run("array-items-range-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesArrayItemsRangeValidationRequest),
		) *handlers.StringTypesStringTypesArrayItemsRangeValidationRequest {
			req := &handlers.StringTypesStringTypesArrayItemsRangeValidationRequest{
				// path
				UnformattedStr:  randomStrings(3, 10, 20),
				CustomFormatStr: randomStrings(3, 20, 30),
				DateStr:         randomDates(3),
				DateTimeStr:     randomDates(3),
				ByteStr:         randomStrings(3, 30, 40),

				// query
				UnformattedStrInQuery:  randomStrings(3, 10, 20),
				CustomFormatStrInQuery: randomStrings(3, 20, 30),
				DateStrInQuery:         randomDates(3),
				DateTimeStrInQuery:     randomDates(3),
				ByteStrInQuery:         randomStrings(3, 30, 40),

				// body
				Payload: &models.StringTypesArrayItemsRangeValidationRequest{
					UnformattedStr:  randomStrings(3, 10, 20),
					CustomFormatStr: randomStrings(3, 20, 30),
					DateStr:         randomDates(3),
					DateTimeStr:     randomDates(3),
					ByteStr:         randomStrings(3, 30, 40),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *handlers.StringTypesStringTypesArrayItemsRangeValidationRequest) url.Values {
			query := url.Values{}
			query["unformattedStrInQuery"] = req.UnformattedStrInQuery
			query["customFormatStrInQuery"] = req.CustomFormatStrInQuery
			query["dateStrInQuery"] = timesToStr(req.DateStrInQuery, time.DateOnly)
			query["dateTimeStrInQuery"] = timesToStr(req.DateTimeStrInQuery, time.RFC3339Nano)
			query["byteStrInQuery"] = req.ByteStrInQuery
			return query
		}

		buildPath := func(req *handlers.StringTypesStringTypesArrayItemsRangeValidationRequest) string {
			return fmt.Sprintf(
				"/string-types/array-items-range-validation/%v/%v/%v/%v/%v",
				strings.Join(req.UnformattedStr, ","),
				strings.Join(req.CustomFormatStr, ","),
				strings.Join(timesToStr(req.DateStr, time.DateOnly), ","),
				strings.Join(timesToStr(req.DateTimeStr, time.RFC3339Nano), ","),
				strings.Join(req.ByteStr, ","),
			)
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = datePartOnly(originalReq.DateStr)
					wantReq.DateStrInQuery = datePartOnly(originalReq.DateStrInQuery)
					assert.Equal(t, &wantReq, testActions.stringTypesArrayItemsRangeValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate min length", setupRouter, func() testCase {
			badIndices := make([]int, 9)
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesArrayItemsRangeValidationRequest) {
				badIndices[0], req.UnformattedStr = injectValueRandomly(fake, req.UnformattedStr, fake.RandomStringWithLength(9))
				badIndices[1], req.CustomFormatStr = injectValueRandomly(fake, req.CustomFormatStr, fake.RandomStringWithLength(19))
				badIndices[2], req.ByteStr = injectValueRandomly(fake, req.ByteStr, fake.RandomStringWithLength(29))

				badIndices[3], req.UnformattedStrInQuery = injectValueRandomly(
					fake, req.UnformattedStrInQuery, fake.RandomStringWithLength(9),
				)
				badIndices[4], req.CustomFormatStrInQuery = injectValueRandomly(
					fake, req.CustomFormatStrInQuery, fake.RandomStringWithLength(19),
				)
				badIndices[5], req.ByteStrInQuery = injectValueRandomly(
					fake, req.ByteStrInQuery, fake.RandomStringWithLength(29),
				)

				badIndices[6], req.Payload.UnformattedStr = injectValueRandomly(
					fake, req.Payload.UnformattedStr, fake.RandomStringWithLength(9),
				)
				badIndices[7], req.Payload.CustomFormatStr = injectValueRandomly(
					fake, req.Payload.CustomFormatStr, fake.RandomStringWithLength(19),
				)
				badIndices[8], req.Payload.ByteStr = injectValueRandomly(
					fake, req.Payload.ByteStr, fake.RandomStringWithLength(29),
				)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						// path
						{Field: strconv.Itoa(badIndices[0]), Location: "path.unformattedStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[1]), Location: "path.customFormatStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[2]), Location: "path.byteStr", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: strconv.Itoa(badIndices[3]), Location: "query.unformattedStrInQuery", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[4]), Location: "query.customFormatStrInQuery", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[5]), Location: "query.byteStrInQuery", Code: "INVALID_OUT_OF_RANGE"},

						// body
						{Field: strconv.Itoa(badIndices[6]), Location: "body.unformattedStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[7]), Location: "body.customFormatStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[8]), Location: "body.byteStr", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})

		runRouteTestCase(t, "should allow inclusive min length", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesArrayItemsRangeValidationRequest) {
				_, req.UnformattedStr = injectValueRandomly(fake, req.UnformattedStr, fake.RandomStringWithLength(10))
				_, req.CustomFormatStr = injectValueRandomly(fake, req.CustomFormatStr, fake.RandomStringWithLength(20))
				_, req.ByteStr = injectValueRandomly(fake, req.ByteStr, fake.RandomStringWithLength(30))

				_, req.UnformattedStrInQuery = injectValueRandomly(
					fake, req.UnformattedStrInQuery, fake.RandomStringWithLength(10),
				)
				_, req.CustomFormatStrInQuery = injectValueRandomly(
					fake, req.CustomFormatStrInQuery, fake.RandomStringWithLength(20),
				)
				_, req.ByteStrInQuery = injectValueRandomly(
					fake, req.ByteStrInQuery, fake.RandomStringWithLength(30),
				)

				_, req.Payload.UnformattedStr = injectValueRandomly(
					fake, req.Payload.UnformattedStr, fake.RandomStringWithLength(10),
				)
				_, req.Payload.CustomFormatStr = injectValueRandomly(
					fake, req.Payload.CustomFormatStr, fake.RandomStringWithLength(20),
				)
				_, req.Payload.ByteStr = injectValueRandomly(
					fake, req.Payload.ByteStr, fake.RandomStringWithLength(30),
				)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = datePartOnly(originalReq.DateStr)
					wantReq.DateStrInQuery = datePartOnly(originalReq.DateStrInQuery)
					assert.Equal(t, &wantReq, testActions.stringTypesArrayItemsRangeValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate max length", setupRouter, func() testCase {
			badIndices := make([]int, 9)
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesArrayItemsRangeValidationRequest) {
				badIndices[0], req.UnformattedStr = injectValueRandomly(fake, req.UnformattedStr, fake.RandomStringWithLength(9))
				badIndices[1], req.CustomFormatStr = injectValueRandomly(fake, req.CustomFormatStr, fake.RandomStringWithLength(19))
				badIndices[2], req.ByteStr = injectValueRandomly(fake, req.ByteStr, fake.RandomStringWithLength(29))

				badIndices[3], req.UnformattedStrInQuery = injectValueRandomly(
					fake, req.UnformattedStrInQuery, fake.RandomStringWithLength(21),
				)
				badIndices[4], req.CustomFormatStrInQuery = injectValueRandomly(
					fake, req.CustomFormatStrInQuery, fake.RandomStringWithLength(31),
				)
				badIndices[5], req.ByteStrInQuery = injectValueRandomly(
					fake, req.ByteStrInQuery, fake.RandomStringWithLength(41),
				)

				badIndices[6], req.Payload.UnformattedStr = injectValueRandomly(
					fake, req.Payload.UnformattedStr, fake.RandomStringWithLength(21),
				)
				badIndices[7], req.Payload.CustomFormatStr = injectValueRandomly(
					fake, req.Payload.CustomFormatStr, fake.RandomStringWithLength(31),
				)
				badIndices[8], req.Payload.ByteStr = injectValueRandomly(
					fake, req.Payload.ByteStr, fake.RandomStringWithLength(41),
				)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						// path
						{Field: strconv.Itoa(badIndices[0]), Location: "path.unformattedStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[1]), Location: "path.customFormatStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[2]), Location: "path.byteStr", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: strconv.Itoa(badIndices[3]), Location: "query.unformattedStrInQuery", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[4]), Location: "query.customFormatStrInQuery", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[5]), Location: "query.byteStrInQuery", Code: "INVALID_OUT_OF_RANGE"},

						// body
						{Field: strconv.Itoa(badIndices[6]), Location: "body.unformattedStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[7]), Location: "body.customFormatStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[8]), Location: "body.byteStr", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})

		runRouteTestCase(t, "should allow inclusive max length", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesArrayItemsRangeValidationRequest) {
				_, req.UnformattedStr = injectValueRandomly(fake, req.UnformattedStr, fake.RandomStringWithLength(10))
				_, req.CustomFormatStr = injectValueRandomly(fake, req.CustomFormatStr, fake.RandomStringWithLength(20))
				_, req.ByteStr = injectValueRandomly(fake, req.ByteStr, fake.RandomStringWithLength(30))

				_, req.UnformattedStrInQuery = injectValueRandomly(
					fake, req.UnformattedStrInQuery, fake.RandomStringWithLength(20),
				)
				_, req.CustomFormatStrInQuery = injectValueRandomly(
					fake, req.CustomFormatStrInQuery, fake.RandomStringWithLength(30),
				)
				_, req.ByteStrInQuery = injectValueRandomly(
					fake, req.ByteStrInQuery, fake.RandomStringWithLength(40),
				)

				_, req.Payload.UnformattedStr = injectValueRandomly(
					fake, req.Payload.UnformattedStr, fake.RandomStringWithLength(20),
				)
				_, req.Payload.CustomFormatStr = injectValueRandomly(
					fake, req.Payload.CustomFormatStr, fake.RandomStringWithLength(30),
				)
				_, req.Payload.ByteStr = injectValueRandomly(
					fake, req.Payload.ByteStr, fake.RandomStringWithLength(40),
				)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = datePartOnly(originalReq.DateStr)
					wantReq.DateStrInQuery = datePartOnly(originalReq.DateStrInQuery)
					assert.Equal(t, &wantReq, testActions.stringTypesArrayItemsRangeValidation.calls[0].params)
				},
			}
		})
	})

	t.Run("nullable-array-items", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesNullableArrayItemsRequest),
		) *handlers.StringTypesStringTypesNullableArrayItemsRequest {
			req := &handlers.StringTypesStringTypesNullableArrayItemsRequest{
				// path
				UnformattedStr:  lo.ToSlicePtr(randomStrings(3, 10, 20)),
				CustomFormatStr: lo.ToSlicePtr(randomStrings(3, 20, 30)),
				DateStr:         lo.ToSlicePtr(randomDates(3)),
				DateTimeStr:     lo.ToSlicePtr(randomDates(3)),
				ByteStr:         lo.ToSlicePtr(randomStrings(3, 30, 40)),

				// query
				UnformattedStrInQuery:  lo.ToSlicePtr(randomStrings(3, 10, 20)),
				CustomFormatStrInQuery: lo.ToSlicePtr(randomStrings(3, 20, 30)),
				DateStrInQuery:         lo.ToSlicePtr(randomDates(3)),
				DateTimeStrInQuery:     lo.ToSlicePtr(randomDates(3)),
				ByteStrInQuery:         lo.ToSlicePtr(randomStrings(3, 30, 40)),

				// body
				Payload: &models.StringTypesNullableArrayItemsRequest{
					UnformattedStr:  lo.ToSlicePtr(randomStrings(3, 10, 20)),
					CustomFormatStr: lo.ToSlicePtr(randomStrings(3, 20, 30)),
					DateStr:         lo.ToSlicePtr(randomDates(3)),
					DateTimeStr:     lo.ToSlicePtr(randomDates(3)),
					ByteStr:         lo.ToSlicePtr(randomStrings(3, 30, 40)),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *handlers.StringTypesStringTypesNullableArrayItemsRequest) url.Values {
			query := url.Values{}
			query["unformattedStrInQuery"] = fromNullableItems(req.UnformattedStrInQuery)
			query["customFormatStrInQuery"] = fromNullableItems(req.CustomFormatStrInQuery)
			query["dateStrInQuery"] = nullableTimesToStr(req.DateStrInQuery, time.DateOnly)
			query["dateTimeStrInQuery"] = nullableTimesToStr(req.DateTimeStrInQuery, time.RFC3339Nano)
			query["byteStrInQuery"] = fromNullableItems(req.ByteStrInQuery)
			return query
		}

		buildPath := func(req *handlers.StringTypesStringTypesNullableArrayItemsRequest) string {
			return fmt.Sprintf(
				"/string-types/nullable-array-items/%v/%v/%v/%v/%v",
				strings.Join(fromNullableItems(req.UnformattedStr), ","),
				strings.Join(fromNullableItems(req.CustomFormatStr), ","),
				strings.Join(nullableTimesToStr(req.DateStr, time.DateOnly), ","),
				strings.Join(nullableTimesToStr(req.DateTimeStr, time.RFC3339Nano), ","),
				strings.Join(fromNullableItems(req.ByteStr), ","),
			)
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = nullableDatePartOnly(originalReq.DateStr)
					wantReq.DateStrInQuery = nullableDatePartOnly(originalReq.DateStrInQuery)
					assert.Equal(t, &wantReq, testActions.stringTypesNullableArrayItems.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should parse and bind valid null values", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesNullableArrayItemsRequest) {
				_, req.UnformattedStr = injectValueRandomly(fake, req.UnformattedStr, nil)
				_, req.CustomFormatStr = injectValueRandomly(fake, req.CustomFormatStr, nil)
				_, req.ByteStr = injectValueRandomly(fake, req.ByteStr, nil)

				_, req.UnformattedStrInQuery = injectValueRandomly(
					fake, req.UnformattedStrInQuery, nil,
				)
				_, req.CustomFormatStrInQuery = injectValueRandomly(
					fake, req.CustomFormatStrInQuery, nil,
				)
				_, req.ByteStrInQuery = injectValueRandomly(
					fake, req.ByteStrInQuery, nil,
				)

				_, req.Payload.UnformattedStr = injectValueRandomly(
					fake, req.Payload.UnformattedStr, nil,
				)
				_, req.Payload.CustomFormatStr = injectValueRandomly(
					fake, req.Payload.CustomFormatStr, nil,
				)
				_, req.Payload.ByteStr = injectValueRandomly(
					fake, req.Payload.ByteStr, nil,
				)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					wantReq := *originalReq
					wantReq.DateStr = nullableDatePartOnly(originalReq.DateStr)
					wantReq.DateStrInQuery = nullableDatePartOnly(originalReq.DateStrInQuery)
					assert.Equal(t, &wantReq, testActions.stringTypesNullableArrayItems.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate bad values", setupRouter, func() testCase {
			badIndices := make([]int, 9)
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesNullableArrayItemsRequest) {
				badIndices[0], req.UnformattedStr = injectValueRandomly(
					fake, req.UnformattedStr, lo.ToPtr(fake.RandomStringWithLength(9)),
				)
				badIndices[1], req.CustomFormatStr = injectValueRandomly(fake,
					req.CustomFormatStr, lo.ToPtr("!"+fake.RandomStringWithLength(21)),
				)
				badIndices[2], req.ByteStr = injectValueRandomly(
					fake, req.ByteStr, lo.ToPtr(fake.RandomStringWithLength(41)),
				)

				badIndices[3], req.UnformattedStrInQuery = injectValueRandomly(
					fake, req.UnformattedStrInQuery, lo.ToPtr(fake.RandomStringWithLength(21)),
				)
				badIndices[4], req.CustomFormatStrInQuery = injectValueRandomly(
					fake, req.CustomFormatStrInQuery, lo.ToPtr(fake.RandomStringWithLength(31)),
				)
				badIndices[5], req.ByteStrInQuery = injectValueRandomly(
					fake, req.ByteStrInQuery, lo.ToPtr(fake.RandomStringWithLength(41)),
				)

				badIndices[6], req.Payload.UnformattedStr = injectValueRandomly(
					fake, req.Payload.UnformattedStr, lo.ToPtr(fake.RandomStringWithLength(21)),
				)
				badIndices[7], req.Payload.CustomFormatStr = injectValueRandomly(
					fake, req.Payload.CustomFormatStr, lo.ToPtr("!"+fake.RandomStringWithLength(21)),
				)
				badIndices[8], req.Payload.ByteStr = injectValueRandomly(
					fake, req.Payload.ByteStr, lo.ToPtr(fake.RandomStringWithLength(41)),
				)
			})
			query := buildQuery(originalReq)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						// path
						{Field: strconv.Itoa(badIndices[0]), Location: "path.unformattedStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[1]), Location: "path.customFormatStr", Code: "INVALID"},
						{Field: strconv.Itoa(badIndices[2]), Location: "path.byteStr", Code: "INVALID_OUT_OF_RANGE"},

						// query
						{Field: strconv.Itoa(badIndices[3]), Location: "query.unformattedStrInQuery", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[4]), Location: "query.customFormatStrInQuery", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[5]), Location: "query.byteStrInQuery", Code: "INVALID_OUT_OF_RANGE"},

						// body
						{Field: strconv.Itoa(badIndices[6]), Location: "body.unformattedStr", Code: "INVALID_OUT_OF_RANGE"},
						{Field: strconv.Itoa(badIndices[7]), Location: "body.customFormatStr", Code: "INVALID"},
						{Field: strconv.Itoa(badIndices[8]), Location: "body.byteStr", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
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

				// body
				Payload: &models.StringTypesRequiredValidationRequest{
					UnformattedStr:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
					CustomFormatStr: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
					DateStr:         fake.Time().Time(time.Now()),
					DateTimeStr:     fake.Time().Time(time.Now()),
					ByteStr:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),

					OptionalUnformattedStr:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
					OptionalCustomFormatStr: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
					OptionalDateStr:         fake.Time().Time(time.Now()),
					OptionalDateTimeStr:     fake.Time().Time(time.Now()),
					OptionalByteStr:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),

					OptionalValidatedUnformattedStr:  fake.BinaryString().BinaryString(fake.IntBetween(10, 20)),
					OptionalValidatedCustomFormatStr: fake.BinaryString().BinaryString(fake.IntBetween(20, 30)),
					OptionalValidatedByteStr:         fake.BinaryString().BinaryString(fake.IntBetween(30, 40)),
				},
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
				method: http.MethodPost,
				path:   "/string-types/required-validation",
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					wantReq.OptionalDateStrInQuery = lo.Must(
						time.Parse(time.DateOnly, originalReq.OptionalDateStrInQuery.Format(time.DateOnly)),
					)
					assert.Equal(t, &wantReq, testActions.stringTypesRequiredValidation.calls[0].params)
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

				req.Payload.OptionalUnformattedStr = ""
				req.Payload.OptionalCustomFormatStr = ""
				req.Payload.OptionalDateStr = time.Time{}
				req.Payload.OptionalDateTimeStr = time.Time{}
				req.Payload.OptionalByteStr = ""
			})

			query := buildQuery(originalReq)
			query.Del("optionalUnformattedStrInQuery")
			query.Del("optionalCustomFormatStrInQuery")
			query.Del("optionalDateStrInQuery")
			query.Del("optionalDateTimeStrInQuery")
			query.Del("optionalByteStrInQuery")

			return testCase{
				method: http.MethodPost,
				path:   "/string-types/required-validation",
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.stringTypesRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate empty params", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesRequiredValidationRequest) {
				req.Payload.OptionalUnformattedStr = ""
				req.Payload.OptionalCustomFormatStr = ""
				req.Payload.OptionalByteStr = ""
				req.Payload.OptionalValidatedUnformattedStr = ""
				req.Payload.OptionalValidatedCustomFormatStr = ""
				req.Payload.OptionalValidatedByteStr = ""
			})

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
				method: http.MethodPost,
				path:   "/string-types/required-validation",
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						// query
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

						// body
						{Field: "optionalValidatedUnformattedStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalValidatedCustomFormatStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalValidatedByteStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
					},
				),
			}
		})

		runRouteTestCase(t, "should ensure required params", setupRouter, func() testCase {
			originalReq := randomReq(func(req *handlers.StringTypesStringTypesRequiredValidationRequest) {
				req.Payload.UnformattedStr = ""
				req.Payload.CustomFormatStr = ""
				req.Payload.DateStr = time.Time{}
				req.Payload.DateTimeStr = time.Time{}
				req.Payload.ByteStr = ""
			})

			return testCase{
				method: http.MethodPost,
				path:   "/string-types/required-validation",
				query:  url.Values{},
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						// query
						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "customFormatStrInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "dateStrInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "dateTimeStrInQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "byteStrInQuery", Location: "query", Code: "INVALID_REQUIRED"},

						// body
						{Field: "unformattedStr", Location: "body", Code: "INVALID_REQUIRED"},
						{Field: "customFormatStr", Location: "body", Code: "INVALID_REQUIRED"},
						{Field: "dateStr", Location: "body", Code: "INVALID_REQUIRED"},
						{Field: "dateTimeStr", Location: "body", Code: "INVALID_REQUIRED"},
						{Field: "byteStr", Location: "body", Code: "INVALID_REQUIRED"},
					},
				),
			}
		})
	})

	t.Run("nullable-required-validation", func(t *testing.T) {
		runRouteTestCase(t, "should allow null values", setupRouter, func() testCase {
			originalReq := &handlers.StringTypesStringTypesNullableRequiredValidationRequest{
				Payload: &models.StringTypesNullableRequiredValidationRequest{},
			}
			query := url.Values{}
			query.Add("unformattedStrInQuery", "null")
			query.Add("optionalUnformattedStrInQuery", "null")
			return testCase{
				method: http.MethodPost,
				path:   "/string-types/nullable-required-validation",
				query:  query,
				body: bytes.NewBufferString(`{
					"unformattedStr": null,
					"optionalUnformattedStr": null
				}`),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}
					assert.Equal(t, originalReq, testActions.stringTypesNullableRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should require nullable values", setupRouter, func() testCase {
			query := url.Values{}
			return testCase{
				method: http.MethodPost,
				path:   "/string-types/nullable-required-validation",
				query:  query,
				body:   bytes.NewBufferString(`{}`),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID_REQUIRED"},
					},
				),
			}
		})

		runRouteTestCase(t, "should validate provided values", setupRouter, func() testCase {
			query := url.Values{}
			query.Add("unformattedStrInQuery", fake.RandomStringWithLength(9))
			query.Add("optionalUnformattedStrInQuery", fake.RandomStringWithLength(9))
			return testCase{
				method: http.MethodPost,
				path:   "/string-types/nullable-required-validation",
				query:  query,
				body: marshalJSONDataAsReader(t, models.StringTypesNullableRequiredValidationRequest{
					UnformattedStr:         lo.ToPtr(fake.RandomStringWithLength(9)),
					OptionalUnformattedStr: lo.ToPtr(fake.RandomStringWithLength(9)),
				}),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalUnformattedStrInQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "unformattedStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
						{Field: "optionalUnformattedStr", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
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

				// body
				Payload: &models.StringTypesPatternValidationRequest{
					UnformattedStr:  strings.Repeat(strconv.Itoa(fake.RandomDigit()), 10),
					CustomFormatStr: strings.Repeat(strconv.Itoa(fake.RandomDigit()), 20),
					DateStr:         fake.Time().Time(time.Now()),
					DateTimeStr:     fake.Time().Time(time.Now()),
				},
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
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/pattern-validation/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano)),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code) {
						t.Fatalf("unexpected response: %v", recorder.Body)
					}

					wantReq := *originalReq
					wantReq.DateStr = lo.Must(time.Parse(time.DateOnly, originalReq.DateStr.Format(time.DateOnly)))
					wantReq.DateStrInQuery = lo.Must(time.Parse(time.DateOnly, originalReq.DateStrInQuery.Format(time.DateOnly)))
					assert.Equal(t, &wantReq, testActions.stringTypesPatternValidation.calls[0].params)
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

				// body
				req.Payload.UnformattedStr = strings.Repeat(strconv.Itoa(fake.RandomDigit()), 11)
				req.Payload.CustomFormatStr = strings.Repeat(strconv.Itoa(fake.RandomDigit()), 21)
			})
			query := buildQuery(originalReq)
			return testCase{
				method: http.MethodPost,
				path: fmt.Sprintf(
					"/string-types/pattern-validation/%v/%v/%v/%v",
					originalReq.UnformattedStr, originalReq.CustomFormatStr, originalReq.DateStr.Format(time.DateOnly),
					originalReq.DateTimeStr.Format(time.RFC3339Nano)),
				query: query,
				body:  marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*stringTypesControllerTestActions](
					[]expectedBindingError{
						{Field: "unformattedStr", Location: "path", Code: "INVALID"},
						{Field: "customFormatStr", Location: "path", Code: "INVALID"},

						{Field: "unformattedStrInQuery", Location: "query", Code: "INVALID"},
						{Field: "customFormatStrInQuery", Location: "query", Code: "INVALID"},

						{Field: "unformattedStr", Location: "body", Code: "INVALID"},
						{Field: "customFormatStr", Location: "body", Code: "INVALID"},
					},
				),
			}
		})
	})

	t.Run("enums", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.StringTypesStringTypesEnumsRequest),
		) *handlers.StringTypesStringTypesEnumsRequest {
			req := &handlers.StringTypesStringTypesEnumsRequest{
				// path
				InlineEnumParam: lo.Shuffle(slices.Clone(handlers.AllowableStringTypesStringTypesEnumsInlineEnumParamValues))[0],
				NullableInlineEnumParam: lo.ToPtr(
					lo.Shuffle(slices.Clone(handlers.AllowableStringTypesStringTypesEnumsNullableInlineEnumParamValues))[0],
				),
				RefEnumParam:         lo.Shuffle(slices.Clone(models.AllowableBasicStringEnumValues))[0],
				NullableRefEnumParam: lo.ToPtr(lo.Shuffle(slices.Clone(models.AllowableNullableStringEnumValues))[0]),

				// query
				InlineEnumParamInQuery: lo.Shuffle(
					slices.Clone(handlers.AllowableStringTypesStringTypesEnumsInlineEnumParamInQueryValues),
				)[0],
				NullableInlineEnumParamInQuery: lo.ToPtr(
					lo.Shuffle(slices.Clone(handlers.AllowableStringTypesStringTypesEnumsNullableInlineEnumParamInQueryValues))[0],
				),
				RefEnumParamInQuery: lo.Shuffle(
					slices.Clone(models.AllowableBasicStringEnumValues),
				)[0],
				NullableRefEnumParamInQuery: lo.ToPtr(lo.Shuffle(slices.Clone(models.AllowableNullableStringEnumValues))[0]),
				OptionalInlineEnumParamInQuery: lo.Shuffle(
					slices.Clone(handlers.AllowableStringTypesStringTypesEnumsOptionalInlineEnumParamInQueryValues),
				)[0],
				OptionalNullableInlineEnumParamInQuery: lo.ToPtr(
					lo.Shuffle(
						slices.Clone(handlers.AllowableStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQueryValues),
					)[0],
				),
				OptionalRefEnumParamInQuery: lo.Shuffle(slices.Clone(models.AllowableBasicStringEnumValues))[0],
				OptionalNullableRefEnumParamInQuery: lo.ToPtr(
					lo.Shuffle(slices.Clone(models.AllowableNullableStringEnumValues))[0]),

				// body
				Payload: &models.StringTypesEnumsRequest{
					InlineEnumProp: lo.Shuffle(slices.Clone(models.AllowableStringTypesEnumsRequestInlineEnumPropValues))[0],
					NullableInlineEnumProp: lo.ToPtr(
						lo.Shuffle(slices.Clone(models.AllowableStringTypesEnumsRequestNullableInlineEnumPropValues))[0]),
					RefEnumProp:         lo.Shuffle(slices.Clone(models.AllowableBasicStringEnumValues))[0],
					NullableRefEnumProp: lo.ToPtr(lo.Shuffle(slices.Clone(models.AllowableNullableStringEnumValues))[0]),
					OptionalInlineEnumProp: lo.Shuffle(
						slices.Clone(models.AllowableStringTypesEnumsRequestOptionalInlineEnumPropValues),
					)[0],
					OptionalNullableInlineEnumProp: lo.ToPtr(
						lo.Shuffle(slices.Clone(models.AllowableStringTypesEnumsRequestOptionalNullableInlineEnumPropValues))[0]),
					OptionalRefEnumProp:         lo.Shuffle(slices.Clone(models.AllowableBasicStringEnumValues))[0],
					OptionalNullableRefEnumProp: lo.ToPtr(lo.Shuffle(slices.Clone(models.AllowableNullableStringEnumValues))[0]),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildPath := func(originalReq *handlers.StringTypesStringTypesEnumsRequest) string {
			return fmt.Sprintf(
				"/string-types/enums/%v/%v/%v/%v",
				originalReq.InlineEnumParam,
				originalReq.NullableInlineEnumParam,
				originalReq.RefEnumParam,
				originalReq.NullableRefEnumParam,
			)
		}

		buildQuery := func(req *handlers.StringTypesStringTypesEnumsRequest) url.Values {
			query := url.Values{}
			query.Add("inlineEnumParamInQuery", req.InlineEnumParamInQuery.String())
			query.Add("nullableInlineEnumParamInQuery", req.NullableInlineEnumParamInQuery.String())
			query.Add("refEnumParamInQuery", req.RefEnumParamInQuery.String())
			query.Add("nullableRefEnumParamInQuery", req.NullableRefEnumParamInQuery.String())
			query.Add("optionalInlineEnumParamInQuery", req.OptionalInlineEnumParamInQuery.String())
			query.Add("optionalNullableInlineEnumParamInQuery", req.OptionalNullableInlineEnumParamInQuery.String())
			query.Add("optionalRefEnumParamInQuery", req.OptionalRefEnumParamInQuery.String())
			query.Add("optionalNullableRefEnumParamInQuery", req.OptionalNullableRefEnumParamInQuery.String())
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			req := randomReq()
			query := buildQuery(req)

			return testCase{
				method: http.MethodPost,
				path:   buildPath(req),
				query:  query,
				body:   marshalJSONDataAsReader(t, req.Payload),
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, req, testActions.stringTypesEnums.calls[0].params)
				},
			}
		})
	})
}
