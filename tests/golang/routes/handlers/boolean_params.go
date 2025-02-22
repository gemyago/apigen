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

func (p *paramsParserBooleanBooleanArrayItems) parse(router httpRouter, req *http.Request) (*BooleanBooleanArrayItemsRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanArrayItemsRequest{}
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

func newParamsParserBooleanBooleanArrayItems(rootHandler *RootHandler) paramsParser[*BooleanBooleanArrayItemsRequest] {
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

func (p *paramsParserBooleanBooleanNullable) parse(router httpRouter, req *http.Request) (*BooleanBooleanNullableRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanNullableRequest{}
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

func newParamsParserBooleanBooleanNullable(rootHandler *RootHandler) paramsParser[*BooleanBooleanNullableRequest] {
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

func (p *paramsParserBooleanBooleanNullableArrayItems) parse(router httpRouter, req *http.Request) (*BooleanBooleanNullableArrayItemsRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanNullableArrayItemsRequest{}
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

func newParamsParserBooleanBooleanNullableArrayItems(rootHandler *RootHandler) paramsParser[*BooleanBooleanNullableArrayItemsRequest] {
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

func (p *paramsParserBooleanBooleanParsing) parse(router httpRouter, req *http.Request) (*BooleanBooleanParsingRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanParsingRequest{}
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

func newParamsParserBooleanBooleanParsing(rootHandler *RootHandler) paramsParser[*BooleanBooleanParsingRequest] {
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

func (p *paramsParserBooleanBooleanRequiredValidation) parse(router httpRouter, req *http.Request) (*BooleanBooleanRequiredValidationRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BooleanBooleanRequiredValidationRequest{}
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

func newParamsParserBooleanBooleanRequiredValidation(rootHandler *RootHandler) paramsParser[*BooleanBooleanRequiredValidationRequest] {
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
