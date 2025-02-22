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
type _ func() BehaviorNamesWithIdData

type paramsParserBehaviorIdNamesBehaviorNamesWithId struct {
	bindId requestParamBinder[string, string]
	bindEndsWithId requestParamBinder[string, string]
	bindTheIdInTheMiddle requestParamBinder[string, string]
	bindQueryEndsWithId requestParamBinder[[]string, string]
	bindQueryTheIdInTheMiddle requestParamBinder[[]string, string]
	bindPayload requestParamBinder[*http.Request, *BehaviorNamesWithIdData]
}

func (p *paramsParserBehaviorIdNamesBehaviorNamesWithId) parse(router httpRouter, req *http.Request) (*BehaviorIdNamesBehaviorNamesWithIdRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BehaviorIdNamesBehaviorNamesWithIdRequest{}
	// path params
	pathParamsCtx := bindingCtx.Fork("path")
	p.bindId(pathParamsCtx.Fork("id"), readPathValue("id", router, req), &reqParams.Id)
	p.bindEndsWithId(pathParamsCtx.Fork("endsWithId"), readPathValue("endsWithId", router, req), &reqParams.EndsWithId)
	p.bindTheIdInTheMiddle(pathParamsCtx.Fork("theIdInTheMiddle"), readPathValue("theIdInTheMiddle", router, req), &reqParams.TheIdInTheMiddle)
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindQueryEndsWithId(queryParamsCtx.Fork("queryEndsWithId"), readQueryValue("queryEndsWithId", query), &reqParams.QueryEndsWithId)
	p.bindQueryTheIdInTheMiddle(queryParamsCtx.Fork("queryTheIdInTheMiddle"), readQueryValue("queryTheIdInTheMiddle", query), &reqParams.QueryTheIdInTheMiddle)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorIdNamesBehaviorNamesWithId(rootHandler *RootHandler) paramsParser[*BehaviorIdNamesBehaviorNamesWithIdRequest] {
	return &paramsParserBehaviorIdNamesBehaviorNamesWithId{
		bindId: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindEndsWithId: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindTheIdInTheMiddle: newRequestParamBinder(binderParams[string, string]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindQueryEndsWithId: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindQueryTheIdInTheMiddle: newRequestParamBinder(binderParams[[]string, string]{
			required: true,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BehaviorNamesWithIdData]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BehaviorNamesWithIdData],
			),
			validateValue: NewBehaviorNamesWithIdDataValidator(),
		}),
	}
}
