package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.NumericTypesNullableRequest{}

type paramsParserNumericTypesNumericTypesNullable struct {
	bindNumberAny requestParamBinder[string, *float32]
	bindNumberFloat requestParamBinder[string, *float32]
	bindNumberDouble requestParamBinder[string, *float64]
	bindNumberInt requestParamBinder[string, *int32]
	bindNumberInt32 requestParamBinder[string, *int32]
	bindNumberInt64 requestParamBinder[string, *int64]
	bindNumberAnyInQuery requestParamBinder[[]string, *float32]
	bindNumberFloatInQuery requestParamBinder[[]string, *float32]
	bindNumberDoubleInQuery requestParamBinder[[]string, *float64]
	bindNumberIntInQuery requestParamBinder[[]string, *int32]
	bindNumberInt32InQuery requestParamBinder[[]string, *int32]
	bindNumberInt64InQuery requestParamBinder[[]string, *int64]
	bindPayload requestParamBinder[*http.Request, *models.NumericTypesNullableRequest]
	bindOptionalNumberAnyInQuery requestParamBinder[[]string, *float32]
}

func (p *paramsParserNumericTypesNumericTypesNullable) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesNullableRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &NumericTypesNumericTypesNullableRequest{}
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
	p.bindOptionalNumberAnyInQuery(&bindingCtx, readQueryValue("optionalNumberAnyInQuery", query), &reqParams.OptionalNumberAnyInQuery)
	p.bindNumberFloatInQuery(&bindingCtx, readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(&bindingCtx, readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(&bindingCtx, readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(&bindingCtx, readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(&bindingCtx, readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesNullable(app *HTTPApp) paramsParser[*NumericTypesNumericTypesNullableRequest] {
	return &paramsParserNumericTypesNumericTypesNullable{
    // isNullable: true
		bindNumberAny: newRequestParamBinder(binderParams[string, *float32]{
			field: "numberAny",
			location: "path",
			required: true,
			parseValue: parseNullableInPath(app.knownParsers.float32InPath),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](100.01, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
    // isNullable: true
		bindNumberFloat: newRequestParamBinder(binderParams[string, *float32]{
			field: "numberFloat",
			location: "path",
			required: true,
			parseValue: parseNullableInPath(app.knownParsers.float32InPath),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](300.03, false, false)),
			),
		}),
    // isNullable: true
		bindNumberDouble: newRequestParamBinder(binderParams[string, *float64]{
			field: "numberDouble",
			location: "path",
			required: true,
			parseValue: parseNullableInPath(app.knownParsers.float64InPath),
			validateValue: internal.NewSimpleFieldValidator[*float64](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](300.03, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](400.04, false, false)),
			),
		}),
    // isNullable: true
		bindNumberInt: newRequestParamBinder(binderParams[string, *int32]{
			field: "numberInt",
			location: "path",
			required: true,
			parseValue: parseNullableInPath(app.knownParsers.int32InPath),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](400, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, false)),
			),
		}),
    // isNullable: true
		bindNumberInt32: newRequestParamBinder(binderParams[string, *int32]{
			field: "numberInt32",
			location: "path",
			required: true,
			parseValue: parseNullableInPath(app.knownParsers.int32InPath),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](600, false, false)),
			),
		}),
    // isNullable: true
		bindNumberInt64: newRequestParamBinder(binderParams[string, *int64]{
			field: "numberInt64",
			location: "path",
			required: true,
			parseValue: parseNullableInPath(app.knownParsers.int64InPath),
			validateValue: internal.NewSimpleFieldValidator[*int64](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](600, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](700, false, false)),
			),
		}),
    // isNullable: true
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			field: "numberAnyInQuery",
			location: "query",
			required: true,
			parseValue: parseNullableInQuery(app.knownParsers.float32InQuery),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](100.01, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
    // isNullable: true
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			field: "numberFloatInQuery",
			location: "query",
			required: true,
			parseValue: parseNullableInQuery(app.knownParsers.float32InQuery),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](300.03, false, false)),
			),
		}),
    // isNullable: true
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, *float64]{
			field: "numberDoubleInQuery",
			location: "query",
			required: true,
			parseValue: parseNullableInQuery(app.knownParsers.float64InQuery),
			validateValue: internal.NewSimpleFieldValidator[*float64](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](300.03, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](400.04, false, false)),
			),
		}),
    // isNullable: true
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, *int32]{
			field: "numberIntInQuery",
			location: "query",
			required: true,
			parseValue: parseNullableInQuery(app.knownParsers.int32InQuery),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](400, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, false)),
			),
		}),
    // isNullable: true
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, *int32]{
			field: "numberInt32InQuery",
			location: "query",
			required: true,
			parseValue: parseNullableInQuery(app.knownParsers.int32InQuery),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](600, false, false)),
			),
		}),
    // isNullable: true
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, *int64]{
			field: "numberInt64InQuery",
			location: "query",
			required: true,
			parseValue: parseNullableInQuery(app.knownParsers.int64InQuery),
			validateValue: internal.NewSimpleFieldValidator[*int64](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](600, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](700, false, false)),
			),
		}),
    // isNullable: false
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesNullableRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.NumericTypesNullableRequest],
			validateValue: internal.NewNumericTypesNullableRequestValidator(),
		}),
    // isNullable: true
		bindOptionalNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			field: "optionalNumberAnyInQuery",
			location: "query",
			required: false,
			parseValue: parseNullableInQuery(app.knownParsers.float32InQuery),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](100.01, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
	}
}

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
	bindPayload requestParamBinder[*http.Request, *models.NumericTypesParsingRequest]
}

