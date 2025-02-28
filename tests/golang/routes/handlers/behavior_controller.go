package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() BehaviorNoParamsWithResponse202Response

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
	// Request type: BehaviorWithParamsAndResponseParams,
	//
	// Response type: BehaviorWithParamsAndResponseResponseBody
	BehaviorWithParamsAndResponse(HandlerBuilder[
		*BehaviorWithParamsAndResponseParams,
		*BehaviorWithParamsAndResponseResponseBody,
	]) http.Handler

	// GET /behavior/with-params-no-response
	//
	// Request type: BehaviorWithParamsNoResponseParams,
	//
	// Response type: none
	BehaviorWithParamsNoResponse(NoResponseHandlerBuilder[
		*BehaviorWithParamsNoResponseParams,
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
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterBehaviorRoutes(controller BehaviorController) *RootHandler {
	builder := newBehaviorControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(builder.BehaviorNoParamsNoResponse))
	rootHandler.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(builder.BehaviorNoParamsWithResponse))
	rootHandler.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(builder.BehaviorNoStatusDefined))
	rootHandler.router.HandleRoute("POST", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(builder.BehaviorWithParamsAndResponse))
	rootHandler.router.HandleRoute("GET", "/behavior/with-params-no-response", controller.BehaviorWithParamsNoResponse(builder.BehaviorWithParamsNoResponse))
	rootHandler.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(builder.BehaviorWithStatusDefined))
	return rootHandler
}