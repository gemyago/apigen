package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint
type _ func() BehaviorNoParamsWithResponse202Response







// BehaviorBehaviorWithParamsAndResponseRequest represents params for behaviorWithParamsAndResponse operation
//
// Request: POST /behavior/with-params-and-response.
type BehaviorBehaviorWithParamsAndResponseRequest struct {
	// QueryParam1 is parsed from request query and declared as queryParam1.
	QueryParam1 string
	// QueryParam2 is parsed from request query and declared as queryParam2.
	QueryParam2 int32
	// Payload is parsed from request body and declared as payload.
	Payload *BehaviorWithParamsAndResponseRequestBody
}

// BehaviorBehaviorWithParamsNoResponseRequest represents params for behaviorWithParamsNoResponse operation
//
// Request: GET /behavior/with-params-no-response.
type BehaviorBehaviorWithParamsNoResponseRequest struct {
	// QueryParam1 is parsed from request query and declared as queryParam1.
	QueryParam1 string
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
		func(context.Context) (error),
		func(http.ResponseWriter, *http.Request) (error),
	]

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse genericHandlerBuilder[
		void,
		*BehaviorNoParamsWithResponse202Response,
		func(context.Context) (*BehaviorNoParamsWithResponse202Response, error),
		func(http.ResponseWriter, *http.Request) (*BehaviorNoParamsWithResponse202Response, error),
	]

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined genericHandlerBuilder[
		void,
		void,
		func(context.Context) (error),
		func(http.ResponseWriter, *http.Request) (error),
	]

	// POST /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorWithParamsAndResponseResponseBody
	BehaviorWithParamsAndResponse genericHandlerBuilder[
		*BehaviorBehaviorWithParamsAndResponseRequest,
		*BehaviorWithParamsAndResponseResponseBody,
		func(context.Context, *BehaviorBehaviorWithParamsAndResponseRequest) (*BehaviorWithParamsAndResponseResponseBody, error),
		func(http.ResponseWriter, *http.Request, *BehaviorBehaviorWithParamsAndResponseRequest) (*BehaviorWithParamsAndResponseResponseBody, error),
	]

	// GET /behavior/with-params-no-response
	//
	// Request type: BehaviorBehaviorWithParamsNoResponseRequest,
	//
	// Response type: none
	BehaviorWithParamsNoResponse genericHandlerBuilder[
		*BehaviorBehaviorWithParamsNoResponseRequest,
		void,
		func(context.Context, *BehaviorBehaviorWithParamsNoResponseRequest) (error),
		func(http.ResponseWriter, *http.Request, *BehaviorBehaviorWithParamsNoResponseRequest) (error),
	]

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined genericHandlerBuilder[
		void,
		void,
		func(context.Context) (error),
		func(http.ResponseWriter, *http.Request) (error),
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
				*BehaviorBehaviorWithParamsAndResponseRequest,
				*BehaviorWithParamsAndResponseResponseBody,
			](),
			newHTTPHandlerAdapter[
				*BehaviorBehaviorWithParamsAndResponseRequest,
				*BehaviorWithParamsAndResponseResponseBody,
			](),
			makeActionBuilderParams[
				*BehaviorBehaviorWithParamsAndResponseRequest,
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
				*BehaviorBehaviorWithParamsNoResponseRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BehaviorBehaviorWithParamsNoResponseRequest,
				void,
			](),
			makeActionBuilderParams[
				*BehaviorBehaviorWithParamsNoResponseRequest,
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

type BehaviorController interface {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse(NoParamsNoResponseHandlerBuilder) http.Handler

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse(NoParamsHandlerBuilder[
		*BehaviorNoParamsWithResponse202Response,
	]) http.Handler

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined(NoParamsNoResponseHandlerBuilder) http.Handler

	// POST /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorWithParamsAndResponseResponseBody
	BehaviorWithParamsAndResponse(HandlerBuilder[
		*BehaviorBehaviorWithParamsAndResponseRequest,
		*BehaviorWithParamsAndResponseResponseBody,
	]) http.Handler

	// GET /behavior/with-params-no-response
	//
	// Request type: BehaviorBehaviorWithParamsNoResponseRequest,
	//
	// Response type: none
	BehaviorWithParamsNoResponse(NoResponseHandlerBuilder[
		*BehaviorBehaviorWithParamsNoResponseRequest,
	]) http.Handler

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined(NoParamsNoResponseHandlerBuilder) http.Handler
}

// RegisterBehaviorRoutes will attach the following routes to the root handler:
// 
// - GET /behavior/no-params-no-response
// 
// - GET /behavior/no-params-with-response
// 
// - GET /behavior/no-status-defined
// 
// - POST /behavior/with-params-and-response
// 
// - GET /behavior/with-params-no-response
// 
// - POST /behavior/with-status-defined
// 
func RegisterBehaviorRoutes(rootHandler *RootHandler, controller BehaviorController) {
	builder := newBehaviorControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(builder.BehaviorNoParamsNoResponse))
	rootHandler.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(builder.BehaviorNoParamsWithResponse))
	rootHandler.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(builder.BehaviorNoStatusDefined))
	rootHandler.router.HandleRoute("POST", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(builder.BehaviorWithParamsAndResponse))
	rootHandler.router.HandleRoute("GET", "/behavior/with-params-no-response", controller.BehaviorWithParamsNoResponse(builder.BehaviorWithParamsNoResponse))
	rootHandler.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(builder.BehaviorWithStatusDefined))
}