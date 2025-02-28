package handlers

import (
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
	. "github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports if that happens.
var _ = BindingContext{}
var _ = http.MethodGet
var _ = time.Time{}
type _ func() NumericTypesArrayItemsRequest

type paramsParserNumericTypesNumericTypesArrayItems struct {
	bindNumberAny requestParamBinder[string, []float32]
	bindNumberFloat requestParamBinder[string, []float32]
	bindNumberDouble requestParamBinder[string, []float64]
	bindNumberInt requestParamBinder[string, []int32]
	bindNumberInt32 requestParamBinder[string, []int32]
	bindNumberInt64 requestParamBinder[string, []int64]
	bindNumberAnyInQuery requestParamBinder[[]string, []float32]
	bindNumberFloatInQuery requestParamBinder[[]string, []float32]
	bindNumberDoubleInQuery requestParamBinder[[]string, []float64]
	bindNumberIntInQuery requestParamBinder[[]string, []int32]
	bindNumberInt32InQuery requestParamBinder[[]string, []int32]
	bindNumberInt64InQuery requestParamBinder[[]string, []int64]
	bindPayload requestParamBinder[*http.Request, *NumericTypesArrayItemsRequest]
}

func (p *paramsParserNumericTypesNumericTypesArrayItems) parse(router httpRouter, req *http.Request) (*NumericTypesArrayItemsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesArrayItemsParams{}
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

func newParamsParserNumericTypesNumericTypesArrayItems(rootHandler *RootHandler) paramsParser[*NumericTypesArrayItemsParams] {
	return &paramsParserNumericTypesNumericTypesArrayItems{
		bindNumberAny: newRequestParamBinder(binderParams[string, []float32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]float32](
				),
				NewSimpleFieldValidator[float32](
					NewMinMaxValueValidator[float32](100.01, false, true),
					NewMinMaxValueValidator[float32](200.02, false, false),
				),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, []float32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]float32](
				),
				NewSimpleFieldValidator[float32](
					NewMinMaxValueValidator[float32](200.02, false, true),
					NewMinMaxValueValidator[float32](300.03, false, false),
				),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, []float64]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]float64](
				),
				NewSimpleFieldValidator[float64](
					NewMinMaxValueValidator[float64](300.03, false, true),
					NewMinMaxValueValidator[float64](400.04, false, false),
				),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, []int32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]int32](
				),
				NewSimpleFieldValidator[int32](
					NewMinMaxValueValidator[int32](400, false, true),
					NewMinMaxValueValidator[int32](500, false, false),
				),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, []int32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]int32](
				),
				NewSimpleFieldValidator[int32](
					NewMinMaxValueValidator[int32](500, false, true),
					NewMinMaxValueValidator[int32](600, false, false),
				),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, []int64]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]int64](
				),
				NewSimpleFieldValidator[int64](
					NewMinMaxValueValidator[int64](600, false, true),
					NewMinMaxValueValidator[int64](700, false, false),
				),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, []float32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]float32](
				),
				NewSimpleFieldValidator[float32](
					NewMinMaxValueValidator[float32](100.01, false, true),
					NewMinMaxValueValidator[float32](200.02, false, false),
				),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, []float32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]float32](
				),
				NewSimpleFieldValidator[float32](
					NewMinMaxValueValidator[float32](200.02, false, true),
					NewMinMaxValueValidator[float32](300.03, false, false),
				),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, []float64]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]float64](
				),
				NewSimpleFieldValidator[float64](
					NewMinMaxValueValidator[float64](300.03, false, true),
					NewMinMaxValueValidator[float64](400.04, false, false),
				),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, []int32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]int32](
				),
				NewSimpleFieldValidator[int32](
					NewMinMaxValueValidator[int32](400, false, true),
					NewMinMaxValueValidator[int32](500, false, false),
				),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, []int32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]int32](
				),
				NewSimpleFieldValidator[int32](
					NewMinMaxValueValidator[int32](500, false, true),
					NewMinMaxValueValidator[int32](600, false, false),
				),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, []int64]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]int64](
				),
				NewSimpleFieldValidator[int64](
					NewMinMaxValueValidator[int64](600, false, true),
					NewMinMaxValueValidator[int64](700, false, false),
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *NumericTypesArrayItemsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*NumericTypesArrayItemsRequest],
			),
			validateValue: NewNumericTypesArrayItemsRequestValidator(),
		}),
	}
}

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
	bindPayload requestParamBinder[*http.Request, *NumericTypesNullableRequest]
	bindOptionalNumberAnyInQuery requestParamBinder[[]string, *float32]
}

