package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type arraysControllerTestActions struct {
	arraysRequiredValidation         mockAction[*models.ArraysArraysRequiredValidationParams]
	arraysNullableRequiredValidation mockAction[*models.ArraysArraysNullableRequiredValidationParams]
	arraysRangeValidation            mockAction[*models.ArraysArraysRangeValidationParams]
}

type arraysController struct {
	testActions *arraysControllerTestActions
}

func (c *arraysController) ArraysRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.ArraysArraysRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.arraysRequiredValidation.action,
	)
}

func (c *arraysController) ArraysNullableRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.ArraysArraysNullableRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.arraysNullableRequiredValidation.action,
	)
}

func (c *arraysController) ArraysRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*models.ArraysArraysRangeValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.arraysRangeValidation.action,
	)
}

var _ handlers.ArraysController = &arraysController{}