func (p *paramsParserNumericTypesNumericTypesParsing) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesParsingRequest, error) {
	bindingCtx := internal.BindingContext{}
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
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesParsing(app *HTTPApp) paramsParser[*NumericTypesNumericTypesParsingRequest] {
	return &paramsParserNumericTypesNumericTypesParsing{
    // isNullable: false
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			field: "numberAny",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float32InPath,
			validateValue: internal.NewSimpleFieldValidator[float32](
			),
		}),
    // isNullable: false
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "numberFloat",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float32InPath,
			validateValue: internal.NewSimpleFieldValidator[float32](
			),
		}),
    // isNullable: false
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "numberDouble",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float64InPath,
			validateValue: internal.NewSimpleFieldValidator[float64](
			),
		}),
    // isNullable: false
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int32InPath,
			validateValue: internal.NewSimpleFieldValidator[int32](
			),
		}),
    // isNullable: false
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt32",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int32InPath,
			validateValue: internal.NewSimpleFieldValidator[int32](
			),
		}),
    // isNullable: false
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			field: "numberInt64",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int64InPath,
			validateValue: internal.NewSimpleFieldValidator[int64](
			),
		}),
    // isNullable: false
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
			),
		}),
    // isNullable: false
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
			),
		}),
    // isNullable: false
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float64InQuery,
			validateValue: internal.NewSimpleFieldValidator[float64](
			),
		}),
    // isNullable: false
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
			),
		}),
    // isNullable: false
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
			),
		}),
    // isNullable: false
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int64InQuery,
			validateValue: internal.NewSimpleFieldValidator[int64](
			),
		}),
    // isNullable: false
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesParsingRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.NumericTypesParsingRequest],
			validateValue: internal.NewNumericTypesParsingRequestValidator(),
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
	bindPayload requestParamBinder[*http.Request, *models.NumericTypesRangeValidationRequest]
}

func (p *paramsParserNumericTypesNumericTypesRangeValidation) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRangeValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
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
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRangeValidation(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRangeValidationRequest] {
	return &paramsParserNumericTypesNumericTypesRangeValidation{
    // isNullable: false
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			field: "numberAny",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float32InPath,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](100.01, false, true),
				internal.NewMinMaxValueValidator[float32](200.02, false, false),
			),
		}),
    // isNullable: false
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "numberFloat",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float32InPath,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](200.02, false, true),
				internal.NewMinMaxValueValidator[float32](300.03, false, false),
			),
		}),
    // isNullable: false
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "numberDouble",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float64InPath,
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.NewMinMaxValueValidator[float64](300.03, false, true),
				internal.NewMinMaxValueValidator[float64](400.04, false, false),
			),
		}),
    // isNullable: false
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int32InPath,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](400, false, true),
				internal.NewMinMaxValueValidator[int32](500, false, false),
			),
		}),
    // isNullable: false
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt32",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int32InPath,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](500, false, true),
				internal.NewMinMaxValueValidator[int32](600, false, false),
			),
		}),
    // isNullable: false
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			field: "numberInt64",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int64InPath,
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.NewMinMaxValueValidator[int64](600, false, true),
				internal.NewMinMaxValueValidator[int64](700, false, false),
			),
		}),
    // isNullable: false
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](100.01, false, true),
				internal.NewMinMaxValueValidator[float32](200.02, false, false),
			),
		}),
    // isNullable: false
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](200.02, false, true),
				internal.NewMinMaxValueValidator[float32](300.03, false, false),
			),
		}),
    // isNullable: false
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float64InQuery,
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.NewMinMaxValueValidator[float64](300.03, false, true),
				internal.NewMinMaxValueValidator[float64](400.04, false, false),
			),
		}),
    // isNullable: false
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](400, false, true),
				internal.NewMinMaxValueValidator[int32](500, false, false),
			),
		}),
    // isNullable: false
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](500, false, true),
				internal.NewMinMaxValueValidator[int32](600, false, false),
			),
		}),
    // isNullable: false
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int64InQuery,
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.NewMinMaxValueValidator[int64](600, false, true),
				internal.NewMinMaxValueValidator[int64](700, false, false),
			),
		}),
    // isNullable: false
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesRangeValidationRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.NumericTypesRangeValidationRequest],
			validateValue: internal.NewNumericTypesRangeValidationRequestValidator(),
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
	bindPayload requestParamBinder[*http.Request, *models.NumericTypesRangeValidationExclusiveRequest]
}