func (p *paramsParserNumericTypesNumericTypesNullable) parse(router httpRouter, req *http.Request) (*NumericTypesNullableParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNullableParams{}
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

func newParamsParserNumericTypesNumericTypesNullable(rootHandler *RootHandler) paramsParser[*NumericTypesNullableParams] {
	return &paramsParserNumericTypesNumericTypesNullable{
		bindNumberAny: newRequestParamBinder(binderParams[string, *float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, *float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, *float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.float64Parser),
			),
			validateValue: NewSimpleFieldValidator[*float64](
				SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, *int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, *int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, *int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.int64Parser),
			),
			validateValue: NewSimpleFieldValidator[*int64](
				SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, *float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.float64Parser),
			),
			validateValue: NewSimpleFieldValidator[*float64](
				SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, *int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, *int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, *int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.int64Parser),
			),
			validateValue: NewSimpleFieldValidator[*int64](
				SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *NumericTypesNullableRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*NumericTypesNullableRequest],
			),
			validateValue: NewNumericTypesNullableRequestValidator(),
		}),
		bindOptionalNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, *float32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
			),
		}),
	}
}

type paramsParserNumericTypesNumericTypesNullableArrayItems struct {
	bindNumberAny requestParamBinder[string, []*float32]
	bindNumberFloat requestParamBinder[string, []*float32]
	bindNumberDouble requestParamBinder[string, []*float64]
	bindNumberInt requestParamBinder[string, []*int32]
	bindNumberInt32 requestParamBinder[string, []*int32]
	bindNumberInt64 requestParamBinder[string, []*int64]
	bindNumberAnyInQuery requestParamBinder[[]string, []*float32]
	bindNumberFloatInQuery requestParamBinder[[]string, []*float32]
	bindNumberDoubleInQuery requestParamBinder[[]string, []*float64]
	bindNumberIntInQuery requestParamBinder[[]string, []*int32]
	bindNumberInt32InQuery requestParamBinder[[]string, []*int32]
	bindNumberInt64InQuery requestParamBinder[[]string, []*int64]
	bindPayload requestParamBinder[*http.Request, *NumericTypesNullableArrayItemsRequest]
}

func (p *paramsParserNumericTypesNumericTypesNullableArrayItems) parse(router httpRouter, req *http.Request) (*NumericTypesNullableArrayItemsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNullableArrayItemsParams{}
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

func newParamsParserNumericTypesNumericTypesNullableArrayItems(rootHandler *RootHandler) paramsParser[*NumericTypesNullableArrayItemsParams] {
	return &paramsParserNumericTypesNumericTypesNullableArrayItems{
		bindNumberAny: newRequestParamBinder(binderParams[string, []*float32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*float32](
				),
				NewSimpleFieldValidator[*float32](
					SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
				),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, []*float32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*float32](
				),
				NewSimpleFieldValidator[*float32](
					SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
				),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, []*float64]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.float64Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*float64](
				),
				NewSimpleFieldValidator[*float64](
					SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
				),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, []*int32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*int32](
				),
				NewSimpleFieldValidator[*int32](
					SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
				),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, []*int32]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*int32](
				),
				NewSimpleFieldValidator[*int32](
					SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
				),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, []*int64]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.int64Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*int64](
				),
				NewSimpleFieldValidator[*int64](
					SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
				),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, []*float32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*float32](
				),
				NewSimpleFieldValidator[*float32](
					SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
				),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, []*float32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.float32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*float32](
				),
				NewSimpleFieldValidator[*float32](
					SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
				),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, []*float64]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.float64Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*float64](
				),
				NewSimpleFieldValidator[*float64](
					SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
				),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, []*int32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*int32](
				),
				NewSimpleFieldValidator[*int32](
					SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
				),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, []*int32]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.int32Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*int32](
				),
				NewSimpleFieldValidator[*int32](
					SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
				),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, []*int64]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.int64Parser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*int64](
				),
				NewSimpleFieldValidator[*int64](
					SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
					SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *NumericTypesNullableArrayItemsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*NumericTypesNullableArrayItemsRequest],
			),
			validateValue: NewNumericTypesNullableArrayItemsRequestValidator(),
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
	bindPayload requestParamBinder[*http.Request, *NumericTypesParsingRequest]
}

