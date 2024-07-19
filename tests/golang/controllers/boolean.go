package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type booleanControllerTestActions struct {
	booleanParsing            mockAction[*handlers.BooleanBooleanParsingRequest]
	booleanRequiredValidation mockAction[*handlers.BooleanBooleanRequiredValidationRequest]
}

func newBooleanController(
	testActions *booleanControllerTestActions,
) *handlers.BooleanController {
	return handlers.BuildBooleanController().
		HandleBooleanParsing.With(testActions.booleanParsing.action).
		HandleBooleanRequiredValidation.With(testActions.booleanRequiredValidation.action).
		Finalize()
}
