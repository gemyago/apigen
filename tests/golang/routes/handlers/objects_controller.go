package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type ObjectsObjectsNestedRequest struct {
	Payload *models.ObjectsNestedRequest
}

type ObjectsController struct {
	// POST /objects/nested
	//
	// Request type: ObjectsObjectsNestedRequest,
	//
	// Response type: none
	ObjectsNested httpHandlerFactory
}

type ObjectsControllerBuilder struct {
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
		ObjectsNested: c.HandleObjectsNested.httpHandlerFactory,
	}
}

func BuildObjectsController() *ObjectsControllerBuilder {
	controllerBuilder := &ObjectsControllerBuilder{}

	// POST /objects/nested
	controllerBuilder.HandleObjectsNested.controllerBuilder = controllerBuilder
	controllerBuilder.HandleObjectsNested.defaultStatusCode = 
	controllerBuilder.HandleObjectsNested.voidResult = true
	controllerBuilder.HandleObjectsNested.paramsParserFactory = newParamsParserObjectsObjectsNested

	return controllerBuilder
}

func RegisterObjectsRoutes(controller *ObjectsController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/objects/nested", controller.ObjectsNested(app))
}
