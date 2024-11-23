package handlers

import (
	"net/http"
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/internal"
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
	bindingCtx := BindingContext{}
	reqParams := &ErrorHandlingErrorHandlingParsingErrorsRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindPathParam1(pathParamsCtx.Fork("pathParam1"), readPathValue("pathParam1", router, req), &reqParams.PathParam1)
	p.bindPathParam2(pathParamsCtx.Fork("pathParam2"), readPathValue("pathParam2", router, req), &reqParams.PathParam2)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindRequiredQuery1(queryParamsCtx.Fork("requiredQuery1"), readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(queryParamsCtx.Fork("requiredQuery2"), readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserErrorHandlingErrorHandlingParsingErrors(app *HTTPApp) paramsParser[*ErrorHandlingErrorHandlingParsingErrorsRequest] {
	return &paramsParserErrorHandlingErrorHandlingParsingErrors{
		bindPathParam1: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindPathParam2: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindRequiredQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindRequiredQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
	}
}

type paramsParserErrorHandlingErrorHandlingValidationErrors struct {
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *paramsParserErrorHandlingErrorHandlingValidationErrors) parse(router httpRouter, req *http.Request) (*ErrorHandlingErrorHandlingValidationErrorsRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ErrorHandlingErrorHandlingValidationErrorsRequest{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindRequiredQuery1(queryParamsCtx.Fork("requiredQuery1"), readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(queryParamsCtx.Fork("requiredQuery2"), readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserErrorHandlingErrorHandlingValidationErrors(app *HTTPApp) paramsParser[*ErrorHandlingErrorHandlingValidationErrorsRequest] {
	return &paramsParserErrorHandlingErrorHandlingValidationErrors{
		bindRequiredQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](10, false, true),
			),
		}),
		bindRequiredQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](10, false, true),
			),
		}),
	}
}
