package handlers

type ErrorHandlingNumberRangeErrorsRequest struct {
	LimitedNum float32
	LimitedFloat float32
	LimitedDouble float64
	LimitedQueryNum float32
	LimitedQueryFloat float32
	LimitedQueryDouble float64
}

type ErrorHandlingParsingErrorsRequest struct {
	PathParam1 float32
	PathParam2 float32
	RequiredQuery1 float32
	RequiredQuery2 float32
}

type ErrorHandlingController struct {
	// GET /error-handling/number-range-errors/{limitedNum}/{limitedFloat}/{limitedDouble}
	//
	// Request type: ErrorHandlingNumberRangeErrorsRequest,
	//
	// Response type: none
	NumberRangeErrors httpHandlerFactory

	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	ParsingErrors httpHandlerFactory
}

type ErrorHandlingControllerBuilder struct {
	// GET /error-handling/number-range-errors/{limitedNum}/{limitedFloat}/{limitedDouble}
	//
	// Request type: ErrorHandlingNumberRangeErrorsRequest,
	//
	// Response type: none
	HandleNumberRangeErrors actionBuilderVoidResult[*ErrorHandlingControllerBuilder, *ErrorHandlingNumberRangeErrorsRequest]

	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	HandleParsingErrors actionBuilderVoidResult[*ErrorHandlingControllerBuilder, *ErrorHandlingParsingErrorsRequest]
}

func (c *ErrorHandlingControllerBuilder) Finalize() *ErrorHandlingController {
	// TODO: panic if any handler is null
	return &ErrorHandlingController{
		NumberRangeErrors: c.HandleNumberRangeErrors.httpHandlerFactory,
		ParsingErrors: c.HandleParsingErrors.httpHandlerFactory,
	}
}

func BuildErrorHandlingController() *ErrorHandlingControllerBuilder {
	controllerBuilder := &ErrorHandlingControllerBuilder{}

	// GET /error-handling/number-range-errors/{limitedNum}/{limitedFloat}/{limitedDouble}
	controllerBuilder.HandleNumberRangeErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumberRangeErrors.defaultStatusCode = 204
	controllerBuilder.HandleNumberRangeErrors.voidResult = true
	controllerBuilder.HandleNumberRangeErrors.paramsParser = newErrorHandlingNumberRangeErrorsParamsParser()

	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	controllerBuilder.HandleParsingErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleParsingErrors.defaultStatusCode = 204
	controllerBuilder.HandleParsingErrors.voidResult = true
	controllerBuilder.HandleParsingErrors.paramsParser = newErrorHandlingParsingErrorsParamsParser()

	return controllerBuilder
}

func MountErrorHandlingRoutes(controller *ErrorHandlingController, app *httpApp) {
	app.router.HandleRoute("GET", "/error-handling/number-range-errors/{limitedNum}/{limitedFloat}/{limitedDouble}", controller.NumberRangeErrors(app))
	app.router.HandleRoute("GET", "/error-handling/parsing-errors/{pathParam1}/{pathParam2}", controller.ParsingErrors(app))
}
