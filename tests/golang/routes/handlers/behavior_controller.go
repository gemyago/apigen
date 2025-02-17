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



type BehaviorControllerBuilder struct {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse ActionBuilder[
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
	BehaviorNoParamsWithResponse ActionBuilder[
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
	BehaviorNoStatusDefined ActionBuilder[
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
	BehaviorWithParamsAndResponse ActionBuilder[
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
	BehaviorWithParamsNoResponse ActionBuilder[
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
	BehaviorWithStatusDefined ActionBuilder[
		void,
		void,
		func(context.Context) (error),
		func(http.ResponseWriter, *http.Request) (error),
	]
}

func newBehaviorControllerBuilder(app *HTTPApp) *BehaviorControllerBuilder {
	return &BehaviorControllerBuilder{
		// GET /behavior/no-params-no-response
		BehaviorNoParamsNoResponse: makeActionBuilder(
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
		BehaviorNoParamsWithResponse: makeActionBuilder(
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
		BehaviorNoStatusDefined: makeActionBuilder(
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
		BehaviorWithParamsAndResponse: makeActionBuilder(
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
		BehaviorWithParamsNoResponse: makeActionBuilder(
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
		BehaviorWithStatusDefined: makeActionBuilder(
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
	BehaviorNoParamsNoResponse(b *BehaviorControllerBuilder) http.Handler

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse(b *BehaviorControllerBuilder) http.Handler

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined(b *BehaviorControllerBuilder) http.Handler

	// POST /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorWithParamsAndResponseResponseBody
	BehaviorWithParamsAndResponse(b *BehaviorControllerBuilder) http.Handler

	// GET /behavior/with-params-no-response
	//
	// Request type: BehaviorBehaviorWithParamsNoResponseRequest,
	//
	// Response type: none
	BehaviorWithParamsNoResponse(b *BehaviorControllerBuilder) http.Handler

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined(b *BehaviorControllerBuilder) http.Handler
}

type BehaviorControllerV3 interface {
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

func RegisterBehaviorRoutes(controller BehaviorController, app *HTTPApp) {
	builder := newBehaviorControllerBuilder(app)
	app.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(builder))
	app.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(builder))
	app.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(builder))
	app.router.HandleRoute("POST", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(builder))
	app.router.HandleRoute("GET", "/behavior/with-params-no-response", controller.BehaviorWithParamsNoResponse(builder))
	app.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(builder))
}

func RegisterBehaviorRoutesV3(controller BehaviorControllerV3, app *HTTPApp) {
	builder := newBehaviorControllerBuilder(app)
	app.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(builder.BehaviorNoParamsNoResponse))
	app.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(builder.BehaviorNoParamsWithResponse))
	app.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(builder.BehaviorNoStatusDefined))
	app.router.HandleRoute("POST", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(builder.BehaviorWithParamsAndResponse))
	app.router.HandleRoute("GET", "/behavior/with-params-no-response", controller.BehaviorWithParamsNoResponse(builder.BehaviorWithParamsNoResponse))
	app.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(builder.BehaviorWithStatusDefined))
}