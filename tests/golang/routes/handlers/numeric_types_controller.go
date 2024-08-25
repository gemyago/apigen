package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint

// NumericTypesNumericTypesArrayItemsRequest represents params for numericTypesArrayItems operation
//
// Request: POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNumericTypesArrayItemsRequest struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny []float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat []float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble []float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt []int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 []int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 []int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery []float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery []float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery []float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery []int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery []int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery []int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesArrayItemsRequest
}

// NumericTypesNumericTypesNullableRequest represents params for numericTypesNullable operation
//
// Request: POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNumericTypesNullableRequest struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny *float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat *float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble *float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt *int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 *int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 *int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery *float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery *float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery *float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery *int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery *int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery *int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesNullableRequest
	// OptionalNumberAnyInQuery is parsed from request query and declared as optionalNumberAnyInQuery.
	OptionalNumberAnyInQuery *float32
}

// NumericTypesNumericTypesNullableArrayItemsRequest represents params for numericTypesNullableArrayItems operation
//
// Request: POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNumericTypesNullableArrayItemsRequest struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny []*float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat []*float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble []*float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt []*int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 []*int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 []*int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery []*float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery []*float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery []*float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery []*int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery []*int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery []*int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesNullableArrayItemsRequest
}

// NumericTypesNumericTypesParsingRequest represents params for numericTypesParsing operation
//
// Request: POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNumericTypesParsingRequest struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesParsingRequest
}

// NumericTypesNumericTypesRangeValidationRequest represents params for numericTypesRangeValidation operation
//
// Request: POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNumericTypesRangeValidationRequest struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesRangeValidationRequest
}

// NumericTypesNumericTypesRangeValidationExclusiveRequest represents params for numericTypesRangeValidationExclusive operation
//
// Request: POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNumericTypesRangeValidationExclusiveRequest struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesRangeValidationExclusiveRequest
}

// NumericTypesNumericTypesRequiredValidationRequest represents params for numericTypesRequiredValidation operation
//
// Request: GET /numeric-types/required-validation.
type NumericTypesNumericTypesRequiredValidationRequest struct {
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// OptionalNumberAnyInQuery is parsed from request query and declared as optionalNumberAnyInQuery.
	OptionalNumberAnyInQuery float32
	// OptionalNumberFloatInQuery is parsed from request query and declared as optionalNumberFloatInQuery.
	OptionalNumberFloatInQuery float32
	// OptionalNumberDoubleInQuery is parsed from request query and declared as optionalNumberDoubleInQuery.
	OptionalNumberDoubleInQuery float64
	// OptionalNumberIntInQuery is parsed from request query and declared as optionalNumberIntInQuery.
	OptionalNumberIntInQuery int32
	// OptionalNumberInt32InQuery is parsed from request query and declared as optionalNumberInt32InQuery.
	OptionalNumberInt32InQuery int32
	// OptionalNumberInt64InQuery is parsed from request query and declared as optionalNumberInt64InQuery.
	OptionalNumberInt64InQuery int64
}

type NumericTypesController struct {
	// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesArrayItemsRequest,
	//
	// Response type: none
	NumericTypesArrayItems httpHandlerFactory

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableRequest,
	//
	// Response type: none
	NumericTypesNullable httpHandlerFactory

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableArrayItemsRequest,
	//
	// Response type: none
	NumericTypesNullableArrayItems httpHandlerFactory

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
	// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesArrayItemsRequest,
	//
	// Response type: none
	HandleNumericTypesArrayItems actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesArrayItemsRequest]

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableRequest,
	//
	// Response type: none
	HandleNumericTypesNullable actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesNullableRequest]

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableArrayItemsRequest,
	//
	// Response type: none
	HandleNumericTypesNullableArrayItems actionBuilderVoidResult[*NumericTypesControllerBuilder, *NumericTypesNumericTypesNullableArrayItemsRequest]

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
		NumericTypesArrayItems: mustInitializeAction("numericTypesArrayItems", c.HandleNumericTypesArrayItems.httpHandlerFactory),
		NumericTypesNullable: mustInitializeAction("numericTypesNullable", c.HandleNumericTypesNullable.httpHandlerFactory),
		NumericTypesNullableArrayItems: mustInitializeAction("numericTypesNullableArrayItems", c.HandleNumericTypesNullableArrayItems.httpHandlerFactory),
		NumericTypesParsing: mustInitializeAction("numericTypesParsing", c.HandleNumericTypesParsing.httpHandlerFactory),
		NumericTypesRangeValidation: mustInitializeAction("numericTypesRangeValidation", c.HandleNumericTypesRangeValidation.httpHandlerFactory),
		NumericTypesRangeValidationExclusive: mustInitializeAction("numericTypesRangeValidationExclusive", c.HandleNumericTypesRangeValidationExclusive.httpHandlerFactory),
		NumericTypesRequiredValidation: mustInitializeAction("numericTypesRequiredValidation", c.HandleNumericTypesRequiredValidation.httpHandlerFactory),
	}
}

func BuildNumericTypesController() *NumericTypesControllerBuilder {
	controllerBuilder := &NumericTypesControllerBuilder{}

	// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesArrayItems.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesArrayItems.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesArrayItems.voidResult = true
	controllerBuilder.HandleNumericTypesArrayItems.paramsParserFactory = newParamsParserNumericTypesNumericTypesArrayItems

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesNullable.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesNullable.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesNullable.voidResult = true
	controllerBuilder.HandleNumericTypesNullable.paramsParserFactory = newParamsParserNumericTypesNumericTypesNullable

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	controllerBuilder.HandleNumericTypesNullableArrayItems.controllerBuilder = controllerBuilder
	controllerBuilder.HandleNumericTypesNullableArrayItems.defaultStatusCode = 204
	controllerBuilder.HandleNumericTypesNullableArrayItems.voidResult = true
	controllerBuilder.HandleNumericTypesNullableArrayItems.paramsParserFactory = newParamsParserNumericTypesNumericTypesNullableArrayItems

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
	app.router.HandleRoute("POST", "/numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesArrayItems(app))
	app.router.HandleRoute("POST", "/numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesNullable(app))
	app.router.HandleRoute("POST", "/numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesNullableArrayItems(app))
	app.router.HandleRoute("POST", "/numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesParsing(app))
	app.router.HandleRoute("POST", "/numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidation(app))
	app.router.HandleRoute("POST", "/numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidationExclusive(app))
	app.router.HandleRoute("GET", "/numeric-types/required-validation", controller.NumericTypesRequiredValidation(app))
}
