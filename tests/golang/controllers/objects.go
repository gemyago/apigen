package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type objectsControllerTestActions struct {
	objectsArrayParsingBodyDirect mockAction[*handlers.ObjectsObjectsArrayBodyDirectRequest]
	objectsArrayParsingBodyNested mockAction[*handlers.ObjectsObjectsArrayBodyNestedRequest]
	objectsDeeplyNested           mockAction[*handlers.ObjectsObjectsDeeplyNestedRequest]
	objectsNullableOptionalBody   mockAction[*handlers.ObjectsObjectsNullableOptionalBodyRequest]
	objectsNullableRequiredBody   mockAction[*handlers.ObjectsObjectsNullableRequiredBodyRequest]
	objectsOptionalBody           mockAction[*handlers.ObjectsObjectsOptionalBodyRequest]
	objectsRequiredBody           mockAction[*handlers.ObjectsObjectsRequiredBodyRequest]
	objectsRequiredNestedObjects  mockAction[*handlers.ObjectsObjectsRequiredNestedObjectsRequest]
}

type objectsController struct {
	testActions *objectsControllerTestActions
}

func (c *objectsController) ObjectsArrayBodyDirect(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsArrayBodyDirectRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsArrayParsingBodyDirect.action,
	)
}

func (c *objectsController) ObjectsArrayBodyNested(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsArrayBodyNestedRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsArrayParsingBodyNested.action,
	)
}

func (c *objectsController) ObjectsDeeplyNested(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsDeeplyNestedRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsDeeplyNested.action,
	)
}

func (c *objectsController) ObjectsNullableOptionalBody(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsNullableOptionalBodyRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsNullableOptionalBody.action,
	)
}

func (c *objectsController) ObjectsNullableRequiredBody(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsNullableRequiredBodyRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsNullableRequiredBody.action,
	)
}

func (c *objectsController) ObjectsOptionalBody(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsOptionalBodyRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsOptionalBody.action,
	)
}

func (c *objectsController) ObjectsRequiredBody(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsRequiredBodyRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsRequiredBody.action,
	)
}

func (c *objectsController) ObjectsRequiredNestedObjects(
	builder handlers.NoResponseHandlerBuilder[*handlers.ObjectsObjectsRequiredNestedObjectsRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsRequiredNestedObjects.action,
	)
}

var _ handlers.ObjectsController = &objectsController{}
