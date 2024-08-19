package controllers

import (
	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type booleanControllerTestActions struct {
	booleanParsing            mockAction[*handlers.BooleanBooleanParsingRequest]
	booleanRequiredValidation mockAction[*handlers.BooleanBooleanRequiredValidationRequest]
	booleanNullable           mockAction[*handlers.BooleanBooleanNullableRequest]
	booleanArrayItems         mockAction[*handlers.BooleanBooleanArrayItemsRequest]
	nullableBooleanArrayItems mockAction[*handlers.BooleanBooleanNullableArrayItemsRequest]
}

func newBooleanController(
	testActions *booleanControllerTestActions,
) *handlers.BooleanController {
	return handlers.BuildBooleanController().
		HandleBooleanParsing.With(testActions.booleanParsing.action).
		HandleBooleanRequiredValidation.With(testActions.booleanRequiredValidation.action).
		HandleBooleanNullable.With(testActions.booleanNullable.action).
		HandleBooleanArrayItems.With(testActions.booleanArrayItems.action).
		HandleBooleanNullableArrayItems.With(testActions.nullableBooleanArrayItems.action).
		Finalize()
}
