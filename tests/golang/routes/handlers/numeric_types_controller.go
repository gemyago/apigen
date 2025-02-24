package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint
type _ func() NumericTypesArrayItemsRequest

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

type numericTypesControllerBuilder struct {
	// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesArrayItemsRequest,
	//
	// Response type: none
	NumericTypesArrayItems genericHandlerBuilder[
		*NumericTypesNumericTypesArrayItemsRequest,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesArrayItemsRequest, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesArrayItemsRequest, void],
	]

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableRequest,
	//
	// Response type: none
	NumericTypesNullable genericHandlerBuilder[
		*NumericTypesNumericTypesNullableRequest,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesNullableRequest, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesNullableRequest, void],
	]

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableArrayItemsRequest,
	//
	// Response type: none
	NumericTypesNullableArrayItems genericHandlerBuilder[
		*NumericTypesNumericTypesNullableArrayItemsRequest,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesNullableArrayItemsRequest, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesNullableArrayItemsRequest, void],
	]

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesParsingRequest,
	//
	// Response type: none
	NumericTypesParsing genericHandlerBuilder[
		*NumericTypesNumericTypesParsingRequest,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesParsingRequest, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesParsingRequest, void],
	]

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationRequest,
	//
	// Response type: none
	NumericTypesRangeValidation genericHandlerBuilder[
		*NumericTypesNumericTypesRangeValidationRequest,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationRequest, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationRequest, void],
	]

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationExclusiveRequest,
	//
	// Response type: none
	NumericTypesRangeValidationExclusive genericHandlerBuilder[
		*NumericTypesNumericTypesRangeValidationExclusiveRequest,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationExclusiveRequest, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesRangeValidationExclusiveRequest, void],
	]

	// GET /numeric-types/required-validation
	//
	// Request type: NumericTypesNumericTypesRequiredValidationRequest,
	//
	// Response type: none
	NumericTypesRequiredValidation genericHandlerBuilder[
		*NumericTypesNumericTypesRequiredValidationRequest,
		void,
		handlerActionFuncNoResponse[*NumericTypesNumericTypesRequiredValidationRequest, void],
		httpHandlerActionFuncNoResponse[*NumericTypesNumericTypesRequiredValidationRequest, void],
	]
}

func newNumericTypesControllerBuilder(app *RootHandler) *numericTypesControllerBuilder {
	return &numericTypesControllerBuilder{
		// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesArrayItemsRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesArrayItemsRequest,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesArrayItemsRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesArrayItems(app),
			},
		),

		// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesNullable: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesNullableRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesNullableRequest,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesNullableRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesNullable(app),
			},
		),

		// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesNullableArrayItems: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesNullableArrayItemsRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesNullableArrayItemsRequest,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesNullableArrayItemsRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesNullableArrayItems(app),
			},
		),

		// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesParsing: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesParsingRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesParsingRequest,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesParsingRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesParsing(app),
			},
		),

		// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRangeValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRangeValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesRangeValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesRangeValidation(app),
			},
		),

		// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
		NumericTypesRangeValidationExclusive: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRangeValidationExclusiveRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRangeValidationExclusiveRequest,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesRangeValidationExclusiveRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesRangeValidationExclusive(app),
			},
		),

		// GET /numeric-types/required-validation
		NumericTypesRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRequiredValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*NumericTypesNumericTypesRequiredValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*NumericTypesNumericTypesRequiredValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserNumericTypesNumericTypesRequiredValidation(app),
			},
		),
	}
}

type NumericTypesController interface {
	// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesArrayItemsRequest,
	//
	// Response type: none
	NumericTypesArrayItems(NoResponseHandlerBuilder[
		*NumericTypesNumericTypesArrayItemsRequest,
	]) http.Handler

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableRequest,
	//
	// Response type: none
	NumericTypesNullable(NoResponseHandlerBuilder[
		*NumericTypesNumericTypesNullableRequest,
	]) http.Handler

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesNullableArrayItemsRequest,
	//
	// Response type: none
	NumericTypesNullableArrayItems(NoResponseHandlerBuilder[
		*NumericTypesNumericTypesNullableArrayItemsRequest,
	]) http.Handler

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesParsingRequest,
	//
	// Response type: none
	NumericTypesParsing(NoResponseHandlerBuilder[
		*NumericTypesNumericTypesParsingRequest,
	]) http.Handler

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationRequest,
	//
	// Response type: none
	NumericTypesRangeValidation(NoResponseHandlerBuilder[
		*NumericTypesNumericTypesRangeValidationRequest,
	]) http.Handler

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNumericTypesRangeValidationExclusiveRequest,
	//
	// Response type: none
	NumericTypesRangeValidationExclusive(NoResponseHandlerBuilder[
		*NumericTypesNumericTypesRangeValidationExclusiveRequest,
	]) http.Handler

	// GET /numeric-types/required-validation
	//
	// Request type: NumericTypesNumericTypesRequiredValidationRequest,
	//
	// Response type: none
	NumericTypesRequiredValidation(NoResponseHandlerBuilder[
		*NumericTypesNumericTypesRequiredValidationRequest,
	]) http.Handler
}

// RegisterNumericTypesRoutes will attach the following routes to the root handler:
// 
// - POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
// 
// - POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
// 
// - POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
// 
// - POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
// 
// - POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
// 
// - POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
// 
// - GET /numeric-types/required-validation
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterNumericTypesRoutes(controller NumericTypesController) *RootHandler {
	builder := newNumericTypesControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesArrayItems(builder.NumericTypesArrayItems))
	rootHandler.router.HandleRoute("POST", "/numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesNullable(builder.NumericTypesNullable))
	rootHandler.router.HandleRoute("POST", "/numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesNullableArrayItems(builder.NumericTypesNullableArrayItems))
	rootHandler.router.HandleRoute("POST", "/numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesParsing(builder.NumericTypesParsing))
	rootHandler.router.HandleRoute("POST", "/numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidation(builder.NumericTypesRangeValidation))
	rootHandler.router.HandleRoute("POST", "/numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}", controller.NumericTypesRangeValidationExclusive(builder.NumericTypesRangeValidationExclusive))
	rootHandler.router.HandleRoute("GET", "/numeric-types/required-validation", controller.NumericTypesRequiredValidation(builder.NumericTypesRequiredValidation))
	return rootHandler
}