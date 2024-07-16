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
}
