package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.StringTypesNullableParsingRequest{}

type paramsParserStringTypesStringTypesNullableParsing struct {
	bindUnformattedStr requestParamBinder[string, *string]
	bindCustomFormatStr requestParamBinder[string, *string]
	bindDateStr requestParamBinder[string, *time.Time]
	bindDateTimeStr requestParamBinder[string, *time.Time]
	bindByteStr requestParamBinder[string, *string]
	bindUnformattedStrInQuery requestParamBinder[[]string, *string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, *string]
	bindDateStrInQuery requestParamBinder[[]string, *time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, *time.Time]
	bindByteStrInQuery requestParamBinder[[]string, *string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesNullableParsingRequest]
}

func (p *paramsParserStringTypesStringTypesNullableParsing) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesNullableParsingRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesNullableParsingRequest{}
	// path params
	p.bindUnformattedStr(&bindingCtx, readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(&bindingCtx, readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(&bindingCtx, readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(&bindingCtx, readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(&bindingCtx, readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(&bindingCtx, readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(&bindingCtx, readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(&bindingCtx, readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(&bindingCtx, readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableParsing(app *HTTPApp) paramsParser[*StringTypesStringTypesNullableParsingRequest] {
	return &paramsParserStringTypesStringTypesNullableParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, *string]{
			field: "unformattedStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStr", Location: "path"},
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, *string]{
			field: "customFormatStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStr", Location: "path"},
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, *time.Time]{
			field: "dateStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseNullableParam(app.knownParsers.dateParser)),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStr", Location: "path"},
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, *time.Time]{
			field: "dateTimeStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseNullableParam(app.knownParsers.timeParser)),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStr", Location: "path"},
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, *string]{
			field: "byteStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "byteStr", Location: "path"},
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStrInQuery", Location: "query"},
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStrInQuery", Location: "query"},
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, *time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.dateParser)),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStrInQuery", Location: "query"},
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, *time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.timeParser)),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStrInQuery", Location: "query"},
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			field: "byteStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "byteStrInQuery", Location: "query"},
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesNullableParsingRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.StringTypesNullableParsingRequest]),
			validateValue: internal.NewStringTypesNullableParsingRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
	}
}

type paramsParserStringTypesStringTypesNullableRequiredValidation struct {
	bindUnformattedStrInQuery requestParamBinder[[]string, *string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesNullableRequiredValidationRequest]
	bindOptionalUnformattedStrInQuery requestParamBinder[[]string, *string]
}

func (p *paramsParserStringTypesStringTypesNullableRequiredValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesNullableRequiredValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesNullableRequiredValidationRequest{}
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindOptionalUnformattedStrInQuery(&bindingCtx, readQueryValue("optionalUnformattedStrInQuery", query), &reqParams.OptionalUnformattedStrInQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableRequiredValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesNullableRequiredValidationRequest] {
	return &paramsParserStringTypesStringTypesNullableRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStrInQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string](10, true)),
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string](100, false)),
				internal.SkipNullValidator(internal.NewPatternValidator[string](".*")),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesNullableRequiredValidationRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.StringTypesNullableRequiredValidationRequest]),
			validateValue: internal.NewStringTypesNullableRequiredValidationRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			field: "optionalUnformattedStrInQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(parseNullableParam(app.knownParsers.stringParser)),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SimpleFieldValidatorParams{Field: "optionalUnformattedStrInQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string](10, true)),
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string](100, false)),
				internal.SkipNullValidator(internal.NewPatternValidator[string](".*")),
			),
		}),
	}
}

type paramsParserStringTypesStringTypesParsing struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindDateStr requestParamBinder[string, time.Time]
	bindDateTimeStr requestParamBinder[string, time.Time]
	bindByteStr requestParamBinder[string, string]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindByteStrInQuery requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesParsingRequest]
}

