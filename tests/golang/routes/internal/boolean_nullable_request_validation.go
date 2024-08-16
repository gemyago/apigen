package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanNullableRequestValidator() FieldValidator[*models.BooleanNullableRequest] {
	validateBoolParam1 := NewSimpleFieldValidator[*bool](
	)
	validateBoolParam2 := NewSimpleFieldValidator[*bool](
	)
	validateOptionalBoolParam1 := NewSimpleFieldValidator[*bool](
	)
	
	return func(bindingCtx *BindingContext, value *models.BooleanNullableRequest) {
		validateBoolParam1(bindingCtx.Fork("boolParam1"), value.BoolParam1)
		validateBoolParam2(bindingCtx.Fork("boolParam2"), value.BoolParam2)
		validateOptionalBoolParam1(bindingCtx.Fork("optionalBoolParam1"), value.OptionalBoolParam1)
	}
}
