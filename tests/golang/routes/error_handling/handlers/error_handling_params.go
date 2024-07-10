package handlers

import (
	"net/http"
)

type ErrorHandlingParsingErrorsParamsParser struct {
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
}

func (p *ErrorHandlingParsingErrorsParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*ErrorHandlingParsingErrorsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &ErrorHandlingParsingErrorsRequest{}
	query := req.URL.Query()
	p.bindRequiredQuery1(&bindingCtx, readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(&bindingCtx, readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newErrorHandlingParsingErrorsParamsParser() *ErrorHandlingParsingErrorsParamsParser {
	return &ErrorHandlingParsingErrorsParamsParser{
		bindRequiredQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			field:      "requiredQuery1",
			location:   "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindRequiredQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			field:      "requiredQuery2",
			location:   "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
	}
}
