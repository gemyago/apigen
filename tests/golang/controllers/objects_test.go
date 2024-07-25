package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/samber/lo"
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

	randomSimpleObject := func() *models.SimpleObject {
		return &models.SimpleObject{
			SimpleField1: fake.Lorem().Word(),
		}
	}

	randomSimpleNullableObject := func() *models.SimpleNullableObject {
		return &models.SimpleNullableObject{
			SimpleField1: fake.Lorem().Word(),
		}
	}

	randomSimpleObjectsContainer := func(
		opts ...func(*models.SimpleObjectsContainer),
	) *models.SimpleObjectsContainer {
		res := &models.SimpleObjectsContainer{
			SimpleField1:                  fake.Lorem().Word(),
			SimpleField2:                  fake.Lorem().Word(),
			SimpleObject1:                 randomSimpleObject(),
			SimpleObject2:                 randomSimpleObject(),
			SimpleNullableObject1:         randomSimpleNullableObject(),
			SimpleNullableObject2:         randomSimpleNullableObject(),
			OptionalSimpleObject1:         randomSimpleObject(),
			OptionalSimpleObject2:         randomSimpleObject(),
			OptionalNullableSimpleObject1: randomSimpleNullableObject(),
			OptionalNullableSimpleObject2: randomSimpleNullableObject(),
		}
		for _, opt := range opts {
			opt(res)
		}
		return res
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

	t.Run("required-body", func(t *testing.T) {
		t.Run("required", func(t *testing.T) {
			runRouteTestCase(t, "should parse required body", setupRouter, func() testCase {
				originalReq := handlers.ObjectsObjectsRequiredBodyRequest{
					Payload: randomSimpleObject(),
				}
				return testCase{
					method: http.MethodPost,
					path:   "/objects/required-body",
					body:   marshalJSONDataAsReader(t, originalReq.Payload),
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, &originalReq, testActions.objectsRequiredBody.calls[0].params)
					},
				}
			})
			runRouteTestCase(t, "should fail if no body provided", setupRouter, func() testCase {
				return testCase{
					method: http.MethodPost,
					path:   "/objects/required-body",
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
			runRouteTestCase(t, "should parse optional body", setupRouter, func() testCase {
				originalReq := handlers.ObjectsObjectsOptionalBodyRequest{
					Payload: randomSimpleObject(),
				}
				return testCase{
					method: http.MethodPut,
					path:   "/objects/required-body",
					body:   marshalJSONDataAsReader(t, originalReq.Payload),
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, &originalReq, testActions.objectsOptionalBody.calls[0].params)
					},
				}
			})
			runRouteTestCase(t, "should allow empty body", setupRouter, func() testCase {
				return testCase{
					method: http.MethodPut,
					path:   "/objects/required-body",
					expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Nil(t, testActions.objectsOptionalBody.calls[0].params.Payload)
					},
				}
			})
		})
	})

	t.Run("required-nested-objects", func(t *testing.T) {
		runRouteTestCase(t, "should parse body", setupRouter, func() testCase {
			originalReq := handlers.ObjectsObjectsRequiredNestedObjectsRequest{
				Payload: randomSimpleObjectsContainer(),
			}
			return testCase{
				method: http.MethodPost,
				path:   "/objects/required-nested-objects",
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, &originalReq, testActions.objectsRequiredNestedObjects.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should allow missing optional fields", setupRouter, func() testCase {
			originalReq := handlers.ObjectsObjectsRequiredNestedObjectsRequest{
				Payload: randomSimpleObjectsContainer(
					func(req *models.SimpleObjectsContainer) {
						req.OptionalSimpleObject1 = nil
						req.OptionalSimpleObject2 = nil
						req.OptionalNullableSimpleObject1 = nil
						req.OptionalNullableSimpleObject2 = nil
					},
				),
			}
			return testCase{
				method: http.MethodPost,
				path:   "/objects/required-nested-objects",
				body: bytes.NewBufferString(fmt.Sprintf(`{
					"simpleField1": "%s",
					"simpleField2": "%s",
					"simpleObject1": %s,
					"simpleObject2": %s,
					"simpleNullableObject1": %s,
					"simpleNullableObject2": %s
				}`,
					originalReq.Payload.SimpleField1,
					originalReq.Payload.SimpleField2,
					lo.Must1(json.Marshal(originalReq.Payload.SimpleObject1)),
					lo.Must1(json.Marshal(originalReq.Payload.SimpleObject2)),
					lo.Must1(json.Marshal(originalReq.Payload.SimpleNullableObject1)),
					lo.Must1(json.Marshal(originalReq.Payload.SimpleNullableObject2)),
				)),
				expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, &originalReq, testActions.objectsRequiredNestedObjects.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should allow null values for nullables", setupRouter, func() testCase {
			originalReq := handlers.ObjectsObjectsRequiredNestedObjectsRequest{
				Payload: randomSimpleObjectsContainer(
					func(req *models.SimpleObjectsContainer) {
						req.SimpleNullableObject1 = nil
						req.SimpleNullableObject2 = nil
						req.OptionalSimpleObject1 = nil
						req.OptionalSimpleObject2 = nil
						req.OptionalNullableSimpleObject1 = nil
						req.OptionalNullableSimpleObject2 = nil
					},
				),
			}
			return testCase{
				method: http.MethodPost,
				path:   "/objects/required-nested-objects",
				body: bytes.NewBufferString(fmt.Sprintf(`{
					"simpleField1": "%s",
					"simpleField2": "%s",
					"simpleObject1": %s,
					"simpleObject2": %s,
					"simpleNullableObject1": null,
					"simpleNullableObject2": null
				}`,
					originalReq.Payload.SimpleField1,
					originalReq.Payload.SimpleField2,
					lo.Must1(json.Marshal(originalReq.Payload.SimpleObject1)),
					lo.Must1(json.Marshal(originalReq.Payload.SimpleObject2)),
				)),
				expect: func(t *testing.T, testActions *objectsControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, &originalReq, testActions.objectsRequiredNestedObjects.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should ensure required objects", setupRouter, func() testCase {
			return testCase{
				method: http.MethodPost,
				path:   "/objects/required-nested-objects",
				body:   bytes.NewBufferString(`{}`),
				expect: expectBindingErrors[*objectsControllerTestActions](
					[]fieldBindingError{
						// body
						{Field: "simpleField1", Location: "body", Code: "INVALID_REQUIRED"},
						{Field: "simpleField2", Location: "body", Code: "INVALID_REQUIRED"},
						{Field: "simpleObject1", Location: "body", Code: "INVALID_REQUIRED"},
						{Field: "simpleObject2", Location: "body", Code: "INVALID_REQUIRED"},
					},
				),
			}
		})
	})
}
