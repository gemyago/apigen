package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsDeeplyNestedRequestValidator(params ModelValidatorParams) FieldValidator[*models.ObjectsDeeplyNestedRequest] {
	validateContainer1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "container1", Location: params.Location, Required: true, Nullable: false},
		NewObjectsDeeplyNestedRequestContainer1Validator(ModelValidatorParams{Location: params.Location + ".container1"}),
	)
	validateContainer2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "container2", Location: params.Location, Required: true, Nullable: false},
		NewObjectsDeeplyNestedRequestContainer2Validator(ModelValidatorParams{Location: params.Location + ".container2"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectsDeeplyNestedRequest) {
		validateContainer1(bindingCtx, value.Container1)
		validateContainer2(bindingCtx, value.Container2)
	}
}
