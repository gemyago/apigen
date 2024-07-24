package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type ObjectsObjectsArrayParsingBodyDirectRequest struct {
	Payload []*models.ObjectArraysSimpleObject
}

type ObjectsObjectsArrayParsingBodyNestedRequest struct {
	Payload *models.ObjectsArrayParsingBodyNestedRequest
}

type ObjectsObjectsNestedRequest struct {
	Payload *models.ObjectsNestedRequest
}

type ObjectsObjectsNullableOptionalBodyRequest struct {
	Payload *models.SimpleObject
}

type ObjectsObjectsNullableRequiredBodyRequest struct {
	Payload *models.SimpleObject
}

type ObjectsObjectsOptionalBodyRequest struct {
	Payload *models.SimpleObject
}

type ObjectsObjectsOptionalNestedObjectsRequest struct {
	Payload *models.ObjectsOptionalNestedObjectsRequest
}

type ObjectsObjectsRequiredBodyRequest struct {
	Payload *models.SimpleObject
}

type ObjectsObjectsRequiredNestedObjectsRequest struct {
	Payload *models.ObjectsRequiredNestedObjectsRequest
}

type ObjectsController struct {
	// POST /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyDirectRequest,
	//
	// Response type: none
	ObjectsArrayParsingBodyDirect httpHandlerFactory

	// PUT /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyNestedRequest,
	//
	// Response type: none
	ObjectsArrayParsingBodyNested httpHandlerFactory

	// POST /objects/nested
	//
	// Request type: ObjectsObjectsNestedRequest,
	//
	// Response type: none
	ObjectsNested httpHandlerFactory

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

	// PUT /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsOptionalNestedObjectsRequest,
	//
	// Response type: none
	ObjectsOptionalNestedObjects httpHandlerFactory

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
	// POST /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyDirectRequest,
	//
	// Response type: none
	HandleObjectsArrayParsingBodyDirect actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsArrayParsingBodyDirectRequest]

	// PUT /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyNestedRequest,
	//
	// Response type: none
	HandleObjectsArrayParsingBodyNested actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsArrayParsingBodyNestedRequest]

	// POST /objects/nested
	//
	// Request type: ObjectsObjectsNestedRequest,
	//
	// Response type: none
	HandleObjectsNested actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsNestedRequest]

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

	// PUT /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsOptionalNestedObjectsRequest,
	//
	// Response type: none
	HandleObjectsOptionalNestedObjects actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsOptionalNestedObjectsRequest]

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
		ObjectsArrayParsingBodyDirect: mustInitializeAction("objectsArrayParsingBodyDirect", c.HandleObjectsArrayParsingBodyDirect.httpHandlerFactory),
		ObjectsArrayParsingBodyNested: mustInitializeAction("objectsArrayParsingBodyNested", c.HandleObjectsArrayParsingBodyNested.httpHandlerFactory),
		ObjectsNested: mustInitializeAction("objectsNested", c.HandleObjectsNested.httpHandlerFactory),
		ObjectsNullableOptionalBody: mustInitializeAction("objectsNullableOptionalBody", c.HandleObjectsNullableOptionalBody.httpHandlerFactory),
		ObjectsNullableRequiredBody: mustInitializeAction("objectsNullableRequiredBody", c.HandleObjectsNullableRequiredBody.httpHandlerFactory),
		ObjectsOptionalBody: mustInitializeAction("objectsOptionalBody", c.HandleObjectsOptionalBody.httpHandlerFactory),
		ObjectsOptionalNestedObjects: mustInitializeAction("objectsOptionalNestedObjects", c.HandleObjectsOptionalNestedObjects.httpHandlerFactory),
		ObjectsRequiredBody: mustInitializeAction("objectsRequiredBody", c.HandleObjectsRequiredBody.httpHandlerFactory),
		ObjectsRequiredNestedObjects: mustInitializeAction("objectsRequiredNestedObjects", c.HandleObjectsRequiredNestedObjects.httpHandlerFactory),
	}
}

