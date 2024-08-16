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
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindNumberAny(pathParamsCtx.Fork("numberAny"), readPathValue("numberAny", router, req), &reqParams.NumberAny)
	p.bindNumberFloat(pathParamsCtx.Fork("numberFloat"), readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	p.bindNumberDouble(pathParamsCtx.Fork("numberDouble"), readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	p.bindNumberInt(pathParamsCtx.Fork("numberInt"), readPathValue("numberInt", router, req), &reqParams.NumberInt)
	p.bindNumberInt32(pathParamsCtx.Fork("numberInt32"), readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	p.bindNumberInt64(pathParamsCtx.Fork("numberInt64"), readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindNumberAnyInQuery(queryParamsCtx.Fork("numberAnyInQuery"), readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindOptionalNumberAnyInQuery(queryParamsCtx.Fork("optionalNumberAnyInQuery"), readQueryValue("optionalNumberAnyInQuery", query), &reqParams.OptionalNumberAnyInQuery)
	p.bindNumberFloatInQuery(queryParamsCtx.Fork("numberFloatInQuery"), readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(queryParamsCtx.Fork("numberDoubleInQuery"), readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(queryParamsCtx.Fork("numberIntInQuery"), readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(queryParamsCtx.Fork("numberInt32InQuery"), readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(queryParamsCtx.Fork("numberInt64InQuery"), readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesNullable(app *HTTPApp) paramsParser[*NumericTypesNumericTypesNullableRequest] {
	return &paramsParserNumericTypesNumericTypesNullable{
		bindNumberAny: newRequestParamBinder(binderParams[string, *float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.float32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SimpleFieldValidatorParams{Field: "numberAny", Location: "path"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](100.01, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, *float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.float32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloat", Location: "path"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](300.03, false, false)),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, *float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.float64Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*float64](
				internal.SimpleFieldValidatorParams{Field: "numberDouble", Location: "path"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](300.03, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](400.04, false, false)),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, *int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.int32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt", Location: "path"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](400, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, false)),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, *int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.int32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32", Location: "path"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](600, false, false)),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, *int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.int64Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64", Location: "path"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](600, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](700, false, false)),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.float32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SimpleFieldValidatorParams{Field: "numberAnyInQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](100.01, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.float32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloatInQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](200.02, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float32](300.03, false, false)),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, *float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.float64Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*float64](
				internal.SimpleFieldValidatorParams{Field: "numberDoubleInQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](300.03, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[float64](400.04, false, false)),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, *int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.int32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SimpleFieldValidatorParams{Field: "numberIntInQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](400, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, false)),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, *int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.int32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32InQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](500, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int32](600, false, false)),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, *int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.int64Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64InQuery", Location: "query"},
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](600, false, true)),
				internal.SkipNullValidator(internal.NewMinMaxValueValidator[int64](700, false, false)),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesNullableRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.NumericTypesNullableRequest],
			),
			validateValue: internal.NewNumericTypesNullableRequestValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
		bindOptionalNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(app.knownParsers.float32Parser),
			),
			validateValue: internal.NewSimpleFieldValidator[*float32](
				internal.SimpleFieldValidatorParams{Field: "optionalNumberAnyInQuery", Location: "query"},
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
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindNumberAny(pathParamsCtx.Fork("numberAny"), readPathValue("numberAny", router, req), &reqParams.NumberAny)
	p.bindNumberFloat(pathParamsCtx.Fork("numberFloat"), readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	p.bindNumberDouble(pathParamsCtx.Fork("numberDouble"), readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	p.bindNumberInt(pathParamsCtx.Fork("numberInt"), readPathValue("numberInt", router, req), &reqParams.NumberInt)
	p.bindNumberInt32(pathParamsCtx.Fork("numberInt32"), readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	p.bindNumberInt64(pathParamsCtx.Fork("numberInt64"), readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindNumberAnyInQuery(queryParamsCtx.Fork("numberAnyInQuery"), readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(queryParamsCtx.Fork("numberFloatInQuery"), readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(queryParamsCtx.Fork("numberDoubleInQuery"), readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(queryParamsCtx.Fork("numberIntInQuery"), readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(queryParamsCtx.Fork("numberInt32InQuery"), readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(queryParamsCtx.Fork("numberInt64InQuery"), readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesParsing(app *HTTPApp) paramsParser[*NumericTypesNumericTypesParsingRequest] {
	return &paramsParserNumericTypesNumericTypesParsing{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberAny", Location: "path"},
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloat", Location: "path"},
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "numberDouble", Location: "path"},
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt", Location: "path"},
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32", Location: "path"},
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64", Location: "path"},
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberAnyInQuery", Location: "query"},
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloatInQuery", Location: "query"},
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "numberDoubleInQuery", Location: "query"},
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberIntInQuery", Location: "query"},
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32InQuery", Location: "query"},
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64InQuery", Location: "query"},
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.NumericTypesParsingRequest],
			),
			validateValue: internal.NewNumericTypesParsingRequestValidator(internal.ModelValidatorParams{Location: "body"}),
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
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindNumberAny(pathParamsCtx.Fork("numberAny"), readPathValue("numberAny", router, req), &reqParams.NumberAny)
	p.bindNumberFloat(pathParamsCtx.Fork("numberFloat"), readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	p.bindNumberDouble(pathParamsCtx.Fork("numberDouble"), readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	p.bindNumberInt(pathParamsCtx.Fork("numberInt"), readPathValue("numberInt", router, req), &reqParams.NumberInt)
	p.bindNumberInt32(pathParamsCtx.Fork("numberInt32"), readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	p.bindNumberInt64(pathParamsCtx.Fork("numberInt64"), readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindNumberAnyInQuery(queryParamsCtx.Fork("numberAnyInQuery"), readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(queryParamsCtx.Fork("numberFloatInQuery"), readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(queryParamsCtx.Fork("numberDoubleInQuery"), readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(queryParamsCtx.Fork("numberIntInQuery"), readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(queryParamsCtx.Fork("numberInt32InQuery"), readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(queryParamsCtx.Fork("numberInt64InQuery"), readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRangeValidation(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRangeValidationRequest] {
	return &paramsParserNumericTypesNumericTypesRangeValidation{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberAny", Location: "path"},
				internal.NewMinMaxValueValidator[float32](100.01, false, true),
				internal.NewMinMaxValueValidator[float32](200.02, false, false),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloat", Location: "path"},
				internal.NewMinMaxValueValidator[float32](200.02, false, true),
				internal.NewMinMaxValueValidator[float32](300.03, false, false),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "numberDouble", Location: "path"},
				internal.NewMinMaxValueValidator[float64](300.03, false, true),
				internal.NewMinMaxValueValidator[float64](400.04, false, false),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt", Location: "path"},
				internal.NewMinMaxValueValidator[int32](400, false, true),
				internal.NewMinMaxValueValidator[int32](500, false, false),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32", Location: "path"},
				internal.NewMinMaxValueValidator[int32](500, false, true),
				internal.NewMinMaxValueValidator[int32](600, false, false),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64", Location: "path"},
				internal.NewMinMaxValueValidator[int64](600, false, true),
				internal.NewMinMaxValueValidator[int64](700, false, false),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberAnyInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float32](100.01, false, true),
				internal.NewMinMaxValueValidator[float32](200.02, false, false),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloatInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float32](200.02, false, true),
				internal.NewMinMaxValueValidator[float32](300.03, false, false),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "numberDoubleInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float64](300.03, false, true),
				internal.NewMinMaxValueValidator[float64](400.04, false, false),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberIntInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int32](400, false, true),
				internal.NewMinMaxValueValidator[int32](500, false, false),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32InQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int32](500, false, true),
				internal.NewMinMaxValueValidator[int32](600, false, false),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64InQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int64](600, false, true),
				internal.NewMinMaxValueValidator[int64](700, false, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.NumericTypesRangeValidationRequest],
			),
			validateValue: internal.NewNumericTypesRangeValidationRequestValidator(internal.ModelValidatorParams{Location: "body"}),
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
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindNumberAny(pathParamsCtx.Fork("numberAny"), readPathValue("numberAny", router, req), &reqParams.NumberAny)
	p.bindNumberFloat(pathParamsCtx.Fork("numberFloat"), readPathValue("numberFloat", router, req), &reqParams.NumberFloat)
	p.bindNumberDouble(pathParamsCtx.Fork("numberDouble"), readPathValue("numberDouble", router, req), &reqParams.NumberDouble)
	p.bindNumberInt(pathParamsCtx.Fork("numberInt"), readPathValue("numberInt", router, req), &reqParams.NumberInt)
	p.bindNumberInt32(pathParamsCtx.Fork("numberInt32"), readPathValue("numberInt32", router, req), &reqParams.NumberInt32)
	p.bindNumberInt64(pathParamsCtx.Fork("numberInt64"), readPathValue("numberInt64", router, req), &reqParams.NumberInt64)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindNumberAnyInQuery(queryParamsCtx.Fork("numberAnyInQuery"), readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(queryParamsCtx.Fork("numberFloatInQuery"), readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(queryParamsCtx.Fork("numberDoubleInQuery"), readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(queryParamsCtx.Fork("numberIntInQuery"), readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(queryParamsCtx.Fork("numberInt32InQuery"), readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(queryParamsCtx.Fork("numberInt64InQuery"), readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRangeValidationExclusive(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRangeValidationExclusiveRequest] {
	return &paramsParserNumericTypesNumericTypesRangeValidationExclusive{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberAny", Location: "path"},
				internal.NewMinMaxValueValidator[float32](100.01, true, true),
				internal.NewMinMaxValueValidator[float32](200.02, true, false),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloat", Location: "path"},
				internal.NewMinMaxValueValidator[float32](200.02, true, true),
				internal.NewMinMaxValueValidator[float32](300.03, true, false),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "numberDouble", Location: "path"},
				internal.NewMinMaxValueValidator[float64](300.03, true, true),
				internal.NewMinMaxValueValidator[float64](400.04, true, false),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt", Location: "path"},
				internal.NewMinMaxValueValidator[int32](400, true, true),
				internal.NewMinMaxValueValidator[int32](500, true, false),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32", Location: "path"},
				internal.NewMinMaxValueValidator[int32](500, true, true),
				internal.NewMinMaxValueValidator[int32](600, true, false),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64", Location: "path"},
				internal.NewMinMaxValueValidator[int64](600, true, true),
				internal.NewMinMaxValueValidator[int64](700, true, false),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberAnyInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float32](100.01, true, true),
				internal.NewMinMaxValueValidator[float32](200.02, true, false),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloatInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float32](200.02, true, true),
				internal.NewMinMaxValueValidator[float32](300.03, true, false),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "numberDoubleInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float64](300.03, true, true),
				internal.NewMinMaxValueValidator[float64](400.04, true, false),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberIntInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int32](400, true, true),
				internal.NewMinMaxValueValidator[int32](500, true, false),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32InQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int32](500, true, true),
				internal.NewMinMaxValueValidator[int32](600, true, false),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64InQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int64](600, true, true),
				internal.NewMinMaxValueValidator[int64](700, true, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.NumericTypesRangeValidationExclusiveRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.NumericTypesRangeValidationExclusiveRequest],
			),
			validateValue: internal.NewNumericTypesRangeValidationExclusiveRequestValidator(internal.ModelValidatorParams{Location: "body"}),
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
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindNumberAnyInQuery(queryParamsCtx.Fork("numberAnyInQuery"), readQueryValue("numberAnyInQuery", query), &reqParams.NumberAnyInQuery)
	p.bindNumberFloatInQuery(queryParamsCtx.Fork("numberFloatInQuery"), readQueryValue("numberFloatInQuery", query), &reqParams.NumberFloatInQuery)
	p.bindNumberDoubleInQuery(queryParamsCtx.Fork("numberDoubleInQuery"), readQueryValue("numberDoubleInQuery", query), &reqParams.NumberDoubleInQuery)
	p.bindNumberIntInQuery(queryParamsCtx.Fork("numberIntInQuery"), readQueryValue("numberIntInQuery", query), &reqParams.NumberIntInQuery)
	p.bindNumberInt32InQuery(queryParamsCtx.Fork("numberInt32InQuery"), readQueryValue("numberInt32InQuery", query), &reqParams.NumberInt32InQuery)
	p.bindNumberInt64InQuery(queryParamsCtx.Fork("numberInt64InQuery"), readQueryValue("numberInt64InQuery", query), &reqParams.NumberInt64InQuery)
	p.bindOptionalNumberAnyInQuery(queryParamsCtx.Fork("optionalNumberAnyInQuery"), readQueryValue("optionalNumberAnyInQuery", query), &reqParams.OptionalNumberAnyInQuery)
	p.bindOptionalNumberFloatInQuery(queryParamsCtx.Fork("optionalNumberFloatInQuery"), readQueryValue("optionalNumberFloatInQuery", query), &reqParams.OptionalNumberFloatInQuery)
	p.bindOptionalNumberDoubleInQuery(queryParamsCtx.Fork("optionalNumberDoubleInQuery"), readQueryValue("optionalNumberDoubleInQuery", query), &reqParams.OptionalNumberDoubleInQuery)
	p.bindOptionalNumberIntInQuery(queryParamsCtx.Fork("optionalNumberIntInQuery"), readQueryValue("optionalNumberIntInQuery", query), &reqParams.OptionalNumberIntInQuery)
	p.bindOptionalNumberInt32InQuery(queryParamsCtx.Fork("optionalNumberInt32InQuery"), readQueryValue("optionalNumberInt32InQuery", query), &reqParams.OptionalNumberInt32InQuery)
	p.bindOptionalNumberInt64InQuery(queryParamsCtx.Fork("optionalNumberInt64InQuery"), readQueryValue("optionalNumberInt64InQuery", query), &reqParams.OptionalNumberInt64InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserNumericTypesNumericTypesRequiredValidation(app *HTTPApp) paramsParser[*NumericTypesNumericTypesRequiredValidationRequest] {
	return &paramsParserNumericTypesNumericTypesRequiredValidation{
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberAnyInQuery", Location: "query"},
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "numberFloatInQuery", Location: "query"},
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "numberDoubleInQuery", Location: "query"},
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberIntInQuery", Location: "query"},
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "numberInt32InQuery", Location: "query"},
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "numberInt64InQuery", Location: "query"},
			),
		}),
		bindOptionalNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "optionalNumberAnyInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float32](100.01, false, true),
			),
		}),
		bindOptionalNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float32](
				internal.SimpleFieldValidatorParams{Field: "optionalNumberFloatInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float32](200.02, false, true),
			),
		}),
		bindOptionalNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.float64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[float64](
				internal.SimpleFieldValidatorParams{Field: "optionalNumberDoubleInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[float64](300.03, false, true),
			),
		}),
		bindOptionalNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "optionalNumberIntInQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int32](400, false, true),
			),
		}),
		bindOptionalNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int32Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int32](
				internal.SimpleFieldValidatorParams{Field: "optionalNumberInt32InQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int32](500, false, true),
			),
		}),
		bindOptionalNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.int64Parser,
			),
			validateValue: internal.NewSimpleFieldValidator[int64](
				internal.SimpleFieldValidatorParams{Field: "optionalNumberInt64InQuery", Location: "query"},
				internal.NewMinMaxValueValidator[int64](600, false, true),
			),
		}),
	}
}
