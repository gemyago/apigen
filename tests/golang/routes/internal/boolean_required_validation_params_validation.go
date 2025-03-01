package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanRequiredValidationParamsValidator() FieldValidator[*BooleanRequiredValidationParams] {
	validateBoolParam1InQuery := NewSimpleFieldValidator[bool](
	)
	validateBoolParam2InQuery := NewSimpleFieldValidator[bool](
	)
	validateOptionalBoolParam1InQuery := NewSimpleFieldValidator[bool](
	)
	validateOptionalBoolParam2InQuery := NewSimpleFieldValidator[bool](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewBooleanRequiredValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *BooleanRequiredValidationParams) {
		validateBoolParam1InQuery(bindingCtx.Fork("boolParam1InQuery"), value.BoolParam1InQuery)
		validateBoolParam2InQuery(bindingCtx.Fork("boolParam2InQuery"), value.BoolParam2InQuery)
		validateOptionalBoolParam1InQuery(bindingCtx.Fork("optionalBoolParam1InQuery"), value.OptionalBoolParam1InQuery)
		validateOptionalBoolParam2InQuery(bindingCtx.Fork("optionalBoolParam2InQuery"), value.OptionalBoolParam2InQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
