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
// Request: GET /behavior/with-params-and-response.
type BehaviorBehaviorWithParamsAndResponseRequest struct {
	// QueryParam1 is parsed from request query and declared as queryParam1.
	QueryParam1 string
}



type BehaviorController struct {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse httpHandlerFactory

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse httpHandlerFactory

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined httpHandlerFactory

	// GET /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorWithParamsAndResponse httpHandlerFactory

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined httpHandlerFactory
}

type BehaviorControllerBuilder struct {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorNoParamsNoResponse actionBuilderNoParamsVoidResult[*BehaviorControllerBuilder]

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	HandleBehaviorNoParamsWithResponse actionBuilderNoParams[*BehaviorControllerBuilder, *BehaviorNoParamsWithResponse202Response]

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorNoStatusDefined actionBuilderNoParamsVoidResult[*BehaviorControllerBuilder]

	// GET /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	HandleBehaviorWithParamsAndResponse actionBuilder[*BehaviorControllerBuilder, *BehaviorBehaviorWithParamsAndResponseRequest, *BehaviorNoParamsWithResponse202Response]

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorWithStatusDefined actionBuilderNoParamsVoidResult[*BehaviorControllerBuilder]
}

func (c *BehaviorControllerBuilder) Finalize() *BehaviorController {
	return &BehaviorController{
		BehaviorNoParamsNoResponse: mustInitializeAction("behaviorNoParamsNoResponse", c.HandleBehaviorNoParamsNoResponse.httpHandlerFactory),
		BehaviorNoParamsWithResponse: mustInitializeAction("behaviorNoParamsWithResponse", c.HandleBehaviorNoParamsWithResponse.httpHandlerFactory),
		BehaviorNoStatusDefined: mustInitializeAction("behaviorNoStatusDefined", c.HandleBehaviorNoStatusDefined.httpHandlerFactory),
		BehaviorWithParamsAndResponse: mustInitializeAction("behaviorWithParamsAndResponse", c.HandleBehaviorWithParamsAndResponse.httpHandlerFactory),
		BehaviorWithStatusDefined: mustInitializeAction("behaviorWithStatusDefined", c.HandleBehaviorWithStatusDefined.httpHandlerFactory),
	}
}

func BuildBehaviorController() *BehaviorControllerBuilder {
	controllerBuilder := &BehaviorControllerBuilder{}

	// GET /behavior/no-params-no-response
	controllerBuilder.HandleBehaviorNoParamsNoResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoParamsNoResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorNoParamsNoResponse.voidResult = true
	controllerBuilder.HandleBehaviorNoParamsNoResponse.paramsParserFactory = makeVoidParamsParser

	// GET /behavior/no-params-with-response
	controllerBuilder.HandleBehaviorNoParamsWithResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoParamsWithResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorNoParamsWithResponse.paramsParserFactory = makeVoidParamsParser

	// GET /behavior/no-status-defined
	controllerBuilder.HandleBehaviorNoStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoStatusDefined.defaultStatusCode = 200
	controllerBuilder.HandleBehaviorNoStatusDefined.voidResult = true
	controllerBuilder.HandleBehaviorNoStatusDefined.paramsParserFactory = makeVoidParamsParser

	// GET /behavior/with-params-and-response
	controllerBuilder.HandleBehaviorWithParamsAndResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorWithParamsAndResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorWithParamsAndResponse.paramsParserFactory = newParamsParserBehaviorBehaviorWithParamsAndResponse

	// POST /behavior/with-status-defined
	controllerBuilder.HandleBehaviorWithStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorWithStatusDefined.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorWithStatusDefined.voidResult = true
	controllerBuilder.HandleBehaviorWithStatusDefined.paramsParserFactory = makeVoidParamsParser

	return controllerBuilder
}

func RegisterBehaviorRoutes(controller *BehaviorController, app *HTTPApp) {
	app.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(app))
	app.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(app))
	app.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(app))
	app.router.HandleRoute("GET", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(app))
	app.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(app))
}

