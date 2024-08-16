package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanRequiredValidationRequestValidator() FieldValidator[*models.BooleanRequiredValidationRequest] {
	validateBoolParam1 := NewSimpleFieldValidator[bool](
	)
	validateBoolParam2 := NewSimpleFieldValidator[bool](
	)
	validateOptionalBoolParam1 := NewSimpleFieldValidator[bool](
	)
	validateOptionalBoolParam2 := NewSimpleFieldValidator[bool](
	)
	
	return func(bindingCtx *BindingContext, value *models.BooleanRequiredValidationRequest) {
		validateBoolParam1(bindingCtx.Fork("boolParam1"), value.BoolParam1)
		validateBoolParam2(bindingCtx.Fork("boolParam2"), value.BoolParam2)
		validateOptionalBoolParam1(bindingCtx.Fork("optionalBoolParam1"), value.OptionalBoolParam1)
		validateOptionalBoolParam2(bindingCtx.Fork("optionalBoolParam2"), value.OptionalBoolParam2)
	}
}
