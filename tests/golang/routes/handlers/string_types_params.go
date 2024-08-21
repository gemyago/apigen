package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.StringTypesArrayItemsRangeValidationRequest{}

type paramsParserStringTypesStringTypesArrayItemsRangeValidation struct {
	bindUnformattedStr requestParamBinder[string, []string]
	bindCustomFormatStr requestParamBinder[string, []string]
	bindDateStr requestParamBinder[string, []time.Time]
	bindDateTimeStr requestParamBinder[string, []time.Time]
	bindByteStr requestParamBinder[string, []string]
	bindUnformattedStrInQuery requestParamBinder[[]string, []string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, []string]
	bindDateStrInQuery requestParamBinder[[]string, []time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, []time.Time]
	bindByteStrInQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesArrayItemsRangeValidationRequest]
}

func (p *paramsParserStringTypesStringTypesArrayItemsRangeValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesArrayItemsRangeValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesArrayItemsRangeValidationRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesArrayItemsRangeValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesArrayItemsRangeValidationRequest] {
	return &paramsParserStringTypesStringTypesArrayItemsRangeValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
					internal.NewMinMaxLengthValidator[string, string](10, true),
					internal.NewMinMaxLengthValidator[string, string](20, false),
				),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
					internal.NewMinMaxLengthValidator[string, string](20, true),
					internal.NewMinMaxLengthValidator[string, string](30, false),
				),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
					internal.NewMinMaxLengthValidator[string, string](30, true),
					internal.NewMinMaxLengthValidator[string, string](40, false),
				),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
					internal.NewMinMaxLengthValidator[string, string](10, true),
					internal.NewMinMaxLengthValidator[string, string](20, false),
				),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
					internal.NewMinMaxLengthValidator[string, string](20, true),
					internal.NewMinMaxLengthValidator[string, string](30, false),
				),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
					internal.NewMinMaxLengthValidator[string, string](30, true),
					internal.NewMinMaxLengthValidator[string, string](40, false),
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesArrayItemsRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesArrayItemsRangeValidationRequest],
			),
			validateValue: internal.NewStringTypesArrayItemsRangeValidationRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesArraysParsing struct {
	bindUnformattedStr requestParamBinder[string, []string]
	bindCustomFormatStr requestParamBinder[string, []string]
	bindDateStr requestParamBinder[string, []time.Time]
	bindDateTimeStr requestParamBinder[string, []time.Time]
	bindByteStr requestParamBinder[string, []string]
	bindUnformattedStrInQuery requestParamBinder[[]string, []string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, []string]
	bindDateStrInQuery requestParamBinder[[]string, []time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, []time.Time]
	bindByteStrInQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesArraysParsingRequest]
}

func (p *paramsParserStringTypesStringTypesArraysParsing) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesArraysParsingRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesArraysParsingRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesArraysParsing(app *HTTPApp) paramsParser[*StringTypesStringTypesArraysParsingRequest] {
	return &paramsParserStringTypesStringTypesArraysParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]time.Time](
				),
				internal.NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]string](
				),
				internal.NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesArraysParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesArraysParsingRequest],
			),
			validateValue: internal.NewStringTypesArraysParsingRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesNullableArrayItems struct {
	bindUnformattedStr requestParamBinder[string, []*string]
	bindCustomFormatStr requestParamBinder[string, []*string]
	bindDateStr requestParamBinder[string, []*time.Time]
	bindDateTimeStr requestParamBinder[string, []*time.Time]
	bindByteStr requestParamBinder[string, []*string]
	bindUnformattedStrInQuery requestParamBinder[[]string, []*string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, []*string]
	bindDateStrInQuery requestParamBinder[[]string, []*time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, []*time.Time]
	bindByteStrInQuery requestParamBinder[[]string, []*string]
	bindPayload requestParamBinder[*http.Request, *models.StringTypesNullableArrayItemsRequest]
}