func (p *paramsParserStringTypesStringTypesParsing) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesParsingRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesParsingRequest{}
	// path params
	p.bindUnformattedStr(&bindingCtx, readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(&bindingCtx, readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(&bindingCtx, readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(&bindingCtx, readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(&bindingCtx, readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(&bindingCtx, readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(&bindingCtx, readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(&bindingCtx, readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(&bindingCtx, readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesParsing(app *HTTPApp) paramsParser[*StringTypesStringTypesParsingRequest] {
	return &paramsParserStringTypesStringTypesParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field: "unformattedStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStr", Location: "path"},
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStr", Location: "path"},
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStr", Location: "path"},
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateTimeStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStr", Location: "path"},
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field: "byteStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "byteStr", Location: "path"},
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStrInQuery", Location: "query"},
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStrInQuery", Location: "query"},
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStrInQuery", Location: "query"},
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStrInQuery", Location: "query"},
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "byteStrInQuery", Location: "query"},
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesParsingRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.StringTypesParsingRequest]),
			validateValue: internal.NewStringTypesParsingRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
	}
}

type paramsParserStringTypesStringTypesPatternValidation struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindDateStr requestParamBinder[string, time.Time]
	bindDateTimeStr requestParamBinder[string, time.Time]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesPatternValidationRequest]
}

func (p *paramsParserStringTypesStringTypesPatternValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesPatternValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesPatternValidationRequest{}
	// path params
	p.bindUnformattedStr(&bindingCtx, readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(&bindingCtx, readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(&bindingCtx, readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(&bindingCtx, readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(&bindingCtx, readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(&bindingCtx, readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(&bindingCtx, readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesPatternValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesPatternValidationRequest] {
	return &paramsParserStringTypesStringTypesPatternValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field: "unformattedStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStr", Location: "path"},
				internal.NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStr", Location: "path"},
				internal.NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStr", Location: "path"},
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateTimeStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStr", Location: "path"},
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStrInQuery", Location: "query"},
				internal.NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStrInQuery", Location: "query"},
				internal.NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStrInQuery", Location: "query"},
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStrInQuery", Location: "query"},
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesPatternValidationRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.StringTypesPatternValidationRequest]),
			validateValue: internal.NewStringTypesPatternValidationRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
	}
}

type paramsParserStringTypesStringTypesRangeValidation struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindDateStr requestParamBinder[string, time.Time]
	bindDateTimeStr requestParamBinder[string, time.Time]
	bindByteStr requestParamBinder[string, string]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindByteStrInQuery requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesRangeValidationRequest]
}

func (p *paramsParserStringTypesStringTypesRangeValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesRangeValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesRangeValidationRequest{}
	// path params
	p.bindUnformattedStr(&bindingCtx, readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(&bindingCtx, readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(&bindingCtx, readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(&bindingCtx, readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(&bindingCtx, readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(&bindingCtx, readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(&bindingCtx, readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(&bindingCtx, readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(&bindingCtx, readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRangeValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesRangeValidationRequest] {
	return &paramsParserStringTypesStringTypesRangeValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field: "unformattedStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStr", Location: "path"},
				internal.NewMinMaxLengthValidator[string](10, true),
				internal.NewMinMaxLengthValidator[string](20, false),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStr", Location: "path"},
				internal.NewMinMaxLengthValidator[string](20, true),
				internal.NewMinMaxLengthValidator[string](30, false),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStr", Location: "path"},
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateTimeStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStr", Location: "path"},
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field: "byteStr",
			location: "path",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "byteStr", Location: "path"},
				internal.NewMinMaxLengthValidator[string](30, true),
				internal.NewMinMaxLengthValidator[string](40, false),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](10, true),
				internal.NewMinMaxLengthValidator[string](20, false),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](20, true),
				internal.NewMinMaxLengthValidator[string](30, false),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStrInQuery", Location: "query"},
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStrInQuery", Location: "query"},
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "byteStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](30, true),
				internal.NewMinMaxLengthValidator[string](40, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesRangeValidationRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.StringTypesRangeValidationRequest]),
			validateValue: internal.NewStringTypesRangeValidationRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
	}
}

type paramsParserStringTypesStringTypesRequiredValidation struct {
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindByteStrInQuery requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesRequiredValidationRequest]
	bindOptionalUnformattedStrInQuery requestParamBinder[[]string, string]
	bindOptionalCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindOptionalDateStrInQuery requestParamBinder[[]string, time.Time]
	bindOptionalDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindOptionalByteStrInQuery requestParamBinder[[]string, string]
}

func (p *paramsParserStringTypesStringTypesRequiredValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesRequiredValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesRequiredValidationRequest{}
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(&bindingCtx, readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(&bindingCtx, readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(&bindingCtx, readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(&bindingCtx, readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	p.bindOptionalUnformattedStrInQuery(&bindingCtx, readQueryValue("optionalUnformattedStrInQuery", query), &reqParams.OptionalUnformattedStrInQuery)
	p.bindOptionalCustomFormatStrInQuery(&bindingCtx, readQueryValue("optionalCustomFormatStrInQuery", query), &reqParams.OptionalCustomFormatStrInQuery)
	p.bindOptionalDateStrInQuery(&bindingCtx, readQueryValue("optionalDateStrInQuery", query), &reqParams.OptionalDateStrInQuery)
	p.bindOptionalDateTimeStrInQuery(&bindingCtx, readQueryValue("optionalDateTimeStrInQuery", query), &reqParams.OptionalDateTimeStrInQuery)
	p.bindOptionalByteStrInQuery(&bindingCtx, readQueryValue("optionalByteStrInQuery", query), &reqParams.OptionalByteStrInQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRequiredValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesRequiredValidationRequest] {
	return &paramsParserStringTypesStringTypesRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "unformattedStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](10, true),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "customFormatStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](20, true),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateStrInQuery", Location: "query"},
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "dateTimeStrInQuery", Location: "query"},
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "byteStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](30, true),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesRequiredValidationRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.StringTypesRequiredValidationRequest]),
			validateValue: internal.NewStringTypesRequiredValidationRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalUnformattedStrInQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "optionalUnformattedStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](10, true),
			),
		}),
		bindOptionalCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalCustomFormatStrInQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "optionalCustomFormatStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](20, true),
			),
		}),
		bindOptionalDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "optionalDateStrInQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.dateParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "optionalDateStrInQuery", Location: "query"},
			),
		}),
		bindOptionalDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "optionalDateTimeStrInQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.timeParser),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
				internal.SimpleFieldValidatorParams{Field: "optionalDateTimeStrInQuery", Location: "query"},
			),
		}),
		bindOptionalByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalByteStrInQuery",
			location: "query",
			required: false,
			parseValue: parseMultiValueParamAsSingleValue(app.knownParsers.stringParser),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "optionalByteStrInQuery", Location: "query"},
				internal.NewMinMaxLengthValidator[string](30, true),
			),
		}),
	}
}
