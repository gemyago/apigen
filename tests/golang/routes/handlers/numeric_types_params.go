package handlers

import (
	"net/http"
)

type NumericTypesNumberAnySimpleParamsParser struct {
	bindPathParam1 requestParamBinder[string, float32]
	bindPathParam2 requestParamBinder[string, float32]
	bindRequiredQuery1 requestParamBinder[[]string, float32]
	bindRequiredQuery2 requestParamBinder[[]string, float32]
	bindOptionalQuery1 requestParamBinder[[]string, float32]
	bindOptionalQuery2 requestParamBinder[[]string, float32]
}

func (p *NumericTypesNumberAnySimpleParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*NumericTypesNumberAnySimpleRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &NumericTypesNumberAnySimpleRequest{}
	query := req.URL.Query()
	
	p.bindPathParam1(&bindingCtx, readPathValue("pathParam1", router, req), &reqParams.PathParam1)
	
	p.bindPathParam2(&bindingCtx, readPathValue("pathParam2", router, req), &reqParams.PathParam2)
	p.bindRequiredQuery1(&bindingCtx, readQueryValue("requiredQuery1", query), &reqParams.RequiredQuery1)
	p.bindRequiredQuery2(&bindingCtx, readQueryValue("requiredQuery2", query), &reqParams.RequiredQuery2)
	p.bindOptionalQuery1(&bindingCtx, readQueryValue("optionalQuery1", query), &reqParams.OptionalQuery1)
	p.bindOptionalQuery2(&bindingCtx, readQueryValue("optionalQuery2", query), &reqParams.OptionalQuery2)
	return reqParams, bindingCtx.AggregatedError()
}

func newNumericTypesNumberAnySimpleParamsParser() *NumericTypesNumberAnySimpleParamsParser {
	return &NumericTypesNumberAnySimpleParamsParser{
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
		bindOptionalQuery1: newRequestParamBinder(binderParams[[]string, float32]{
			field: "optionalQuery1",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
			),
		}),
		bindOptionalQuery2: newRequestParamBinder(binderParams[[]string, float32]{
			field: "optionalQuery2",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
			),
		}),
	}
}
