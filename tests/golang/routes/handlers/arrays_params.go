package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.ArraysNullableRequiredValidationRequest{}

type paramsParserArraysArraysNullableRequiredValidation struct {
	bindSimpleItems1 requestParamBinder[string, []string]
	bindSimpleItems2 requestParamBinder[string, []string]
	bindSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindSimpleItems2InQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *models.ArraysNullableRequiredValidationRequest]
	bindOptionalSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindOptionalSimpleItems2InQuery requestParamBinder[[]string, []string]
}

func (p *paramsParserArraysArraysNullableRequiredValidation) parse(router httpRouter, req *http.Request) (*ArraysArraysNullableRequiredValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ArraysArraysNullableRequiredValidationRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindSimpleItems1(pathParamsCtx.Fork("simpleItems1"), readPathValue("simpleItems1", router, req), &reqParams.SimpleItems1)
	p.bindSimpleItems2(pathParamsCtx.Fork("simpleItems2"), readPathValue("simpleItems2", router, req), &reqParams.SimpleItems2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindSimpleItems1InQuery(queryParamsCtx.Fork("simpleItems1InQuery"), readQueryValue("simpleItems1InQuery", query), &reqParams.SimpleItems1InQuery)
	p.bindSimpleItems2InQuery(queryParamsCtx.Fork("simpleItems2InQuery"), readQueryValue("simpleItems2InQuery", query), &reqParams.SimpleItems2InQuery)
	p.bindOptionalSimpleItems1InQuery(queryParamsCtx.Fork("optionalSimpleItems1InQuery"), readQueryValue("optionalSimpleItems1InQuery", query), &reqParams.OptionalSimpleItems1InQuery)
	p.bindOptionalSimpleItems2InQuery(queryParamsCtx.Fork("optionalSimpleItems2InQuery"), readQueryValue("optionalSimpleItems2InQuery", query), &reqParams.OptionalSimpleItems2InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserArraysArraysNullableRequiredValidation(app *HTTPApp) paramsParser[*ArraysArraysNullableRequiredValidationRequest] {
	return &paramsParserArraysArraysNullableRequiredValidation{
		bindSimpleItems1: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ArraysNullableRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.ArraysNullableRequiredValidationRequest],
			),
			validateValue: internal.NewArraysNullableRequiredValidationRequestValidator(),
		}),
		bindOptionalSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindOptionalSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
	}
}

type paramsParserArraysArraysRangeValidation struct {
	bindSimpleItems1 requestParamBinder[string, []string]
	bindSimpleItems2 requestParamBinder[string, []string]
	bindSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindSimpleItems2InQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *models.ArraysRangeValidationRequest]
	bindOptionalSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindOptionalSimpleItems2InQuery requestParamBinder[[]string, []string]
}

func (p *paramsParserArraysArraysRangeValidation) parse(router httpRouter, req *http.Request) (*ArraysArraysRangeValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ArraysArraysRangeValidationRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindSimpleItems1(pathParamsCtx.Fork("simpleItems1"), readPathValue("simpleItems1", router, req), &reqParams.SimpleItems1)
	p.bindSimpleItems2(pathParamsCtx.Fork("simpleItems2"), readPathValue("simpleItems2", router, req), &reqParams.SimpleItems2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindSimpleItems1InQuery(queryParamsCtx.Fork("simpleItems1InQuery"), readQueryValue("simpleItems1InQuery", query), &reqParams.SimpleItems1InQuery)
	p.bindSimpleItems2InQuery(queryParamsCtx.Fork("simpleItems2InQuery"), readQueryValue("simpleItems2InQuery", query), &reqParams.SimpleItems2InQuery)
	p.bindOptionalSimpleItems1InQuery(queryParamsCtx.Fork("optionalSimpleItems1InQuery"), readQueryValue("optionalSimpleItems1InQuery", query), &reqParams.OptionalSimpleItems1InQuery)
	p.bindOptionalSimpleItems2InQuery(queryParamsCtx.Fork("optionalSimpleItems2InQuery"), readQueryValue("optionalSimpleItems2InQuery", query), &reqParams.OptionalSimpleItems2InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserArraysArraysRangeValidation(app *HTTPApp) paramsParser[*ArraysArraysRangeValidationRequest] {
	return &paramsParserArraysArraysRangeValidation{
		bindSimpleItems1: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ArraysRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.ArraysRangeValidationRequest],
			),
			validateValue: internal.NewArraysRangeValidationRequestValidator(),
		}),
		bindOptionalSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindOptionalSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
	}
}

type paramsParserArraysArraysRequiredValidation struct {
	bindSimpleItems1 requestParamBinder[string, []string]
	bindSimpleItems2 requestParamBinder[string, []string]
	bindSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindSimpleItems2InQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *models.ArraysRequiredValidationRequest]
	bindOptionalSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindOptionalSimpleItems2InQuery requestParamBinder[[]string, []string]
}

func (p *paramsParserArraysArraysRequiredValidation) parse(router httpRouter, req *http.Request) (*ArraysArraysRequiredValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ArraysArraysRequiredValidationRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindSimpleItems1(pathParamsCtx.Fork("simpleItems1"), readPathValue("simpleItems1", router, req), &reqParams.SimpleItems1)
	p.bindSimpleItems2(pathParamsCtx.Fork("simpleItems2"), readPathValue("simpleItems2", router, req), &reqParams.SimpleItems2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindSimpleItems1InQuery(queryParamsCtx.Fork("simpleItems1InQuery"), readQueryValue("simpleItems1InQuery", query), &reqParams.SimpleItems1InQuery)
	p.bindSimpleItems2InQuery(queryParamsCtx.Fork("simpleItems2InQuery"), readQueryValue("simpleItems2InQuery", query), &reqParams.SimpleItems2InQuery)
	p.bindOptionalSimpleItems1InQuery(queryParamsCtx.Fork("optionalSimpleItems1InQuery"), readQueryValue("optionalSimpleItems1InQuery", query), &reqParams.OptionalSimpleItems1InQuery)
	p.bindOptionalSimpleItems2InQuery(queryParamsCtx.Fork("optionalSimpleItems2InQuery"), readQueryValue("optionalSimpleItems2InQuery", query), &reqParams.OptionalSimpleItems2InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserArraysArraysRequiredValidation(app *HTTPApp) paramsParser[*ArraysArraysRequiredValidationRequest] {
	return &paramsParserArraysArraysRequiredValidation{
		bindSimpleItems1: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ArraysRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.ArraysRequiredValidationRequest],
			),
			validateValue: internal.NewArraysRequiredValidationRequestValidator(),
		}),
		bindOptionalSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindOptionalSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
	}
}
