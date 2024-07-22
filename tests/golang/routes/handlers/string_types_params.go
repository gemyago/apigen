package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/internal"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type paramsParserStringTypesStringTypesParsing struct {
	bindUnformattedStr         requestParamBinder[string, string]
	bindCustomFormatStr        requestParamBinder[string, string]
	bindDateStr                requestParamBinder[string, time.Time]
	bindDateTimeStr            requestParamBinder[string, time.Time]
	bindByteStr                requestParamBinder[string, string]
	bindUnformattedStrInQuery  requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery         requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery     requestParamBinder[[]string, time.Time]
	bindByteStrInQuery         requestParamBinder[[]string, string]
	bindPayload                requestParamBinder[*http.Request, *models.StringTypesParsingRequest]
}

func (p *paramsParserStringTypesStringTypesParsing) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesParsingRequest, error) {
	bindingCtx := bindingContext{}
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
	p.bindPayload(&bindingCtx, internal.OptionalVal[*http.Request]{Value: req, Assigned: true}, &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesParsing(app *HTTPApp) paramsParser[*StringTypesStringTypesParsingRequest] {
	return &paramsParserStringTypesStringTypesParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field:      "unformattedStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field:      "customFormatStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field:      "dateStr",
			location:   "path",
			parseValue: app.knownParsers.dateInPath,
			validateValue: internal.NewCompositeValidator[string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field:      "dateTimeStr",
			location:   "path",
			parseValue: app.knownParsers.timeInPath,
			validateValue: internal.NewCompositeValidator[string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field:      "byteStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "unformattedStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "customFormatStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateTimeStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "byteStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.StringTypesParsingRequest]{
			field:      "payload",
			location:   "body",
			parseValue: parseJSONPayload[*models.StringTypesParsingRequest],
			validateValue: internal.NewCompositeValidator[*http.Request, *models.StringTypesParsingRequest](
				internal.ValidateNonEmpty,
				internal.NewModelParamValidator[*http.Request, *models.StringTypesParsingRequest](models.NewStringTypesParsingRequestValidator()),
			),
		}),
	}
}

type paramsParserStringTypesStringTypesPatternValidation struct {
	bindUnformattedStr         requestParamBinder[string, string]
	bindCustomFormatStr        requestParamBinder[string, string]
	bindDateStr                requestParamBinder[string, time.Time]
	bindDateTimeStr            requestParamBinder[string, time.Time]
	bindUnformattedStrInQuery  requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery         requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery     requestParamBinder[[]string, time.Time]
}

func (p *paramsParserStringTypesStringTypesPatternValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesPatternValidationRequest, error) {
	bindingCtx := bindingContext{}
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
			field:      "unformattedStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
				internal.NewPatternValidator[string, string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field:      "customFormatStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
				internal.NewPatternValidator[string, string]("^\\d{20}$"),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field:      "dateStr",
			location:   "path",
			parseValue: app.knownParsers.dateInPath,
			validateValue: internal.NewCompositeValidator[string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field:      "dateTimeStr",
			location:   "path",
			parseValue: app.knownParsers.timeInPath,
			validateValue: internal.NewCompositeValidator[string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "unformattedStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewPatternValidator[[]string, string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "customFormatStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewPatternValidator[[]string, string]("^\\d{20}$"),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateTimeStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
	}
}

type paramsParserStringTypesStringTypesRangeValidation struct {
	bindUnformattedStr         requestParamBinder[string, string]
	bindCustomFormatStr        requestParamBinder[string, string]
	bindDateStr                requestParamBinder[string, time.Time]
	bindDateTimeStr            requestParamBinder[string, time.Time]
	bindByteStr                requestParamBinder[string, string]
	bindUnformattedStrInQuery  requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery         requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery     requestParamBinder[[]string, time.Time]
	bindByteStrInQuery         requestParamBinder[[]string, string]
}

func (p *paramsParserStringTypesStringTypesRangeValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesRangeValidationRequest, error) {
	bindingCtx := bindingContext{}
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
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRangeValidation(app *HTTPApp) paramsParser[*StringTypesStringTypesRangeValidationRequest] {
	return &paramsParserStringTypesStringTypesRangeValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field:      "unformattedStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[string, string](10, true),
				internal.NewMinMaxLengthValidator[string, string](20, false),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field:      "customFormatStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[string, string](20, true),
				internal.NewMinMaxLengthValidator[string, string](30, false),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			field:      "dateStr",
			location:   "path",
			parseValue: app.knownParsers.dateInPath,
			validateValue: internal.NewCompositeValidator[string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			field:      "dateTimeStr",
			location:   "path",
			parseValue: app.knownParsers.timeInPath,
			validateValue: internal.NewCompositeValidator[string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field:      "byteStr",
			location:   "path",
			parseValue: app.knownParsers.stringInPath,
			validateValue: internal.NewCompositeValidator[string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[string, string](30, true),
				internal.NewMinMaxLengthValidator[string, string](40, false),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "unformattedStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[[]string, string](10, true),
				internal.NewMinMaxLengthValidator[[]string, string](20, false),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "customFormatStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[[]string, string](20, true),
				internal.NewMinMaxLengthValidator[[]string, string](30, false),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateTimeStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "byteStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[[]string, string](30, true),
				internal.NewMinMaxLengthValidator[[]string, string](40, false),
			),
		}),
	}
}

type paramsParserStringTypesStringTypesRequiredValidation struct {
	bindUnformattedStrInQuery          requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery         requestParamBinder[[]string, string]
	bindDateStrInQuery                 requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery             requestParamBinder[[]string, time.Time]
	bindByteStrInQuery                 requestParamBinder[[]string, string]
	bindOptionalUnformattedStrInQuery  requestParamBinder[[]string, string]
	bindOptionalCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindOptionalDateStrInQuery         requestParamBinder[[]string, time.Time]
	bindOptionalDateTimeStrInQuery     requestParamBinder[[]string, time.Time]
	bindOptionalByteStrInQuery         requestParamBinder[[]string, string]
}

func (p *paramsParserStringTypesStringTypesRequiredValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesRequiredValidationRequest, error) {
	bindingCtx := bindingContext{}
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
			field:      "unformattedStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[[]string, string](10, true),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "customFormatStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[[]string, string](20, true),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.dateInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:      "dateTimeStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.timeInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](
				internal.ValidateNonEmpty,
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "byteStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.ValidateNonEmpty,
				internal.NewMinMaxLengthValidator[[]string, string](30, true),
			),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "optionalUnformattedStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.NewMinMaxLengthValidator[[]string, string](10, true),
			),
		}),
		bindOptionalCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "optionalCustomFormatStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.NewMinMaxLengthValidator[[]string, string](20, true),
			),
		}),
		bindOptionalDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:         "optionalDateStrInQuery",
			location:      "query",
			parseValue:    app.knownParsers.dateInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](),
		}),
		bindOptionalDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			field:         "optionalDateTimeStrInQuery",
			location:      "query",
			parseValue:    app.knownParsers.timeInQuery,
			validateValue: internal.NewCompositeValidator[[]string, time.Time](),
		}),
		bindOptionalByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field:      "optionalByteStrInQuery",
			location:   "query",
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewCompositeValidator[[]string, string](
				internal.NewMinMaxLengthValidator[[]string, string](30, true),
			),
		}),
	}
}
