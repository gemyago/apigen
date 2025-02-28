package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type objectsControllerTestActions struct {
	objectsArrayParsingBodyDirect mockAction[*models.ObjectsArrayBodyDirectParams]
	objectsArrayParsingBodyNested mockAction[*models.ObjectsArrayBodyNestedParams]
	objectsDeeplyNested           mockAction[*models.ObjectsDeeplyNestedParams]
	objectsNullableOptionalBody   mockAction[*models.ObjectsNullableOptionalBodyParams]
	objectsNullableRequiredBody   mockAction[*models.ObjectsNullableRequiredBodyParams]
	objectsOptionalBody           mockAction[*models.ObjectsOptionalBodyParams]
	objectsRequiredBody           mockAction[*models.ObjectsRequiredBodyParams]
	objectsRequiredNestedObjects  mockAction[*models.ObjectsRequiredNestedObjectsParams]
}

type objectsController struct {
	testActions *objectsControllerTestActions
}

func (c *objectsController) ObjectsArrayBodyDirect(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsArrayBodyDirectParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsArrayParsingBodyDirect.action,
	)
}

func (c *objectsController) ObjectsArrayBodyNested(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsArrayBodyNestedParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsArrayParsingBodyNested.action,
	)
}

func (c *objectsController) ObjectsDeeplyNested(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsDeeplyNestedParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsDeeplyNested.action,
	)
}

func (c *objectsController) ObjectsNullableOptionalBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsNullableOptionalBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsNullableOptionalBody.action,
	)
}

func (c *objectsController) ObjectsNullableRequiredBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsNullableRequiredBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsNullableRequiredBody.action,
	)
}

func (c *objectsController) ObjectsOptionalBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsOptionalBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsOptionalBody.action,
	)
}

func (c *objectsController) ObjectsRequiredBody(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsRequiredBodyParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsRequiredBody.action,
	)
}

func (c *objectsController) ObjectsRequiredNestedObjects(
	builder handlers.NoResponseHandlerBuilder[*models.ObjectsRequiredNestedObjectsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.objectsRequiredNestedObjects.action,
	)
}

var _ handlers.ObjectsController = &objectsController{}
