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

func (p *paramsParserBehaviorBehaviorWithParamsAndResponse) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorWithParamsAndResponseParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &BehaviorBehaviorWithParamsAndResponseParams{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindQueryParam1(queryParamsCtx.Fork("queryParam1"), readQueryValue("queryParam1", query), &reqParams.QueryParam1)
	p.bindQueryParam2(queryParamsCtx.Fork("queryParam2"), readQueryValue("queryParam2", query), &reqParams.QueryParam2)
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorWithParamsAndResponse(rootHandler *RootHandler) paramsParser[*BehaviorBehaviorWithParamsAndResponseParams] {
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

func (p *paramsParserBehaviorBehaviorWithParamsNoResponse) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorWithParamsNoResponseParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &BehaviorBehaviorWithParamsNoResponseParams{}
	// query params
	query := req.URL.Query()
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindQueryParam1(queryParamsCtx.Fork("queryParam1"), readQueryValue("queryParam1", query), &reqParams.QueryParam1)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorWithParamsNoResponse(rootHandler *RootHandler) paramsParser[*BehaviorBehaviorWithParamsNoResponseParams] {
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



type behaviorControllerBuilder struct {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse genericHandlerBuilder[
		void,
		void,
		handlerActionFuncNoParamsNoResponse[void, void],
		httpHandlerActionFuncNoParamsNoResponse[void, void],
	]

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse genericHandlerBuilder[
		void,
		*BehaviorNoParamsWithResponse202Response,
		handlerActionFuncNoParams[void, *BehaviorNoParamsWithResponse202Response],
		httpHandlerActionFuncNoParams[void, *BehaviorNoParamsWithResponse202Response],
	]

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined genericHandlerBuilder[
		void,
		void,
		handlerActionFuncNoParamsNoResponse[void, void],
		httpHandlerActionFuncNoParamsNoResponse[void, void],
	]

	// POST /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseParams,
	//
	// Response type: BehaviorWithParamsAndResponseResponseBody
	BehaviorWithParamsAndResponse genericHandlerBuilder[
		*BehaviorBehaviorWithParamsAndResponseParams,
		*BehaviorWithParamsAndResponseResponseBody,
		handlerActionFunc[*BehaviorBehaviorWithParamsAndResponseParams, *BehaviorWithParamsAndResponseResponseBody],
		httpHandlerActionFunc[*BehaviorBehaviorWithParamsAndResponseParams, *BehaviorWithParamsAndResponseResponseBody],
	]

	// GET /behavior/with-params-no-response
	//
	// Request type: BehaviorBehaviorWithParamsNoResponseParams,
	//
	// Response type: none
	BehaviorWithParamsNoResponse genericHandlerBuilder[
		*BehaviorBehaviorWithParamsNoResponseParams,
		void,
		handlerActionFuncNoResponse[*BehaviorBehaviorWithParamsNoResponseParams, void],
		httpHandlerActionFuncNoResponse[*BehaviorBehaviorWithParamsNoResponseParams, void],
	]

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined genericHandlerBuilder[
		void,
		void,
		handlerActionFuncNoParamsNoResponse[void, void],
		httpHandlerActionFuncNoParamsNoResponse[void, void],
	]
}

func newBehaviorControllerBuilder(app *RootHandler) *behaviorControllerBuilder {
	return &behaviorControllerBuilder{
		// GET /behavior/no-params-no-response
		BehaviorNoParamsNoResponse: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			makeActionBuilderParams[
				void,
				void,
			]{
				defaultStatus: 202,
				voidResult:    true,
				paramsParser:  makeVoidParamsParser(app),
			},
		),

		// GET /behavior/no-params-with-response
		BehaviorNoParamsWithResponse: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoParams[
				void,
				*BehaviorNoParamsWithResponse202Response,
			](),
			newHTTPHandlerAdapterNoParams[
				void,
				*BehaviorNoParamsWithResponse202Response,
			](),
			makeActionBuilderParams[
				void,
				*BehaviorNoParamsWithResponse202Response,
			]{
				defaultStatus: 202,
				paramsParser:  makeVoidParamsParser(app),
			},
		),

		// GET /behavior/no-status-defined
		BehaviorNoStatusDefined: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			makeActionBuilderParams[
				void,
				void,
			]{
				defaultStatus: 200,
				voidResult:    true,
				paramsParser:  makeVoidParamsParser(app),
			},
		),

		// POST /behavior/with-params-and-response
		BehaviorWithParamsAndResponse: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*BehaviorBehaviorWithParamsAndResponseParams,
				*BehaviorWithParamsAndResponseResponseBody,
			](),
			newHTTPHandlerAdapter[
				*BehaviorBehaviorWithParamsAndResponseParams,
				*BehaviorWithParamsAndResponseResponseBody,
			](),
			makeActionBuilderParams[
				*BehaviorBehaviorWithParamsAndResponseParams,
				*BehaviorWithParamsAndResponseResponseBody,
			]{
				defaultStatus: 202,
				paramsParser:  newParamsParserBehaviorBehaviorWithParamsAndResponse(app),
			},
		),

		// GET /behavior/with-params-no-response
		BehaviorWithParamsNoResponse: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*BehaviorBehaviorWithParamsNoResponseParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BehaviorBehaviorWithParamsNoResponseParams,
				void,
			](),
			makeActionBuilderParams[
				*BehaviorBehaviorWithParamsNoResponseParams,
				void,
			]{
				defaultStatus: 202,
				voidResult:    true,
				paramsParser:  newParamsParserBehaviorBehaviorWithParamsNoResponse(app),
			},
		),

		// POST /behavior/with-status-defined
		BehaviorWithStatusDefined: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			makeActionBuilderParams[
				void,
				void,
			]{
				defaultStatus: 202,
				voidResult:    true,
				paramsParser:  makeVoidParamsParser(app),
			},
		),
	}
}
