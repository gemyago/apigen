package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanNullableParamsValidator() FieldValidator[*BooleanNullableParams] {
	validateBoolParam1 := NewSimpleFieldValidator[*bool](
	)
	validateBoolParam2 := NewSimpleFieldValidator[*bool](
	)
	validateBoolParam1InQuery := NewSimpleFieldValidator[*bool](
	)
	validateBoolParam2InQuery := NewSimpleFieldValidator[*bool](
	)
	validateOptionalBoolParam1InQuery := NewSimpleFieldValidator[*bool](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewBooleanNullableRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *BooleanNullableParams) {
		validateBoolParam1(bindingCtx.Fork("boolParam1"), value.BoolParam1)
		validateBoolParam2(bindingCtx.Fork("boolParam2"), value.BoolParam2)
		validateBoolParam1InQuery(bindingCtx.Fork("boolParam1InQuery"), value.BoolParam1InQuery)
		validateBoolParam2InQuery(bindingCtx.Fork("boolParam2InQuery"), value.BoolParam2InQuery)
		validateOptionalBoolParam1InQuery(bindingCtx.Fork("optionalBoolParam1InQuery"), value.OptionalBoolParam1InQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
