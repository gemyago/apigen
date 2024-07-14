package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestNumericTypes(t *testing.T) {
	fake := faker.New()

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
		}
	}

	setupRouter := func(testActions *numericTypesControllerTestActions) *routerAdapter {
		controller := newNumericTypesController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.MountNumericTypesRoutes(controller, handlers.NewHttpApp(router, handlers.WithLogger(newLogger())))
		return router
	}

	t.Run("number", func(t *testing.T) {
		t.Run("should parse and bind valid values", func(t *testing.T) {
			testActions := &numericTypesControllerTestActions{}
			router := setupRouter(testActions)
			wantReq := randomReq()
			testReq := httptest.NewRequest(
				"GET",
				fmt.Sprintf("/numeric-types/parsing/%v/%v/%v/%v/%v/%v", wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32, wantReq.NumberInt64),
				http.NoBody,
			)
			query := url.Values{}
			query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
			query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
			query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
			query.Add("numberIntInQuery", fmt.Sprint(wantReq.NumberIntInQuery))
			query.Add("numberInt32InQuery", fmt.Sprint(wantReq.NumberInt32InQuery))
			query.Add("numberInt64InQuery", fmt.Sprint(wantReq.NumberInt64InQuery))
			testReq.URL.RawQuery = query.Encode()
			recorder := httptest.NewRecorder()
			router.mux.ServeHTTP(recorder, testReq)

			assert.Equal(t, 204, recorder.Code)
			assert.Len(t, testActions.numericTypesParsing, 1)
			assert.Equal(t, wantReq, testActions.numericTypesParsing[0].params)
		})
	})
}
