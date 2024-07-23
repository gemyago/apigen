package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type ObjectsObjectsArrayParsingBodyDirectRequest struct {
	Payload *models.ObjectsArrayParsingBodyDirectRequest
}

type ObjectsObjectsArrayParsingBodyDirect0Request struct {
	Payload *models.[]ObjectArraysSimpleObject
}

type ObjectsObjectsNestedRequest struct {
	Payload *models.ObjectsNestedRequest
}

type ObjectsController struct {
	// PUT /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyDirectRequest,
	//
	// Response type: none
	ObjectsArrayParsingBodyDirect httpHandlerFactory

	// POST /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyDirect0Request,
	//
	// Response type: none
	ObjectsArrayParsingBodyDirect_0 httpHandlerFactory

	// POST /objects/nested
	//
	// Request type: ObjectsObjectsNestedRequest,
	//
	// Response type: none
	ObjectsNested httpHandlerFactory
}

type ObjectsControllerBuilder struct {
	// PUT /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyDirectRequest,
	//
	// Response type: none
	HandleObjectsArrayParsingBodyDirect actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsArrayParsingBodyDirectRequest]

	// POST /objects/arrays-parsing
	//
	// Request type: ObjectsObjectsArrayParsingBodyDirect0Request,
	//
	// Response type: none
	HandleObjectsArrayParsingBodyDirect_0 actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsArrayParsingBodyDirect_0Request]

	// POST /objects/nested
	//
	// Request type: ObjectsObjectsNestedRequest,
	//
	// Response type: none
	HandleObjectsNested actionBuilderVoidResult[*ObjectsControllerBuilder, *ObjectsObjectsNestedRequest]
}

func (c *ObjectsControllerBuilder) Finalize() *ObjectsController {
	// TODO: panic if any handler is null
	return &ObjectsController{
		ObjectsArrayParsingBodyDirect: c.HandleObjectsArrayParsingBodyDirect.httpHandlerFactory,
		ObjectsArrayParsingBodyDirect_0: c.HandleObjectsArrayParsingBodyDirect_0.httpHandlerFactory,
		ObjectsNested: c.HandleObjectsNested.httpHandlerFactory,
	}
}

func BuildObjectsController() *ObjectsControllerBuilder {
	controllerBuilder := &ObjectsControllerBuilder{}

	// PUT /objects/arrays-parsing
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.defaultStatusCode = 200
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.voidResult = true
	controllerBuilder.HandleObjectsArrayParsingBodyDirect.paramsParserFactory = newParamsParserObjectsObjectsArrayParsingBodyDirect

	// POST /objects/arrays-parsing
	controllerBuilder.HandleObjectsArrayParsingBodyDirect_0.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsArrayParsingBodyDirect_0.defaultStatusCode = 200
	controllerBuilder.HandleObjectsArrayParsingBodyDirect_0.voidResult = true
	controllerBuilder.HandleObjectsArrayParsingBodyDirect_0.paramsParserFactory = newParamsParserObjectsObjectsArrayParsingBodyDirect_0

	// POST /objects/nested
	controllerBuilder.HandleObjectsNested.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsNested.defaultStatusCode = 200
	controllerBuilder.HandleObjectsNested.voidResult = true
	controllerBuilder.HandleObjectsNested.paramsParserFactory = newParamsParserObjectsObjectsNested

	return controllerBuilder
}

func RegisterObjectsRoutes(controller *ObjectsController, app *HTTPApp) {
	app.router.HandleRoute("PUT", "/objects/arrays-parsing", controller.ObjectsArrayParsingBodyDirect(app))
	app.router.HandleRoute("POST", "/objects/arrays-parsing", controller.ObjectsArrayParsingBodyDirect_0(app))
	app.router.HandleRoute("POST", "/objects/nested", controller.ObjectsNested(app))
}
