package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestArrays(t *testing.T) {
	fake := faker.New()

	type testCase = routeTestCase[*arraysControllerTestActions]
	setupRouter := func(_ testCase) (*arraysControllerTestActions, http.Handler) {
		testActions := &arraysControllerTestActions{}
		controller := &arraysController{testActions}
		rootHandler := handlers.
			NewRootHandler(&routerAdapter{mux: http.NewServeMux()}, handlers.WithLogger(newLogger())).
			RegisterArraysRoutes(controller)
		return testActions, rootHandler
	}

	randomStrings := func(minItems, maxItems int) []string {
		n := fake.IntBetween(minItems, maxItems)
		strs := make([]string, n)
		for i := range n {
			strs[i] = fake.Lorem().Word()
		}
		return strs
	}

	t.Run("required-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*models.ArraysRequiredValidationParams),
		) *models.ArraysRequiredValidationParams {
			req := &models.ArraysRequiredValidationParams{
				// path
				SimpleItems1: randomStrings(3, 5),
				SimpleItems2: randomStrings(3, 5),

				// query
				SimpleItems1InQuery:         randomStrings(3, 5),
				SimpleItems2InQuery:         randomStrings(3, 5),
				OptionalSimpleItems1InQuery: randomStrings(3, 5),
				OptionalSimpleItems2InQuery: randomStrings(3, 5),

				// body
				Payload: &models.ArraysRequiredValidationRequest{
					SimpleItems1:         randomStrings(3, 5),
					SimpleItems2:         randomStrings(3, 5),
					OptionalSimpleItems1: randomStrings(3, 5),
					OptionalSimpleItems2: randomStrings(3, 5),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *models.ArraysRequiredValidationParams) url.Values {
			query := url.Values{}
			query["simpleItems1InQuery"] = req.SimpleItems1InQuery
			query["simpleItems2InQuery"] = req.SimpleItems2InQuery
			query["optionalSimpleItems1InQuery"] = req.OptionalSimpleItems1InQuery
			query["optionalSimpleItems2InQuery"] = req.OptionalSimpleItems2InQuery
			return query
		}

		buildPath := func(req *models.ArraysRequiredValidationParams) string {
			return fmt.Sprintf(
				"/arrays/required-validation/%v/%v",
				strings.Join(req.SimpleItems1, ","),
				strings.Join(req.SimpleItems2, ","),
			)
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  buildQuery(originalReq),
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *arraysControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Equal(t, originalReq, testActions.arraysRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should allow optional params", setupRouter, func() testCase {
			originalReq := randomReq(
				func(req *models.ArraysRequiredValidationParams) {
					req.OptionalSimpleItems1InQuery = nil
					req.OptionalSimpleItems2InQuery = nil
					req.Payload.OptionalSimpleItems1 = nil
					req.Payload.OptionalSimpleItems2 = nil
				},
			)
			query := buildQuery(originalReq)
			query.Del("optionalSimpleItems1InQuery")
			query.Del("optionalSimpleItems2InQuery")
			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *arraysControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Equal(t, originalReq, testActions.arraysRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate required params", setupRouter, func() testCase {
			originalReq := randomReq(
				func(req *models.ArraysRequiredValidationParams) {
					req.Payload.SimpleItems1 = nil
					req.Payload.SimpleItems2 = nil
				},
			)
			query := buildQuery(originalReq)
			query.Del("simpleItems1InQuery")
			query.Del("simpleItems2InQuery")
			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*arraysControllerTestActions]([]expectedBindingError{
					{Field: "simpleItems1InQuery", Location: "query", Code: "INVALID_REQUIRED"},
					{Field: "simpleItems2InQuery", Location: "query", Code: "INVALID_REQUIRED"},
					{Field: "simpleItems1", Location: "body", Code: "INVALID_REQUIRED"},
					{Field: "simpleItems2", Location: "body", Code: "INVALID_REQUIRED"},
				}),
			}
		})
	})

	t.Run("nullable-required-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*models.ArraysNullableRequiredValidationParams),
		) *models.ArraysNullableRequiredValidationParams {
			req := &models.ArraysNullableRequiredValidationParams{
				// path
				SimpleItems1: randomStrings(2, 5),
				SimpleItems2: randomStrings(5, 10),

				// query
				SimpleItems1InQuery:         randomStrings(2, 5),
				SimpleItems2InQuery:         randomStrings(5, 10),
				OptionalSimpleItems1InQuery: randomStrings(15, 20),
				OptionalSimpleItems2InQuery: randomStrings(20, 25),

				// body
				Payload: &models.ArraysNullableRequiredValidationRequest{
					SimpleItems1:         randomStrings(2, 5),
					SimpleItems2:         randomStrings(5, 10),
					OptionalSimpleItems1: randomStrings(15, 20),
					OptionalSimpleItems2: randomStrings(20, 25),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *models.ArraysNullableRequiredValidationParams) url.Values {
			query := url.Values{}
			query["simpleItems1InQuery"] = req.SimpleItems1InQuery
			query["simpleItems2InQuery"] = req.SimpleItems2InQuery
			query["optionalSimpleItems1InQuery"] = req.OptionalSimpleItems1InQuery
			query["optionalSimpleItems2InQuery"] = req.OptionalSimpleItems2InQuery
			return query
		}

		buildPath := func(req *models.ArraysNullableRequiredValidationParams) string {
			return fmt.Sprintf(
				"/arrays/nullable-required-validation/%v/%v",
				strings.Join(req.SimpleItems1, ","),
				strings.Join(req.SimpleItems2, ","),
			)
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  buildQuery(originalReq),
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *arraysControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Equal(t, originalReq, testActions.arraysNullableRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should allow optional params", setupRouter, func() testCase {
			originalReq := randomReq(
				func(req *models.ArraysNullableRequiredValidationParams) {
					req.OptionalSimpleItems1InQuery = nil
					req.OptionalSimpleItems2InQuery = nil
					req.Payload.OptionalSimpleItems1 = nil
					req.Payload.OptionalSimpleItems2 = nil
				},
			)
			query := buildQuery(originalReq)
			query.Del("optionalSimpleItems1InQuery")
			query.Del("optionalSimpleItems2InQuery")
			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *arraysControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Equal(t, originalReq, testActions.arraysNullableRequiredValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate required params", setupRouter, func() testCase {
			originalReq := randomReq(
				func(req *models.ArraysNullableRequiredValidationParams) {
					req.Payload.SimpleItems1 = nil
					req.Payload.SimpleItems2 = nil
				},
			)
			query := buildQuery(originalReq)
			query.Del("simpleItems1InQuery")
			query.Del("simpleItems2InQuery")
			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  query,
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*arraysControllerTestActions]([]expectedBindingError{
					{Field: "simpleItems1InQuery", Location: "query", Code: "INVALID_REQUIRED"},
					{Field: "simpleItems2InQuery", Location: "query", Code: "INVALID_REQUIRED"},
					{Field: "simpleItems1", Location: "body", Code: "INVALID_REQUIRED"},
					{Field: "simpleItems2", Location: "body", Code: "INVALID_REQUIRED"},
				}),
			}
		})
	})

	t.Run("range-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*models.ArraysRangeValidationParams),
		) *models.ArraysRangeValidationParams {
			req := &models.ArraysRangeValidationParams{
				// path
				SimpleItems1: randomStrings(5, 10),
				SimpleItems2: randomStrings(10, 15),

				// query
				SimpleItems1InQuery:         randomStrings(5, 10),
				SimpleItems2InQuery:         randomStrings(10, 15),
				OptionalSimpleItems1InQuery: randomStrings(15, 20),
				OptionalSimpleItems2InQuery: randomStrings(20, 25),

				// body
				Payload: &models.ArraysRangeValidationRequest{
					SimpleItems1:         randomStrings(5, 10),
					SimpleItems2:         randomStrings(10, 15),
					OptionalSimpleItems1: randomStrings(15, 20),
					OptionalSimpleItems2: randomStrings(20, 25),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *models.ArraysRangeValidationParams) url.Values {
			query := url.Values{}
			query["simpleItems1InQuery"] = req.SimpleItems1InQuery
			query["simpleItems2InQuery"] = req.SimpleItems2InQuery
			query["optionalSimpleItems1InQuery"] = req.OptionalSimpleItems1InQuery
			query["optionalSimpleItems2InQuery"] = req.OptionalSimpleItems2InQuery
			return query
		}

		buildPath := func(req *models.ArraysRangeValidationParams) string {
			return fmt.Sprintf(
				"/arrays/range-validation/%v/%v",
				strings.Join(req.SimpleItems1, ","),
				strings.Join(req.SimpleItems2, ","),
			)
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter, func() testCase {
			originalReq := randomReq()

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  buildQuery(originalReq),
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: func(t *testing.T, testActions *arraysControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
						return
					}

					assert.Equal(t, originalReq, testActions.arraysRangeValidation.calls[0].params)
				},
			}
		})

		runRouteTestCase(t, "should validate minItems", setupRouter, func() testCase {
			originalReq := randomReq(func(req *models.ArraysRangeValidationParams) {
				req.SimpleItems1 = randomStrings(1, 4)
				req.SimpleItems2 = randomStrings(1, 9)
				req.SimpleItems1InQuery = randomStrings(1, 4)
				req.SimpleItems2InQuery = randomStrings(1, 9)
				req.OptionalSimpleItems1InQuery = randomStrings(1, 14)
				req.OptionalSimpleItems2InQuery = randomStrings(1, 19)

				req.Payload.SimpleItems1 = randomStrings(1, 4)
				req.Payload.SimpleItems2 = randomStrings(1, 9)
				req.Payload.OptionalSimpleItems1 = randomStrings(1, 14)
				req.Payload.OptionalSimpleItems2 = randomStrings(1, 19)
			})

			return testCase{
				method: http.MethodPost,
				path:   buildPath(originalReq),
				query:  buildQuery(originalReq),
				body:   marshalJSONDataAsReader(t, originalReq.Payload),
				expect: expectBindingErrors[*arraysControllerTestActions]([]expectedBindingError{
					{Field: "simpleItems1", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "simpleItems2", Location: "path", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "simpleItems1InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "simpleItems2InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "optionalSimpleItems1InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "optionalSimpleItems2InQuery", Location: "query", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "simpleItems1", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "simpleItems2", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "optionalSimpleItems1", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
					{Field: "optionalSimpleItems2", Location: "body", Code: "INVALID_OUT_OF_RANGE"},
				}),
			}
		})
	})
}
