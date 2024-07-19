package handlers

import (
	"net/http"
	"time"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type paramsParserNumericTypesNumericTypesParsing struct {
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

func (p *paramsParserNumericTypesNumericTypesParsing) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesParsingRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &NumericTypesNumericTypesParsingRequest{}
	// path params
	p.bindNumberAny(&bindingCtx, readPathValue("numberAny", router, req), &reqParams.NumberAny)
	p.bindNumberFloat(&bindingCtx, readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	p.bindNumberDouble(&bindingCtx, readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	p.bindNumberInt(&bindingCtx, readPathValue("numberInt", router, req), &reqParams.NumberInt)
	p.bindNumberInt32(&bindingCtx, readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	p.bindNumberInt64(&bindingCtx, readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	// query params
	query := req.URL.Query()
	p.bindNumberAnyInQuery(&bindingCtx, readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(&bindingCtx, readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(&bindingCtx, readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(&bindingCtx, readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(&bindingCtx, readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(&bindingCtx, readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesParsing(app *HTTPApp) paramsParser[*NumericTypesNumericTypesParsingRequest] {
	return &paramsParserNumericTypesNumericTypesParsing{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			field: "numberAny",
			location: "path",
			parseValue: app.knownParsers.float32InPath,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "numberFloat",
			location: "path",
			parseValue: app.knownParsers.float32InPath,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "numberDouble",
			location: "path",
			parseValue: app.knownParsers.float64InPath,
			validateValue: newCompositeValidator[string, float64](
				validateNonEmpty,
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt",
			location: "path",
			parseValue: app.knownParsers.int32InPath,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt32",
			location: "path",
			parseValue: app.knownParsers.int32InPath,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			field: "numberInt64",
			location: "path",
			parseValue: app.knownParsers.int64InPath,
			validateValue: newCompositeValidator[string, int64](
				validateNonEmpty,
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			parseValue: app.knownParsers.float64InQuery,
			validateValue: newCompositeValidator[[]string, float64](
				validateNonEmpty,
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			parseValue: app.knownParsers.int64InQuery,
			validateValue: newCompositeValidator[[]string, int64](
				validateNonEmpty,
			),
		}),
	}
}

type paramsParserNumericTypesNumericTypesRangeValidation struct {
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

func (p *paramsParserNumericTypesNumericTypesRangeValidation) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRangeValidationRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &NumericTypesNumericTypesRangeValidationRequest{}
	// path params
	p.bindNumberAny(&bindingCtx, readPathValue("numberAny", router, req), &reqParams.NumberAny)
	p.bindNumberFloat(&bindingCtx, readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	p.bindNumberDouble(&bindingCtx, readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	p.bindNumberInt(&bindingCtx, readPathValue("numberInt", router, req), &reqParams.NumberInt)
	p.bindNumberInt32(&bindingCtx, readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	p.bindNumberInt64(&bindingCtx, readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	// query params
	query := req.URL.Query()
	p.bindNumberAnyInQuery(&bindingCtx, readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(&bindingCtx, readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(&bindingCtx, readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(&bindingCtx, readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(&bindingCtx, readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(&bindingCtx, readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRangeValidation(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRangeValidationRequest] {
	return &paramsParserNumericTypesNumericTypesRangeValidation{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			field: "numberAny",
			location: "path",
			parseValue: app.knownParsers.float32InPath,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[string, float32](100.01, false, true),
				newMinMaxValueValidator[string, float32](200.02, false, false),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "numberFloat",
			location: "path",
			parseValue: app.knownParsers.float32InPath,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[string, float32](200.02, false, true),
				newMinMaxValueValidator[string, float32](300.03, false, false),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "numberDouble",
			location: "path",
			parseValue: app.knownParsers.float64InPath,
			validateValue: newCompositeValidator[string, float64](
				validateNonEmpty,
				newMinMaxValueValidator[string, float64](300.03, false, true),
				newMinMaxValueValidator[string, float64](400.04, false, false),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt",
			location: "path",
			parseValue: app.knownParsers.int32InPath,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[string, int32](400, false, true),
				newMinMaxValueValidator[string, int32](500, false, false),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt32",
			location: "path",
			parseValue: app.knownParsers.int32InPath,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[string, int32](500, false, true),
				newMinMaxValueValidator[string, int32](600, false, false),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			field: "numberInt64",
			location: "path",
			parseValue: app.knownParsers.int64InPath,
			validateValue: newCompositeValidator[string, int64](
				validateNonEmpty,
				newMinMaxValueValidator[string, int64](600, false, true),
				newMinMaxValueValidator[string, int64](700, false, false),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float32](100.01, false, true),
				newMinMaxValueValidator[[]string, float32](200.02, false, false),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float32](200.02, false, true),
				newMinMaxValueValidator[[]string, float32](300.03, false, false),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			parseValue: app.knownParsers.float64InQuery,
			validateValue: newCompositeValidator[[]string, float64](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float64](300.03, false, true),
				newMinMaxValueValidator[[]string, float64](400.04, false, false),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int32](400, false, true),
				newMinMaxValueValidator[[]string, int32](500, false, false),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int32](500, false, true),
				newMinMaxValueValidator[[]string, int32](600, false, false),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			parseValue: app.knownParsers.int64InQuery,
			validateValue: newCompositeValidator[[]string, int64](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int64](600, false, true),
				newMinMaxValueValidator[[]string, int64](700, false, false),
			),
		}),
	}
}

type paramsParserNumericTypesNumericTypesRangeValidationExclusive struct {
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

func (p *paramsParserNumericTypesNumericTypesRangeValidationExclusive) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRangeValidationExclusiveRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &NumericTypesNumericTypesRangeValidationExclusiveRequest{}
	// path params
	p.bindNumberAny(&bindingCtx, readPathValue("numberAny", router, req), &reqParams.NumberAny)
	p.bindNumberFloat(&bindingCtx, readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	p.bindNumberDouble(&bindingCtx, readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	p.bindNumberInt(&bindingCtx, readPathValue("numberInt", router, req), &reqParams.NumberInt)
	p.bindNumberInt32(&bindingCtx, readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	p.bindNumberInt64(&bindingCtx, readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	// query params
	query := req.URL.Query()
	p.bindNumberAnyInQuery(&bindingCtx, readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(&bindingCtx, readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(&bindingCtx, readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(&bindingCtx, readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(&bindingCtx, readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(&bindingCtx, readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRangeValidationExclusive(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRangeValidationExclusiveRequest] {
	return &paramsParserNumericTypesNumericTypesRangeValidationExclusive{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			field: "numberAny",
			location: "path",
			parseValue: app.knownParsers.float32InPath,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[string, float32](100.01, true, true),
				newMinMaxValueValidator[string, float32](200.02, true, false),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "numberFloat",
			location: "path",
			parseValue: app.knownParsers.float32InPath,
			validateValue: newCompositeValidator[string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[string, float32](200.02, true, true),
				newMinMaxValueValidator[string, float32](300.03, true, false),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "numberDouble",
			location: "path",
			parseValue: app.knownParsers.float64InPath,
			validateValue: newCompositeValidator[string, float64](
				validateNonEmpty,
				newMinMaxValueValidator[string, float64](300.03, true, true),
				newMinMaxValueValidator[string, float64](400.04, true, false),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt",
			location: "path",
			parseValue: app.knownParsers.int32InPath,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[string, int32](400, true, true),
				newMinMaxValueValidator[string, int32](500, true, false),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt32",
			location: "path",
			parseValue: app.knownParsers.int32InPath,
			validateValue: newCompositeValidator[string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[string, int32](500, true, true),
				newMinMaxValueValidator[string, int32](600, true, false),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			field: "numberInt64",
			location: "path",
			parseValue: app.knownParsers.int64InPath,
			validateValue: newCompositeValidator[string, int64](
				validateNonEmpty,
				newMinMaxValueValidator[string, int64](600, true, true),
				newMinMaxValueValidator[string, int64](700, true, false),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float32](100.01, true, true),
				newMinMaxValueValidator[[]string, float32](200.02, true, false),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float32](200.02, true, true),
				newMinMaxValueValidator[[]string, float32](300.03, true, false),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			parseValue: app.knownParsers.float64InQuery,
			validateValue: newCompositeValidator[[]string, float64](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, float64](300.03, true, true),
				newMinMaxValueValidator[[]string, float64](400.04, true, false),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int32](400, true, true),
				newMinMaxValueValidator[[]string, int32](500, true, false),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int32](500, true, true),
				newMinMaxValueValidator[[]string, int32](600, true, false),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			parseValue: app.knownParsers.int64InQuery,
			validateValue: newCompositeValidator[[]string, int64](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int64](600, true, true),
				newMinMaxValueValidator[[]string, int64](700, true, false),
			),
		}),
	}
}

type paramsParserNumericTypesNumericTypesRequiredValidation struct {
	bindNumberAnyInQuery requestParamBinder[[]string, float32]
	bindNumberFloatInQuery requestParamBinder[[]string, float32]
	bindNumberDoubleInQuery requestParamBinder[[]string, float64]
	bindNumberIntInQuery requestParamBinder[[]string, int32]
	bindNumberInt32InQuery requestParamBinder[[]string, int32]
	bindNumberInt64InQuery requestParamBinder[[]string, int64]
	bindOptionalNumberAnyInQuery requestParamBinder[[]string, float32]
	bindOptionalNumberFloatInQuery requestParamBinder[[]string, float32]
	bindOptionalNumberDoubleInQuery requestParamBinder[[]string, float64]
	bindOptionalNumberIntInQuery requestParamBinder[[]string, int32]
	bindOptionalNumberInt32InQuery requestParamBinder[[]string, int32]
	bindOptionalNumberInt64InQuery requestParamBinder[[]string, int64]
}

func (p *paramsParserNumericTypesNumericTypesRequiredValidation) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRequiredValidationRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &NumericTypesNumericTypesRequiredValidationRequest{}
	// query params
	query := req.URL.Query()
	p.bindNumberAnyInQuery(&bindingCtx, readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(&bindingCtx, readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(&bindingCtx, readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(&bindingCtx, readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(&bindingCtx, readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(&bindingCtx, readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	p.bindOptionalNumberAnyInQuery(&bindingCtx, readQueryValue("optionalNumberAnyInQuery", query), &reqParams.OptionalNumberAnyInQuery)
	p.bindOptionalNumberFloatInQuery(&bindingCtx, readQueryValue("optionalNumberFloatInQuery", query), &reqParams.OptionalNumberFloatInQuery)
	p.bindOptionalNumberDoubleInQuery(&bindingCtx, readQueryValue("optionalNumberDoubleInQuery", query), &reqParams.OptionalNumberDoubleInQuery)
	p.bindOptionalNumberIntInQuery(&bindingCtx, readQueryValue("optionalNumberIntInQuery", query), &reqParams.OptionalNumberIntInQuery)
	p.bindOptionalNumberInt32InQuery(&bindingCtx, readQueryValue("optionalNumberInt32InQuery", query), &reqParams.OptionalNumberInt32InQuery)
	p.bindOptionalNumberInt64InQuery(&bindingCtx, readQueryValue("optionalNumberInt64InQuery", query), &reqParams.OptionalNumberInt64InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRequiredValidation(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRequiredValidationRequest] {
	return &paramsParserNumericTypesNumericTypesRequiredValidation{
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				validateNonEmpty,
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			parseValue: app.knownParsers.float64InQuery,
			validateValue: newCompositeValidator[[]string, float64](
				validateNonEmpty,
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				validateNonEmpty,
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			parseValue: app.knownParsers.int64InQuery,
			validateValue: newCompositeValidator[[]string, int64](
				validateNonEmpty,
			),
		}),
		bindOptionalNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "optionalNumberAnyInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				newMinMaxValueValidator[[]string, float32](100.01, false, true),
			),
		}),
		bindOptionalNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "optionalNumberFloatInQuery",
			location: "query",
			parseValue: app.knownParsers.float32InQuery,
			validateValue: newCompositeValidator[[]string, float32](
				newMinMaxValueValidator[[]string, float32](200.02, false, true),
			),
		}),
		bindOptionalNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "optionalNumberDoubleInQuery",
			location: "query",
			parseValue: app.knownParsers.float64InQuery,
			validateValue: newCompositeValidator[[]string, float64](
				newMinMaxValueValidator[[]string, float64](300.03, false, true),
			),
		}),
		bindOptionalNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "optionalNumberIntInQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				newMinMaxValueValidator[[]string, int32](400, false, true),
			),
		}),
		bindOptionalNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "optionalNumberInt32InQuery",
			location: "query",
			parseValue: app.knownParsers.int32InQuery,
			validateValue: newCompositeValidator[[]string, int32](
				newMinMaxValueValidator[[]string, int32](500, false, true),
			),
		}),
		bindOptionalNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "optionalNumberInt64InQuery",
			location: "query",
			parseValue: app.knownParsers.int64InQuery,
			validateValue: newCompositeValidator[[]string, int64](
				newMinMaxValueValidator[[]string, int64](600, false, true),
			),
		}),
	}
}
