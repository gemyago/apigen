package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() ArraysNullableRequiredValidationRequest

type ArraysController interface {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysNullableRequiredValidationParams,
	//
	// Response type: none
	ArraysNullableRequiredValidation(NoResponseHandlerBuilder[
		*ArraysArraysNullableRequiredValidationParams,
	]) http.Handler

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRangeValidationParams,
	//
	// Response type: none
	ArraysRangeValidation(NoResponseHandlerBuilder[
		*ArraysArraysRangeValidationParams,
	]) http.Handler

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRequiredValidationParams,
	//
	// Response type: none
	ArraysRequiredValidation(NoResponseHandlerBuilder[
		*ArraysArraysRequiredValidationParams,
	]) http.Handler
}

// RegisterArraysRoutes will attach the following routes to the root handler:
// 
// - POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
// 
// - POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
// 
// - POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterArraysRoutes(controller ArraysController) *RootHandler {
	builder := newArraysControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysNullableRequiredValidation(builder.ArraysNullableRequiredValidation))
	rootHandler.router.HandleRoute("POST", "/arrays/range-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRangeValidation(builder.ArraysRangeValidation))
	rootHandler.router.HandleRoute("POST", "/arrays/required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRequiredValidation(builder.ArraysRequiredValidation))
	return rootHandler
}