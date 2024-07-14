package handlers

type NumericTypesErrorHandlingParsingErrorsRequest struct {
	NumberAny float32
	NumberFloat float32
	NumberDouble float64
	NumberInt int32
	NumberInt32 int32
	NumberInt64 int64
	NumberAnyInQuery float32
	NumberFloatInQuery float32
	NumberDoubleInQuery float64
	NumberIntInQuery int32
	NumberInt32InQuery int32
	NumberInt64InQuery int64
}

type NumericTypesController struct {
	// GET /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	ErrorHandlingParsingErrors httpHandlerFactory
}

type NumericTypesControllerBuilder struct {
	// GET /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	HandleErrorHandlingParsingErrors actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesErrorHandlingParsingErrorsRequest]
}

func (c *NumericTypesControllerBuilder) Finalize() *NumericTypesController {
	// TODO: panic if any handler is null
	return &NumericTypesController{
		ErrorHandlingParsingErrors: c.HandleErrorHandlingParsingErrors.httpHandlerFactory,
	}
}

func BuildNumericTypesController() *NumericTypesControllerBuilder {
	controllerBuilder := &NumericTypesControllerBuilder{}

	// GET /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleErrorHandlingParsingErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleErrorHandlingParsingErrors.defaultStatusCode = 204
	controllerBuilder.HandleErrorHandlingParsingErrors.voidResult = true
	controllerBuilder.HandleErrorHandlingParsingErrors.paramsParser = newNumericTypesErrorHandlingParsingErrorsParamsParser()

	return controllerBuilder
}

func MountNumericTypesRoutes(controller *NumericTypesController, app *httpApp) {
	app.router.HandleRoute("GET", "/numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.ErrorHandlingParsingErrors(app))
}
