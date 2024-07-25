package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestNestedObject2Validator(params ModelValidatorParams) FieldValidator[*models.ObjectsNestedRequestNestedObject2] {
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "simpleRequiredField1", Location: params.Location},
		EnsureNonDefault[string],
	)
	validateNestedObject21 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "nestedObject21", Location: params.Location, Required: true, Nullable: false},
		NewObjectsNestedRequestNestedObject1NestedObject11Validator(ModelValidatorParams{Location: params.Location + ".nestedObject21"}),
	)
	validateNestedObject22 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "nestedObject22", Location: params.Location, Required: true, Nullable: false},
		NewObjectsNestedRequestNestedObject1NestedObject11Validator(ModelValidatorParams{Location: params.Location + ".nestedObject22"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectsNestedRequestNestedObject2) {
		validateSimpleRequiredField1(bindingCtx, value.SimpleRequiredField1)
		validateNestedObject21(bindingCtx, value.NestedObject21)
		validateNestedObject22(bindingCtx, value.NestedObject22)
	}
}
