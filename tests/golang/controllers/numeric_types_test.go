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

	randomReq := func() *handlers.NumericTypesNumberAnySimpleRequest {
		return &handlers.NumericTypesNumberAnySimpleRequest{
			PathParam1:     fake.Float32(2, 10, 1000),
			PathParam2:     fake.Float32(2, 10, 1000),
			RequiredQuery1: fake.Float32(2, 10, 1000),
			RequiredQuery2: fake.Float32(2, 10, 1000),
			OptionalQuery1: fake.Float32(2, 10, 1000),
			OptionalQuery2: fake.Float32(2, 10, 1000),
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
				fmt.Sprintf("/numeric-types/number/any/%v/%v", wantReq.PathParam1, wantReq.PathParam2),
				http.NoBody,
			)
			query := url.Values{}
			query.Add("requiredQuery1", fmt.Sprint(wantReq.RequiredQuery1))
			query.Add("requiredQuery2", fmt.Sprint(wantReq.RequiredQuery2))
			query.Add("optionalQuery1", fmt.Sprint(wantReq.OptionalQuery1))
			query.Add("optionalQuery2", fmt.Sprint(wantReq.OptionalQuery2))
			testReq.URL.RawQuery = query.Encode()
			recorder := httptest.NewRecorder()
			router.mux.ServeHTTP(recorder, testReq)

			assert.Equal(t, 204, recorder.Code)
			assert.Len(t, testActions.getNumberAnySimpleCalls, 1)
			assert.Equal(t, wantReq, testActions.getNumberAnySimpleCalls[0].params)
		})
		t.Run("should fail if required values are missing", func(t *testing.T) {
			testActions := &numericTypesControllerTestActions{}
			router := setupRouter(testActions)
			wantReq := randomReq()
			testReq := httptest.NewRequest(
				"GET",
				fmt.Sprintf("/numeric-types/number/any/%v/%v", wantReq.PathParam1, wantReq.PathParam2),
				http.NoBody,
			)
			recorder := httptest.NewRecorder()
			router.mux.ServeHTTP(recorder, testReq)

			assert.Equal(t, 400, recorder.Code)
			assert.Len(t, testActions.getNumberAnySimpleCalls, 0)

			gotErrors := unmarshalBindingErrors(t, recorder.Body)

			assertFieldError(t, gotErrors, "query", "requiredQuery1", handlers.ErrValueRequired)
			assertFieldError(t, gotErrors, "query", "requiredQuery2", handlers.ErrValueRequired)
		})
	})
}
