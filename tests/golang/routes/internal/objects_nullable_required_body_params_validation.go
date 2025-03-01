package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNullableRequiredBodyParamsValidator() FieldValidator[*ObjectsNullableRequiredBodyParams] {
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: true},
		NewSimpleNullableObjectValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsNullableRequiredBodyParams) {
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
