package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.BooleanNullableRequest{}

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
	p.bindBoolParam1(&bindingCtx, readPathValue("boolParam1", router, req), &reqParams.BoolParam1)
	p.bindBoolParam2(&bindingCtx, readPathValue("boolParam2", router, req), &reqParams.BoolParam2)
	// query params
	query := req.URL.Query()
	p.bindBoolParam1InQuery(&bindingCtx, readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(&bindingCtx, readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	p.bindOptionalBoolParam1InQuery(&bindingCtx, readQueryValue("optionalBoolParam1InQuery", query), &reqParams.OptionalBoolParam1InQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanNullable(app *HTTPApp) paramsParser[*BooleanBooleanNullableRequest] {
	return &paramsParserBooleanBooleanNullable{
		bindBoolParam1: newRequestParamBinder(binderParams[string, *bool]{
			field: "boolParam1",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseNullableParam(app.knownParsers.boolParser)),
			validateValue: internal.NewSimpleFieldValidator[*bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam1", Location: "path"},
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, *bool]{
			field: "boolParam2",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseNullableParam(app.knownParsers.boolParser)),
			validateValue: internal.NewSimpleFieldValidator[*bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam2", Location: "path"},
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			field: "boolParam1InQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.boolParser)),
			validateValue: internal.NewSimpleFieldValidator[*bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam1InQuery", Location: "query"},
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			field: "boolParam2InQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.boolParser)),
			validateValue: internal.NewSimpleFieldValidator[*bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam2InQuery", Location: "query"},
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.BooleanNullableRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.BooleanNullableRequest]),
			validateValue: internal.NewBooleanNullableRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
		bindOptionalBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, *bool]{
			field: "optionalBoolParam1InQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.boolParser)),
			validateValue: internal.NewSimpleFieldValidator[*bool](
				internal.SimpleFieldValidatorParams{Field: "optionalBoolParam1InQuery", Location: "query"},
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
	p.bindBoolParam1(&bindingCtx, readPathValue("boolParam1", router, req), &reqParams.BoolParam1)
	p.bindBoolParam2(&bindingCtx, readPathValue("boolParam2", router, req), &reqParams.BoolParam2)
	// query params
	query := req.URL.Query()
	p.bindBoolParam1InQuery(&bindingCtx, readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(&bindingCtx, readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanParsing(app *HTTPApp) paramsParser[*BooleanBooleanParsingRequest] {
	return &paramsParserBooleanBooleanParsing{
		bindBoolParam1: newRequestParamBinder(binderParams[string, bool]{
			field: "boolParam1",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam1", Location: "path"},
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, bool]{
			field: "boolParam2",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam2", Location: "path"},
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam1InQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam1InQuery", Location: "query"},
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam2InQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam2InQuery", Location: "query"},
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.BooleanParsingRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.BooleanParsingRequest]),
			validateValue: internal.NewBooleanParsingRequestValidator(internal.ModelValidatorParams{Location: "body"}),
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
	p.bindBoolParam1InQuery(&bindingCtx, readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(&bindingCtx, readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	p.bindOptionalBoolParam1InQuery(&bindingCtx, readQueryValue("optionalBoolParam1InQuery", query), &reqParams.OptionalBoolParam1InQuery)
	p.bindOptionalBoolParam2InQuery(&bindingCtx, readQueryValue("optionalBoolParam2InQuery", query), &reqParams.OptionalBoolParam2InQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanRequiredValidation(app *HTTPApp) paramsParser[*BooleanBooleanRequiredValidationRequest] {
	return &paramsParserBooleanBooleanRequiredValidation{
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam1InQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam1InQuery", Location: "query"},
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam2InQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "boolParam2InQuery", Location: "query"},
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.BooleanRequiredValidationRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.BooleanRequiredValidationRequest]),
			validateValue: internal.NewBooleanRequiredValidationRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
		bindOptionalBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "optionalBoolParam1InQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "optionalBoolParam1InQuery", Location: "query"},
			),
		}),
		bindOptionalBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "optionalBoolParam2InQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.boolParser),
			validateValue: internal.NewSimpleFieldValidator[bool](
				internal.SimpleFieldValidatorParams{Field: "optionalBoolParam2InQuery", Location: "query"},
			),
		}),
	}
}