func (p *paramsParserNumericTypesNumericTypesRangeValidationExclusive) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRangeValidationExclusiveRequest, error) {
	bindingCtx := internal.BindingContext{}
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
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRangeValidationExclusive(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRangeValidationExclusiveRequest] {
	return &paramsParserNumericTypesNumericTypesRangeValidationExclusive{
    // isNullable: false
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			field: "numberAny",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float32InPath,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](100.01, true, true),
				internal.NewMinMaxValueValidator[float32](200.02, true, false),
			),
		}),
    // isNullable: false
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			field: "numberFloat",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float32InPath,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](200.02, true, true),
				internal.NewMinMaxValueValidator[float32](300.03, true, false),
			),
		}),
    // isNullable: false
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			field: "numberDouble",
			location: "path",
			required: true,
			parseValue: app.knownParsers.float64InPath,
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.NewMinMaxValueValidator[float64](300.03, true, true),
				internal.NewMinMaxValueValidator[float64](400.04, true, false),
			),
		}),
    // isNullable: false
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int32InPath,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](400, true, true),
				internal.NewMinMaxValueValidator[int32](500, true, false),
			),
		}),
    // isNullable: false
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			field: "numberInt32",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int32InPath,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](500, true, true),
				internal.NewMinMaxValueValidator[int32](600, true, false),
			),
		}),
    // isNullable: false
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			field: "numberInt64",
			location: "path",
			required: true,
			parseValue: app.knownParsers.int64InPath,
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.NewMinMaxValueValidator[int64](600, true, true),
				internal.NewMinMaxValueValidator[int64](700, true, false),
			),
		}),
    // isNullable: false
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](100.01, true, true),
				internal.NewMinMaxValueValidator[float32](200.02, true, false),
			),
		}),
    // isNullable: false
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](200.02, true, true),
				internal.NewMinMaxValueValidator[float32](300.03, true, false),
			),
		}),
    // isNullable: false
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float64InQuery,
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.NewMinMaxValueValidator[float64](300.03, true, true),
				internal.NewMinMaxValueValidator[float64](400.04, true, false),
			),
		}),
    // isNullable: false
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](400, true, true),
				internal.NewMinMaxValueValidator[int32](500, true, false),
			),
		}),
    // isNullable: false
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](500, true, true),
				internal.NewMinMaxValueValidator[int32](600, true, false),
			),
		}),
    // isNullable: false
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int64InQuery,
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.NewMinMaxValueValidator[int64](600, true, true),
				internal.NewMinMaxValueValidator[int64](700, true, false),
			),
		}),
    // isNullable: false
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesRangeValidationExclusiveRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.NumericTypesRangeValidationExclusiveRequest],
			validateValue: internal.NewNumericTypesRangeValidationExclusiveRequestValidator(),
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
	bindingCtx := internal.BindingContext{}
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
    // isNullable: false
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberAnyInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
			),
		}),
    // isNullable: false
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "numberFloatInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
			),
		}),
    // isNullable: false
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "numberDoubleInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.float64InQuery,
			validateValue: internal.NewSimpleFieldValidator[float64](
			),
		}),
    // isNullable: false
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberIntInQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
			),
		}),
    // isNullable: false
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "numberInt32InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
			),
		}),
    // isNullable: false
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "numberInt64InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.int64InQuery,
			validateValue: internal.NewSimpleFieldValidator[int64](
			),
		}),
    // isNullable: false
		bindOptionalNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "optionalNumberAnyInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](100.01, false, true),
			),
		}),
    // isNullable: false
		bindOptionalNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			field: "optionalNumberFloatInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.float32InQuery,
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.NewMinMaxValueValidator[float32](200.02, false, true),
			),
		}),
    // isNullable: false
		bindOptionalNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			field: "optionalNumberDoubleInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.float64InQuery,
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.NewMinMaxValueValidator[float64](300.03, false, true),
			),
		}),
    // isNullable: false
		bindOptionalNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "optionalNumberIntInQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](400, false, true),
			),
		}),
    // isNullable: false
		bindOptionalNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			field: "optionalNumberInt32InQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.int32InQuery,
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.NewMinMaxValueValidator[int32](500, false, true),
			),
		}),
    // isNullable: false
		bindOptionalNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			field: "optionalNumberInt64InQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.int64InQuery,
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.NewMinMaxValueValidator[int64](600, false, true),
			),
		}),
	}
}
