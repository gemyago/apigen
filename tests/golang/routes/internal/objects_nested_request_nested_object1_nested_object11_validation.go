package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestNestedObject1NestedObject11Validator() FieldValidator[*models.ObjectsNestedRequestNestedObject1NestedObject11] {
	
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault,
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectsNestedRequestNestedObject1NestedObject11) {
		validateSimpleRequiredField1(bindingCtx, "simpleRequiredField1", location, value.SimpleRequiredField1)
	}
}
