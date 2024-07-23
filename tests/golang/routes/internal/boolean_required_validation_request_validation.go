package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanRequiredValidationRequestValidator() FieldValidator[*models.BooleanRequiredValidationRequest] {
	validateBoolParam1 := NewSimpleFieldValidator[bool](
		EnsureNonDefault,
	)
	validateBoolParam2 := NewSimpleFieldValidator[bool](
		EnsureNonDefault,
	)
	validateOptionalBoolParam1 := NewSimpleFieldValidator[bool](
		EnsureNonDefault,
	)
	validateOptionalBoolParam2 := NewSimpleFieldValidator[bool](
		EnsureNonDefault,
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.BooleanRequiredValidationRequest) {
		validateBoolParam1(bindingCtx, "boolParam1", location, value.BoolParam1)
		validateBoolParam2(bindingCtx, "boolParam2", location, value.BoolParam2)
		validateOptionalBoolParam1(bindingCtx, "optionalBoolParam1", location, value.OptionalBoolParam1)
		validateOptionalBoolParam2(bindingCtx, "optionalBoolParam2", location, value.OptionalBoolParam2)
	}
}
