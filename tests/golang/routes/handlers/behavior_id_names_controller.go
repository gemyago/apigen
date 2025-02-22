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
type _ func() BehaviorNamesWithIdData

// BehaviorIdNamesBehaviorNamesWithIDRequest represents params for behaviorNamesWithId operation
//
// Request: POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}.
type BehaviorIdNamesBehaviorNamesWithIDRequest struct {
	// ID is parsed from request path and declared as id.
	ID string
	// EndsWithID is parsed from request path and declared as endsWithId.
	EndsWithID string
	// TheIDInTheMiddle is parsed from request path and declared as theIdInTheMiddle.
	TheIDInTheMiddle string
	// QueryEndsWithID is parsed from request query and declared as queryEndsWithId.
	QueryEndsWithID string
	// QueryTheIDInTheMiddle is parsed from request query and declared as queryTheIdInTheMiddle.
	QueryTheIDInTheMiddle string
	// Payload is parsed from request body and declared as payload.
	Payload *BehaviorNamesWithIdData
}

type behaviorIDNamesControllerBuilder struct {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIdNamesBehaviorNamesWithIDRequest,
	//
	// Response type: BehaviorNamesWithIdData
	BehaviorNamesWithID genericHandlerBuilder[
		*BehaviorIdNamesBehaviorNamesWithIDRequest,
		*BehaviorNamesWithIdData,
		func(context.Context, *BehaviorIdNamesBehaviorNamesWithIDRequest) (*BehaviorNamesWithIdData, error),
		func(http.ResponseWriter, *http.Request, *BehaviorIdNamesBehaviorNamesWithIDRequest) (*BehaviorNamesWithIdData, error),
	]
}

func newBehaviorIdNamesControllerBuilder(app *RootHandler) *behaviorIDNamesControllerBuilder {
	return &behaviorIDNamesControllerBuilder{
		// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
		BehaviorNamesWithID: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*BehaviorIdNamesBehaviorNamesWithIDRequest,
				*BehaviorNamesWithIdData,
			](),
			newHTTPHandlerAdapter[
				*BehaviorIdNamesBehaviorNamesWithIDRequest,
				*BehaviorNamesWithIdData,
			](),
			makeActionBuilderParams[
				*BehaviorIdNamesBehaviorNamesWithIDRequest,
				*BehaviorNamesWithIdData,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserBehaviorIdNamesBehaviorNamesWithID(app),
			},
		),
	}
}

type BehaviorIdNamesController interface {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIdNamesBehaviorNamesWithIDRequest,
	//
	// Response type: BehaviorNamesWithIdData
	BehaviorNamesWithID(HandlerBuilder[
		*BehaviorIdNamesBehaviorNamesWithIDRequest,
		*BehaviorNamesWithIdData,
	]) http.Handler
}

// RegisterBehaviorIdNamesRoutes will attach the following routes to the root handler:
// 
// - POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterBehaviorIdNamesRoutes(controller BehaviorIdNamesController) *RootHandler {
	builder := newBehaviorIdNamesControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}", controller.BehaviorNamesWithID(builder.BehaviorNamesWithID))
	return rootHandler
}