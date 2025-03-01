package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsRequiredBodyParamsValidator() FieldValidator[*ObjectsRequiredBodyParams] {
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewSimpleObjectValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsRequiredBodyParams) {
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