func (p *paramsParserStringTypesStringTypesNullableArrayItems) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesNullableArrayItemsRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesNullableArrayItemsRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableArrayItems(app *HTTPApp) paramsParser[*StringTypesStringTypesNullableArrayItemsRequest] {
	return &paramsParserStringTypesStringTypesNullableArrayItems{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, []*string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*string](
				),
				internal.NewSimpleFieldValidator[*string](
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](10, true)),
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](20, false)),
					internal.SkipNullValidator(internal.NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, []*string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*string](
				),
				internal.NewSimpleFieldValidator[*string](
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](20, true)),
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](30, false)),
					internal.SkipNullValidator(internal.NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, []*time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(app.knownParsers.dateParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*time.Time](
				),
				internal.NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, []*time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(app.knownParsers.timeParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*time.Time](
				),
				internal.NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, []*string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*string](
				),
				internal.NewSimpleFieldValidator[*string](
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](30, true)),
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](40, false)),
				),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, []*string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*string](
				),
				internal.NewSimpleFieldValidator[*string](
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](10, true)),
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](20, false)),
					internal.SkipNullValidator(internal.NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, []*string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*string](
				),
				internal.NewSimpleFieldValidator[*string](
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](20, true)),
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](30, false)),
					internal.SkipNullValidator(internal.NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, []*time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(app.knownParsers.dateParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*time.Time](
				),
				internal.NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, []*time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(app.knownParsers.timeParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*time.Time](
				),
				internal.NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, []*string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewArrayValidator(
				internal.NewSimpleFieldValidator[[]*string](
				),
				internal.NewSimpleFieldValidator[*string](
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](30, true)),
					internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](40, false)),
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesNullableArrayItemsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesNullableArrayItemsRequest],
			),
			validateValue: internal.NewStringTypesNullableArrayItemsRequestValidator(),
		}),
	}
}

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
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableParsing(app *HTTPApp) paramsParser[*StringTypesStringTypesNullableParsingRequest] {
	return &paramsParserStringTypesStringTypesNullableParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, *string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, *string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, *time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.dateParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, *time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.timeParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, *string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, *time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.dateParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, *time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.timeParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesNullableParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesNullableParsingRequest],
			),
			validateValue: internal.NewStringTypesNullableParsingRequestValidator(),
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
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindOptionalUnformattedStrInQuery(queryParamsCtx.Fork("optionalUnformattedStrInQuery"), readQueryValue("optionalUnformattedStrInQuery", query), &reqParams.OptionalUnformattedStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableRequiredValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesNullableRequiredValidationRequest] {
	return &paramsParserStringTypesStringTypesNullableRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](10, true)),
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](100, false)),
				internal.SkipNullValidator(internal.NewPatternValidator[string](".*")),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesNullableRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesNullableRequiredValidationRequest],
			),
			validateValue: internal.NewStringTypesNullableRequiredValidationRequestValidator(),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.stringParser),
			),
			validateValue: internal.NewSimpleFieldValidator[*string](
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](10, true)),
				internal.SkipNullValidator(internal.NewMinMaxLengthValidator[string, string](100, false)),
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
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesParsing(app *HTTPApp) paramsParser[*StringTypesStringTypesParsingRequest] {
	return &paramsParserStringTypesStringTypesParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesParsingRequest],
			),
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
	bindPayload requestParamBinder[*http.Request, *models.StringTypesPatternValidationRequest]
}

func (p *paramsParserStringTypesStringTypesPatternValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesPatternValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &StringTypesStringTypesPatternValidationRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesPatternValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesPatternValidationRequest] {
	return &paramsParserStringTypesStringTypesPatternValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesPatternValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesPatternValidationRequest],
			),
			validateValue: internal.NewStringTypesPatternValidationRequestValidator(),
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
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRangeValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesRangeValidationRequest] {
	return &paramsParserStringTypesStringTypesRangeValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](10, true),
				internal.NewMinMaxLengthValidator[string, string](20, false),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](20, true),
				internal.NewMinMaxLengthValidator[string, string](30, false),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](30, true),
				internal.NewMinMaxLengthValidator[string, string](40, false),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](10, true),
				internal.NewMinMaxLengthValidator[string, string](20, false),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](20, true),
				internal.NewMinMaxLengthValidator[string, string](30, false),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](30, true),
				internal.NewMinMaxLengthValidator[string, string](40, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesRangeValidationRequest],
			),
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
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	p.bindOptionalUnformattedStrInQuery(queryParamsCtx.Fork("optionalUnformattedStrInQuery"), readQueryValue("optionalUnformattedStrInQuery", query), &reqParams.OptionalUnformattedStrInQuery)
	p.bindOptionalCustomFormatStrInQuery(queryParamsCtx.Fork("optionalCustomFormatStrInQuery"), readQueryValue("optionalCustomFormatStrInQuery", query), &reqParams.OptionalCustomFormatStrInQuery)
	p.bindOptionalDateStrInQuery(queryParamsCtx.Fork("optionalDateStrInQuery"), readQueryValue("optionalDateStrInQuery", query), &reqParams.OptionalDateStrInQuery)
	p.bindOptionalDateTimeStrInQuery(queryParamsCtx.Fork("optionalDateTimeStrInQuery"), readQueryValue("optionalDateTimeStrInQuery", query), &reqParams.OptionalDateTimeStrInQuery)
	p.bindOptionalByteStrInQuery(queryParamsCtx.Fork("optionalByteStrInQuery"), readQueryValue("optionalByteStrInQuery", query), &reqParams.OptionalByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRequiredValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesRequiredValidationRequest] {
	return &paramsParserStringTypesStringTypesRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](10, true),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](20, true),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](30, true),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.StringTypesRequiredValidationRequest],
			),
			validateValue: internal.NewStringTypesRequiredValidationRequestValidator(),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](10, true),
			),
		}),
		bindOptionalCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](20, true),
			),
		}),
		bindOptionalDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.dateParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindOptionalDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.timeParser,
			),
			validateValue: internal.NewSimpleFieldValidator[time.Time](
			),
		}),
		bindOptionalByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.NewMinMaxLengthValidator[string, string](30, true),
			),
		}),
	}
}
