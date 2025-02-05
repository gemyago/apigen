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

type ObjectsController struct {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectRequest,
	//
	// Response type: none
	ObjectsArrayBodyDirect httpHandlerFactory

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedRequest,
	//
	// Response type: none
	ObjectsArrayBodyNested httpHandlerFactory

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedRequest,
	//
	// Response type: none
	ObjectsDeeplyNested httpHandlerFactory

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyRequest,
	//
	// Response type: none
	ObjectsNullableOptionalBody httpHandlerFactory

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyRequest,
	//
	// Response type: none
	ObjectsNullableRequiredBody httpHandlerFactory

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyRequest,
	//
	// Response type: none
	ObjectsOptionalBody httpHandlerFactory

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyRequest,
	//
	// Response type: none
	ObjectsRequiredBody httpHandlerFactory

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsRequest,
	//
	// Response type: none
	ObjectsRequiredNestedObjects httpHandlerFactory
}

type ObjectsControllerBuilder struct {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectRequest,
	//
	// Response type: none
	HandleObjectsArrayBodyDirect actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsArrayBodyDirectRequest]

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedRequest,
	//
	// Response type: none
	HandleObjectsArrayBodyNested actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsArrayBodyNestedRequest]

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedRequest,
	//
	// Response type: none
	HandleObjectsDeeplyNested actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsDeeplyNestedRequest]

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyRequest,
	//
	// Response type: none
	HandleObjectsNullableOptionalBody actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsNullableOptionalBodyRequest]

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyRequest,
	//
	// Response type: none
	HandleObjectsNullableRequiredBody actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsNullableRequiredBodyRequest]

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyRequest,
	//
	// Response type: none
	HandleObjectsOptionalBody actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsOptionalBodyRequest]

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyRequest,
	//
	// Response type: none
	HandleObjectsRequiredBody actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsRequiredBodyRequest]

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsRequest,
	//
	// Response type: none
	HandleObjectsRequiredNestedObjects actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsRequiredNestedObjectsRequest]
}

func (c *ObjectsControllerBuilder) Finalize() *ObjectsController {
	return &ObjectsController{
		ObjectsArrayBodyDirect: mustInitializeAction("objectsArrayBodyDirect", c.HandleObjectsArrayBodyDirect.httpHandlerFactory),
		ObjectsArrayBodyNested: mustInitializeAction("objectsArrayBodyNested", c.HandleObjectsArrayBodyNested.httpHandlerFactory),
		ObjectsDeeplyNested: mustInitializeAction("objectsDeeplyNested", c.HandleObjectsDeeplyNested.httpHandlerFactory),
		ObjectsNullableOptionalBody: mustInitializeAction("objectsNullableOptionalBody", c.HandleObjectsNullableOptionalBody.httpHandlerFactory),
		ObjectsNullableRequiredBody: mustInitializeAction("objectsNullableRequiredBody", c.HandleObjectsNullableRequiredBody.httpHandlerFactory),
		ObjectsOptionalBody: mustInitializeAction("objectsOptionalBody", c.HandleObjectsOptionalBody.httpHandlerFactory),
		ObjectsRequiredBody: mustInitializeAction("objectsRequiredBody", c.HandleObjectsRequiredBody.httpHandlerFactory),
		ObjectsRequiredNestedObjects: mustInitializeAction("objectsRequiredNestedObjects", c.HandleObjectsRequiredNestedObjects.httpHandlerFactory),
	}
}

