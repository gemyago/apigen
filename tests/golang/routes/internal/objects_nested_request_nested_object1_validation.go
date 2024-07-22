package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestNestedObject1Validator() FieldValidator[*models.ObjectsNestedRequestNestedObject1] {
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault,
	)
	validateNestedObject11 := NewSimpleFieldValidator[ObjectsNestedRequestNestedObject1NestedObject11](
		EnsureNonDefault,
	)
	validateNestedObject12 := NewSimpleFieldValidator[ObjectsNestedRequestNestedObject1NestedObject11](
		EnsureNonDefault,
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectsNestedRequestNestedObject1) {
		validateSimpleRequiredField1(bindingCtx, "simpleRequiredField1", location, value.SimpleRequiredField1)
		validateNestedObject11(bindingCtx, "nestedObject11", location, value.NestedObject11)
		validateNestedObject12(bindingCtx, "nestedObject12", location, value.NestedObject12)
	}
}