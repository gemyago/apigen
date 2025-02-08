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
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsArrayBodyDirect.HandleWith(
		c.testActions.objectsArrayParsingBodyDirect.action,
	)
}

func (c *objectsController) ObjectsArrayBodyNested(
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsArrayBodyNested.HandleWith(
		c.testActions.objectsArrayParsingBodyNested.action,
	)
}

func (c *objectsController) ObjectsDeeplyNested(
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsDeeplyNested.HandleWith(
		c.testActions.objectsDeeplyNested.action,
	)
}

func (c *objectsController) ObjectsNullableOptionalBody(
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsNullableOptionalBody.HandleWith(
		c.testActions.objectsNullableOptionalBody.action,
	)
}

func (c *objectsController) ObjectsNullableRequiredBody(
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsNullableRequiredBody.HandleWith(
		c.testActions.objectsNullableRequiredBody.action,
	)
}

func (c *objectsController) ObjectsOptionalBody(
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsOptionalBody.HandleWith(
		c.testActions.objectsOptionalBody.action,
	)
}

func (c *objectsController) ObjectsRequiredBody(
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsRequiredBody.HandleWith(
		c.testActions.objectsRequiredBody.action,
	)
}

func (c *objectsController) ObjectsRequiredNestedObjects(
	builder *handlers.ObjectsControllerBuilder,
) http.Handler {
	return builder.ObjectsRequiredNestedObjects.HandleWith(
		c.testActions.objectsRequiredNestedObjects.action,
	)
}

var _ handlers.ObjectsController = &objectsController{}
