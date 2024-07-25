package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsDeeplyNestedRequestContainer2Validator(params ModelValidatorParams) FieldValidator[*models.ObjectsDeeplyNestedRequestContainer2] {
	validateContainer21 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "container21", Location: params.Location, Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(ModelValidatorParams{Location: params.Location + ".container21"}),
	)
	validateContainer22 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "container22", Location: params.Location, Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(ModelValidatorParams{Location: params.Location + ".container22"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectsDeeplyNestedRequestContainer2) {
		validateContainer21(bindingCtx, value.Container21)
		validateContainer22(bindingCtx, value.Container22)
	}
}
