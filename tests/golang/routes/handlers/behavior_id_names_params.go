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
type _ func() BehaviorNamesWithIDData

type paramsParserBehaviorIdNamesBehaviorNamesWithID struct {
	bindID requestParamBinder[string, string]
	bindEndsWithID requestParamBinder[string, string]
	bindTheIDInTheMiddle requestParamBinder[string, string]
	bindQueryEndsWithID requestParamBinder[[]string, string]
	bindQueryTheIDInTheMiddle requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *BehaviorNamesWithIDData]
}

func (p *paramsParserBehaviorIdNamesBehaviorNamesWithID) parse(router httpRouter, req *http.Request) (*BehaviorIDNamesBehaviorNamesWithIDRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BehaviorIDNamesBehaviorNamesWithIDRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindID(pathParamsCtx.Fork("id"), readPathValue("id", router, req), &reqParams.ID)
	p.bindEndsWithID(pathParamsCtx.Fork("endsWithId"), readPathValue("endsWithId", router, req), &reqParams.EndsWithID)
	p.bindTheIDInTheMiddle(pathParamsCtx.Fork("theIdInTheMiddle"), readPathValue("theIdInTheMiddle", router, req), &reqParams.TheIDInTheMiddle)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindQueryEndsWithID(queryParamsCtx.Fork("queryEndsWithId"), readQueryValue("queryEndsWithId", query), &reqParams.QueryEndsWithID)
	p.bindQueryTheIDInTheMiddle(queryParamsCtx.Fork("queryTheIdInTheMiddle"), readQueryValue("queryTheIdInTheMiddle", query), &reqParams.QueryTheIDInTheMiddle)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorIdNamesBehaviorNamesWithID(rootHandler *RootHandler) paramsParser[*BehaviorIDNamesBehaviorNamesWithIDRequest] {
	return &paramsParserBehaviorIdNamesBehaviorNamesWithID{
		bindID: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindEndsWithID: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindTheIDInTheMiddle: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindQueryEndsWithID: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindQueryTheIDInTheMiddle: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BehaviorNamesWithIDData]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BehaviorNamesWithIDData],
			),
			validateValue: NewBehaviorNamesWithIDDataValidator(),
		}),
	}
}