func BuildObjectsController() *ObjectsControllerBuilder {
	controllerBuilder := &ObjectsControllerBuilder{}

	// POST /objects/arrays
	controllerBuilder.HandleObjectsArrayBodyDirect.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsArrayBodyDirect.defaultStatusCode = 204
	controllerBuilder.HandleObjectsArrayBodyDirect.voidResult = true
	controllerBuilder.HandleObjectsArrayBodyDirect.paramsParserFactory = newParamsParserObjectsObjectsArrayBodyDirect

	// PUT /objects/arrays
	controllerBuilder.HandleObjectsArrayBodyNested.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsArrayBodyNested.defaultStatusCode = 204
	controllerBuilder.HandleObjectsArrayBodyNested.voidResult = true
	controllerBuilder.HandleObjectsArrayBodyNested.paramsParserFactory = newParamsParserObjectsObjectsArrayBodyNested

	// POST /objects/deeply-nested
	controllerBuilder.HandleObjectsDeeplyNested.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsDeeplyNested.defaultStatusCode = 204
	controllerBuilder.HandleObjectsDeeplyNested.voidResult = true
	controllerBuilder.HandleObjectsDeeplyNested.paramsParserFactory = newParamsParserObjectsObjectsDeeplyNested

	// PUT /objects/nullable-body
	controllerBuilder.HandleObjectsNullableOptionalBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsNullableOptionalBody.defaultStatusCode = 204
	controllerBuilder.HandleObjectsNullableOptionalBody.voidResult = true
	controllerBuilder.HandleObjectsNullableOptionalBody.paramsParserFactory = newParamsParserObjectsObjectsNullableOptionalBody

	// POST /objects/nullable-body
	controllerBuilder.HandleObjectsNullableRequiredBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsNullableRequiredBody.defaultStatusCode = 204
	controllerBuilder.HandleObjectsNullableRequiredBody.voidResult = true
	controllerBuilder.HandleObjectsNullableRequiredBody.paramsParserFactory = newParamsParserObjectsObjectsNullableRequiredBody

	// PUT /objects/required-body
	controllerBuilder.HandleObjectsOptionalBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsOptionalBody.defaultStatusCode = 204
	controllerBuilder.HandleObjectsOptionalBody.voidResult = true
	controllerBuilder.HandleObjectsOptionalBody.paramsParserFactory = newParamsParserObjectsObjectsOptionalBody

	// POST /objects/required-body
	controllerBuilder.HandleObjectsRequiredBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsRequiredBody.defaultStatusCode = 204
	controllerBuilder.HandleObjectsRequiredBody.voidResult = true
	controllerBuilder.HandleObjectsRequiredBody.paramsParserFactory = newParamsParserObjectsObjectsRequiredBody

	// POST /objects/required-nested-objects
	controllerBuilder.HandleObjectsRequiredNestedObjects.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsRequiredNestedObjects.defaultStatusCode = 204
	controllerBuilder.HandleObjectsRequiredNestedObjects.voidResult = true
	controllerBuilder.HandleObjectsRequiredNestedObjects.paramsParserFactory = newParamsParserObjectsObjectsRequiredNestedObjects

	return controllerBuilder
}

func RegisterObjectsRoutes(controller *ObjectsController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/objects/arrays", controller.ObjectsArrayBodyDirect(app))
	app.router.HandleRoute("PUT", "/objects/arrays", controller.ObjectsArrayBodyNested(app))
	app.router.HandleRoute("POST", "/objects/deeply-nested", controller.ObjectsDeeplyNested(app))
	app.router.HandleRoute("PUT", "/objects/nullable-body", controller.ObjectsNullableOptionalBody(app))
	app.router.HandleRoute("POST", "/objects/nullable-body", controller.ObjectsNullableRequiredBody(app))
	app.router.HandleRoute("PUT", "/objects/required-body", controller.ObjectsOptionalBody(app))
	app.router.HandleRoute("POST", "/objects/required-body", controller.ObjectsRequiredBody(app))
	app.router.HandleRoute("POST", "/objects/required-nested-objects", controller.ObjectsRequiredNestedObjects(app))
}

