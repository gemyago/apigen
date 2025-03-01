package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsRequiredNestedObjectsParamsValidator() FieldValidator[*ObjectsRequiredNestedObjectsParams] {
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewSimpleObjectsContainerValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsRequiredNestedObjectsParams) {
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
