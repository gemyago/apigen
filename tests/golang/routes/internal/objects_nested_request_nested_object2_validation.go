package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestNestedObject2Validator() FieldValidator[*models.ObjectsNestedRequestNestedObject2] {
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault,
	)
	validateNestedObject21 := NewSimpleFieldValidator[ObjectsNestedRequestNestedObject1NestedObject11](
		EnsureNonDefault,
	)
	validateNestedObject22 := NewSimpleFieldValidator[ObjectsNestedRequestNestedObject1NestedObject11](
		EnsureNonDefault,
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectsNestedRequestNestedObject2) {
		validateSimpleRequiredField1(bindingCtx, "simpleRequiredField1", location, value.SimpleRequiredField1)
		validateNestedObject21(bindingCtx, "nestedObject21", location, value.NestedObject21)
		validateNestedObject22(bindingCtx, "nestedObject22", location, value.NestedObject22)
	}
}
