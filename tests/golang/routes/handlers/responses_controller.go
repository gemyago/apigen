package handlers

import (
	"time"
)

// Below is to workaround unused imports.
var _ = time.Time{}





type ResponsesController struct {
	// POST /responses/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	ResponsesNoStatusDefined httpHandlerFactory

	// POST /responses/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	ResponsesWithStatusDefined httpHandlerFactory
}

type ResponsesControllerBuilder struct {
	// POST /responses/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	HandleResponsesNoStatusDefined actionBuilderVoidResult[*ResponsesControllerBuilder, *ResponsesResponsesNoStatusDefinedRequest]

	// POST /responses/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	HandleResponsesWithStatusDefined actionBuilderVoidResult[*ResponsesControllerBuilder, *ResponsesResponsesWithStatusDefinedRequest]
}

func (c *ResponsesControllerBuilder) Finalize() *ResponsesController {
	// TODO: panic if any handler is null
	return &ResponsesController{
		ResponsesNoStatusDefined: c.HandleResponsesNoStatusDefined.httpHandlerFactory,
		ResponsesWithStatusDefined: c.HandleResponsesWithStatusDefined.httpHandlerFactory,
	}
}

func BuildResponsesController() *ResponsesControllerBuilder {
	controllerBuilder := &ResponsesControllerBuilder{}

	// POST /responses/no-status-defined
	controllerBuilder.HandleResponsesNoStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleResponsesNoStatusDefined.defaultStatusCode = 
	controllerBuilder.HandleResponsesNoStatusDefined.voidResult = true
	controllerBuilder.HandleResponsesNoStatusDefined.paramsParserFactory = newParamsParserResponsesResponsesNoStatusDefined

	// POST /responses/with-status-defined
	controllerBuilder.HandleResponsesWithStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleResponsesWithStatusDefined.defaultStatusCode = 202
	controllerBuilder.HandleResponsesWithStatusDefined.voidResult = true
	controllerBuilder.HandleResponsesWithStatusDefined.paramsParserFactory = newParamsParserResponsesResponsesWithStatusDefined

	return controllerBuilder
}

func RegisterResponsesRoutes(controller *ResponsesController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/responses/no-status-defined", controller.ResponsesNoStatusDefined(app))
	app.router.HandleRoute("POST", "/responses/with-status-defined", controller.ResponsesWithStatusDefined(app))
}
