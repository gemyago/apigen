package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsDeeplyNestedRequestContainer2Validator() FieldValidator[*ObjectsDeeplyNestedRequestContainer2] {
	validateContainer21 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(),
	)
	validateContainer22 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsDeeplyNestedRequestContainer2) {
		validateContainer21(bindingCtx.Fork("container21"), value.Container21)
		validateContainer22(bindingCtx.Fork("container22"), value.Container22)
	}
}
