package handlers

type ErrorHandlingParsingErrorsRequest struct {
	PathParam1 float32
	PathParam2 float32
	RequiredQuery1 float32
	RequiredQuery2 float32
}

type ErrorHandlingValidationErrorsRequest struct {
	RequiredQuery1 float32
	RequiredQuery2 float32
}

type ErrorHandlingController struct {
	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	ParsingErrors httpHandlerFactory

	// GET /error-handling/validation-errors
	//
	// Request type: ErrorHandlingValidationErrorsRequest,
	//
	// Response type: none
	ValidationErrors httpHandlerFactory
}

type ErrorHandlingControllerBuilder struct {
	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	HandleParsingErrors actionBuilderVoidResult[*ErrorHandlingControllerBuilder, *ErrorHandlingParsingErrorsRequest]

	// GET /error-handling/validation-errors
	//
	// Request type: ErrorHandlingValidationErrorsRequest,
	//
	// Response type: none
	HandleValidationErrors actionBuilderVoidResult[*ErrorHandlingControllerBuilder, *ErrorHandlingValidationErrorsRequest]
}

func (c *ErrorHandlingControllerBuilder) Finalize() *ErrorHandlingController {
	// TODO: panic if any handler is null
	return &ErrorHandlingController{
		ParsingErrors: c.HandleParsingErrors.httpHandlerFactory,
		ValidationErrors: c.HandleValidationErrors.httpHandlerFactory,
	}
}

func BuildErrorHandlingController() *ErrorHandlingControllerBuilder {
	controllerBuilder := &ErrorHandlingControllerBuilder{}

	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	controllerBuilder.HandleParsingErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleParsingErrors.defaultStatusCode = 204
	controllerBuilder.HandleParsingErrors.voidResult = true
	controllerBuilder.HandleParsingErrors.paramsParser = newErrorHandlingParsingErrorsParamsParser()

	// GET /error-handling/validation-errors
	controllerBuilder.HandleValidationErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleValidationErrors.defaultStatusCode = 204
	controllerBuilder.HandleValidationErrors.voidResult = true
	controllerBuilder.HandleValidationErrors.paramsParser = newErrorHandlingValidationErrorsParamsParser()

	return controllerBuilder
}

func MountErrorHandlingRoutes(controller *ErrorHandlingController, app *httpApp) {
	app.router.HandleRoute("GET", "/error-handling/parsing-errors/{pathParam1}/{pathParam2}", controller.ParsingErrors(app))
	app.router.HandleRoute("GET", "/error-handling/validation-errors", controller.ValidationErrors(app))
}
