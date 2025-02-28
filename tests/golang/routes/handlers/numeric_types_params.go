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

func (p *paramsParserNumericTypesNumericTypesArrayItems) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesArrayItemsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNumericTypesArrayItemsParams{}
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

func newParamsParserNumericTypesNumericTypesArrayItems(rootHandler *RootHandler) paramsParser[*NumericTypesNumericTypesArrayItemsParams] {
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

func (p *paramsParserNumericTypesNumericTypesNullable) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesNullableParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNumericTypesNullableParams{}
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

func newParamsParserNumericTypesNumericTypesNullable(rootHandler *RootHandler) paramsParser[*NumericTypesNumericTypesNullableParams] {
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

func (p *paramsParserNumericTypesNumericTypesNullableArrayItems) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesNullableArrayItemsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNumericTypesNullableArrayItemsParams{}
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

func newParamsParserNumericTypesNumericTypesNullableArrayItems(rootHandler *RootHandler) paramsParser[*NumericTypesNumericTypesNullableArrayItemsParams] {
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

func (p *paramsParserNumericTypesNumericTypesParsing) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesParsingParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNumericTypesParsingParams{}
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

func newParamsParserNumericTypesNumericTypesParsing(rootHandler *RootHandler) paramsParser[*NumericTypesNumericTypesParsingParams] {
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

func (p *paramsParserNumericTypesNumericTypesRangeValidation) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRangeValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNumericTypesRangeValidationParams{}
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

func newParamsParserNumericTypesNumericTypesRangeValidation(rootHandler *RootHandler) paramsParser[*NumericTypesNumericTypesRangeValidationParams] {
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

func (p *paramsParserNumericTypesNumericTypesRangeValidationExclusive) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRangeValidationExclusiveParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNumericTypesRangeValidationExclusiveParams{}
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

func newParamsParserNumericTypesNumericTypesRangeValidationExclusive(rootHandler *RootHandler) paramsParser[*NumericTypesNumericTypesRangeValidationExclusiveParams] {
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

func (p *paramsParserNumericTypesNumericTypesRequiredValidation) parse(router httpRouter, req *http.Request) (*NumericTypesNumericTypesRequiredValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &NumericTypesNumericTypesRequiredValidationParams{}
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

func newParamsParserNumericTypesNumericTypesRequiredValidation(rootHandler *RootHandler) paramsParser[*NumericTypesNumericTypesRequiredValidationParams] {
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
	// Request type: NumericTypesNumericTypesArrayItemsParams,
	//
	// Response type: none
	NumericTypesArrayItems genericHandlerBuilder[
		*NumericTypesNumericTypesArrayItemsParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesArrayItemsParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesArrayItemsParams, void],
	]

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableParams,
	//
	// Response type: none
	NumericTypesNullable genericHandlerBuilder[
		*NumericTypesNumericTypesNullableParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesNullableParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesNullableParams, void],
	]

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableArrayItemsParams,
	//
	// Response type: none
	NumericTypesNullableArrayItems genericHandlerBuilder[
		*NumericTypesNumericTypesNullableArrayItemsParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesNullableArrayItemsParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesNullableArrayItemsParams, void],
	]

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesParsingParams,
	//
	// Response type: none
	NumericTypesParsing genericHandlerBuilder[
		*NumericTypesNumericTypesParsingParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesParsingParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesParsingParams, void],
	]

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationParams,
	//
	// Response type: none
	NumericTypesRangeValidation genericHandlerBuilder[
		*NumericTypesNumericTypesRangeValidationParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationParams, void],
	]

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationExclusiveParams,
	//
	// Response type: none
	NumericTypesRangeValidationExclusive genericHandlerBuilder[
		*NumericTypesNumericTypesRangeValidationExclusiveParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationExclusiveParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationExclusiveParams, void],
	]

	// GET /numeric-types/required-validation
	//
	// Request type: NumericTypesNumericTypesRequiredValidationParams,
	//
	// Response type: none
	NumericTypesRequiredValidation genericHandlerBuilder[
		*NumericTypesNumericTypesRequiredValidationParams,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesRequiredValidationParams, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesRequiredValidationParams, void],
	]
}

func newNumericTypesControllerBuilder(app *RootHandler) *numericTypesControllerBuilder {
	return &numericTypesControllerBuilder{
		// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesArrayItemsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesArrayItemsParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesArrayItemsParams,
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
				*NumericTypesNumericTypesNullableParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesNullableParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesNullableParams,
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
				*NumericTypesNumericTypesNullableArrayItemsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesNullableArrayItemsParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesNullableArrayItemsParams,
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
				*NumericTypesNumericTypesParsingParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesParsingParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesParsingParams,
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
				*NumericTypesNumericTypesRangeValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRangeValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesRangeValidationParams,
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
				*NumericTypesNumericTypesRangeValidationExclusiveParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRangeValidationExclusiveParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesRangeValidationExclusiveParams,
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
				*NumericTypesNumericTypesRequiredValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRequiredValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesRequiredValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesRequiredValidation(app),
			},
		),
	}
}
