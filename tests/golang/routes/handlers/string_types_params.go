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
type _ func() BasicStringEnum

type paramsParserStringTypesStringTypesArrayItemsRangeValidation struct {
	bindUnformattedStr requestParamBinder[string, []string]
	bindCustomFormatStr requestParamBinder[string, []string]
	bindDateStr requestParamBinder[string, []time.Time]
	bindDateTimeStr requestParamBinder[string, []time.Time]
	bindByteStr requestParamBinder[string, []string]
	bindUnformattedStrInQuery requestParamBinder[[]string, []string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, []string]
	bindDateStrInQuery requestParamBinder[[]string, []time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, []time.Time]
	bindByteStrInQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *StringTypesArrayItemsRangeValidationRequest]
}

func (p *paramsParserStringTypesStringTypesArrayItemsRangeValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesArrayItemsRangeValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesArrayItemsRangeValidationParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesArrayItemsRangeValidation(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesArrayItemsRangeValidationParams] {
	return &paramsParserStringTypesStringTypesArrayItemsRangeValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
					NewMinMaxLengthValidator[string, string](10, true),
					NewMinMaxLengthValidator[string, string](20, false),
				),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
					NewMinMaxLengthValidator[string, string](20, true),
					NewMinMaxLengthValidator[string, string](30, false),
				),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
					NewMinMaxLengthValidator[string, string](30, true),
					NewMinMaxLengthValidator[string, string](40, false),
				),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
					NewMinMaxLengthValidator[string, string](10, true),
					NewMinMaxLengthValidator[string, string](20, false),
				),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
					NewMinMaxLengthValidator[string, string](20, true),
					NewMinMaxLengthValidator[string, string](30, false),
				),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
					NewMinMaxLengthValidator[string, string](30, true),
					NewMinMaxLengthValidator[string, string](40, false),
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesArrayItemsRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesArrayItemsRangeValidationRequest],
			),
			validateValue: NewStringTypesArrayItemsRangeValidationRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesArraysParsing struct {
	bindUnformattedStr requestParamBinder[string, []string]
	bindCustomFormatStr requestParamBinder[string, []string]
	bindDateStr requestParamBinder[string, []time.Time]
	bindDateTimeStr requestParamBinder[string, []time.Time]
	bindByteStr requestParamBinder[string, []string]
	bindUnformattedStrInQuery requestParamBinder[[]string, []string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, []string]
	bindDateStrInQuery requestParamBinder[[]string, []time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, []time.Time]
	bindByteStrInQuery requestParamBinder[[]string, []string]
	bindPayload requestParamBinder[*http.Request, *StringTypesArraysParsingRequest]
}

func (p *paramsParserStringTypesStringTypesArraysParsing) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesArraysParsingParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesArraysParsingParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesArraysParsing(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesArraysParsingParams] {
	return &paramsParserStringTypesStringTypesArraysParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, []time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, []string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, []time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]time.Time](
				),
				NewSimpleFieldValidator[time.Time](
				),
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, []string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]string](
				),
				NewSimpleFieldValidator[string](
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesArraysParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesArraysParsingRequest],
			),
			validateValue: NewStringTypesArraysParsingRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesEnums struct {
	bindInlineEnumParam requestParamBinder[string, StringTypesStringTypesEnumsInlineEnumParam]
	bindNullableInlineEnumParam requestParamBinder[string, *StringTypesStringTypesEnumsNullableInlineEnumParam]
	bindRefEnumParam requestParamBinder[string, BasicStringEnum]
	bindNullableRefEnumParam requestParamBinder[string, *NullableStringEnum]
	bindInlineEnumParamInQuery requestParamBinder[[]string, StringTypesStringTypesEnumsInlineEnumParamInQuery]
	bindNullableInlineEnumParamInQuery requestParamBinder[[]string, *StringTypesStringTypesEnumsNullableInlineEnumParamInQuery]
	bindRefEnumParamInQuery requestParamBinder[[]string, BasicStringEnum]
	bindNullableRefEnumParamInQuery requestParamBinder[[]string, *NullableStringEnum]
	bindPayload requestParamBinder[*http.Request, *StringTypesEnumsRequest]
	bindOptionalInlineEnumParamInQuery requestParamBinder[[]string, StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery]
	bindOptionalNullableInlineEnumParamInQuery requestParamBinder[[]string, *StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery]
	bindOptionalRefEnumParamInQuery requestParamBinder[[]string, BasicStringEnum]
	bindOptionalNullableRefEnumParamInQuery requestParamBinder[[]string, *NullableStringEnum]
}

