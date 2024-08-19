package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type NumericTypesNumericTypesArraysParsingRequest struct {
	NumberAny []float32
	NumberFloat []float32
	NumberDouble []float64
	NumberInt []int32
	NumberInt32 []int32
	NumberInt64 []int64
	NumberAnyInQuery []float32
	NumberFloatInQuery []float32
	NumberDoubleInQuery []float64
	NumberIntInQuery []int32
	NumberInt32InQuery []int32
	NumberInt64InQuery []int64
	Payload *models.NumericTypesArraysParsingRequest
}

type NumericTypesNumericTypesNullableRequest struct {
	NumberAny *float32
	NumberFloat *float32
	NumberDouble *float64
	NumberInt *int32
	NumberInt32 *int32
	NumberInt64 *int64
	NumberAnyInQuery *float32
	NumberFloatInQuery *float32
	NumberDoubleInQuery *float64
	NumberIntInQuery *int32
	NumberInt32InQuery *int32
	NumberInt64InQuery *int64
	Payload *models.NumericTypesNullableRequest
	OptionalNumberAnyInQuery *float32
}

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
	Payload *models.NumericTypesParsingRequest
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
	Payload *models.NumericTypesRangeValidationRequest
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
	Payload *models.NumericTypesRangeValidationExclusiveRequest
}

type NumericTypesNumericTypesRequiredValidationRequest struct {
	NumberAnyInQuery float32
	NumberFloatInQuery float32
	NumberDoubleInQuery float64
	NumberIntInQuery int32
	NumberInt32InQuery int32
	NumberInt64InQuery int64
	OptionalNumberAnyInQuery float32
	OptionalNumberFloatInQuery float32
	OptionalNumberDoubleInQuery float64
	OptionalNumberIntInQuery int32
	OptionalNumberInt32InQuery int32
	OptionalNumberInt64InQuery int64
}

type NumericTypesController struct {
	// POST /numeric-types/arrays-parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesArraysParsingRequest,
	//
	// Response type: none
	NumericTypesArraysParsing httpHandlerFactory

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableRequest,
	//
	// Response type: none
	NumericTypesNullable httpHandlerFactory

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesParsingRequest,
	//
	// Response type: none
	NumericTypesParsing httpHandlerFactory

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationRequest,
	//
	// Response type: none
	NumericTypesRangeValidation httpHandlerFactory

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationExclusiveRequest,
	//
	// Response type: none
	NumericTypesRangeValidationExclusive httpHandlerFactory

	// GET /numeric-types/required-validation
	//
	// Request type: NumericTypesNumericTypesRequiredValidationRequest,
	//
	// Response type: none
	NumericTypesRequiredValidation httpHandlerFactory
}

type NumericTypesControllerBuilder struct {
	// POST /numeric-types/arrays-parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesArraysParsingRequest,
	//
	// Response type: none
	HandleNumericTypesArraysParsing actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesArraysParsingRequest]

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableRequest,
	//
	// Response type: none
	HandleNumericTypesNullable actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesNullableRequest]

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesParsingRequest,
	//
	// Response type: none
	HandleNumericTypesParsing actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesParsingRequest]

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationRequest,
	//
	// Response type: none
	HandleNumericTypesRangeValidation actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesRangeValidationRequest]

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationExclusiveRequest,
	//
	// Response type: none
	HandleNumericTypesRangeValidationExclusive actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesRangeValidationExclusiveRequest]

	// GET /numeric-types/required-validation
	//
	// Request type: NumericTypesNumericTypesRequiredValidationRequest,
	//
	// Response type: none
	HandleNumericTypesRequiredValidation actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesRequiredValidationRequest]
}

func (c *NumericTypesControllerBuilder) Finalize() *NumericTypesController {
	return &NumericTypesController{
		NumericTypesArraysParsing: mustInitializeAction("numericTypesArraysParsing", c.HandleNumericTypesArraysParsing.httpHandlerFactory),
		NumericTypesNullable: mustInitializeAction("numericTypesNullable", c.HandleNumericTypesNullable.httpHandlerFactory),
		NumericTypesParsing: mustInitializeAction("numericTypesParsing", c.HandleNumericTypesParsing.httpHandlerFactory),
		NumericTypesRangeValidation: mustInitializeAction("numericTypesRangeValidation", c.HandleNumericTypesRangeValidation.httpHandlerFactory),
		NumericTypesRangeValidationExclusive: mustInitializeAction("numericTypesRangeValidationExclusive", c.HandleNumericTypesRangeValidationExclusive.httpHandlerFactory),
		NumericTypesRequiredValidation: mustInitializeAction("numericTypesRequiredValidation", c.HandleNumericTypesRequiredValidation.httpHandlerFactory),
	}
}

func BuildNumericTypesController() *NumericTypesControllerBuilder {
	controllerBuilder := &NumericTypesControllerBuilder{}

	// POST /numeric-types/arrays-parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesArraysParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesArraysParsing.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesArraysParsing.voidResult = true
	controllerBuilder.HandleNumericTypesArraysParsing.paramsParserFactory = newParamsParserNumericTypesNumericTypesArraysParsing

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesNullable.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesNullable.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesNullable.voidResult = true
	controllerBuilder.HandleNumericTypesNullable.paramsParserFactory = newParamsParserNumericTypesNumericTypesNullable

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesParsing.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesParsing.voidResult = true
	controllerBuilder.HandleNumericTypesParsing.paramsParserFactory = newParamsParserNumericTypesNumericTypesParsing

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesRangeValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesRangeValidation.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesRangeValidation.voidResult = true
	controllerBuilder.HandleNumericTypesRangeValidation.paramsParserFactory = newParamsParserNumericTypesNumericTypesRangeValidation

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.voidResult = true
	controllerBuilder.HandleNumericTypesRangeValidationExclusive.paramsParserFactory = newParamsParserNumericTypesNumericTypesRangeValidationExclusive

	// GET /numeric-types/required-validation
	controllerBuilder.HandleNumericTypesRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesRequiredValidation.voidResult = true
	controllerBuilder.HandleNumericTypesRequiredValidation.paramsParserFactory = newParamsParserNumericTypesNumericTypesRequiredValidation

	return controllerBuilder
}

func RegisterNumericTypesRoutes(controller *NumericTypesController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/numeric-types/arrays-parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesArraysParsing(app))
	app.router.HandleRoute("POST", "/numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesNullable(app))
	app.router.HandleRoute("POST", "/numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesParsing(app))
	app.router.HandleRoute("POST", "/numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidation(app))
	app.router.HandleRoute("POST", "/numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidationExclusive(app))
	app.router.HandleRoute("GET", "/numeric-types/required-validation", controller.NumericTypesRequiredValidation(app))
}
