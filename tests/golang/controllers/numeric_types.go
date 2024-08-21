package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type numericTypesControllerTestActions struct {
	numericTypesParsing                  mockAction[*handlers.NumericTypesNumericTypesParsingRequest]
	numericTypesArrayItems               mockAction[*handlers.NumericTypesNumericTypesArrayItemsRequest]
	numericTypesRangeValidation          mockAction[*handlers.NumericTypesNumericTypesRangeValidationRequest]
	numericTypesRangeValidationExclusive mockAction[*handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest]
	numericTypesRequiredValidation       mockAction[*handlers.NumericTypesNumericTypesRequiredValidationRequest]
	numericTypesNullable                 mockAction[*handlers.NumericTypesNumericTypesNullableRequest]
	numericTypesNullableArrayItems       mockAction[*handlers.NumericTypesNumericTypesNullableArrayItemsRequest]
}

func newNumericTypesController(
	testActions *numericTypesControllerTestActions,
) *handlers.NumericTypesController {
	return handlers.BuildNumericTypesController().
		HandleNumericTypesParsing.With(testActions.numericTypesParsing.action).
		HandleNumericTypesArrayItems.With(testActions.numericTypesArrayItems.action).
		HandleNumericTypesRangeValidation.With(testActions.numericTypesRangeValidation.action).
		HandleNumericTypesRangeValidationExclusive.With(testActions.numericTypesRangeValidationExclusive.action).
		HandleNumericTypesRequiredValidation.With(testActions.numericTypesRequiredValidation.action).
		HandleNumericTypesNullable.With(testActions.numericTypesNullable.action).
		HandleNumericTypesNullableArrayItems.With(testActions.numericTypesNullableArrayItems.action).
		Finalize()
}
