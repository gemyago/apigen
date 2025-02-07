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
	// QueryParam2 is parsed from request query and declared as queryParam2.
	QueryParam2 int32
}

// BehaviorBehaviorWithParamsNoResponseRequest represents params for behaviorWithParamsNoResponse operation
//
// Request: GET /behavior/with-params-no-response.
type BehaviorBehaviorWithParamsNoResponseRequest struct {
	// QueryParam1 is parsed from request query and declared as queryParam1.
	QueryParam1 string
}



type BehaviorControllerBuilderV2 struct {
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

	// GET /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: BehaviorNoParamsWithResponse202Response
	BehaviorWithParamsAndResponse ActionBuilder[
		*BehaviorBehaviorWithParamsAndResponseRequest,
		*BehaviorNoParamsWithResponse202Response,
		func(context.Context, *BehaviorBehaviorWithParamsAndResponseRequest) (*BehaviorNoParamsWithResponse202Response, error),
		func(http.ResponseWriter, *http.Request, *BehaviorBehaviorWithParamsAndResponseRequest) (*BehaviorNoParamsWithResponse202Response, error),
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

func newBehaviorControllerBuilderV2(app *HTTPApp) *BehaviorControllerBuilderV2 {
	return &BehaviorControllerBuilderV2{
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
				paramsParser:  makeVoidParamsParserV2(app),
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
				paramsParser:  makeVoidParamsParserV2(app),
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
				paramsParser:  makeVoidParamsParserV2(app),
			},
		),

		// GET /behavior/with-params-and-response
		BehaviorWithParamsAndResponse: makeActionBuilder(
			app,
			newHandlerAdapter[
				*BehaviorBehaviorWithParamsAndResponseRequest,
				*BehaviorNoParamsWithResponse202Response,
			](),
			newHTTPHandlerAdapter[
				*BehaviorBehaviorWithParamsAndResponseRequest,
				*BehaviorNoParamsWithResponse202Response,
			](),
			makeActionBuilderParams[
				*BehaviorBehaviorWithParamsAndResponseRequest,
				*BehaviorNoParamsWithResponse202Response,
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

	// GET /behavior/with-params-no-response
	//
	// Request type: BehaviorBehaviorWithParamsNoResponseRequest,
	//
	// Response type: none
	BehaviorWithParamsNoResponse(b *BehaviorControllerBuilderV2) http.Handler

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
	app.router.HandleRoute("GET", "/behavior/with-params-no-response", controller.BehaviorWithParamsNoResponse(builder))
	app.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(builder))
}