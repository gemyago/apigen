package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestValidator(params ModelValidatorParams) FieldValidator[*models.ObjectsNestedRequest] {
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "simpleRequiredField1", Location: params.Location},
	)
	validateNestedObject1 := NewObjectsNestedRequestNestedObject1Validator(ModelValidatorParams{Location: params.Location + ".nestedObject1"})
	validateNestedObject2 := NewObjectsNestedRequestNestedObject2Validator(ModelValidatorParams{Location: params.Location + ".nestedObject2"})
	
	return func(bindingCtx *BindingContext, value *models.ObjectsNestedRequest) {
		validateSimpleRequiredField1(bindingCtx, value.SimpleRequiredField1)
		validateNestedObject1(bindingCtx, &value.NestedObject1)
		validateNestedObject2(bindingCtx, &value.NestedObject2)
	}
}
