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
	builder *handlers.ArraysControllerBuilderV2,
) http.Handler {
	return builder.ArraysRequiredValidation.HandleWith(
		c.testActions.arraysRequiredValidation.action,
	)
}

func (c *arraysController) ArraysNullableRequiredValidation(
	builder *handlers.ArraysControllerBuilderV2,
) http.Handler {
	return builder.ArraysNullableRequiredValidation.HandleWith(
		c.testActions.arraysNullableRequiredValidation.action,
	)
}

func (c *arraysController) ArraysRangeValidation(
	builder *handlers.ArraysControllerBuilderV2,
) http.Handler {
	return builder.ArraysRangeValidation.HandleWith(
		c.testActions.arraysRangeValidation.action,
	)
}

var _ handlers.ArraysControllerV2 = &arraysController{}
