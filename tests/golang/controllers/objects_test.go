package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestObjects(t *testing.T) {
	fake := faker.New()

	setupRouter := func() (*objectsControllerTestActions, http.Handler) {
		testActions := &objectsControllerTestActions{}
		controller := newObjectsController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterObjectsRoutes(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	randomSimpleNullableObject := func() *models.SimpleNullableObject {
		return &models.SimpleNullableObject{
			SimpleField1: fake.Lorem().Word(),
		}
	}

	type testCase = routeTestCase[*objectsControllerTestActions]

	t.Run("nullable-body", func(t *testing.T) {
		t.Run("required", func(t *testing.T) {
			runRouteTestCase(t, "should parse nullable body", setupRouter, func() testCase {
				originalReq := handlers.ObjectsObjectsNullableRequiredBodyRequest{
					Payload: randomSimpleNullableObject(),
				}
				return testCase{
					method: http.MethodPost,
					path:   "/objects/nullable-body",
					body:   marshalJSONDataAsReader(t, originalReq.Payload),
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, &originalReq, testActions.objectsNullableRequiredBody.calls[0].params)
					},
				}
			})
			runRouteTestCase(t, "should parse nullable body with null value", setupRouter, func() testCase {
				return testCase{
					method: http.MethodPost,
					path:   "/objects/nullable-body",
					body:   bytes.NewReader([]byte("null")),
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Nil(t, testActions.objectsNullableRequiredBody.calls[0].params.Payload)
					},
				}
			})
			runRouteTestCase(t, "should fail if no body provided", setupRouter, func() testCase {
				return testCase{
					method: http.MethodPost,
					path:   "/objects/nullable-body",
					expect: expectBindingErrors[*objectsControllerTestActions](
						[]fieldBindingError{
							// body
							{Field: "payload", Location: "body", Code: "INVALID_REQUIRED"},
						},
					),
				}
			})
		})
		t.Run("optional", func(t *testing.T) {
			runRouteTestCase(t, "should parse nullable body", setupRouter, func() testCase {
				originalReq := handlers.ObjectsObjectsNullableOptionalBodyRequest{
					Payload: randomSimpleNullableObject(),
				}
				return testCase{
					method: http.MethodPut,
					path:   "/objects/nullable-body",
					body:   marshalJSONDataAsReader(t, originalReq.Payload),
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, &originalReq, testActions.objectsNullableOptionalBody.calls[0].params)
					},
				}
			})
			runRouteTestCase(t, "should parse nullable body with null value", setupRouter, func() testCase {
				return testCase{
					method: http.MethodPut,
					path:   "/objects/nullable-body",
					body:   bytes.NewReader([]byte("null")),
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Nil(t, testActions.objectsNullableOptionalBody.calls[0].params.Payload)
					},
				}
			})
			runRouteTestCase(t, "should allow empty body", setupRouter, func() testCase {
				return testCase{
					method: http.MethodPut,
					path:   "/objects/nullable-body",
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Nil(t, testActions.objectsNullableOptionalBody.calls[0].params.Payload)
					},
				}
			})
		})
	})
}
