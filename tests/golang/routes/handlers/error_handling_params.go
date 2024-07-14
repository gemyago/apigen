package handlers

import (
	"net/http"
)

type ErrorHandlingParsingErrorsParamsParser struct {
	bindPathParam1 requestParamBinder[string, float32]
	bindPathParam2 requestParamBinder[string, float32]
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *ErrorHandlingParsingErrorsParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*ErrorHandlingParsingErrorsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &ErrorHandlingParsingErrorsRequest{}
	query := req.URL.Query()
	
	p.bindPathParam1(&bindingCtx, readPathValue("pathParam1", router, req), &reqParams.PathParam1)
	
	p.bindPathParam2(&bindingCtx, readPathValue("pathParam2", router, req), &reqParams.PathParam2)
	p.bindRequiredQuery1(&bindingCtx, readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(&bindingCtx, readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newErrorHandlingParsingErrorsParamsParser() *ErrorHandlingParsingErrorsParamsParser {
	return &ErrorHandlingParsingErrorsParamsParser{
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

type ErrorHandlingValidationErrorsParamsParser struct {
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *ErrorHandlingValidationErrorsParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*ErrorHandlingValidationErrorsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &ErrorHandlingValidationErrorsRequest{}
	query := req.URL.Query()
	p.bindRequiredQuery1(&bindingCtx, readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(&bindingCtx, readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newErrorHandlingValidationErrorsParamsParser() *ErrorHandlingValidationErrorsParamsParser {
	return &ErrorHandlingValidationErrorsParamsParser{
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
