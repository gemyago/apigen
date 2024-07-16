package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type stringTypesControllerTestActions struct {
	stringTypesParsing            mockAction[*handlers.StringTypesStringTypesParsingRequest]
	StringTypesRangeValidation    mockAction[*handlers.StringTypesStringTypesRangeValidationRequest]
	StringTypesPatternValidation  mockAction[*handlers.StringTypesStringTypesPatternValidationRequest]
	StringTypesRequiredValidation mockAction[*handlers.StringTypesStringTypesRequiredValidationRequest]
}

func newStringTypesController(
	testActions *stringTypesControllerTestActions,
) *handlers.StringTypesController {
	return handlers.BuildStringTypesController().
		HandleStringTypesParsing.With(testActions.stringTypesParsing.action).
		HandleStringTypesRangeValidation.With(testActions.StringTypesRangeValidation.action).
		HandleStringTypesPatternValidation.With(testActions.StringTypesPatternValidation.action).
		HandleStringTypesRequiredValidation.With(testActions.StringTypesRequiredValidation.action).
		Finalize()
}
