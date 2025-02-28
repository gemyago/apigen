package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() Error

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