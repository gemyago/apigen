package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type numericTypesControllerTestActions struct {
	numericTypesParsing                  mockAction[*handlers.NumericTypesNumericTypesParsingRequest]
	numericTypesArraysParsing            mockAction[*handlers.NumericTypesNumericTypesArraysParsingRequest]
	numericTypesRangeValidation          mockAction[*handlers.NumericTypesNumericTypesRangeValidationRequest]
	numericTypesRangeValidationExclusive mockAction[*handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest]
	numericTypesRequiredValidation       mockAction[*handlers.NumericTypesNumericTypesRequiredValidationRequest]
	numericTypesNullable                 mockAction[*handlers.NumericTypesNumericTypesNullableRequest]
}

func newNumericTypesController(
	testActions *numericTypesControllerTestActions,
) *handlers.NumericTypesController {
	return handlers.BuildNumericTypesController().
		HandleNumericTypesParsing.With(testActions.numericTypesParsing.action).
		HandleNumericTypesArraysParsing.With(testActions.numericTypesArraysParsing.action).
		HandleNumericTypesRangeValidation.With(testActions.numericTypesRangeValidation.action).
		HandleNumericTypesRangeValidationExclusive.With(testActions.numericTypesRangeValidationExclusive.action).
		HandleNumericTypesRequiredValidation.With(testActions.numericTypesRequiredValidation.action).
		HandleNumericTypesNullable.With(testActions.numericTypesNullable.action).
		Finalize()
}
