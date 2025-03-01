package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsDeeplyNestedParamsValidator() FieldValidator[*ObjectsDeeplyNestedParams] {
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewObjectsDeeplyNestedRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsDeeplyNestedParams) {
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
