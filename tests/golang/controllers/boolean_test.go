package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

func TestBoolean(t *testing.T) {
	fake := faker.New()

	setupRouter := func() (*booleanControllerTestActions, http.Handler) {
		testActions := &booleanControllerTestActions{}
		controller := newBooleanController(testActions)
		router := &routerAdapter{
			mux: http.NewServeMux(),
		}
		handlers.RegisterBooleanRoutes(controller, handlers.NewHTTPApp(router, handlers.WithLogger(newLogger())))
		return testActions, router.mux
	}

	type testCase = routeTestCase[*booleanControllerTestActions]

	t.Run("parsing", func(t *testing.T) {
		randomReq := func() *handlers.BooleanBooleanParsingRequest {
			return &handlers.BooleanBooleanParsingRequest{
				// path
				BoolParam1: fake.Bool(),
				BoolParam2: fake.Bool(),

				// query
				BoolParam1InQuery: fake.Bool(),
				BoolParam2InQuery: fake.Bool(),
			}
		}

		buildQuery := func(wantReq *handlers.BooleanBooleanParsingRequest) url.Values {
			query := url.Values{}
			query.Add("boolParam1InQuery", strconv.FormatBool(wantReq.BoolParam1InQuery))
			query.Add("boolParam2InQuery", strconv.FormatBool(wantReq.BoolParam2InQuery))
			return query
		}

		runRouteTestCase(t, "should parse and bind valid values", setupRouter,
			func() routeTestCase[*booleanControllerTestActions] {
				wantReq := randomReq()
				query := buildQuery(wantReq)

				return routeTestCase[*booleanControllerTestActions]{
					path: fmt.Sprintf("/boolean/parsing/%v/%v",
						wantReq.BoolParam1, wantReq.BoolParam2),
					query: query,
					expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
						if !assert.Equal(t, 204, recorder.Code, "Unexpected response: %v", recorder.Body) {
							return
						}
						assert.Equal(t, wantReq, testActions.booleanParsing.calls[0].params)
					},
				}
			})

		nonBooleanValues := []string{
			"True",
			"False",
			"0",
			"1",
			fake.Lorem().Word(),
		}

		for _, val := range nonBooleanValues {
			runRouteTestCase(t, "should fail if non boolean value "+val, setupRouter, func() testCase {
				query := url.Values{}
				query.Add("boolParam1InQuery", val)
				query.Add("boolParam2InQuery", val)
				query.Add("optionalBoolParam1InQuery", val)
				query.Add("optionalBoolParam2InQuery", val)

				return testCase{
					path:  "/boolean/required-validation",
					query: query,
					expect: expectBindingErrors[*booleanControllerTestActions](
						[]handlers.FieldBindingError{
							{Field: "boolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "boolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "optionalBoolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
							{Field: "optionalBoolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},
						},
					),
				}
			})
		}
	})

	t.Run("required-validation", func(t *testing.T) {
		buildQuery := func(wantReq *handlers.BooleanBooleanRequiredValidationRequest) url.Values {
			query := url.Values{}
			query.Add("boolParam1InQuery", strconv.FormatBool(wantReq.BoolParam1InQuery))
			query.Add("boolParam2InQuery", strconv.FormatBool(wantReq.BoolParam2InQuery))
			query.Add("optionalBoolParam1InQuery", strconv.FormatBool(wantReq.OptionalBoolParam1InQuery))
			query.Add("optionalBoolParam2InQuery", strconv.FormatBool(wantReq.OptionalBoolParam2InQuery))
			return query
		}

		runRouteTestCase(t, "should ignore missing optional params", setupRouter, func() testCase {
			wantReq := &handlers.BooleanBooleanRequiredValidationRequest{
				// query
				BoolParam1InQuery: fake.Bool(),
				BoolParam2InQuery: fake.Bool(),
			}

			query := buildQuery(wantReq)
			query.Del("optionalBoolParam1InQuery")
			query.Del("optionalBoolParam2InQuery")

			return testCase{
				path:  "/boolean/required-validation",
				query: query,
				expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, wantReq, testActions.booleanRequiredValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should parse optional params", setupRouter, func() testCase {
			wantReq := &handlers.BooleanBooleanRequiredValidationRequest{
				// query
				BoolParam1InQuery:         fake.Bool(),
				BoolParam2InQuery:         fake.Bool(),
				OptionalBoolParam1InQuery: fake.Bool(),
				OptionalBoolParam2InQuery: fake.Bool(),
			}

			return testCase{
				path:  "/boolean/required-validation",
				query: buildQuery(wantReq),
				expect: func(t *testing.T, testActions *booleanControllerTestActions, recorder *httptest.ResponseRecorder) {
					if !assert.Equal(t, 204, recorder.Code, "Got unexpected response: %v", recorder.Body) {
						return
					}
					assert.Equal(t, wantReq, testActions.booleanRequiredValidation.calls[0].params)
				},
			}
		})
		runRouteTestCase(t, "should validate empty params", setupRouter, func() testCase {
			query := url.Values{}
			query.Add("boolParam1InQuery", "")
			query.Add("boolParam2InQuery", "")
			query.Add("optionalBoolParam1InQuery", "")
			query.Add("optionalBoolParam2InQuery", "")

			return testCase{
				path:  "/boolean/required-validation",
				query: query,
				expect: expectBindingErrors[*booleanControllerTestActions](
					[]handlers.FieldBindingError{
						{Field: "boolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "boolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalBoolParam1InQuery", Location: "query", Code: "BAD_FORMAT"},
						{Field: "optionalBoolParam2InQuery", Location: "query", Code: "BAD_FORMAT"},
					},
				),
			}
		})
		runRouteTestCase(t, "should ensure required params", setupRouter, func() testCase {
			query := url.Values{}
			query.Del("boolParam1InQuery")
			query.Del("boolParam2InQuery")

			return testCase{
				path:  "/boolean/required-validation",
				query: query,
				expect: expectBindingErrors[*booleanControllerTestActions](
					[]handlers.FieldBindingError{
						{Field: "boolParam1InQuery", Location: "query", Code: "INVALID_REQUIRED"},
						{Field: "boolParam2InQuery", Location: "query", Code: "INVALID_REQUIRED"},
					},
				),
			}
		})
	})
}
