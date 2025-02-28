package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

type BehaviorIDNamesController interface {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIDNamesBehaviorNamesWithIDParams,
	//
	// Response type: BehaviorNamesWithIDData
	BehaviorNamesWithID(HandlerBuilder[
		*BehaviorIDNamesBehaviorNamesWithIDParams,
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