func (p *paramsParserStringTypesStringTypesEnums) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesEnumsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesEnumsParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindInlineEnumParam(pathParamsCtx.Fork("inlineEnumParam"), readPathValue("inlineEnumParam", router, req), &reqParams.InlineEnumParam)
	p.bindNullableInlineEnumParam(pathParamsCtx.Fork("nullableInlineEnumParam"), readPathValue("nullableInlineEnumParam", router, req), &reqParams.NullableInlineEnumParam)
	p.bindRefEnumParam(pathParamsCtx.Fork("refEnumParam"), readPathValue("refEnumParam", router, req), &reqParams.RefEnumParam)
	p.bindNullableRefEnumParam(pathParamsCtx.Fork("nullableRefEnumParam"), readPathValue("nullableRefEnumParam", router, req), &reqParams.NullableRefEnumParam)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindInlineEnumParamInQuery(queryParamsCtx.Fork("inlineEnumParamInQuery"), readQueryValue("inlineEnumParamInQuery", query), &reqParams.InlineEnumParamInQuery)
	p.bindOptionalInlineEnumParamInQuery(queryParamsCtx.Fork("optionalInlineEnumParamInQuery"), readQueryValue("optionalInlineEnumParamInQuery", query), &reqParams.OptionalInlineEnumParamInQuery)
	p.bindNullableInlineEnumParamInQuery(queryParamsCtx.Fork("nullableInlineEnumParamInQuery"), readQueryValue("nullableInlineEnumParamInQuery", query), &reqParams.NullableInlineEnumParamInQuery)
	p.bindOptionalNullableInlineEnumParamInQuery(queryParamsCtx.Fork("optionalNullableInlineEnumParamInQuery"), readQueryValue("optionalNullableInlineEnumParamInQuery", query), &reqParams.OptionalNullableInlineEnumParamInQuery)
	p.bindRefEnumParamInQuery(queryParamsCtx.Fork("refEnumParamInQuery"), readQueryValue("refEnumParamInQuery", query), &reqParams.RefEnumParamInQuery)
	p.bindNullableRefEnumParamInQuery(queryParamsCtx.Fork("nullableRefEnumParamInQuery"), readQueryValue("nullableRefEnumParamInQuery", query), &reqParams.NullableRefEnumParamInQuery)
	p.bindOptionalRefEnumParamInQuery(queryParamsCtx.Fork("optionalRefEnumParamInQuery"), readQueryValue("optionalRefEnumParamInQuery", query), &reqParams.OptionalRefEnumParamInQuery)
	p.bindOptionalNullableRefEnumParamInQuery(queryParamsCtx.Fork("optionalNullableRefEnumParamInQuery"), readQueryValue("optionalNullableRefEnumParamInQuery", query), &reqParams.OptionalNullableRefEnumParamInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesEnums(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesEnumsParams] {
	return &paramsParserStringTypesStringTypesEnums{
		bindInlineEnumParam: newRequestParamBinder(binderParams[string, StringTypesStringTypesEnumsInlineEnumParam]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				ParseStringTypesStringTypesEnumsInlineEnumParam,
			),
			validateValue: NewSimpleFieldValidator[StringTypesStringTypesEnumsInlineEnumParam](
			),
		}),
		bindNullableInlineEnumParam: newRequestParamBinder(binderParams[string, *StringTypesStringTypesEnumsNullableInlineEnumParam]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(ParseStringTypesStringTypesEnumsNullableInlineEnumParam),
			),
			validateValue: NewSimpleFieldValidator[*StringTypesStringTypesEnumsNullableInlineEnumParam](
			),
		}),
		bindRefEnumParam: newRequestParamBinder(binderParams[string, BasicStringEnum]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				ParseBasicStringEnum,
			),
			validateValue: NewSimpleFieldValidator[BasicStringEnum](
			),
		}),
		bindNullableRefEnumParam: newRequestParamBinder(binderParams[string, *NullableStringEnum]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(ParseNullableStringEnum),
			),
			validateValue: NewSimpleFieldValidator[*NullableStringEnum](
			),
		}),
		bindInlineEnumParamInQuery: newRequestParamBinder(binderParams[[]string, StringTypesStringTypesEnumsInlineEnumParamInQuery]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				ParseStringTypesStringTypesEnumsInlineEnumParamInQuery,
			),
			validateValue: NewSimpleFieldValidator[StringTypesStringTypesEnumsInlineEnumParamInQuery](
			),
		}),
		bindNullableInlineEnumParamInQuery: newRequestParamBinder(binderParams[[]string, *StringTypesStringTypesEnumsNullableInlineEnumParamInQuery]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(ParseStringTypesStringTypesEnumsNullableInlineEnumParamInQuery),
			),
			validateValue: NewSimpleFieldValidator[*StringTypesStringTypesEnumsNullableInlineEnumParamInQuery](
			),
		}),
		bindRefEnumParamInQuery: newRequestParamBinder(binderParams[[]string, BasicStringEnum]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				ParseBasicStringEnum,
			),
			validateValue: NewSimpleFieldValidator[BasicStringEnum](
			),
		}),
		bindNullableRefEnumParamInQuery: newRequestParamBinder(binderParams[[]string, *NullableStringEnum]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(ParseNullableStringEnum),
			),
			validateValue: NewSimpleFieldValidator[*NullableStringEnum](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesEnumsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesEnumsRequest],
			),
			validateValue: NewStringTypesEnumsRequestValidator(),
		}),
		bindOptionalInlineEnumParamInQuery: newRequestParamBinder(binderParams[[]string, StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				ParseStringTypesStringTypesEnumsOptionalInlineEnumParamInQuery,
			),
			validateValue: NewSimpleFieldValidator[StringTypesStringTypesEnumsOptionalInlineEnumParamInQuery](
			),
		}),
		bindOptionalNullableInlineEnumParamInQuery: newRequestParamBinder(binderParams[[]string, *StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(ParseStringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery),
			),
			validateValue: NewSimpleFieldValidator[*StringTypesStringTypesEnumsOptionalNullableInlineEnumParamInQuery](
			),
		}),
		bindOptionalRefEnumParamInQuery: newRequestParamBinder(binderParams[[]string, BasicStringEnum]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				ParseBasicStringEnum,
			),
			validateValue: NewSimpleFieldValidator[BasicStringEnum](
			),
		}),
		bindOptionalNullableRefEnumParamInQuery: newRequestParamBinder(binderParams[[]string, *NullableStringEnum]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(ParseNullableStringEnum),
			),
			validateValue: NewSimpleFieldValidator[*NullableStringEnum](
			),
		}),
	}
}

