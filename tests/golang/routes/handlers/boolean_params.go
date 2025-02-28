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
type _ func() BooleanArrayItemsRequest

type paramsParserBooleanBooleanArrayItems struct {
	bindBoolParam1 requestParamBinder[string, []bool]
	bindBoolParam2 requestParamBinder[string, []bool]
	bindBoolParam1InQuery requestParamBinder[[]string, []bool]
	bindBoolParam2InQuery requestParamBinder[[]string, []bool]
	bindPayload requestParamBinder[*http.Request, *BooleanArrayItemsRequest]
}

func (p *paramsParserBooleanBooleanArrayItems) parse(router httpRouter, req *http.Request) (*BooleanBooleanArrayItemsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanArrayItemsParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindBoolParam1(pathParamsCtx.Fork("boolParam1"), readPathValue("boolParam1", router, req), &reqParams.BoolParam1)
	p.bindBoolParam2(pathParamsCtx.Fork("boolParam2"), readPathValue("boolParam2", router, req), &reqParams.BoolParam2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindBoolParam1InQuery(queryParamsCtx.Fork("boolParam1InQuery"), readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(queryParamsCtx.Fork("boolParam2InQuery"), readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanArrayItems(rootHandler *RootHandler) paramsParser[*BooleanBooleanArrayItemsParams] {
	return &paramsParserBooleanBooleanArrayItems{
		bindBoolParam1: newRequestParamBinder(binderParams[string, []bool]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]bool](
				),
				NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, []bool]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]bool](
				),
				NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, []bool]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]bool](
				),
				NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, []bool]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]bool](
				),
				NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BooleanArrayItemsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BooleanArrayItemsRequest],
			),
			validateValue: NewBooleanArrayItemsRequestValidator(),
		}),
	}
}

type paramsParserBooleanBooleanNullable struct {
	bindBoolParam1 requestParamBinder[string, *bool]
	bindBoolParam2 requestParamBinder[string, *bool]
	bindBoolParam1InQuery requestParamBinder[[]string, *bool]
	bindBoolParam2InQuery requestParamBinder[[]string, *bool]
	bindPayload requestParamBinder[*http.Request, *BooleanNullableRequest]
	bindOptionalBoolParam1InQuery requestParamBinder[[]string, *bool]
}

func (p *paramsParserBooleanBooleanNullable) parse(router httpRouter, req *http.Request) (*BooleanBooleanNullableParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanNullableParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindBoolParam1(pathParamsCtx.Fork("boolParam1"), readPathValue("boolParam1", router, req), &reqParams.BoolParam1)
	p.bindBoolParam2(pathParamsCtx.Fork("boolParam2"), readPathValue("boolParam2", router, req), &reqParams.BoolParam2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindBoolParam1InQuery(queryParamsCtx.Fork("boolParam1InQuery"), readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(queryParamsCtx.Fork("boolParam2InQuery"), readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	p.bindOptionalBoolParam1InQuery(queryParamsCtx.Fork("optionalBoolParam1InQuery"), readQueryValue("optionalBoolParam1InQuery", query), &reqParams.OptionalBoolParam1InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanNullable(rootHandler *RootHandler) paramsParser[*BooleanBooleanNullableParams] {
	return &paramsParserBooleanBooleanNullable{
		bindBoolParam1: newRequestParamBinder(binderParams[string, *bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewSimpleFieldValidator[*bool](
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, *bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewSimpleFieldValidator[*bool](
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewSimpleFieldValidator[*bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewSimpleFieldValidator[*bool](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BooleanNullableRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BooleanNullableRequest],
			),
			validateValue: NewBooleanNullableRequestValidator(),
		}),
		bindOptionalBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewSimpleFieldValidator[*bool](
			),
		}),
	}
}

type paramsParserBooleanBooleanNullableArrayItems struct {
	bindBoolParam1 requestParamBinder[string, []*bool]
	bindBoolParam2 requestParamBinder[string, []*bool]
	bindBoolParam1InQuery requestParamBinder[[]string, []*bool]
	bindBoolParam2InQuery requestParamBinder[[]string, []*bool]
	bindPayload requestParamBinder[*http.Request, *BooleanNullableArrayItemsRequest]
}

func (p *paramsParserBooleanBooleanNullableArrayItems) parse(router httpRouter, req *http.Request) (*BooleanBooleanNullableArrayItemsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanNullableArrayItemsParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindBoolParam1(pathParamsCtx.Fork("boolParam1"), readPathValue("boolParam1", router, req), &reqParams.BoolParam1)
	p.bindBoolParam2(pathParamsCtx.Fork("boolParam2"), readPathValue("boolParam2", router, req), &reqParams.BoolParam2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindBoolParam1InQuery(queryParamsCtx.Fork("boolParam1InQuery"), readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(queryParamsCtx.Fork("boolParam2InQuery"), readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanNullableArrayItems(rootHandler *RootHandler) paramsParser[*BooleanBooleanNullableArrayItemsParams] {
	return &paramsParserBooleanBooleanNullableArrayItems{
		bindBoolParam1: newRequestParamBinder(binderParams[string, []*bool]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*bool](
				),
				NewSimpleFieldValidator[*bool](
				),
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, []*bool]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*bool](
				),
				NewSimpleFieldValidator[*bool](
				),
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, []*bool]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*bool](
				),
				NewSimpleFieldValidator[*bool](
				),
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, []*bool]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.boolParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*bool](
				),
				NewSimpleFieldValidator[*bool](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BooleanNullableArrayItemsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BooleanNullableArrayItemsRequest],
			),
			validateValue: NewBooleanNullableArrayItemsRequestValidator(),
		}),
	}
}

