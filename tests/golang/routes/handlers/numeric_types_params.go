package handlers

import (
	"net/http"
)

type NumericTypesNumericTypesParsingParamsParser struct {
	bindNumberAny requestParamBinder[string, float32]
	bindNumberFloat requestParamBinder[string, float32]
	bindNumberDouble requestParamBinder[string, float64]
	bindNumberInt requestParamBinder[string, int32]
	bindNumberInt32 requestParamBinder[string, int32]
	bindNumberInt64 requestParamBinder[string, int64]
	bindNumberAnyInQuery requestParamBinder[[]string, float32]
	bindNumberFloatInQuery requestParamBinder[[]string, float32]
	bindNumberDoubleInQuery requestParamBinder[[]string, float64]
	bindNumberIntInQuery requestParamBinder[[]string, int32]
	bindNumberInt32InQuery requestParamBinder[[]string, int32]
	bindNumberInt64InQuery requestParamBinder[[]string, int64]
}

func (p *NumericTypesNumericTypesParsingParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*NumericTypesNumericTypesParsingRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &NumericTypesNumericTypesParsingRequest{}
	query := req.URL.Query()
	
	p.bindNumberAny(&bindingCtx, readPathValue("numberAny", router, req), &reqParams.NumberAny)
	
	p.bindNumberFloat(&bindingCtx, readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	
	p.bindNumberDouble(&bindingCtx, readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	
	p.bindNumberInt(&bindingCtx, readPathValue("numberInt", router, req), &reqParams.NumberInt)
	
	p.bindNumberInt32(&bindingCtx, readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	
	p.bindNumberInt64(&bindingCtx, readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	p.bindNumberAnyInQuery(&bindingCtx, readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(&bindingCtx, readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(&bindingCtx, readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(&bindingCtx, readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(&bindingCtx, readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(&bindingCtx, readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newNumericTypesNumericTypesParsingParamsParser() *NumericTypesNumericTypesParsingParamsParser {
	return &NumericTypesNumericTypesParsingParamsParser{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			field: "numberAny",
			location: "path",
			parseValue: knownParsers.float32_in_path,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "numberFloat",
			location: "path",
			parseValue: knownParsers.float32_in_path,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "numberDouble",
			location: "path",
			parseValue: knownParsers.float64_in_path,
			validateValue: newCompositeValidator[string, float64](
				validateNonEmpty,
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt",
			location: "path",
			parseValue: knownParsers.int32_in_path,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt32",
			location: "path",
			parseValue: knownParsers.int32_in_path,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			field: "numberInt64",
			location: "path",
			parseValue: knownParsers.int64_in_path,
			validateValue: newCompositeValidator[string, int64](
				validateNonEmpty,
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			parseValue: knownParsers.float32_in_query,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			parseValue: knownParsers.float64_in_query,
			validateValue: newCompositeValidator[[]string, float64](
				validateNonEmpty,
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			parseValue: knownParsers.int32_in_query,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			parseValue: knownParsers.int32_in_query,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			parseValue: knownParsers.int64_in_query,
			validateValue: newCompositeValidator[[]string, int64](
				validateNonEmpty,
			),
		}),
	}
}
