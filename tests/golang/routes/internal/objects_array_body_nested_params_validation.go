package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsArrayBodyNestedParamsValidator() FieldValidator[*ObjectsArrayBodyNestedParams] {
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewObjectsArrayBodyNestedRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsArrayBodyNestedParams) {
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
