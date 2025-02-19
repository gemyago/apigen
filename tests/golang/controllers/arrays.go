package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type arraysControllerTestActions struct {
	arraysRequiredValidation         mockAction[*handlers.ArraysArraysRequiredValidationRequest]
	arraysNullableRequiredValidation mockAction[*handlers.ArraysArraysNullableRequiredValidationRequest]
	arraysRangeValidation            mockAction[*handlers.ArraysArraysRangeValidationRequest]
}

type arraysController struct {
	testActions *arraysControllerTestActions
}

func (c *arraysController) ArraysRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.ArraysArraysRequiredValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.arraysRequiredValidation.action,
	)
}

func (c *arraysController) ArraysNullableRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.ArraysArraysNullableRequiredValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.arraysNullableRequiredValidation.action,
	)
}

func (c *arraysController) ArraysRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.ArraysArraysRangeValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.arraysRangeValidation.action,
	)
}

var _ handlers.ArraysController = &arraysController{}
