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

	setupRouter := func(testActions *numericTypesControllerTestActions) *routerAdapter {
		controller := newNumericTypesController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.MountNumericTypesRoutes(controller, handlers.NewHttpApp(router, handlers.WithLogger(newLogger())))
		return router
	}

	type expectFn func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder)

	type testCase struct {
		path   string
		query  url.Values
		expect expectFn
	}

	// expectErrors := func(wantErrors []handlers.BindingError) expectFn {
	// 	return func(
	// 		t *testing.T,
	// 		testActions *numericTypesControllerTestActions,
	// 		recorder *httptest.ResponseRecorder,
	// 	) {
	// 		assert.Equal(t, 400, recorder.Code)
	// 		assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("content-type"))
	// 		gotErrors := unmarshalBindingErrors(t, recorder.Body)
	// 		if !assert.Len(t, gotErrors.Errors, len(wantErrors)) {
	// 			return
	// 		}
	// 		for _, fe := range wantErrors {
	// 			assertFieldError(t, gotErrors, fe.Location, fe.Field, fe.Code)
	// 		}
	// 	}
	// }

	runTestCase := func(t *testing.T, name string, tc func() testCase) {
		t.Run(name, func(t *testing.T) {
			tc := tc()
			testActions := &numericTypesControllerTestActions{}
			router := setupRouter(testActions)
			testReq := httptest.NewRequest(
				"GET",
				tc.path,
				http.NoBody,
			)
			if len(tc.query) != 0 {
				testReq.URL.RawQuery = tc.query.Encode()
			}
			recorder := httptest.NewRecorder()
			router.mux.ServeHTTP(recorder, testReq)
			tc.expect(t, testActions, recorder)
		})
	}

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
			}
		}

		runTestCase(t, "should parse and bind valid values", func() testCase {
			wantReq := randomReq()
			query := url.Values{}
			query.Add("numberAnyInQuery", fmt.Sprint(wantReq.NumberAnyInQuery))
			query.Add("numberFloatInQuery", fmt.Sprint(wantReq.NumberFloatInQuery))
			query.Add("numberDoubleInQuery", fmt.Sprint(wantReq.NumberDoubleInQuery))
			query.Add("numberIntInQuery", fmt.Sprint(wantReq.NumberIntInQuery))
			query.Add("numberInt32InQuery", fmt.Sprint(wantReq.NumberInt32InQuery))
			query.Add("numberInt64InQuery", fmt.Sprint(wantReq.NumberInt64InQuery))

			return testCase{
				path:  fmt.Sprintf("/numeric-types/parsing/%v/%v/%v/%v/%v/%v", wantReq.NumberAny, wantReq.NumberFloat, wantReq.NumberDouble, wantReq.NumberInt, wantReq.NumberInt32, wantReq.NumberInt64),
				query: query,
				expect: func(t *testing.T, testActions *numericTypesControllerTestActions, recorder *httptest.ResponseRecorder) {
					assert.Equal(t, 204, recorder.Code)
					assert.Equal(t, wantReq, testActions.numericTypesParsing[0].params)
				},
			}
		})
	})
}
