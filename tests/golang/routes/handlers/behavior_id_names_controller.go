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
type _ func() BehaviorNamesWithIDData

// BehaviorIDNamesBehaviorNamesWithIDRequest represents params for behaviorNamesWithId operation
//
// Request: POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}.
type BehaviorIDNamesBehaviorNamesWithIDRequest struct {
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
	Payload *BehaviorNamesWithIDData
}

type behaviorIDNamesControllerBuilder struct {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIDNamesBehaviorNamesWithIDRequest,
	//
	// Response type: BehaviorNamesWithIDData
	BehaviorNamesWithID genericHandlerBuilder[
		*BehaviorIDNamesBehaviorNamesWithIDRequest,
		*BehaviorNamesWithIDData,
		func(context.Context, *BehaviorIDNamesBehaviorNamesWithIDRequest) (*BehaviorNamesWithIDData, error),
		func(http.ResponseWriter, *http.Request, *BehaviorIDNamesBehaviorNamesWithIDRequest) (*BehaviorNamesWithIDData, error),
	]
}

func newBehaviorIDNamesControllerBuilder(app *RootHandler) *behaviorIDNamesControllerBuilder {
	return &behaviorIDNamesControllerBuilder{
		// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
		BehaviorNamesWithID: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*BehaviorIDNamesBehaviorNamesWithIDRequest,
				*BehaviorNamesWithIDData,
			](),
			newHTTPHandlerAdapter[
				*BehaviorIDNamesBehaviorNamesWithIDRequest,
				*BehaviorNamesWithIDData,
			](),
			makeActionBuilderParams[
				*BehaviorIDNamesBehaviorNamesWithIDRequest,
				*BehaviorNamesWithIDData,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserBehaviorIDNamesBehaviorNamesWithID(app),
			},
		),
	}
}

type BehaviorIDNamesController interface {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIDNamesBehaviorNamesWithIDRequest,
	//
	// Response type: BehaviorNamesWithIDData
	BehaviorNamesWithID(HandlerBuilder[
		*BehaviorIDNamesBehaviorNamesWithIDRequest,
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