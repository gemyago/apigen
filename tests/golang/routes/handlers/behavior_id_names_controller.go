package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() BehaviorNamesWithIDData

type BehaviorIDNamesController interface {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorNamesWithIDParams,
	//
	// Response type: BehaviorNamesWithIDData
	BehaviorNamesWithID(HandlerBuilder[
		*BehaviorNamesWithIDParams,
		*BehaviorNamesWithIDData,
	]) http.Handler
}

// RegisterBehaviorIDNamesRoutes will attach the following routes to the root handler:
// 
// - POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterBehaviorIDNamesRoutes(controller BehaviorIDNamesController) *RootHandler {
	builder := newBehaviorIDNamesControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}", controller.BehaviorNamesWithID(builder.BehaviorNamesWithID))
	return rootHandler
}