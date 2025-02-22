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
type _ func() BehaviorNamesWithId200Response

// BehaviorIdNamesBehaviorNamesWithIdRequest represents params for behaviorNamesWithId operation
//
// Request: GET /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}.
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
}

type behaviorIdNamesControllerBuilder struct {
	// GET /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIdNamesBehaviorNamesWithIdRequest,
	//
	// Response type: BehaviorNamesWithId200Response
	BehaviorNamesWithId genericHandlerBuilder[
		*BehaviorIdNamesBehaviorNamesWithIdRequest,
		*BehaviorNamesWithId200Response,
		func(context.Context, *BehaviorIdNamesBehaviorNamesWithIdRequest) (*BehaviorNamesWithId200Response, error),
		func(http.ResponseWriter, *http.Request, *BehaviorIdNamesBehaviorNamesWithIdRequest) (*BehaviorNamesWithId200Response, error),
	]
}

func newBehaviorIdNamesControllerBuilder(app *RootHandler) *behaviorIdNamesControllerBuilder {
	return &behaviorIdNamesControllerBuilder{
		// GET /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
		BehaviorNamesWithId: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*BehaviorIdNamesBehaviorNamesWithIdRequest,
				*BehaviorNamesWithId200Response,
			](),
			newHTTPHandlerAdapter[
				*BehaviorIdNamesBehaviorNamesWithIdRequest,
				*BehaviorNamesWithId200Response,
			](),
			makeActionBuilderParams[
				*BehaviorIdNamesBehaviorNamesWithIdRequest,
				*BehaviorNamesWithId200Response,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserBehaviorIdNamesBehaviorNamesWithId(app),
			},
		),
	}
}

type BehaviorIdNamesController interface {
	// GET /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
	//
	// Request type: BehaviorIdNamesBehaviorNamesWithIdRequest,
	//
	// Response type: BehaviorNamesWithId200Response
	BehaviorNamesWithId(HandlerBuilder[
		*BehaviorIdNamesBehaviorNamesWithIdRequest,
		*BehaviorNamesWithId200Response,
	]) http.Handler
}

// RegisterBehaviorIdNamesRoutes will attach the following routes to the root handler:
// 
// - GET /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterBehaviorIdNamesRoutes(controller BehaviorIdNamesController) *RootHandler {
	builder := newBehaviorIdNamesControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("GET", "/behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}", controller.BehaviorNamesWithId(builder.BehaviorNamesWithId))
	return rootHandler
}