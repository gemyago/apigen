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
}

func (c *ObjectsControllerBuilder) Finalize() *ObjectsController {
	return &ObjectsController{
		ObjectsArrayParsingBodyDirect: mustInitializeAction("objectsArrayParsingBodyDirect", c.HandleObjectsArrayParsingBodyDirect.httpHandlerFactory),
		ObjectsArrayParsingBodyNested: mustInitializeAction("objectsArrayParsingBodyNested", c.HandleObjectsArrayParsingBodyNested.httpHandlerFactory),
		ObjectsNested: mustInitializeAction("objectsNested", c.HandleObjectsNested.httpHandlerFactory),
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

	return controllerBuilder
}

func RegisterObjectsRoutes(controller *ObjectsController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/objects/arrays-parsing", controller.ObjectsArrayParsingBodyDirect(app))
	app.router.HandleRoute("PUT", "/objects/arrays-parsing", controller.ObjectsArrayParsingBodyNested(app))
	app.router.HandleRoute("POST", "/objects/nested", controller.ObjectsNested(app))
}
