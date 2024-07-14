package handlers

type NumericTypesNumericTypesParsingRequest struct {
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

type NumericTypesNumericTypesRangeValidationRequest struct {
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

type NumericTypesNumericTypesRangeValidationExclusiveRequest struct {
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
	// Request type: NumericTypesNumericTypesParsingRequest,
	//
	// Response type: none
	NumericTypesParsing httpHandlerFactory

	// GET /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationRequest,
	//
	// Response type: none
	NumericTypesRangeValidation httpHandlerFactory

	// GET /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationExclusiveRequest,
	//
	// Response type: none
	NumericTypesRangeValidationExclusive httpHandlerFactory
}

type NumericTypesControllerBuilder struct {
	// GET /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesParsingRequest,
	//
	// Response type: none
	HandleNumericTypesParsing actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesParsingRequest]

	// GET /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationRequest,
	//
	// Response type: none
	HandleNumericTypesRangeValidation actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesRangeValidationRequest]

	// GET /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationExclusiveRequest,
	//
	// Response type: none
	HandleNumericTypesRangeValidationExclusive actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesRangeValidationExclusiveRequest]
}

func (c *NumericTypesControllerBuilder) Finalize() *NumericTypesController {
	// TODO: panic if any handler is null
	return &NumericTypesController{
		NumericTypesParsing: c.HandleNumericTypesParsing.httpHandlerFactory,
		NumericTypesRangeValidation: c.HandleNumericTypesRangeValidation.httpHandlerFactory,
		NumericTypesRangeValidationExclusive: c.HandleNumericTypesRangeValidationExclusive.httpHandlerFactory,
	}
}

func BuildNumericTypesController() *NumericTypesControllerBuilder {
	controllerBuilder := &NumericTypesControllerBuilder{}

	// GET /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesParsing.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesParsing.voidResult = true
	controllerBuilder.HandleNumericTypesParsing.paramsParser = newNumericTypesNumericTypesParsingParamsParser()

	// GET /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesRangeValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesRangeValidation.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesRangeValidation.voidResult = true
	controllerBuilder.HandleNumericTypesRangeValidation.paramsParser = newNumericTypesNumericTypesRangeValidationParamsParser()

	// GET /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.voidResult = true
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.paramsParser = newNumericTypesNumericTypesRangeValidationExclusiveParamsParser()

	return controllerBuilder
}

func MountNumericTypesRoutes(controller *NumericTypesController, app *httpApp) {
	app.router.HandleRoute("GET", "/numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesParsing(app))
	app.router.HandleRoute("GET", "/numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidation(app))
	app.router.HandleRoute("GET", "/numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidationExclusive(app))
}
