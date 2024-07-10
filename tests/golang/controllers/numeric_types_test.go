package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/numeric_types/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

type routerAdapter struct {
	mux           *http.ServeMux
	handledErrors []error
}

func (r *routerAdapter) PathValue(req *http.Request, paramName string) string {
	return req.PathValue(paramName)
}

func (r *routerAdapter) HandleRoute(method, pathPattern string, h http.Handler) {
	r.mux.Handle(method+" "+pathPattern, h)
}

func (r *routerAdapter) HandleError(req *http.Request, w http.ResponseWriter, err error) {
	r.handledErrors = append(r.handledErrors, err)
	w.WriteHeader(http.StatusInternalServerError)
}

func TestNumericTypesController(t *testing.T) {
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
		handlers.MountNumericTypesRoutes(controller, handlers.NewHttpApp(router))
		return router
	}

	t.Run("number", func(t *testing.T) {
		t.Run("params parsing and binding", func(t *testing.T) {
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
	})
}
