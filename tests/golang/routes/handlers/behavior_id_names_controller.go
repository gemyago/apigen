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

// BehaviorIdNamesBehaviorNamesWithIdRequest represents params for behaviorNamesWithId operation
//
// Request: POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}.
type BehaviorIdNamesBehaviorNamesWithIdRequest struct {
	// Id is parsed from request path and declared as id.
	Id string
	// EndsWithId is parsed from request path and declared as endsWithId.
	EndsWithId string
	// TheIdInTheMiddle is parsed from request path and declared as theIdInTheMiddle.
	TheIdInTheMiddle string
	// QueryEndsWithId is parsed from request query and declared as queryEndsWithId.
	QueryEndsWithId string
	// QueryTheIdInTheMiddle is parsed from request query and declared as queryTheIdInTheMiddle.
	QueryTheIdInTheMiddle string
	// Payload is parsed from request body and declared as payload.
	Payload *BehaviorNamesWithIdData
}

type behaviorIdNamesControllerBuilder struct {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIdNamesBehaviorNamesWithIdRequest,
	//
	// Response type: BehaviorNamesWithIdData
	BehaviorNamesWithId genericHandlerBuilder[
		*BehaviorIdNamesBehaviorNamesWithIdRequest,
		*BehaviorNamesWithIdData,
		func(context.Context, *BehaviorIdNamesBehaviorNamesWithIdRequest) (*BehaviorNamesWithIdData, error),
		func(http.ResponseWriter, *http.Request, *BehaviorIdNamesBehaviorNamesWithIdRequest) (*BehaviorNamesWithIdData, error),
	]
}

func newBehaviorIdNamesControllerBuilder(app *RootHandler) *behaviorIdNamesControllerBuilder {
	return &behaviorIdNamesControllerBuilder{
		// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
		BehaviorNamesWithId: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*BehaviorIdNamesBehaviorNamesWithIdRequest,
				*BehaviorNamesWithIdData,
			](),
			newHTTPHandlerAdapter[
				*BehaviorIdNamesBehaviorNamesWithIdRequest,
				*BehaviorNamesWithIdData,
			](),
			makeActionBuilderParams[
				*BehaviorIdNamesBehaviorNamesWithIdRequest,
				*BehaviorNamesWithIdData,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserBehaviorIdNamesBehaviorNamesWithId(app),
			},
		),
	}
}

type BehaviorIdNamesController interface {
	// POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIdNamesBehaviorNamesWithIdRequest,
	//
	// Response type: BehaviorNamesWithIdData
	BehaviorNamesWithId(HandlerBuilder[
		*BehaviorIdNamesBehaviorNamesWithIdRequest,
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
	rootHandler.router.HandleRoute("POST", "/behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}", controller.BehaviorNamesWithId(builder.BehaviorNamesWithId))
	return rootHandler
}