func (p *paramsParserNumericTypesNumericTypesParsing) parse(router httpRouter, req *http.Request) (*NumericTypesParsingParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesParsingParams{}
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

func newParamsParserNumericTypesNumericTypesParsing(rootHandler *RootHandler) paramsParser[*NumericTypesParsingParams] {
	return &paramsParserNumericTypesNumericTypesParsing{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *NumericTypesParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*NumericTypesParsingRequest],
			),
			validateValue: NewNumericTypesParsingRequestValidator(),
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
	bindPayload requestParamBinder[*http.Request, *NumericTypesRangeValidationRequest]
}

func (p *paramsParserNumericTypesNumericTypesRangeValidation) parse(router httpRouter, req *http.Request) (*NumericTypesRangeValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesRangeValidationParams{}
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

func newParamsParserNumericTypesNumericTypesRangeValidation(rootHandler *RootHandler) paramsParser[*NumericTypesRangeValidationParams] {
	return &paramsParserNumericTypesNumericTypesRangeValidation{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](100.01, false, true),
				NewMinMaxValueValidator[float32](200.02, false, false),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](200.02, false, true),
				NewMinMaxValueValidator[float32](300.03, false, false),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
				NewMinMaxValueValidator[float64](300.03, false, true),
				NewMinMaxValueValidator[float64](400.04, false, false),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](400, false, true),
				NewMinMaxValueValidator[int32](500, false, false),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](500, false, true),
				NewMinMaxValueValidator[int32](600, false, false),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](600, false, true),
				NewMinMaxValueValidator[int64](700, false, false),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](100.01, false, true),
				NewMinMaxValueValidator[float32](200.02, false, false),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](200.02, false, true),
				NewMinMaxValueValidator[float32](300.03, false, false),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
				NewMinMaxValueValidator[float64](300.03, false, true),
				NewMinMaxValueValidator[float64](400.04, false, false),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](400, false, true),
				NewMinMaxValueValidator[int32](500, false, false),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](500, false, true),
				NewMinMaxValueValidator[int32](600, false, false),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](600, false, true),
				NewMinMaxValueValidator[int64](700, false, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *NumericTypesRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*NumericTypesRangeValidationRequest],
			),
			validateValue: NewNumericTypesRangeValidationRequestValidator(),
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
	bindPayload requestParamBinder[*http.Request, *NumericTypesRangeValidationExclusiveRequest]
}

func (p *paramsParserNumericTypesNumericTypesRangeValidationExclusive) parse(router httpRouter, req *http.Request) (*NumericTypesRangeValidationExclusiveParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesRangeValidationExclusiveParams{}
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

func newParamsParserNumericTypesNumericTypesRangeValidationExclusive(rootHandler *RootHandler) paramsParser[*NumericTypesRangeValidationExclusiveParams] {
	return &paramsParserNumericTypesNumericTypesRangeValidationExclusive{
		bindNumberAny: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](100.01, true, true),
				NewMinMaxValueValidator[float32](200.02, true, false),
			),
		}),
		bindNumberFloat: newRequestParamBinder(binderParams[string, float32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](200.02, true, true),
				NewMinMaxValueValidator[float32](300.03, true, false),
			),
		}),
		bindNumberDouble: newRequestParamBinder(binderParams[string, float64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
				NewMinMaxValueValidator[float64](300.03, true, true),
				NewMinMaxValueValidator[float64](400.04, true, false),
			),
		}),
		bindNumberInt: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](400, true, true),
				NewMinMaxValueValidator[int32](500, true, false),
			),
		}),
		bindNumberInt32: newRequestParamBinder(binderParams[string, int32]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](500, true, true),
				NewMinMaxValueValidator[int32](600, true, false),
			),
		}),
		bindNumberInt64: newRequestParamBinder(binderParams[string, int64]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](600, true, true),
				NewMinMaxValueValidator[int64](700, true, false),
			),
		}),
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](100.01, true, true),
				NewMinMaxValueValidator[float32](200.02, true, false),
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](200.02, true, true),
				NewMinMaxValueValidator[float32](300.03, true, false),
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
				NewMinMaxValueValidator[float64](300.03, true, true),
				NewMinMaxValueValidator[float64](400.04, true, false),
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](400, true, true),
				NewMinMaxValueValidator[int32](500, true, false),
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](500, true, true),
				NewMinMaxValueValidator[int32](600, true, false),
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](600, true, true),
				NewMinMaxValueValidator[int64](700, true, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *NumericTypesRangeValidationExclusiveRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*NumericTypesRangeValidationExclusiveRequest],
			),
			validateValue: NewNumericTypesRangeValidationExclusiveRequestValidator(),
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

