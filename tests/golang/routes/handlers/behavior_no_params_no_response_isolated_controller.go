package handlers

import (
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
type _ func() Error



type behaviorNoParamsNoResponseIsolatedControllerBuilder struct {
	// GET /behavior/no-params-no-response-isolated
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
}

func newBehaviorNoParamsNoResponseIsolatedControllerBuilder(app *RootHandler) *behaviorNoParamsNoResponseIsolatedControllerBuilder {
	return &behaviorNoParamsNoResponseIsolatedControllerBuilder{
		// GET /behavior/no-params-no-response-isolated
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
	}
}

type BehaviorNoParamsNoResponseIsolatedController interface {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse(NoParamsNoResponseHandlerBuilder) http.Handler
}

// RegisterBehaviorNoParamsNoResponseIsolatedRoutes will attach the following routes to the root handler:
// 
// - GET /behavior/no-params-no-response-isolated
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterBehaviorNoParamsNoResponseIsolatedRoutes(controller BehaviorNoParamsNoResponseIsolatedController) *RootHandler {
	builder := newBehaviorNoParamsNoResponseIsolatedControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("GET", "/behavior/no-params-no-response-isolated", controller.BehaviorNoParamsNoResponse(builder.BehaviorNoParamsNoResponse))
	return rootHandler
}