type BehaviorControllerBuilderV2 struct {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse ActionBuilder[
	  Void,
	  Void,
	  func(context.Context) (error),
	  func(http.ResponseWriter, *http.Request) (error),
	]

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse ActionBuilder[
	  Void,
	  BehaviorNoParamsWithResponse202Response,
	  func(context.Context) (*BehaviorNoParamsWithResponse202Response, error),
	  func(http.ResponseWriter, *http.Request) (*BehaviorNoParamsWithResponse202Response, error),
	]

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined ActionBuilder[
	  Void,
	  Void,
	  func(context.Context) (error),
	  func(http.ResponseWriter, *http.Request) (error),
	]

	// GET /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorWithParamsAndResponse ActionBuilder[
	  BehaviorBehaviorWithParamsAndResponseRequest,
	  BehaviorNoParamsWithResponse202Response,
	  func(context.Context, *BehaviorBehaviorWithParamsAndResponseRequest) (*BehaviorNoParamsWithResponse202Response, error),
	  func(http.ResponseWriter, *http.Request, *BehaviorBehaviorWithParamsAndResponseRequest) (*BehaviorNoParamsWithResponse202Response, error),
	]

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined ActionBuilder[
	  Void,
	  Void,
	  func(context.Context) (error),
	  func(http.ResponseWriter, *http.Request) (error),
	]
}

func newBehaviorControllerBuilderV2(app *HTTPApp) *BehaviorControllerBuilderV2 {
	return &BehaviorControllerBuilderV2{
		// GET /behavior/no-params-no-response
		BehaviorNoParamsNoResponse: makeActionBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			makeActionBuilderParams[
				Void,
				Void,
			]{
				defaultStatus: 202,
				voidResult:    true,
				paramsParser:  makeVoidParamsParserV2(app),
			},
		),

		// GET /behavior/no-params-with-response
		BehaviorNoParamsWithResponse: makeActionBuilder(
			app,
			newHandlerAdapterNoParams[
				Void,
				BehaviorNoParamsWithResponse202Response,
			](),
			newHTTPHandlerAdapterNoParams[
				Void,
				BehaviorNoParamsWithResponse202Response,
			](),
			makeActionBuilderParams[
				Void,
				BehaviorNoParamsWithResponse202Response,
			]{
				defaultStatus: 202,
				paramsParser:  makeVoidParamsParserV2(app),
			},
		),

		// GET /behavior/no-status-defined
		BehaviorNoStatusDefined: makeActionBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			makeActionBuilderParams[
				Void,
				Void,
			]{
				defaultStatus: 200,
				voidResult:    true,
				paramsParser:  makeVoidParamsParserV2(app),
			},
		),

		// GET /behavior/with-params-and-response
		BehaviorWithParamsAndResponse: makeActionBuilder(
			app,
			newHandlerAdapter[
				BehaviorBehaviorWithParamsAndResponseRequest,
				BehaviorNoParamsWithResponse202Response,
			](),
			newHTTPHandlerAdapter[
				BehaviorBehaviorWithParamsAndResponseRequest,
				BehaviorNoParamsWithResponse202Response,
			](),
			makeActionBuilderParams[
				BehaviorBehaviorWithParamsAndResponseRequest,
				BehaviorNoParamsWithResponse202Response,
			]{
				defaultStatus: 202,
				paramsParser:  newParamsParserBehaviorBehaviorWithParamsAndResponse(app),
			},
		),

		// POST /behavior/with-status-defined
		BehaviorWithStatusDefined: makeActionBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			makeActionBuilderParams[
				Void,
				Void,
			]{
				defaultStatus: 202,
				voidResult:    true,
				paramsParser:  makeVoidParamsParserV2(app),
			},
		),
	}
}

type BehaviorControllerV2 interface {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse(b *BehaviorControllerBuilderV2) http.Handler

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse(b *BehaviorControllerBuilderV2) http.Handler

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined(b *BehaviorControllerBuilderV2) http.Handler

	// GET /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorWithParamsAndResponse(b *BehaviorControllerBuilderV2) http.Handler

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined(b *BehaviorControllerBuilderV2) http.Handler
}

func RegisterBehaviorRoutesV2(controller BehaviorControllerV2, app *HTTPApp) {
	builder := newBehaviorControllerBuilderV2(app)
	app.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(builder))
	app.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(builder))
	app.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(builder))
	app.router.HandleRoute("GET", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(builder))
	app.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(builder))
}