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
type _ func() ObjectArraysSimpleObject

// ObjectsObjectsArrayBodyDirectRequest represents params for objectsArrayBodyDirect operation
//
// Request: POST /objects/arrays.
type ObjectsObjectsArrayBodyDirectRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload []*ObjectArraysSimpleObject
}

// ObjectsObjectsArrayBodyNestedRequest represents params for objectsArrayBodyNested operation
//
// Request: PUT /objects/arrays.
type ObjectsObjectsArrayBodyNestedRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *ObjectsArrayBodyNestedRequest
}

// ObjectsObjectsDeeplyNestedRequest represents params for objectsDeeplyNested operation
//
// Request: POST /objects/deeply-nested.
type ObjectsObjectsDeeplyNestedRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *ObjectsDeeplyNestedRequest
}

// ObjectsObjectsNullableOptionalBodyRequest represents params for objectsNullableOptionalBody operation
//
// Request: PUT /objects/nullable-body.
type ObjectsObjectsNullableOptionalBodyRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleNullableObject
}

// ObjectsObjectsNullableRequiredBodyRequest represents params for objectsNullableRequiredBody operation
//
// Request: POST /objects/nullable-body.
type ObjectsObjectsNullableRequiredBodyRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleNullableObject
}

// ObjectsObjectsOptionalBodyRequest represents params for objectsOptionalBody operation
//
// Request: PUT /objects/required-body.
type ObjectsObjectsOptionalBodyRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObject
}

// ObjectsObjectsRequiredBodyRequest represents params for objectsRequiredBody operation
//
// Request: POST /objects/required-body.
type ObjectsObjectsRequiredBodyRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObject
}

// ObjectsObjectsRequiredNestedObjectsRequest represents params for objectsRequiredNestedObjects operation
//
// Request: POST /objects/required-nested-objects.
type ObjectsObjectsRequiredNestedObjectsRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObjectsContainer
}

type objectsControllerBuilder struct {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectRequest,
	//
	// Response type: none
	ObjectsArrayBodyDirect genericHandlerBuilder[
		*ObjectsObjectsArrayBodyDirectRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsArrayBodyDirectRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsArrayBodyDirectRequest, void],
	]

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedRequest,
	//
	// Response type: none
	ObjectsArrayBodyNested genericHandlerBuilder[
		*ObjectsObjectsArrayBodyNestedRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsArrayBodyNestedRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsArrayBodyNestedRequest, void],
	]

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedRequest,
	//
	// Response type: none
	ObjectsDeeplyNested genericHandlerBuilder[
		*ObjectsObjectsDeeplyNestedRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsDeeplyNestedRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsDeeplyNestedRequest, void],
	]

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyRequest,
	//
	// Response type: none
	ObjectsNullableOptionalBody genericHandlerBuilder[
		*ObjectsObjectsNullableOptionalBodyRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsNullableOptionalBodyRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsNullableOptionalBodyRequest, void],
	]

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyRequest,
	//
	// Response type: none
	ObjectsNullableRequiredBody genericHandlerBuilder[
		*ObjectsObjectsNullableRequiredBodyRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsNullableRequiredBodyRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsNullableRequiredBodyRequest, void],
	]

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyRequest,
	//
	// Response type: none
	ObjectsOptionalBody genericHandlerBuilder[
		*ObjectsObjectsOptionalBodyRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsOptionalBodyRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsOptionalBodyRequest, void],
	]

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyRequest,
	//
	// Response type: none
	ObjectsRequiredBody genericHandlerBuilder[
		*ObjectsObjectsRequiredBodyRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsRequiredBodyRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsRequiredBodyRequest, void],
	]

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsRequest,
	//
	// Response type: none
	ObjectsRequiredNestedObjects genericHandlerBuilder[
		*ObjectsObjectsRequiredNestedObjectsRequest,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsRequiredNestedObjectsRequest, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsRequiredNestedObjectsRequest, void],
	]
}