func BuildObjectsController() *ObjectsControllerBuilder {
	controllerBuilder := &ObjectsControllerBuilder{}

	// POST /objects/arrays-parsing
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.defaultStatusCode = 200
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.voidResult = true
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.paramsParserFactory = newParamsParserObjectsObjectsArrayParsingBodyDirect

	// PUT /objects/arrays-parsing
	controllerBuilder.HandleObjectsArrayParsingBodyNested.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsArrayParsingBodyNested.defaultStatusCode = 200
	controllerBuilder.HandleObjectsArrayParsingBodyNested.voidResult = true
	controllerBuilder.HandleObjectsArrayParsingBodyNested.paramsParserFactory = newParamsParserObjectsObjectsArrayParsingBodyNested

	// POST /objects/nested
	controllerBuilder.HandleObjectsNested.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsNested.defaultStatusCode = 200
	controllerBuilder.HandleObjectsNested.voidResult = true
	controllerBuilder.HandleObjectsNested.paramsParserFactory = newParamsParserObjectsObjectsNested

	// PUT /objects/nullable-body
	controllerBuilder.HandleObjectsNullableOptionalBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsNullableOptionalBody.defaultStatusCode = 200
	controllerBuilder.HandleObjectsNullableOptionalBody.voidResult = true
	controllerBuilder.HandleObjectsNullableOptionalBody.paramsParserFactory = newParamsParserObjectsObjectsNullableOptionalBody

	// POST /objects/nullable-body
	controllerBuilder.HandleObjectsNullableRequiredBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsNullableRequiredBody.defaultStatusCode = 200
	controllerBuilder.HandleObjectsNullableRequiredBody.voidResult = true
	controllerBuilder.HandleObjectsNullableRequiredBody.paramsParserFactory = newParamsParserObjectsObjectsNullableRequiredBody

	// PUT /objects/required-body
	controllerBuilder.HandleObjectsOptionalBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsOptionalBody.defaultStatusCode = 200
	controllerBuilder.HandleObjectsOptionalBody.voidResult = true
	controllerBuilder.HandleObjectsOptionalBody.paramsParserFactory = newParamsParserObjectsObjectsOptionalBody

	// PUT /objects/required-nested-objects
	controllerBuilder.HandleObjectsOptionalNestedObjects.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsOptionalNestedObjects.defaultStatusCode = 200
	controllerBuilder.HandleObjectsOptionalNestedObjects.voidResult = true
	controllerBuilder.HandleObjectsOptionalNestedObjects.paramsParserFactory = newParamsParserObjectsObjectsOptionalNestedObjects

	// POST /objects/required-body
	controllerBuilder.HandleObjectsRequiredBody.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsRequiredBody.defaultStatusCode = 200
	controllerBuilder.HandleObjectsRequiredBody.voidResult = true
	controllerBuilder.HandleObjectsRequiredBody.paramsParserFactory = newParamsParserObjectsObjectsRequiredBody

	// POST /objects/required-nested-objects
	controllerBuilder.HandleObjectsRequiredNestedObjects.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsRequiredNestedObjects.defaultStatusCode = 200
	controllerBuilder.HandleObjectsRequiredNestedObjects.voidResult = true
	controllerBuilder.HandleObjectsRequiredNestedObjects.paramsParserFactory = newParamsParserObjectsObjectsRequiredNestedObjects

	return controllerBuilder
}

func RegisterObjectsRoutes(controller *ObjectsController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/objects/arrays-parsing", controller.ObjectsArrayParsingBodyDirect(app))
	app.router.HandleRoute("PUT", "/objects/arrays-parsing", controller.ObjectsArrayParsingBodyNested(app))
	app.router.HandleRoute("POST", "/objects/nested", controller.ObjectsNested(app))
	app.router.HandleRoute("PUT", "/objects/nullable-body", controller.ObjectsNullableOptionalBody(app))
	app.router.HandleRoute("POST", "/objects/nullable-body", controller.ObjectsNullableRequiredBody(app))
	app.router.HandleRoute("PUT", "/objects/required-body", controller.ObjectsOptionalBody(app))
	app.router.HandleRoute("PUT", "/objects/required-nested-objects", controller.ObjectsOptionalNestedObjects(app))
	app.router.HandleRoute("POST", "/objects/required-body", controller.ObjectsRequiredBody(app))
	app.router.HandleRoute("POST", "/objects/required-nested-objects", controller.ObjectsRequiredNestedObjects(app))
}
