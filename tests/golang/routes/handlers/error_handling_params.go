package handlers

import (
	"net/http"
)

type ErrorHandlingNumberRangeErrorsParamsParser struct {
	bindLimitedNum requestParamBinder[string, float32]
	bindLimitedFloat requestParamBinder[string, float32]
	bindLimitedDouble requestParamBinder[string, float64]
	bindLimitedQueryNum requestParamBinder[[]string, float32]
	bindLimitedQueryFloat requestParamBinder[[]string, float32]
	bindLimitedQueryDouble requestParamBinder[[]string, float64]
}

func (p *ErrorHandlingNumberRangeErrorsParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*ErrorHandlingNumberRangeErrorsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &ErrorHandlingNumberRangeErrorsRequest{}
	query := req.URL.Query()
	
	p.bindLimitedNum(&bindingCtx, readPathValue("limitedNum", router, req), &reqParams.LimitedNum)
	
	p.bindLimitedFloat(&bindingCtx, readPathValue("limitedFloat", router, req), &reqParams.LimitedFloat)
	
	p.bindLimitedDouble(&bindingCtx, readPathValue("limitedDouble", router, req), &reqParams.LimitedDouble)
	p.bindLimitedQueryNum(&bindingCtx, readQueryValue("limitedQueryNum", query), &reqParams.LimitedQueryNum)
	p.bindLimitedQueryFloat(&bindingCtx, readQueryValue("limitedQueryFloat", query), &reqParams.LimitedQueryFloat)
	p.bindLimitedQueryDouble(&bindingCtx, readQueryValue("limitedQueryDouble", query), &reqParams.LimitedQueryDouble)
	return reqParams, bindingCtx.AggregatedError()
}

func newErrorHandlingNumberRangeErrorsParamsParser() *ErrorHandlingNumberRangeErrorsParamsParser {
	return &ErrorHandlingNumberRangeErrorsParamsParser{
		bindLimitedNum: newRequestParamBinder(binderParams[string, float32]{
			field: "limitedNum",
			location: "path",
			parseValue: knownParsers.float32_in_path,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindLimitedFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "limitedFloat",
			location: "path",
			parseValue: knownParsers.float32_in_path,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindLimitedDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "limitedDouble",
			location: "path",
			parseValue: knownParsers.float64_in_path,
			validateValue: newCompositeValidator[string, float64](
				validateNonEmpty,
			),
		}),
		bindLimitedQueryNum: newRequestParamBinder(binderParams[[]string, float32]{
			field: "limitedQueryNum",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindLimitedQueryFloat: newRequestParamBinder(binderParams[[]string, float32]{
			field: "limitedQueryFloat",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindLimitedQueryDouble: newRequestParamBinder(binderParams[[]string, float64]{
			field: "limitedQueryDouble",
			location: "query",
			parseValue: knownParsers.float64_in_query,
			validateValue: newCompositeValidator[[]string, float64](
				validateNonEmpty,
			),
		}),
	}
}

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
