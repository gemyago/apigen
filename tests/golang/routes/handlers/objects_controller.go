package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() ObjectArraysSimpleObject

type ObjectsController interface {
	// POST /objects/arrays
	//
	// Request type: ObjectsArrayBodyDirectParams,
	//
	// Response type: none
	ObjectsArrayBodyDirect(NoResponseHandlerBuilder[
		*ObjectsArrayBodyDirectParams,
	]) http.Handler

	// PUT /objects/arrays
	//
	// Request type: ObjectsArrayBodyNestedParams,
	//
	// Response type: none
	ObjectsArrayBodyNested(NoResponseHandlerBuilder[
		*ObjectsArrayBodyNestedParams,
	]) http.Handler

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsDeeplyNestedParams,
	//
	// Response type: none
	ObjectsDeeplyNested(NoResponseHandlerBuilder[
		*ObjectsDeeplyNestedParams,
	]) http.Handler

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsNullableOptionalBodyParams,
	//
	// Response type: none
	ObjectsNullableOptionalBody(NoResponseHandlerBuilder[
		*ObjectsNullableOptionalBodyParams,
	]) http.Handler

	// POST /objects/nullable-body
	//
	// Request type: ObjectsNullableRequiredBodyParams,
	//
	// Response type: none
	ObjectsNullableRequiredBody(NoResponseHandlerBuilder[
		*ObjectsNullableRequiredBodyParams,
	]) http.Handler

	// PUT /objects/required-body
	//
	// Request type: ObjectsOptionalBodyParams,
	//
	// Response type: none
	ObjectsOptionalBody(NoResponseHandlerBuilder[
		*ObjectsOptionalBodyParams,
	]) http.Handler

	// POST /objects/required-body
	//
	// Request type: ObjectsRequiredBodyParams,
	//
	// Response type: none
	ObjectsRequiredBody(NoResponseHandlerBuilder[
		*ObjectsRequiredBodyParams,
	]) http.Handler

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsRequiredNestedObjectsParams,
	//
	// Response type: none
	ObjectsRequiredNestedObjects(NoResponseHandlerBuilder[
		*ObjectsRequiredNestedObjectsParams,
	]) http.Handler
}

// RegisterObjectsRoutes will attach the following routes to the root handler:
// 
// - POST /objects/arrays
// 
// - PUT /objects/arrays
// 
// - POST /objects/deeply-nested
// 
// - PUT /objects/nullable-body
// 
// - POST /objects/nullable-body
// 
// - PUT /objects/required-body
// 
// - POST /objects/required-body
// 
// - POST /objects/required-nested-objects
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterObjectsRoutes(controller ObjectsController) *RootHandler {
	builder := newObjectsControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/objects/arrays", controller.ObjectsArrayBodyDirect(builder.ObjectsArrayBodyDirect))
	rootHandler.router.HandleRoute("PUT", "/objects/arrays", controller.ObjectsArrayBodyNested(builder.ObjectsArrayBodyNested))
	rootHandler.router.HandleRoute("POST", "/objects/deeply-nested", controller.ObjectsDeeplyNested(builder.ObjectsDeeplyNested))
	rootHandler.router.HandleRoute("PUT", "/objects/nullable-body", controller.ObjectsNullableOptionalBody(builder.ObjectsNullableOptionalBody))
	rootHandler.router.HandleRoute("POST", "/objects/nullable-body", controller.ObjectsNullableRequiredBody(builder.ObjectsNullableRequiredBody))
	rootHandler.router.HandleRoute("PUT", "/objects/required-body", controller.ObjectsOptionalBody(builder.ObjectsOptionalBody))
	rootHandler.router.HandleRoute("POST", "/objects/required-body", controller.ObjectsRequiredBody(builder.ObjectsRequiredBody))
	rootHandler.router.HandleRoute("POST", "/objects/required-nested-objects", controller.ObjectsRequiredNestedObjects(builder.ObjectsRequiredNestedObjects))
	return rootHandler
}