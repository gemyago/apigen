package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNullableOptionalBodyParamsValidator() FieldValidator[*ObjectsNullableOptionalBodyParams] {
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: true},
		NewSimpleNullableObjectValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsNullableOptionalBodyParams) {
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