type paramsParserBooleanBooleanParsing struct {
	bindBoolParam1 requestParamBinder[string, bool]
	bindBoolParam2 requestParamBinder[string, bool]
	bindBoolParam1InQuery requestParamBinder[[]string, bool]
	bindBoolParam2InQuery requestParamBinder[[]string, bool]
	bindPayload requestParamBinder[*http.Request, *BooleanParsingRequest]
}

func (p *paramsParserBooleanBooleanParsing) parse(router httpRouter, req *http.Request) (*BooleanBooleanParsingParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanParsingParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindBoolParam1(pathParamsCtx.Fork("boolParam1"), readPathValue("boolParam1", router, req), &reqParams.BoolParam1)
	p.bindBoolParam2(pathParamsCtx.Fork("boolParam2"), readPathValue("boolParam2", router, req), &reqParams.BoolParam2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindBoolParam1InQuery(queryParamsCtx.Fork("boolParam1InQuery"), readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(queryParamsCtx.Fork("boolParam2InQuery"), readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanParsing(rootHandler *RootHandler) paramsParser[*BooleanBooleanParsingParams] {
	return &paramsParserBooleanBooleanParsing{
		bindBoolParam1: newRequestParamBinder(binderParams[string, bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BooleanParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BooleanParsingRequest],
			),
			validateValue: NewBooleanParsingRequestValidator(),
		}),
	}
}

type paramsParserBooleanBooleanRequiredValidation struct {
	bindBoolParam1InQuery requestParamBinder[[]string, bool]
	bindBoolParam2InQuery requestParamBinder[[]string, bool]
	bindPayload requestParamBinder[*http.Request, *BooleanRequiredValidationRequest]
	bindOptionalBoolParam1InQuery requestParamBinder[[]string, bool]
	bindOptionalBoolParam2InQuery requestParamBinder[[]string, bool]
}

func (p *paramsParserBooleanBooleanRequiredValidation) parse(router httpRouter, req *http.Request) (*BooleanBooleanRequiredValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanRequiredValidationParams{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindBoolParam1InQuery(queryParamsCtx.Fork("boolParam1InQuery"), readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(queryParamsCtx.Fork("boolParam2InQuery"), readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	p.bindOptionalBoolParam1InQuery(queryParamsCtx.Fork("optionalBoolParam1InQuery"), readQueryValue("optionalBoolParam1InQuery", query), &reqParams.OptionalBoolParam1InQuery)
	p.bindOptionalBoolParam2InQuery(queryParamsCtx.Fork("optionalBoolParam2InQuery"), readQueryValue("optionalBoolParam2InQuery", query), &reqParams.OptionalBoolParam2InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanRequiredValidation(rootHandler *RootHandler) paramsParser[*BooleanBooleanRequiredValidationParams] {
	return &paramsParserBooleanBooleanRequiredValidation{
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BooleanRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BooleanRequiredValidationRequest],
			),
			validateValue: NewBooleanRequiredValidationRequestValidator(),
		}),
		bindOptionalBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
		bindOptionalBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.boolParser,
			),
			validateValue: NewSimpleFieldValidator[bool](
			),
		}),
	}
}

type booleanControllerBuilder struct {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsParams,
	//
	// Response type: none
	BooleanArrayItems genericHandlerBuilder[
		*BooleanBooleanArrayItemsParams,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanArrayItemsParams, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanArrayItemsParams, void],
	]

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableParams,
	//
	// Response type: none
	BooleanNullable genericHandlerBuilder[
		*BooleanBooleanNullableParams,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanNullableParams, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanNullableParams, void],
	]

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsParams,
	//
	// Response type: none
	BooleanNullableArrayItems genericHandlerBuilder[
		*BooleanBooleanNullableArrayItemsParams,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanNullableArrayItemsParams, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanNullableArrayItemsParams, void],
	]

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingParams,
	//
	// Response type: none
	BooleanParsing genericHandlerBuilder[
		*BooleanBooleanParsingParams,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanParsingParams, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanParsingParams, void],
	]

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationParams,
	//
	// Response type: none
	BooleanRequiredValidation genericHandlerBuilder[
		*BooleanBooleanRequiredValidationParams,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanRequiredValidationParams, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanRequiredValidationParams, void],
	]
}

func newBooleanControllerBuilder(app *RootHandler) *booleanControllerBuilder {
	return &booleanControllerBuilder{
		// POST /boolean/array-items/{boolParam1}/{boolParam2}
		BooleanArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*BooleanBooleanArrayItemsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanArrayItemsParams,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanArrayItemsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanArrayItems(app),
			},
		),

		// POST /boolean/nullable/{boolParam1}/{boolParam2}
		BooleanNullable: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*BooleanBooleanNullableParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanNullableParams,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanNullableParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanNullable(app),
			},
		),

		// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
		BooleanNullableArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*BooleanBooleanNullableArrayItemsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanNullableArrayItemsParams,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanNullableArrayItemsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanNullableArrayItems(app),
			},
		),

		// POST /boolean/parsing/{boolParam1}/{boolParam2}
		BooleanParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*BooleanBooleanParsingParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanParsingParams,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanParsingParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanParsing(app),
			},
		),

		// POST /boolean/required-validation
		BooleanRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*BooleanBooleanRequiredValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanRequiredValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanRequiredValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanRequiredValidation(app),
			},
		),
	}
}