type paramsParserStringTypesStringTypesNullableArrayItems struct {
	bindUnformattedStr requestParamBinder[string, []*string]
	bindCustomFormatStr requestParamBinder[string, []*string]
	bindDateStr requestParamBinder[string, []*time.Time]
	bindDateTimeStr requestParamBinder[string, []*time.Time]
	bindByteStr requestParamBinder[string, []*string]
	bindUnformattedStrInQuery requestParamBinder[[]string, []*string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, []*string]
	bindDateStrInQuery requestParamBinder[[]string, []*time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, []*time.Time]
	bindByteStrInQuery requestParamBinder[[]string, []*string]
	bindPayload requestParamBinder[*http.Request, *StringTypesNullableArrayItemsRequest]
}

func (p *paramsParserStringTypesStringTypesNullableArrayItems) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesNullableArrayItemsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesNullableArrayItemsParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableArrayItems(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesNullableArrayItemsParams] {
	return &paramsParserStringTypesStringTypesNullableArrayItems{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, []*string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*string](
				),
				NewSimpleFieldValidator[*string](
					SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
					SkipNullValidator(NewMinMaxLengthValidator[string, string](20, false)),
					SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, []*string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*string](
				),
				NewSimpleFieldValidator[*string](
					SkipNullValidator(NewMinMaxLengthValidator[string, string](20, true)),
					SkipNullValidator(NewMinMaxLengthValidator[string, string](30, false)),
					SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, []*time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.dateParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*time.Time](
				),
				NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, []*time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.timeParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*time.Time](
				),
				NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, []*string]{
			required: true,
			parseValue: parseSoloValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*string](
				),
				NewSimpleFieldValidator[*string](
					SkipNullValidator(NewMinMaxLengthValidator[string, string](30, true)),
					SkipNullValidator(NewMinMaxLengthValidator[string, string](40, false)),
				),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, []*string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*string](
				),
				NewSimpleFieldValidator[*string](
					SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
					SkipNullValidator(NewMinMaxLengthValidator[string, string](20, false)),
					SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, []*string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*string](
				),
				NewSimpleFieldValidator[*string](
					SkipNullValidator(NewMinMaxLengthValidator[string, string](20, true)),
					SkipNullValidator(NewMinMaxLengthValidator[string, string](30, false)),
					SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
				),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, []*time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.dateParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*time.Time](
				),
				NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, []*time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.timeParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*time.Time](
				),
				NewSimpleFieldValidator[*time.Time](
				),
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, []*string]{
			required: true,
			parseValue: parseMultiValueParamAsSlice(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*string](
				),
				NewSimpleFieldValidator[*string](
					SkipNullValidator(NewMinMaxLengthValidator[string, string](30, true)),
					SkipNullValidator(NewMinMaxLengthValidator[string, string](40, false)),
				),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesNullableArrayItemsRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesNullableArrayItemsRequest],
			),
			validateValue: NewStringTypesNullableArrayItemsRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesNullableParsing struct {
	bindUnformattedStr requestParamBinder[string, *string]
	bindCustomFormatStr requestParamBinder[string, *string]
	bindDateStr requestParamBinder[string, *time.Time]
	bindDateTimeStr requestParamBinder[string, *time.Time]
	bindByteStr requestParamBinder[string, *string]
	bindUnformattedStrInQuery requestParamBinder[[]string, *string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, *string]
	bindDateStrInQuery requestParamBinder[[]string, *time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, *time.Time]
	bindByteStrInQuery requestParamBinder[[]string, *string]
	bindPayload requestParamBinder[*http.Request, *StringTypesNullableParsingRequest]
}

