package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() NumericTypesArrayItemsRequest

type NumericTypesController interface {
	// POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesArrayItemsParams,
	//
	// Response type: none
	NumericTypesArrayItems(NoResponseHandlerBuilder[
		*NumericTypesArrayItemsParams,
	]) http.Handler

	// POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNullableParams,
	//
	// Response type: none
	NumericTypesNullable(NoResponseHandlerBuilder[
		*NumericTypesNullableParams,
	]) http.Handler

	// POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesNullableArrayItemsParams,
	//
	// Response type: none
	NumericTypesNullableArrayItems(NoResponseHandlerBuilder[
		*NumericTypesNullableArrayItemsParams,
	]) http.Handler

	// POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesParsingParams,
	//
	// Response type: none
	NumericTypesParsing(NoResponseHandlerBuilder[
		*NumericTypesParsingParams,
	]) http.Handler

	// POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesRangeValidationParams,
	//
	// Response type: none
	NumericTypesRangeValidation(NoResponseHandlerBuilder[
		*NumericTypesRangeValidationParams,
	]) http.Handler

	// POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}
	//
	// Request type: NumericTypesRangeValidationExclusiveParams,
	//
	// Response type: none
	NumericTypesRangeValidationExclusive(NoResponseHandlerBuilder[
		*NumericTypesRangeValidationExclusiveParams,
	]) http.Handler

	// GET /numeric-types/required-validation
	//
	// Request type: NumericTypesRequiredValidationParams,
	//
	// Response type: none
	NumericTypesRequiredValidation(NoResponseHandlerBuilder[
		*NumericTypesRequiredValidationParams,
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