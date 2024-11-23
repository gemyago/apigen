package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsDeeplyNestedRequestContainer1Validator() FieldValidator[*ObjectsDeeplyNestedRequestContainer1] {
	validateContainer11 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(),
	)
	validateContainer12 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsDeeplyNestedRequestContainer1) {
		validateContainer11(bindingCtx.Fork("container11"), value.Container11)
		validateContainer12(bindingCtx.Fork("container12"), value.Container12)
	}
}
