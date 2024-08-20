package controllers

import "github.com/gemyago/apigen/tests/golang/routes/handlers"

type arraysControllerTestActions struct {
	arraysRequiredValidation         mockAction[*handlers.ArraysArraysRequiredValidationRequest]
	arraysNullableRequiredValidation mockAction[*handlers.ArraysArraysNullableRequiredValidationRequest]
	arraysRangeValidation            mockAction[*handlers.ArraysArraysRangeValidationRequest]
}

func newArraysController(
	testActions *arraysControllerTestActions,
) *handlers.ArraysController {
	return handlers.BuildArraysController().
		HandleArraysRequiredValidation.With(testActions.arraysRequiredValidation.action).
		HandleArraysNullableRequiredValidation.With(testActions.arraysNullableRequiredValidation.action).
		HandleArraysRangeValidation.With(testActions.arraysRangeValidation.action).
		Finalize()
}