func newObjectsControllerBuilder(app *RootHandler) *objectsControllerBuilder {
	return &objectsControllerBuilder{
		// POST /objects/arrays
		ObjectsArrayBodyDirect: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyDirectRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyDirectRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsArrayBodyDirectRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsArrayBodyDirect(app),
			},
		),

		// PUT /objects/arrays
		ObjectsArrayBodyNested: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyNestedRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyNestedRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsArrayBodyNestedRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsArrayBodyNested(app),
			},
		),

		// POST /objects/deeply-nested
		ObjectsDeeplyNested: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsDeeplyNestedRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsDeeplyNestedRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsDeeplyNestedRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsDeeplyNested(app),
			},
		),

		// PUT /objects/nullable-body
		ObjectsNullableOptionalBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsNullableOptionalBodyRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsNullableOptionalBodyRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsNullableOptionalBodyRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsNullableOptionalBody(app),
			},
		),

		// POST /objects/nullable-body
		ObjectsNullableRequiredBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsNullableRequiredBodyRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsNullableRequiredBodyRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsNullableRequiredBodyRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsNullableRequiredBody(app),
			},
		),

		// PUT /objects/required-body
		ObjectsOptionalBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsOptionalBodyRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsOptionalBodyRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsOptionalBodyRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsOptionalBody(app),
			},
		),

		// POST /objects/required-body
		ObjectsRequiredBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredBodyRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredBodyRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsRequiredBodyRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsRequiredBody(app),
			},
		),

		// POST /objects/required-nested-objects
		ObjectsRequiredNestedObjects: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredNestedObjectsRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredNestedObjectsRequest,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsRequiredNestedObjectsRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsRequiredNestedObjects(app),
			},
		),
	}
}

type ObjectsController interface {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectRequest,
	//
	// Response type: none
	ObjectsArrayBodyDirect(NoResponseHandlerBuilder[
		*ObjectsObjectsArrayBodyDirectRequest,
	]) http.Handler

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedRequest,
	//
	// Response type: none
	ObjectsArrayBodyNested(NoResponseHandlerBuilder[
		*ObjectsObjectsArrayBodyNestedRequest,
	]) http.Handler

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedRequest,
	//
	// Response type: none
	ObjectsDeeplyNested(NoResponseHandlerBuilder[
		*ObjectsObjectsDeeplyNestedRequest,
	]) http.Handler

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyRequest,
	//
	// Response type: none
	ObjectsNullableOptionalBody(NoResponseHandlerBuilder[
		*ObjectsObjectsNullableOptionalBodyRequest,
	]) http.Handler

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyRequest,
	//
	// Response type: none
	ObjectsNullableRequiredBody(NoResponseHandlerBuilder[
		*ObjectsObjectsNullableRequiredBodyRequest,
	]) http.Handler

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyRequest,
	//
	// Response type: none
	ObjectsOptionalBody(NoResponseHandlerBuilder[
		*ObjectsObjectsOptionalBodyRequest,
	]) http.Handler

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyRequest,
	//
	// Response type: none
	ObjectsRequiredBody(NoResponseHandlerBuilder[
		*ObjectsObjectsRequiredBodyRequest,
	]) http.Handler

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsRequest,
	//
	// Response type: none
	ObjectsRequiredNestedObjects(NoResponseHandlerBuilder[
		*ObjectsObjectsRequiredNestedObjectsRequest,
	]) http.Handler
}

// RegisterObjectsRoutes will attach the following routes to the root handler:
// 
// - POST /objects/arrays
// 
// - PUT /objects/arrays
// 
// - POST /objects/deeply-nested
// 
// - PUT /objects/nullable-body
// 
// - POST /objects/nullable-body
// 
// - PUT /objects/required-body
// 
// - POST /objects/required-body
// 
// - POST /objects/required-nested-objects
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterObjectsRoutes(controller ObjectsController) *RootHandler {
	builder := newObjectsControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/objects/arrays", controller.ObjectsArrayBodyDirect(builder.ObjectsArrayBodyDirect))
	rootHandler.router.HandleRoute("PUT", "/objects/arrays", controller.ObjectsArrayBodyNested(builder.ObjectsArrayBodyNested))
	rootHandler.router.HandleRoute("POST", "/objects/deeply-nested", controller.ObjectsDeeplyNested(builder.ObjectsDeeplyNested))
	rootHandler.router.HandleRoute("PUT", "/objects/nullable-body", controller.ObjectsNullableOptionalBody(builder.ObjectsNullableOptionalBody))
	rootHandler.router.HandleRoute("POST", "/objects/nullable-body", controller.ObjectsNullableRequiredBody(builder.ObjectsNullableRequiredBody))
	rootHandler.router.HandleRoute("PUT", "/objects/required-body", controller.ObjectsOptionalBody(builder.ObjectsOptionalBody))
	rootHandler.router.HandleRoute("POST", "/objects/required-body", controller.ObjectsRequiredBody(builder.ObjectsRequiredBody))
	rootHandler.router.HandleRoute("POST", "/objects/required-nested-objects", controller.ObjectsRequiredNestedObjects(builder.ObjectsRequiredNestedObjects))
	return rootHandler
}