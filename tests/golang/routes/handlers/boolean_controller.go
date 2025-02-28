package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() BooleanArrayItemsRequest

type BooleanController interface {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanArrayItemsParams,
	//
	// Response type: none
	BooleanArrayItems(NoResponseHandlerBuilder[
		*BooleanArrayItemsParams,
	]) http.Handler

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanNullableParams,
	//
	// Response type: none
	BooleanNullable(NoResponseHandlerBuilder[
		*BooleanNullableParams,
	]) http.Handler

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanNullableArrayItemsParams,
	//
	// Response type: none
	BooleanNullableArrayItems(NoResponseHandlerBuilder[
		*BooleanNullableArrayItemsParams,
	]) http.Handler

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanParsingParams,
	//
	// Response type: none
	BooleanParsing(NoResponseHandlerBuilder[
		*BooleanParsingParams,
	]) http.Handler

	// POST /boolean/required-validation
	//
	// Request type: BooleanRequiredValidationParams,
	//
	// Response type: none
	BooleanRequiredValidation(NoResponseHandlerBuilder[
		*BooleanRequiredValidationParams,
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