func (p *paramsParserStringTypesStringTypesNullableParsing) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesNullableParsingParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesNullableParsingParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableParsing(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesNullableParsingParams] {
	return &paramsParserStringTypesStringTypesNullableParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, *string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, *string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, *time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.dateParser),
			),
			validateValue: NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, *time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.timeParser),
			),
			validateValue: NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, *string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, *time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.dateParser),
			),
			validateValue: NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, *time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.timeParser),
			),
			validateValue: NewSimpleFieldValidator[*time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesNullableParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesNullableParsingRequest],
			),
			validateValue: NewStringTypesNullableParsingRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesNullableRequiredValidation struct {
	bindUnformattedStrInQuery requestParamBinder[[]string, *string]
	bindPayload requestParamBinder[*http.Request, *StringTypesNullableRequiredValidationRequest]
	bindOptionalUnformattedStrInQuery requestParamBinder[[]string, *string]
}

func (p *paramsParserStringTypesStringTypesNullableRequiredValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesNullableRequiredValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesNullableRequiredValidationParams{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindOptionalUnformattedStrInQuery(queryParamsCtx.Fork("optionalUnformattedStrInQuery"), readQueryValue("optionalUnformattedStrInQuery", query), &reqParams.OptionalUnformattedStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesNullableRequiredValidation(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesNullableRequiredValidationParams] {
	return &paramsParserStringTypesStringTypesNullableRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](100, false)),
				SkipNullValidator(NewPatternValidator[string](".*")),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesNullableRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesNullableRequiredValidationRequest],
			),
			validateValue: NewStringTypesNullableRequiredValidationRequestValidator(),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, *string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				parseNullableParam(rootHandler.knownParsers.stringParser),
			),
			validateValue: NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](100, false)),
				SkipNullValidator(NewPatternValidator[string](".*")),
			),
		}),
	}
}

type paramsParserStringTypesStringTypesParsing struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindDateStr requestParamBinder[string, time.Time]
	bindDateTimeStr requestParamBinder[string, time.Time]
	bindByteStr requestParamBinder[string, string]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindByteStrInQuery requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *StringTypesParsingRequest]
}

func (p *paramsParserStringTypesStringTypesParsing) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesParsingParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesParsingParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesParsing(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesParsingParams] {
	return &paramsParserStringTypesStringTypesParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesParsingRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesParsingRequest],
			),
			validateValue: NewStringTypesParsingRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesPatternValidation struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindDateStr requestParamBinder[string, time.Time]
	bindDateTimeStr requestParamBinder[string, time.Time]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindPayload requestParamBinder[*http.Request, *StringTypesPatternValidationRequest]
}

func (p *paramsParserStringTypesStringTypesPatternValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesPatternValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesPatternValidationParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesPatternValidation(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesPatternValidationParams] {
	return &paramsParserStringTypesStringTypesPatternValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewPatternValidator[string]("^\\d{10}$"),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewPatternValidator[string]("^\\d{20}$"),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesPatternValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesPatternValidationRequest],
			),
			validateValue: NewStringTypesPatternValidationRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesRangeValidation struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindDateStr requestParamBinder[string, time.Time]
	bindDateTimeStr requestParamBinder[string, time.Time]
	bindByteStr requestParamBinder[string, string]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindByteStrInQuery requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *StringTypesRangeValidationRequest]
}

