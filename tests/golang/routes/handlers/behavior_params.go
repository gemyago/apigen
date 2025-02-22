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
type _ func() BehaviorNoParamsWithResponse202Response

type paramsParserBehaviorBehaviorWithParamsAndResponse struct {
	bindQueryParam1 requestParamBinder[[]string, string]
	bindQueryParam2 requestParamBinder[[]string, int32]
	bindPayload requestParamBinder[*http.Request, *BehaviorWithParamsAndResponseRequestBody]
}

func (p *paramsParserBehaviorBehaviorWithParamsAndResponse) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorWithParamsAndResponseRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BehaviorBehaviorWithParamsAndResponseRequest{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindQueryParam1(queryParamsCtx.Fork("queryParam1"), readQueryValue("queryParam1", query), &reqParams.QueryParam1)
	p.bindQueryParam2(queryParamsCtx.Fork("queryParam2"), readQueryValue("queryParam2", query), &reqParams.QueryParam2)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorWithParamsAndResponse(rootHandler *RootHandler) paramsParser[*BehaviorBehaviorWithParamsAndResponseRequest] {
	return &paramsParserBehaviorBehaviorWithParamsAndResponse{
		bindQueryParam1: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
		bindQueryParam2: newRequestParamBinder(binderParams[[]string, int32]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.int32Parser,
			),
			validateValue: NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](0, false, true),
				NewMinMaxValueValidator[int32](5000, false, false),
			),
		}),
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *BehaviorWithParamsAndResponseRequestBody]{
			required: false,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*BehaviorWithParamsAndResponseRequestBody],
			),
			validateValue: NewBehaviorWithParamsAndResponseRequestBodyValidator(),
		}),
	}
}

type paramsParserBehaviorBehaviorWithParamsNoResponse struct {
	bindQueryParam1 requestParamBinder[[]string, string]
}

func (p *paramsParserBehaviorBehaviorWithParamsNoResponse) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorWithParamsNoResponseRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &BehaviorBehaviorWithParamsNoResponseRequest{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindQueryParam1(queryParamsCtx.Fork("queryParam1"), readQueryValue("queryParam1", query), &reqParams.QueryParam1)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorWithParamsNoResponse(rootHandler *RootHandler) paramsParser[*BehaviorBehaviorWithParamsNoResponseRequest] {
	return &paramsParserBehaviorBehaviorWithParamsNoResponse{
		bindQueryParam1: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				rootHandler.knownParsers.stringParser,
			),
			validateValue: NewSimpleFieldValidator[string](
			),
		}),
	}
}


