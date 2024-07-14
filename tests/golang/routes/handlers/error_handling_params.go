package handlers

import (
	"net/http"
)

type paramsParserErrorHandlingErrorHandlingParsingErrors struct {
	bindPathParam1 requestParamBinder[string, float32]
	bindPathParam2 requestParamBinder[string, float32]
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *paramsParserErrorHandlingErrorHandlingParsingErrors) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*ErrorHandlingErrorHandlingParsingErrorsRequest, error) {
	bindingCtx := bindingContext{}
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

func newParamsParserErrorHandlingErrorHandlingParsingErrors() *paramsParserErrorHandlingErrorHandlingParsingErrors {
	return &paramsParserErrorHandlingErrorHandlingParsingErrors{
		bindPathParam1: newRequestParamBinder(binderParams[string, float32]{
			field: "pathParam1",
			location: "path",
			parseValue: knownParsers.float32_in_path,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindPathParam2: newRequestParamBinder(binderParams[string, float32]{
			field: "pathParam2",
			location: "path",
			parseValue: knownParsers.float32_in_path,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindRequiredQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery1",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindRequiredQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery2",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
	}
}

type paramsParserErrorHandlingErrorHandlingValidationErrors struct {
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *paramsParserErrorHandlingErrorHandlingValidationErrors) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*ErrorHandlingErrorHandlingValidationErrorsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &ErrorHandlingErrorHandlingValidationErrorsRequest{}
	// query params
	query := req.URL.Query()
	p.bindRequiredQuery1(&bindingCtx, readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(&bindingCtx, readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserErrorHandlingErrorHandlingValidationErrors() *paramsParserErrorHandlingErrorHandlingValidationErrors {
	return &paramsParserErrorHandlingErrorHandlingValidationErrors{
		bindRequiredQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery1",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float32](10, false, true),
			),
		}),
		bindRequiredQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			field: "requiredQuery2",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float32](10, false, true),
			),
		}),
	}
}
