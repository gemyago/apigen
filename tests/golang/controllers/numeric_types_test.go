package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
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
	testActions := &numericTypesControllerTestActions{}
	controller := newNumericTypesController(testActions)
	router := &routerAdapter{
		mux: http.NewServeMux(),
	}
	handlers.MountNumericTypesRoutes(controller, router)

	testReq := httptest.NewRequest("GET", "/numeric-types/number/any/{pathParam1}/{pathParam2}", http.NoBody)
	recorder := httptest.NewRecorder()
	router.mux.ServeHTTP(recorder, testReq)

	assert.Equal(t, 204, recorder.Code)
	assert.Len(t, testActions, 1)
}
