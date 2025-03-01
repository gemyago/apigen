package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorNamesWithIDParamsValidator() FieldValidator[*BehaviorNamesWithIDParams] {
	validateID := NewSimpleFieldValidator[string](
	)
	validateEndsWithID := NewSimpleFieldValidator[string](
	)
	validateTheIDInTheMiddle := NewSimpleFieldValidator[string](
	)
	validateQueryEndsWithID := NewSimpleFieldValidator[string](
	)
	validateQueryTheIDInTheMiddle := NewSimpleFieldValidator[string](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewBehaviorNamesWithIDDataValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorNamesWithIDParams) {
		validateID(bindingCtx.Fork("id"), value.ID)
		validateEndsWithID(bindingCtx.Fork("endsWithId"), value.EndsWithID)
		validateTheIDInTheMiddle(bindingCtx.Fork("theIdInTheMiddle"), value.TheIDInTheMiddle)
		validateQueryEndsWithID(bindingCtx.Fork("queryEndsWithId"), value.QueryEndsWithID)
		validateQueryTheIDInTheMiddle(bindingCtx.Fork("queryTheIdInTheMiddle"), value.QueryTheIDInTheMiddle)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
