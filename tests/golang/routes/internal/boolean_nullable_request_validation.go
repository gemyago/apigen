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
	
	return func(bindingCtx *BindingContext, field, location string, value *models.BooleanNullableRequest) {
		validateBoolParam1(bindingCtx, "boolParam1", location, value.BoolParam1)
		validateBoolParam2(bindingCtx, "boolParam2", location, value.BoolParam2)
		validateOptionalBoolParam1(bindingCtx, "optionalBoolParam1", location, value.OptionalBoolParam1)
	}
}
