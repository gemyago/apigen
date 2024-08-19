package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.BooleanArrayItemsRequest{}

type paramsParserBooleanBooleanArrayItems struct {
	bindBoolParam1 requestParamBinder[string, []bool]
	bindBoolParam2 requestParamBinder[string, []bool]
	bindBoolParam1InQuery requestParamBinder[[]string, []bool]
	bindBoolParam2InQuery requestParamBinder[[]string, []bool]
	bindPayload requestParamBinder[*http.Request, *models.BooleanArrayItemsRequest]
}

func (p *paramsParserBooleanBooleanArrayItems) parse(router httpRouter, req *http.Request) (*BooleanBooleanArrayItemsRequest, error) {
	bindingCtx := internal.BindingContext{}
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

func newParamsParserBooleanBooleanArrayItems(app *HTTPApp) paramsParser[*BooleanBooleanArrayItemsRequest] {
	return &paramsParserBooleanBooleanArrayItems{
		bindBoolParam1: newRequestParamBinder(binderParams[string, []bool]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, []bool]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, []bool]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, []bool]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[bool](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.BooleanArrayItemsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.BooleanArrayItemsRequest],
			),
			validateValue: internal.NewBooleanArrayItemsRequestValidator(),
		}),
	}
}

type paramsParserBooleanBooleanNullable struct {
	bindBoolParam1 requestParamBinder[string, *bool]
	bindBoolParam2 requestParamBinder[string, *bool]
	bindBoolParam1InQuery requestParamBinder[[]string, *bool]
	bindBoolParam2InQuery requestParamBinder[[]string, *bool]
	bindPayload requestParamBinder[*http.Request, *models.BooleanNullableRequest]
	bindOptionalBoolParam1InQuery requestParamBinder[[]string, *bool]
}

func (p *paramsParserBooleanBooleanNullable) parse(router httpRouter, req *http.Request) (*BooleanBooleanNullableRequest, error) {
	bindingCtx := internal.BindingContext{}
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

func newParamsParserBooleanBooleanNullable(app *HTTPApp) paramsParser[*BooleanBooleanNullableRequest] {
	return &paramsParserBooleanBooleanNullable{
		bindBoolParam1: newRequestParamBinder(binderParams[string, *bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.boolParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*bool](
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, *bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.boolParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*bool](
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.boolParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.boolParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*bool](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.BooleanNullableRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.BooleanNullableRequest],
			),
			validateValue: internal.NewBooleanNullableRequestValidator(),
		}),
		bindOptionalBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.boolParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*bool](
			),
		}),
	}
}

type paramsParserBooleanBooleanParsing struct {
	bindBoolParam1 requestParamBinder[string, bool]
	bindBoolParam2 requestParamBinder[string, bool]
	bindBoolParam1InQuery requestParamBinder[[]string, bool]
	bindBoolParam2InQuery requestParamBinder[[]string, bool]
	bindPayload requestParamBinder[*http.Request, *models.BooleanParsingRequest]
}

func (p *paramsParserBooleanBooleanParsing) parse(router httpRouter, req *http.Request) (*BooleanBooleanParsingRequest, error) {
	bindingCtx := internal.BindingContext{}
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

func newParamsParserBooleanBooleanParsing(app *HTTPApp) paramsParser[*BooleanBooleanParsingRequest] {
	return &paramsParserBooleanBooleanParsing{
		bindBoolParam1: newRequestParamBinder(binderParams[string, bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, bool]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.BooleanParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.BooleanParsingRequest],
			),
			validateValue: internal.NewBooleanParsingRequestValidator(),
		}),
	}
}

type paramsParserBooleanBooleanRequiredValidation struct {
	bindBoolParam1InQuery requestParamBinder[[]string, bool]
	bindBoolParam2InQuery requestParamBinder[[]string, bool]
	bindPayload requestParamBinder[*http.Request, *models.BooleanRequiredValidationRequest]
	bindOptionalBoolParam1InQuery requestParamBinder[[]string, bool]
	bindOptionalBoolParam2InQuery requestParamBinder[[]string, bool]
}

func (p *paramsParserBooleanBooleanRequiredValidation) parse(router httpRouter, req *http.Request) (*BooleanBooleanRequiredValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
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

func newParamsParserBooleanBooleanRequiredValidation(app *HTTPApp) paramsParser[*BooleanBooleanRequiredValidationRequest] {
	return &paramsParserBooleanBooleanRequiredValidation{
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.BooleanRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.BooleanRequiredValidationRequest],
			),
			validateValue: internal.NewBooleanRequiredValidationRequestValidator(),
		}),
		bindOptionalBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindOptionalBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.boolParser,
			),
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
	}
}
