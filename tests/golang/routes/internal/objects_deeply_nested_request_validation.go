package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsDeeplyNestedRequestValidator() FieldValidator[*ObjectsDeeplyNestedRequest] {
	validateContainer1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewObjectsDeeplyNestedRequestContainer1Validator(),
	)
	validateContainer2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewObjectsDeeplyNestedRequestContainer2Validator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsDeeplyNestedRequest) {
		validateContainer1(bindingCtx.Fork("container1"), value.Container1)
		validateContainer2(bindingCtx.Fork("container2"), value.Container2)
	}
}