func (p *paramsParserNumericTypesNumericTypesRequiredValidation) parse(router httpRouter, req *http.Request) (*NumericTypesRequiredValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesRequiredValidationParams{}
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

func newParamsParserNumericTypesNumericTypesRequiredValidation(rootHandler *RootHandler) paramsParser[*NumericTypesRequiredValidationParams] {
	return &paramsParserNumericTypesNumericTypesRequiredValidation{
		bindNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
			),
		}),
		bindNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
			),
		}),
		bindNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
			),
		}),
		bindNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
			),
		}),
		bindNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
			),
		}),
		bindOptionalNumberAnyInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](100.01, false, true),
			),
		}),
		bindOptionalNumberFloatInQuery: newRequestParamBinder(binderParams[[]string, float32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float32Parser,
			),
			validateValue: NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](200.02, false, true),
			),
		}),
		bindOptionalNumberDoubleInQuery: newRequestParamBinder(binderParams[[]string, float64]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.float64Parser,
			),
			validateValue: NewSimpleFieldValidator[float64](
				NewMinMaxValueValidator[float64](300.03, false, true),
			),
		}),
		bindOptionalNumberIntInQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](400, false, true),
			),
		}),
		bindOptionalNumberInt32InQuery: newRequestParamBinder(binderParams[[]string, int32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](500, false, true),
			),
		}),
		bindOptionalNumberInt64InQuery: newRequestParamBinder(binderParams[[]string, int64]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int64Parser,
			),
			validateValue: NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](600, false, true),
			),
		}),
	}
}

type numericTypesControllerBuilder struct {
	// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesArrayItemsParams,
	//
	// Response type: none
	NumericTypesArrayItems genericHandlerBuilder[
		*NumericTypesArrayItemsParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesArrayItemsParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesArrayItemsParams, void],
	]

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNullableParams,
	//
	// Response type: none
	NumericTypesNullable genericHandlerBuilder[
		*NumericTypesNullableParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNullableParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNullableParams, void],
	]

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNullableArrayItemsParams,
	//
	// Response type: none
	NumericTypesNullableArrayItems genericHandlerBuilder[
		*NumericTypesNullableArrayItemsParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNullableArrayItemsParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNullableArrayItemsParams, void],
	]

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesParsingParams,
	//
	// Response type: none
	NumericTypesParsing genericHandlerBuilder[
		*NumericTypesParsingParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesParsingParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesParsingParams, void],
	]

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesRangeValidationParams,
	//
	// Response type: none
	NumericTypesRangeValidation genericHandlerBuilder[
		*NumericTypesRangeValidationParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesRangeValidationParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesRangeValidationParams, void],
	]

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesRangeValidationExclusiveParams,
	//
	// Response type: none
	NumericTypesRangeValidationExclusive genericHandlerBuilder[
		*NumericTypesRangeValidationExclusiveParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesRangeValidationExclusiveParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesRangeValidationExclusiveParams, void],
	]

	// GET /numeric-types/required-validation
	//
	// Request type: NumericTypesRequiredValidationParams,
	//
	// Response type: none
	NumericTypesRequiredValidation genericHandlerBuilder[
		*NumericTypesRequiredValidationParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesRequiredValidationParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesRequiredValidationParams, void],
	]
}

func newNumericTypesControllerBuilder(app *RootHandler) *numericTypesControllerBuilder {
	return &numericTypesControllerBuilder{
		// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesArrayItemsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesArrayItemsParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesArrayItemsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesArrayItems(app),
			},
		),

		// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesNullable: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNullableParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNullableParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNullableParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesNullable(app),
			},
		),

		// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesNullableArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNullableArrayItemsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNullableArrayItemsParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNullableArrayItemsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesNullableArrayItems(app),
			},
		),

		// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesParsingParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesParsingParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesParsingParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesParsing(app),
			},
		),

		// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesRangeValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesRangeValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesRangeValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesRangeValidation(app),
			},
		),

		// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesRangeValidationExclusive: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesRangeValidationExclusiveParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesRangeValidationExclusiveParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesRangeValidationExclusiveParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesRangeValidationExclusive(app),
			},
		),

		// GET /numeric-types/required-validation
		NumericTypesRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesRequiredValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesRequiredValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesRequiredValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesRequiredValidation(app),
			},
		),
	}
}
