package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type numericTypesControllerTestActions struct {
	numericTypesParsing                  mockAction[*handlers.NumericTypesNumericTypesParsingRequest]
	numericTypesRangeValidation          mockAction[*handlers.NumericTypesNumericTypesRangeValidationRequest]
	numericTypesRangeValidationExclusive mockAction[*handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest]
	numericTypesRequiredValidation       mockAction[*handlers.NumericTypesNumericTypesRequiredValidationRequest]
}

func newNumericTypesController(
	testActions *numericTypesControllerTestActions,
) *handlers.NumericTypesController {
	return handlers.BuildNumericTypesController().
		HandleNumericTypesParsing.With(testActions.numericTypesParsing.action).
		HandleNumericTypesRangeValidation.With(testActions.numericTypesRangeValidation.action).
		HandleNumericTypesRangeValidationExclusive.With(testActions.numericTypesRangeValidationExclusive.action).
		HandleNumericTypesRequiredValidation.With(testActions.numericTypesRequiredValidation.action).
		Finalize()
}
