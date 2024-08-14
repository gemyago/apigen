package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type stringTypesControllerTestActions struct {
	stringTypesParsing                    mockAction[*handlers.StringTypesStringTypesParsingRequest]
	stringTypesArraysParsing              mockAction[*handlers.StringTypesStringTypesArraysParsingRequest]
	stringTypesNullableParsing            mockAction[*handlers.StringTypesStringTypesNullableParsingRequest]
	stringTypesRangeValidation            mockAction[*handlers.StringTypesStringTypesRangeValidationRequest]
	stringTypesPatternValidation          mockAction[*handlers.StringTypesStringTypesPatternValidationRequest]
	stringTypesRequiredValidation         mockAction[*handlers.StringTypesStringTypesRequiredValidationRequest]
	stringTypesNullableRequiredValidation mockAction[*handlers.StringTypesStringTypesNullableRequiredValidationRequest]
}

func newStringTypesController(
	testActions *stringTypesControllerTestActions,
) *handlers.StringTypesController {
	return handlers.BuildStringTypesController().
		HandleStringTypesParsing.With(testActions.stringTypesParsing.action).
		HandleStringTypesArraysParsing.With(testActions.stringTypesArraysParsing.action).
		HandleStringTypesNullableParsing.With(testActions.stringTypesNullableParsing.action).
		HandleStringTypesRangeValidation.With(testActions.stringTypesRangeValidation.action).
		HandleStringTypesPatternValidation.With(testActions.stringTypesPatternValidation.action).
		HandleStringTypesRequiredValidation.With(testActions.stringTypesRequiredValidation.action).
		HandleStringTypesNullableRequiredValidation.With(testActions.stringTypesNullableRequiredValidation.action).
		Finalize()
}
