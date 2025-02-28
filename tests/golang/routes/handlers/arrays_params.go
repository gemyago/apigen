package handlers

import (
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
	. "github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports if that happens.
var _ = BindingContext{}
var _ = http.MethodGet
var _ = time.Time{}
type _ func() ArraysNullableRequiredValidationRequest

type paramsParserArraysArraysNullableRequiredValidation struct {
	bindSimpleItems1 requestParamBinder[string, []string]
	bindSimpleItems2 requestParamBinder[string, []string]
	bindSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindSimpleItems2InQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *ArraysNullableRequiredValidationRequest]
	bindOptionalSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindOptionalSimpleItems2InQuery requestParamBinder[[]string, []string]
}

func (p *paramsParserArraysArraysNullableRequiredValidation) parse(router httpRouter, req *http.Request) (*ArraysNullableRequiredValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ArraysNullableRequiredValidationParams{}
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

func newParamsParserArraysArraysNullableRequiredValidation(rootHandler *RootHandler) paramsParser[*ArraysNullableRequiredValidationParams] {
	return &paramsParserArraysArraysNullableRequiredValidation{
		bindSimpleItems1: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *ArraysNullableRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*ArraysNullableRequiredValidationRequest],
			),
			validateValue: NewArraysNullableRequiredValidationRequestValidator(),
		}),
		bindOptionalSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindOptionalSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
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
	bindPayload requestParamBinder[*http.Request, *ArraysRangeValidationRequest]
	bindOptionalSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindOptionalSimpleItems2InQuery requestParamBinder[[]string, []string]
}

func (p *paramsParserArraysArraysRangeValidation) parse(router httpRouter, req *http.Request) (*ArraysRangeValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ArraysRangeValidationParams{}
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

func newParamsParserArraysArraysRangeValidation(rootHandler *RootHandler) paramsParser[*ArraysRangeValidationParams] {
	return &paramsParserArraysArraysRangeValidation{
		bindSimpleItems1: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
					NewMinMaxLengthValidator[string, []string](5, true),
					NewMinMaxLengthValidator[string, []string](10, false),
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
					NewMinMaxLengthValidator[string, []string](10, true),
					NewMinMaxLengthValidator[string, []string](15, false),
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
					NewMinMaxLengthValidator[string, []string](5, true),
					NewMinMaxLengthValidator[string, []string](10, false),
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
					NewMinMaxLengthValidator[string, []string](10, true),
					NewMinMaxLengthValidator[string, []string](15, false),
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *ArraysRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*ArraysRangeValidationRequest],
			),
			validateValue: NewArraysRangeValidationRequestValidator(),
		}),
		bindOptionalSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
					NewMinMaxLengthValidator[string, []string](15, true),
					NewMinMaxLengthValidator[string, []string](20, false),
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindOptionalSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
					NewMinMaxLengthValidator[string, []string](20, true),
					NewMinMaxLengthValidator[string, []string](25, false),
				),
				NewSimpleFieldValidator[string](
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
	bindPayload requestParamBinder[*http.Request, *ArraysRequiredValidationRequest]
	bindOptionalSimpleItems1InQuery requestParamBinder[[]string, []string]
	bindOptionalSimpleItems2InQuery requestParamBinder[[]string, []string]
}

func (p *paramsParserArraysArraysRequiredValidation) parse(router httpRouter, req *http.Request) (*ArraysRequiredValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ArraysRequiredValidationParams{}
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

func newParamsParserArraysArraysRequiredValidation(rootHandler *RootHandler) paramsParser[*ArraysRequiredValidationParams] {
	return &paramsParserArraysArraysRequiredValidation{
		bindSimpleItems1: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *ArraysRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*ArraysRequiredValidationRequest],
			),
			validateValue: NewArraysRequiredValidationRequestValidator(),
		}),
		bindOptionalSimpleItems1InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindOptionalSimpleItems2InQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: false,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
	}
}

type arraysControllerBuilder struct {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysNullableRequiredValidationParams,
	//
	// Response type: none
	ArraysNullableRequiredValidation genericHandlerBuilder[
		*ArraysNullableRequiredValidationParams,
		void,
		handlerActionFuncNoResponse[*ArraysNullableRequiredValidationParams, void],
		httpHandlerActionFuncNoResponse[*ArraysNullableRequiredValidationParams, void],
	]

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysRangeValidationParams,
	//
	// Response type: none
	ArraysRangeValidation genericHandlerBuilder[
		*ArraysRangeValidationParams,
		void,
		handlerActionFuncNoResponse[*ArraysRangeValidationParams, void],
		httpHandlerActionFuncNoResponse[*ArraysRangeValidationParams, void],
	]

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysRequiredValidationParams,
	//
	// Response type: none
	ArraysRequiredValidation genericHandlerBuilder[
		*ArraysRequiredValidationParams,
		void,
		handlerActionFuncNoResponse[*ArraysRequiredValidationParams, void],
		httpHandlerActionFuncNoResponse[*ArraysRequiredValidationParams, void],
	]
}

func newArraysControllerBuilder(app *RootHandler) *arraysControllerBuilder {
	return &arraysControllerBuilder{
		// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
		ArraysNullableRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ArraysNullableRequiredValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ArraysNullableRequiredValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*ArraysNullableRequiredValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysNullableRequiredValidation(app),
			},
		),

		// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
		ArraysRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ArraysRangeValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ArraysRangeValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*ArraysRangeValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysRangeValidation(app),
			},
		),

		// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
		ArraysRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ArraysRequiredValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ArraysRequiredValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*ArraysRequiredValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysRequiredValidation(app),
			},
		),
	}
}
