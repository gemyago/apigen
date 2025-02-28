package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

type ObjectsController interface {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectParams,
	//
	// Response type: none
	ObjectsArrayBodyDirect(NoResponseHandlerBuilder[
		*ObjectsObjectsArrayBodyDirectParams,
	]) http.Handler

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedParams,
	//
	// Response type: none
	ObjectsArrayBodyNested(NoResponseHandlerBuilder[
		*ObjectsObjectsArrayBodyNestedParams,
	]) http.Handler

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedParams,
	//
	// Response type: none
	ObjectsDeeplyNested(NoResponseHandlerBuilder[
		*ObjectsObjectsDeeplyNestedParams,
	]) http.Handler

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyParams,
	//
	// Response type: none
	ObjectsNullableOptionalBody(NoResponseHandlerBuilder[
		*ObjectsObjectsNullableOptionalBodyParams,
	]) http.Handler

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyParams,
	//
	// Response type: none
	ObjectsNullableRequiredBody(NoResponseHandlerBuilder[
		*ObjectsObjectsNullableRequiredBodyParams,
	]) http.Handler

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyParams,
	//
	// Response type: none
	ObjectsOptionalBody(NoResponseHandlerBuilder[
		*ObjectsObjectsOptionalBodyParams,
	]) http.Handler

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyParams,
	//
	// Response type: none
	ObjectsRequiredBody(NoResponseHandlerBuilder[
		*ObjectsObjectsRequiredBodyParams,
	]) http.Handler

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsParams,
	//
	// Response type: none
	ObjectsRequiredNestedObjects(NoResponseHandlerBuilder[
		*ObjectsObjectsRequiredNestedObjectsParams,
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