func (p *paramsParserStringTypesStringTypesRangeValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesRangeValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesRangeValidationParams{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindUnformattedStr(pathParamsCtx.Fork("unformattedStr"), readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(pathParamsCtx.Fork("customFormatStr"), readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(pathParamsCtx.Fork("dateStr"), readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(pathParamsCtx.Fork("dateTimeStr"), readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(pathParamsCtx.Fork("byteStr"), readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRangeValidation(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesRangeValidationParams] {
	return &paramsParserStringTypesStringTypesRangeValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](10, true),
				NewMinMaxLengthValidator[string, string](20, false),
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](20, true),
				NewMinMaxLengthValidator[string, string](30, false),
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, time.Time]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](30, true),
				NewMinMaxLengthValidator[string, string](40, false),
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](10, true),
				NewMinMaxLengthValidator[string, string](20, false),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](20, true),
				NewMinMaxLengthValidator[string, string](30, false),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](30, true),
				NewMinMaxLengthValidator[string, string](40, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesRangeValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesRangeValidationRequest],
			),
			validateValue: NewStringTypesRangeValidationRequestValidator(),
		}),
	}
}

type paramsParserStringTypesStringTypesRequiredValidation struct {
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, time.Time]
	bindDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindByteStrInQuery requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *StringTypesRequiredValidationRequest]
	bindOptionalUnformattedStrInQuery requestParamBinder[[]string, string]
	bindOptionalCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindOptionalDateStrInQuery requestParamBinder[[]string, time.Time]
	bindOptionalDateTimeStrInQuery requestParamBinder[[]string, time.Time]
	bindOptionalByteStrInQuery requestParamBinder[[]string, string]
}

func (p *paramsParserStringTypesStringTypesRequiredValidation) parse(router httpRouter, req *http.Request) (*StringTypesStringTypesRequiredValidationParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &StringTypesStringTypesRequiredValidationParams{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindUnformattedStrInQuery(queryParamsCtx.Fork("unformattedStrInQuery"), readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(queryParamsCtx.Fork("customFormatStrInQuery"), readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(queryParamsCtx.Fork("dateStrInQuery"), readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(queryParamsCtx.Fork("dateTimeStrInQuery"), readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(queryParamsCtx.Fork("byteStrInQuery"), readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	p.bindOptionalUnformattedStrInQuery(queryParamsCtx.Fork("optionalUnformattedStrInQuery"), readQueryValue("optionalUnformattedStrInQuery", query), &reqParams.OptionalUnformattedStrInQuery)
	p.bindOptionalCustomFormatStrInQuery(queryParamsCtx.Fork("optionalCustomFormatStrInQuery"), readQueryValue("optionalCustomFormatStrInQuery", query), &reqParams.OptionalCustomFormatStrInQuery)
	p.bindOptionalDateStrInQuery(queryParamsCtx.Fork("optionalDateStrInQuery"), readQueryValue("optionalDateStrInQuery", query), &reqParams.OptionalDateStrInQuery)
	p.bindOptionalDateTimeStrInQuery(queryParamsCtx.Fork("optionalDateTimeStrInQuery"), readQueryValue("optionalDateTimeStrInQuery", query), &reqParams.OptionalDateTimeStrInQuery)
	p.bindOptionalByteStrInQuery(queryParamsCtx.Fork("optionalByteStrInQuery"), readQueryValue("optionalByteStrInQuery", query), &reqParams.OptionalByteStrInQuery)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRequiredValidation(rootHandler *RootHandler) paramsParser[*StringTypesStringTypesRequiredValidationParams] {
	return &paramsParserStringTypesStringTypesRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](10, true),
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](20, true),
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](30, true),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *StringTypesRequiredValidationRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*StringTypesRequiredValidationRequest],
			),
			validateValue: NewStringTypesRequiredValidationRequestValidator(),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](10, true),
			),
		}),
		bindOptionalCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](20, true),
			),
		}),
		bindOptionalDateStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.dateParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindOptionalDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, time.Time]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.timeParser,
			),
			validateValue: NewSimpleFieldValidator[time.Time](
			),
		}),
		bindOptionalByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](30, true),
			),
		}),
	}
}

