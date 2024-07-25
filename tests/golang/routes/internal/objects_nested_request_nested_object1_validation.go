package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestNestedObject1Validator(params ModelValidatorParams) FieldValidator[*models.ObjectsNestedRequestNestedObject1] {
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "simpleRequiredField1", Location: params.Location},
		EnsureNonDefault[string],
	)
	validateNestedObject11 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "nestedObject11", Location: params.Location, Required: true, Nullable: false},
		NewObjectsNestedRequestNestedObject1NestedObject11Validator(ModelValidatorParams{Location: params.Location + ".nestedObject11"}),
	)
	validateNestedObject12 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "nestedObject12", Location: params.Location, Required: true, Nullable: false},
		NewObjectsNestedRequestNestedObject1NestedObject11Validator(ModelValidatorParams{Location: params.Location + ".nestedObject12"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectsNestedRequestNestedObject1) {
		validateSimpleRequiredField1(bindingCtx, value.SimpleRequiredField1)
		validateNestedObject11(bindingCtx, value.NestedObject11)
		validateNestedObject12(bindingCtx, value.NestedObject12)
	}
}
