package handlers

import (
	"net/http"
)

type paramsParserStringTypesStringTypesParsing struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindDateStr requestParamBinder[string, string]
	bindDateTimeStr requestParamBinder[string, string]
	bindByteStr requestParamBinder[string, string]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, string]
	bindDateTimeStrInQuery requestParamBinder[[]string, string]
	bindByteStrInQuery requestParamBinder[[]string, string]
}

func (p *paramsParserStringTypesStringTypesParsing) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*StringTypesStringTypesParsingRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &StringTypesStringTypesParsingRequest{}
	// path params
	p.bindUnformattedStr(&bindingCtx, readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(&bindingCtx, readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindDateStr(&bindingCtx, readPathValue("dateStr", router, req), &reqParams.DateStr)
	p.bindDateTimeStr(&bindingCtx, readPathValue("dateTimeStr", router, req), &reqParams.DateTimeStr)
	p.bindByteStr(&bindingCtx, readPathValue("byteStr", router, req), &reqParams.ByteStr)
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(&bindingCtx, readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(&bindingCtx, readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(&bindingCtx, readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(&bindingCtx, readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesParsing() *paramsParserStringTypesStringTypesParsing {
	return &paramsParserStringTypesStringTypesParsing{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field: "unformattedStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindDateStr: newRequestParamBinder(binderParams[string, string]{
			field: "dateStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindDateTimeStr: newRequestParamBinder(binderParams[string, string]{
			field: "dateTimeStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field: "byteStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "dateStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "dateTimeStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
	}
}

type paramsParserStringTypesStringTypesPatternValidation struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[string, string]
}

func (p *paramsParserStringTypesStringTypesPatternValidation) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*StringTypesStringTypesPatternValidationRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &StringTypesStringTypesPatternValidationRequest{}
	// path params
	p.bindUnformattedStr(&bindingCtx, readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(&bindingCtx, readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindCustomFormatStrInQuery(&bindingCtx, readPathValue("customFormatStrInQuery", router, req), &reqParams.CustomFormatStrInQuery)
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesPatternValidation() *paramsParserStringTypesStringTypesPatternValidation {
	return &paramsParserStringTypesStringTypesPatternValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field: "unformattedStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStrInQuery",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
	}
}

type paramsParserStringTypesStringTypesRangeValidation struct {
	bindUnformattedStr requestParamBinder[string, string]
	bindCustomFormatStr requestParamBinder[string, string]
	bindByteStr requestParamBinder[string, string]
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[string, string]
	bindByteStrInQuery requestParamBinder[[]string, string]
}

func (p *paramsParserStringTypesStringTypesRangeValidation) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*StringTypesStringTypesRangeValidationRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &StringTypesStringTypesRangeValidationRequest{}
	// path params
	p.bindUnformattedStr(&bindingCtx, readPathValue("unformattedStr", router, req), &reqParams.UnformattedStr)
	p.bindCustomFormatStr(&bindingCtx, readPathValue("customFormatStr", router, req), &reqParams.CustomFormatStr)
	p.bindByteStr(&bindingCtx, readPathValue("byteStr", router, req), &reqParams.ByteStr)
	p.bindCustomFormatStrInQuery(&bindingCtx, readPathValue("customFormatStrInQuery", router, req), &reqParams.CustomFormatStrInQuery)
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindByteStrInQuery(&bindingCtx, readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRangeValidation() *paramsParserStringTypesStringTypesRangeValidation {
	return &paramsParserStringTypesStringTypesRangeValidation{
		bindUnformattedStr: newRequestParamBinder(binderParams[string, string]{
			field: "unformattedStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindCustomFormatStr: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindByteStr: newRequestParamBinder(binderParams[string, string]{
			field: "byteStr",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[string, string]{
			field: "customFormatStrInQuery",
			location: "path",
			parseValue: knownParsers.string_in_path,
			validateValue: newCompositeValidator[string, string](
				validateNonEmpty,
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
	}
}

type paramsParserStringTypesStringTypesRequiredValidation struct {
	bindUnformattedStrInQuery requestParamBinder[[]string, string]
	bindCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindDateStrInQuery requestParamBinder[[]string, string]
	bindDateTimeStrInQuery requestParamBinder[[]string, string]
	bindByteStrInQuery requestParamBinder[[]string, string]
	bindOptionalUnformattedStrInQuery requestParamBinder[[]string, string]
	bindOptionalCustomFormatStrInQuery requestParamBinder[[]string, string]
	bindOptionalDateStrInQuery requestParamBinder[[]string, string]
	bindOptionalDateTimeStrInQuery requestParamBinder[[]string, string]
	bindOptionalByteStrInQuery requestParamBinder[[]string, string]
}

func (p *paramsParserStringTypesStringTypesRequiredValidation) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*StringTypesStringTypesRequiredValidationRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &StringTypesStringTypesRequiredValidationRequest{}
	// query params
	query := req.URL.Query()
	p.bindUnformattedStrInQuery(&bindingCtx, readQueryValue("unformattedStrInQuery", query), &reqParams.UnformattedStrInQuery)
	p.bindCustomFormatStrInQuery(&bindingCtx, readQueryValue("customFormatStrInQuery", query), &reqParams.CustomFormatStrInQuery)
	p.bindDateStrInQuery(&bindingCtx, readQueryValue("dateStrInQuery", query), &reqParams.DateStrInQuery)
	p.bindDateTimeStrInQuery(&bindingCtx, readQueryValue("dateTimeStrInQuery", query), &reqParams.DateTimeStrInQuery)
	p.bindByteStrInQuery(&bindingCtx, readQueryValue("byteStrInQuery", query), &reqParams.ByteStrInQuery)
	p.bindOptionalUnformattedStrInQuery(&bindingCtx, readQueryValue("optionalUnformattedStrInQuery", query), &reqParams.OptionalUnformattedStrInQuery)
	p.bindOptionalCustomFormatStrInQuery(&bindingCtx, readQueryValue("optionalCustomFormatStrInQuery", query), &reqParams.OptionalCustomFormatStrInQuery)
	p.bindOptionalDateStrInQuery(&bindingCtx, readQueryValue("optionalDateStrInQuery", query), &reqParams.OptionalDateStrInQuery)
	p.bindOptionalDateTimeStrInQuery(&bindingCtx, readQueryValue("optionalDateTimeStrInQuery", query), &reqParams.OptionalDateTimeStrInQuery)
	p.bindOptionalByteStrInQuery(&bindingCtx, readQueryValue("optionalByteStrInQuery", query), &reqParams.OptionalByteStrInQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserStringTypesStringTypesRequiredValidation() *paramsParserStringTypesStringTypesRequiredValidation {
	return &paramsParserStringTypesStringTypesRequiredValidation{
		bindUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "unformattedStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "customFormatStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindDateStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "dateStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "dateTimeStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "byteStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
				validateNonEmpty,
			),
		}),
		bindOptionalUnformattedStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalUnformattedStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
			),
		}),
		bindOptionalCustomFormatStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalCustomFormatStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
			),
		}),
		bindOptionalDateStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalDateStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
			),
		}),
		bindOptionalDateTimeStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalDateTimeStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
			),
		}),
		bindOptionalByteStrInQuery: newRequestParamBinder(binderParams[[]string, string]{
			field: "optionalByteStrInQuery",
			location: "query",
			parseValue: knownParsers.string_in_query,
			validateValue: newCompositeValidator[[]string, string](
			),
		}),
	}
}