type stringTypesControllerBuilder struct {
	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArrayItemsRangeValidationParams,
	//
	// Response type: none
	StringTypesArrayItemsRangeValidation genericHandlerBuilder[
		*StringTypesStringTypesArrayItemsRangeValidationParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesArrayItemsRangeValidationParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesArrayItemsRangeValidationParams, void],
	]

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArraysParsingParams,
	//
	// Response type: none
	StringTypesArraysParsing genericHandlerBuilder[
		*StringTypesStringTypesArraysParsingParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesArraysParsingParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesArraysParsingParams, void],
	]

	// POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
	//
	// Request type: StringTypesStringTypesEnumsParams,
	//
	// Response type: none
	StringTypesEnums genericHandlerBuilder[
		*StringTypesStringTypesEnumsParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesEnumsParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesEnumsParams, void],
	]

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableArrayItemsParams,
	//
	// Response type: none
	StringTypesNullableArrayItems genericHandlerBuilder[
		*StringTypesStringTypesNullableArrayItemsParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesNullableArrayItemsParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesNullableArrayItemsParams, void],
	]

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingParams,
	//
	// Response type: none
	StringTypesNullableParsing genericHandlerBuilder[
		*StringTypesStringTypesNullableParsingParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesNullableParsingParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesNullableParsingParams, void],
	]

	// POST /string-types/nullable-required-validation
	//
	// Request type: StringTypesStringTypesNullableRequiredValidationParams,
	//
	// Response type: none
	StringTypesNullableRequiredValidation genericHandlerBuilder[
		*StringTypesStringTypesNullableRequiredValidationParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesNullableRequiredValidationParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesNullableRequiredValidationParams, void],
	]

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingParams,
	//
	// Response type: none
	StringTypesParsing genericHandlerBuilder[
		*StringTypesStringTypesParsingParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesParsingParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesParsingParams, void],
	]

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationParams,
	//
	// Response type: none
	StringTypesPatternValidation genericHandlerBuilder[
		*StringTypesStringTypesPatternValidationParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesPatternValidationParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesPatternValidationParams, void],
	]

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationParams,
	//
	// Response type: none
	StringTypesRangeValidation genericHandlerBuilder[
		*StringTypesStringTypesRangeValidationParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesRangeValidationParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesRangeValidationParams, void],
	]

	// POST /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationParams,
	//
	// Response type: none
	StringTypesRequiredValidation genericHandlerBuilder[
		*StringTypesStringTypesRequiredValidationParams,
		void,
		handlerActionFuncNoResponse[*StringTypesStringTypesRequiredValidationParams, void],
		httpHandlerActionFuncNoResponse[*StringTypesStringTypesRequiredValidationParams, void],
	]
}

func newStringTypesControllerBuilder(app *RootHandler) *stringTypesControllerBuilder {
	return &stringTypesControllerBuilder{
		// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesArrayItemsRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesArrayItemsRangeValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesArrayItemsRangeValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesArrayItemsRangeValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesArrayItemsRangeValidation(app),
			},
		),

		// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesArraysParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesArraysParsingParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesArraysParsingParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesArraysParsingParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesArraysParsing(app),
			},
		),

		// POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
		StringTypesEnums: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesEnumsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesEnumsParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesEnumsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesEnums(app),
			},
		),

		// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesNullableArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableArrayItemsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableArrayItemsParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesNullableArrayItemsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesNullableArrayItems(app),
			},
		),

		// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesNullableParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableParsingParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableParsingParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesNullableParsingParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesNullableParsing(app),
			},
		),

		// POST /string-types/nullable-required-validation
		StringTypesNullableRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableRequiredValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesNullableRequiredValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesNullableRequiredValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesNullableRequiredValidation(app),
			},
		),

		// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesParsingParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesParsingParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesParsingParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesParsing(app),
			},
		),

		// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
		StringTypesPatternValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesPatternValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesPatternValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesPatternValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesPatternValidation(app),
			},
		),

		// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
		StringTypesRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesRangeValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesRangeValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesRangeValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesRangeValidation(app),
			},
		),

		// POST /string-types/required-validation
		StringTypesRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*StringTypesStringTypesRequiredValidationParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*StringTypesStringTypesRequiredValidationParams,
				void,
			](),
			makeActionBuilderParams[
				*StringTypesStringTypesRequiredValidationParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserStringTypesStringTypesRequiredValidation(app),
			},
		),
	}
}
