package handlers

import (
	"net/http"
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type paramsParserErrorHandlingErrorHandlingParsingErrors struct {
	bindPathParam1 requestParamBinder[string, float32]
	bindPathParam2 requestParamBinder[string, float32]
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *paramsParserErrorHandlingErrorHandlingParsingErrors) parse(router httpRouter, req *http.Request) (*ErrorHandlingErrorHandlingParsingErrorsRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ErrorHandlingErrorHandlingParsingErrorsRequest{}
	// path params
	p.bindPathParam1(&bindingCtx, readPathValue("pathParam1", router, req), &reqParams.PathParam1)
	p.bindPathParam2(&bindingCtx, readPathValue("pathParam2", router, req), &reqParams.PathParam2)
	// query params
	query := req.URL.Query()
	p.bindRequiredQuery1(&bindingCtx, readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(&bindingCtx, readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserErrorHandlingErrorHandlingParsingErrors(app *HTTPApp) paramsParser[*ErrorHandlingErrorHandlingParsingErrorsRequest] {
	return &paramsParserErrorHandlingErrorHandlingParsingErrors{
		bindPathParam1: newRequestParamBinder(binderParams[string, float32]{
			field: "pathParam1",
			location: "path",
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "pathParam1", Location: "path"},
			),
		}),
		bindPathParam2: newRequestParamBinder(binderParams[string, float32]{
			field: "pathParam2",
			location: "path",
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "pathParam2", Location: "path"},
			),
		}),
		bindRequiredQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery1",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "requiredQuery1", Location: "query"},
			),
		}),
		bindRequiredQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery2",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "requiredQuery2", Location: "query"},
			),
		}),
	}
}

type paramsParserErrorHandlingErrorHandlingValidationErrors struct {
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *paramsParserErrorHandlingErrorHandlingValidationErrors) parse(router httpRouter, req *http.Request) (*ErrorHandlingErrorHandlingValidationErrorsRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ErrorHandlingErrorHandlingValidationErrorsRequest{}
	// query params
	query := req.URL.Query()
	p.bindRequiredQuery1(&bindingCtx, readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(&bindingCtx, readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserErrorHandlingErrorHandlingValidationErrors(app *HTTPApp) paramsParser[*ErrorHandlingErrorHandlingValidationErrorsRequest] {
	return &paramsParserErrorHandlingErrorHandlingValidationErrors{
		bindRequiredQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery1",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "requiredQuery1", Location: "query"},
				internal.NewMinMaxValueValidator[float32](10, false, true),
			),
		}),
		bindRequiredQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery2",
			location: "query",
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "requiredQuery2", Location: "query"},
				internal.NewMinMaxValueValidator[float32](10, false, true),
			),
		}),
	}
}
