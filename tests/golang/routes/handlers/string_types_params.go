package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}

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
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.dateInPath,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateTimeStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.timeInPath,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field: "byteStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesParsingRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.StringTypesParsingRequest],
			validateValue: internal.NewStringTypesParsingRequestValidator(),
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
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesPatternValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesPatternValidationRequest] {
	return &paramsParserStringTypesStringTypesPatternValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field: "unformattedStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.dateInPath,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateTimeStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.timeInPath,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
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
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](10, true),
				internal.NewMinMaxLengthValidator[string](20, false),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](20, true),
				internal.NewMinMaxLengthValidator[string](30, false),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.dateInPath,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field: "dateTimeStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.timeInPath,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field: "byteStr",
			location: "path",
			required: true,
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](30, true),
				internal.NewMinMaxLengthValidator[string](40, false),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](10, true),
				internal.NewMinMaxLengthValidator[string](20, false),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](20, true),
				internal.NewMinMaxLengthValidator[string](30, false),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](30, true),
				internal.NewMinMaxLengthValidator[string](40, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesRangeValidationRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.StringTypesRangeValidationRequest],
			validateValue: internal.NewStringTypesRangeValidationRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesRequiredValidation struct {
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindByteStrInQuery requestParamBinder[[]string, string]
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
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRequiredValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesRequiredValidationRequest] {
	return &paramsParserStringTypesStringTypesRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](10, true),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](20, true),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "dateTimeStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](30, true),
			),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalUnformattedStrInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](10, true),
			),
		}),
		bindOptionalCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalCustomFormatStrInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](20, true),
			),
		}),
		bindOptionalDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "optionalDateStrInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindOptionalDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field: "optionalDateTimeStrInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindOptionalByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalByteStrInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string](30, true),
			),
		}),
	}
}
