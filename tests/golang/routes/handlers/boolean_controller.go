package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

type BooleanController interface {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsParams,
	//
	// Response type: none
	BooleanArrayItems(NoResponseHandlerBuilder[
		*BooleanBooleanArrayItemsParams,
	]) http.Handler

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableParams,
	//
	// Response type: none
	BooleanNullable(NoResponseHandlerBuilder[
		*BooleanBooleanNullableParams,
	]) http.Handler

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsParams,
	//
	// Response type: none
	BooleanNullableArrayItems(NoResponseHandlerBuilder[
		*BooleanBooleanNullableArrayItemsParams,
	]) http.Handler

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingParams,
	//
	// Response type: none
	BooleanParsing(NoResponseHandlerBuilder[
		*BooleanBooleanParsingParams,
	]) http.Handler

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationParams,
	//
	// Response type: none
	BooleanRequiredValidation(NoResponseHandlerBuilder[
		*BooleanBooleanRequiredValidationParams,
	]) http.Handler
}

// RegisterBooleanRoutes will attach the following routes to the root handler:
// 
// - POST /boolean/array-items/{boolParam1}/{boolParam2}
// 
// - POST /boolean/nullable/{boolParam1}/{boolParam2}
// 
// - POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
// 
// - POST /boolean/parsing/{boolParam1}/{boolParam2}
// 
// - POST /boolean/required-validation
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterBooleanRoutes(controller BooleanController) *RootHandler {
	builder := newBooleanControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/boolean/array-items/{boolParam1}/{boolParam2}", controller.BooleanArrayItems(builder.BooleanArrayItems))
	rootHandler.router.HandleRoute("POST", "/boolean/nullable/{boolParam1}/{boolParam2}", controller.BooleanNullable(builder.BooleanNullable))
	rootHandler.router.HandleRoute("POST", "/boolean/nullable-array-items/{boolParam1}/{boolParam2}", controller.BooleanNullableArrayItems(builder.BooleanNullableArrayItems))
	rootHandler.router.HandleRoute("POST", "/boolean/parsing/{boolParam1}/{boolParam2}", controller.BooleanParsing(builder.BooleanParsing))
	rootHandler.router.HandleRoute("POST", "/boolean/required-validation", controller.BooleanRequiredValidation(builder.BooleanRequiredValidation))
	return rootHandler
}