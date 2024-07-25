package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsDeeplyNestedRequestContainer1Validator(params ModelValidatorParams) FieldValidator[*models.ObjectsDeeplyNestedRequestContainer1] {
	validateContainer11 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "container11", Location: params.Location, Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(ModelValidatorParams{Location: params.Location + ".container11"}),
	)
	validateContainer12 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "container12", Location: params.Location, Required: true, Nullable: false},
		NewSimpleObjectsContainerValidator(ModelValidatorParams{Location: params.Location + ".container12"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectsDeeplyNestedRequestContainer1) {
		validateContainer11(bindingCtx, value.Container11)
		validateContainer12(bindingCtx, value.Container12)
	}
}
