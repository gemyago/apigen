package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type objectsControllerTestActions struct {
	objectsArrayParsingBodyDirect mockAction[*models.ObjectsObjectsArrayBodyDirectParams]
	objectsArrayParsingBodyNested mockAction[*models.ObjectsObjectsArrayBodyNestedParams]
	objectsDeeplyNested           mockAction[*models.ObjectsObjectsDeeplyNestedParams]
	objectsNullableOptionalBody   mockAction[*models.ObjectsObjectsNullableOptionalBodyParams]
	objectsNullableRequiredBody   mockAction[*models.ObjectsObjectsNullableRequiredBodyParams]
	objectsOptionalBody           mockAction[*models.ObjectsObjectsOptionalBodyParams]
	objectsRequiredBody           mockAction[*models.ObjectsObjectsRequiredBodyParams]
	objectsRequiredNestedObjects  mockAction[*models.ObjectsObjectsRequiredNestedObjectsParams]
}

type objectsController struct {
	testActions *objectsControllerTestActions
}

func (c *objectsController) ObjectsArrayBodyDirect(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsArrayBodyDirectParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsArrayParsingBodyDirect.action,
	)
}

func (c *objectsController) ObjectsArrayBodyNested(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsArrayBodyNestedParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsArrayParsingBodyNested.action,
	)
}

func (c *objectsController) ObjectsDeeplyNested(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsDeeplyNestedParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsDeeplyNested.action,
	)
}

func (c *objectsController) ObjectsNullableOptionalBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsNullableOptionalBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsNullableOptionalBody.action,
	)
}

func (c *objectsController) ObjectsNullableRequiredBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsNullableRequiredBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsNullableRequiredBody.action,
	)
}

func (c *objectsController) ObjectsOptionalBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsOptionalBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsOptionalBody.action,
	)
}

func (c *objectsController) ObjectsRequiredBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsRequiredBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsRequiredBody.action,
	)
}

func (c *objectsController) ObjectsRequiredNestedObjects(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsObjectsRequiredNestedObjectsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsRequiredNestedObjects.action,
	)
}

var _ handlers.ObjectsController = &objectsController{}
