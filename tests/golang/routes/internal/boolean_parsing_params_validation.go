package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanParsingParamsValidator() FieldValidator[*BooleanParsingParams] {
	validateBoolParam1 := NewSimpleFieldValidator[bool](
	)
	validateBoolParam2 := NewSimpleFieldValidator[bool](
	)
	validateBoolParam1InQuery := NewSimpleFieldValidator[bool](
	)
	validateBoolParam2InQuery := NewSimpleFieldValidator[bool](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewBooleanParsingRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *BooleanParsingParams) {
		validateBoolParam1(bindingCtx.Fork("boolParam1"), value.BoolParam1)
		validateBoolParam2(bindingCtx.Fork("boolParam2"), value.BoolParam2)
		validateBoolParam1InQuery(bindingCtx.Fork("boolParam1InQuery"), value.BoolParam1InQuery)
		validateBoolParam2InQuery(bindingCtx.Fork("boolParam2InQuery"), value.BoolParam2InQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
