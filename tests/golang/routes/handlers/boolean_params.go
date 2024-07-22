package handlers

import (
	"net/http"
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type paramsParserBooleanBooleanParsing struct {
	bindBoolParam1 requestParamBinder[string, bool]
	bindBoolParam2 requestParamBinder[string, bool]
	bindBoolParam1InQuery requestParamBinder[[]string, bool]
	bindBoolParam2InQuery requestParamBinder[[]string, bool]
}

func (p *paramsParserBooleanBooleanParsing) parse(router httpRouter, req *http.Request) (*BooleanBooleanParsingRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &BooleanBooleanParsingRequest{}
	// path params
	p.bindBoolParam1(&bindingCtx, readPathValue("boolParam1", router, req), &reqParams.BoolParam1)
	p.bindBoolParam2(&bindingCtx, readPathValue("boolParam2", router, req), &reqParams.BoolParam2)
	// query params
	query := req.URL.Query()
	p.bindBoolParam1InQuery(&bindingCtx, readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(&bindingCtx, readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanParsing(app *HTTPApp) paramsParser[*BooleanBooleanParsingRequest] {
	return &paramsParserBooleanBooleanParsing{
		bindBoolParam1: newRequestParamBinder(binderParams[string, bool]{
			field: "boolParam1",
			location: "path",
			required: true,
			parseValue: app.knownParsers.boolInPath,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2: newRequestParamBinder(binderParams[string, bool]{
			field: "boolParam2",
			location: "path",
			required: true,
			parseValue: app.knownParsers.boolInPath,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam1InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.boolInQuery,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam2InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.boolInQuery,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
	}
}

type paramsParserBooleanBooleanRequiredValidation struct {
	bindBoolParam1InQuery requestParamBinder[[]string, bool]
	bindBoolParam2InQuery requestParamBinder[[]string, bool]
	bindOptionalBoolParam1InQuery requestParamBinder[[]string, bool]
	bindOptionalBoolParam2InQuery requestParamBinder[[]string, bool]
}

func (p *paramsParserBooleanBooleanRequiredValidation) parse(router httpRouter, req *http.Request) (*BooleanBooleanRequiredValidationRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &BooleanBooleanRequiredValidationRequest{}
	// query params
	query := req.URL.Query()
	p.bindBoolParam1InQuery(&bindingCtx, readQueryValue("boolParam1InQuery", query), &reqParams.BoolParam1InQuery)
	p.bindBoolParam2InQuery(&bindingCtx, readQueryValue("boolParam2InQuery", query), &reqParams.BoolParam2InQuery)
	p.bindOptionalBoolParam1InQuery(&bindingCtx, readQueryValue("optionalBoolParam1InQuery", query), &reqParams.OptionalBoolParam1InQuery)
	p.bindOptionalBoolParam2InQuery(&bindingCtx, readQueryValue("optionalBoolParam2InQuery", query), &reqParams.OptionalBoolParam2InQuery)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBooleanBooleanRequiredValidation(app *HTTPApp) paramsParser[*BooleanBooleanRequiredValidationRequest] {
	return &paramsParserBooleanBooleanRequiredValidation{
		bindBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam1InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.boolInQuery,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "boolParam2InQuery",
			location: "query",
			required: true,
			parseValue: app.knownParsers.boolInQuery,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindOptionalBoolParam1InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "optionalBoolParam1InQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.boolInQuery,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
		bindOptionalBoolParam2InQuery: newRequestParamBinder(binderParams[[]string, bool]{
			field: "optionalBoolParam2InQuery",
			location: "query",
			required: false,
			parseValue: app.knownParsers.boolInQuery,
			validateValue: internal.NewSimpleFieldValidator[bool](
			),
		}),
	}
}