type ObjectsControllerBuilderV2 struct {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectRequest,
	//
	// Response type: none
	ObjectsArrayBodyDirect ActionBuilder[
	  ObjectsObjectsArrayBodyDirectRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsArrayBodyDirectRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsArrayBodyDirectRequest) (error),
	]

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedRequest,
	//
	// Response type: none
	ObjectsArrayBodyNested ActionBuilder[
	  ObjectsObjectsArrayBodyNestedRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsArrayBodyNestedRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsArrayBodyNestedRequest) (error),
	]

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedRequest,
	//
	// Response type: none
	ObjectsDeeplyNested ActionBuilder[
	  ObjectsObjectsDeeplyNestedRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsDeeplyNestedRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsDeeplyNestedRequest) (error),
	]

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyRequest,
	//
	// Response type: none
	ObjectsNullableOptionalBody ActionBuilder[
	  ObjectsObjectsNullableOptionalBodyRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsNullableOptionalBodyRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsNullableOptionalBodyRequest) (error),
	]

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyRequest,
	//
	// Response type: none
	ObjectsNullableRequiredBody ActionBuilder[
	  ObjectsObjectsNullableRequiredBodyRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsNullableRequiredBodyRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsNullableRequiredBodyRequest) (error),
	]

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyRequest,
	//
	// Response type: none
	ObjectsOptionalBody ActionBuilder[
	  ObjectsObjectsOptionalBodyRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsOptionalBodyRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsOptionalBodyRequest) (error),
	]

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyRequest,
	//
	// Response type: none
	ObjectsRequiredBody ActionBuilder[
	  ObjectsObjectsRequiredBodyRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsRequiredBodyRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsRequiredBodyRequest) (error),
	]

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsRequest,
	//
	// Response type: none
	ObjectsRequiredNestedObjects ActionBuilder[
	  ObjectsObjectsRequiredNestedObjectsRequest,
	  Void,
	  func(context.Context, *ObjectsObjectsRequiredNestedObjectsRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ObjectsObjectsRequiredNestedObjectsRequest) (error),
	]
}

type ObjectsControllerV2 interface {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectRequest,
	//
	// Response type: none
	ObjectsArrayBodyDirect(b *ObjectsControllerBuilderV2) http.Handler

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedRequest,
	//
	// Response type: none
	ObjectsArrayBodyNested(b *ObjectsControllerBuilderV2) http.Handler

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedRequest,
	//
	// Response type: none
	ObjectsDeeplyNested(b *ObjectsControllerBuilderV2) http.Handler

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyRequest,
	//
	// Response type: none
	ObjectsNullableOptionalBody(b *ObjectsControllerBuilderV2) http.Handler

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyRequest,
	//
	// Response type: none
	ObjectsNullableRequiredBody(b *ObjectsControllerBuilderV2) http.Handler

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyRequest,
	//
	// Response type: none
	ObjectsOptionalBody(b *ObjectsControllerBuilderV2) http.Handler

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyRequest,
	//
	// Response type: none
	ObjectsRequiredBody(b *ObjectsControllerBuilderV2) http.Handler

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsRequest,
	//
	// Response type: none
	ObjectsRequiredNestedObjects(b *ObjectsControllerBuilderV2) http.Handler
}

func RegisterObjectsRoutesV2(controller ObjectsControllerV2, app *HTTPApp) {
  builder := ObjectsControllerBuilderV2{}
	app.router.HandleRoute("POST", "/objects/arrays", controller.ObjectsArrayBodyDirect(&builder))
	app.router.HandleRoute("PUT", "/objects/arrays", controller.ObjectsArrayBodyNested(&builder))
	app.router.HandleRoute("POST", "/objects/deeply-nested", controller.ObjectsDeeplyNested(&builder))
	app.router.HandleRoute("PUT", "/objects/nullable-body", controller.ObjectsNullableOptionalBody(&builder))
	app.router.HandleRoute("POST", "/objects/nullable-body", controller.ObjectsNullableRequiredBody(&builder))
	app.router.HandleRoute("PUT", "/objects/required-body", controller.ObjectsOptionalBody(&builder))
	app.router.HandleRoute("POST", "/objects/required-body", controller.ObjectsRequiredBody(&builder))
	app.router.HandleRoute("POST", "/objects/required-nested-objects", controller.ObjectsRequiredNestedObjects(&builder))
}