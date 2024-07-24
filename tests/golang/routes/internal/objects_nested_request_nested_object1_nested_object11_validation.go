package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestNestedObject1NestedObject11Validator(params ModelValidatorParams) FieldValidator[*models.ObjectsNestedRequestNestedObject1NestedObject11] {
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "simpleRequiredField1", Location: params.Location},
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectsNestedRequestNestedObject1NestedObject11) {
		validateSimpleRequiredField1(bindingCtx, value.SimpleRequiredField1)
	}
}
