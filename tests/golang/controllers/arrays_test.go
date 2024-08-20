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

	setupRouter := func() (*arraysControllerTestActions, http.Handler) {
		testActions := &arraysControllerTestActions{}
		controller := newArraysController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterArraysRoutes(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	randomStrings := func(n, minLength, maxLength int) []string {
		strs := make([]string, n)
		for i := range n {
			length := fake.IntBetween(minLength, maxLength)
			strs[i] = fake.RandomStringWithLength(length)
		}
		return strs
	}

	type testCase = routeTestCase[*arraysControllerTestActions]

	t.Run("required-validation", func(t *testing.T) {
		randomReq := func(
			opts ...func(*handlers.ArraysArraysRequiredValidationRequest),
		) *handlers.ArraysArraysRequiredValidationRequest {
			req := &handlers.ArraysArraysRequiredValidationRequest{
				// path
				SimpleItems1: randomStrings(3, 1, 10),
				SimpleItems2: randomStrings(3, 1, 10),

				// query
				SimpleItems1InQuery:         randomStrings(3, 1, 10),
				SimpleItems2InQuery:         randomStrings(3, 1, 10),
				OptionalSimpleItems1InQuery: randomStrings(3, 1, 10),
				OptionalSimpleItems2InQuery: randomStrings(3, 1, 10),

				// body
				Payload: &models.ArraysRequiredValidationRequest{
					SimpleItems1:         randomStrings(3, 1, 10),
					SimpleItems2:         randomStrings(3, 1, 10),
					OptionalSimpleItems1: randomStrings(3, 1, 10),
					OptionalSimpleItems2: randomStrings(3, 1, 10),
				},
			}
			for _, opt := range opts {
				opt(req)
			}
			return req
		}

		buildQuery := func(req *handlers.ArraysArraysRequiredValidationRequest) url.Values {
			query := url.Values{}
			query["simpleItems1InQuery"] = req.SimpleItems1InQuery
			query["simpleItems2InQuery"] = req.SimpleItems2InQuery
			query["optionalSimpleItems1InQuery"] = req.OptionalSimpleItems1InQuery
			query["optionalSimpleItems2InQuery"] = req.OptionalSimpleItems2InQuery
			return query
		}

		buildPath := func(req *handlers.ArraysArraysRequiredValidationRequest) string {
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
				func(req *handlers.ArraysArraysRequiredValidationRequest) {
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
				func(req *handlers.ArraysArraysRequiredValidationRequest) {
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
}
