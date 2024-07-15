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
				DateStr:         fake.Time().Faker.Time().ISO8601(time.Now()),
				DateTimeStr:     fake.Time().Faker.Time().ISO8601(time.Now()),
				ByteStr:         fake.BinaryString().BinaryString(10),

				// query
				UnformattedStrInQuery:  fake.Lorem().Word(),
				CustomFormatStrInQuery: fake.Lorem().Word(),
				DateStrInQuery:         fake.Time().Faker.Time().ISO8601(time.Now()),
				DateTimeStrInQuery:     fake.Time().Faker.Time().ISO8601(time.Now()),
				ByteStrInQuery:         fake.BinaryString().BinaryString(10),
			}
		}
		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			wantReq := randomReq()
			query := url.Values{}
			query.Add("unformattedStrInQuery", wantReq.UnformattedStrInQuery)
			query.Add("customFormatStrInQuery", wantReq.CustomFormatStrInQuery)
			query.Add("dateStrInQuery", wantReq.DateStrInQuery)
			query.Add("dateTimeStrInQuery", wantReq.DateTimeStrInQuery)
			query.Add("byteStrInQuery", wantReq.ByteStrInQuery)

			return testCase{
				path:  fmt.Sprintf("/string-types/parsing/%v/%v/%v/%v/%v", wantReq.UnformattedStr, wantReq.CustomFormatStr, wantReq.DateStr, wantReq.DateTimeStr, wantReq.ByteStr),
				query: query,
				expect: func(t *testing.T, testActions *stringTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					assert.Equal(t, 204, recorder.Code)
					assert.Equal(t, wantReq, testActions.stringTypesParsing.calls[0].params)
				},
			}
		})
	})
}
