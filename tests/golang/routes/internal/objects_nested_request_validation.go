package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsNestedRequestValidator() FieldValidator[*models.ObjectsNestedRequest] {
	
	validateSimpleRequiredField1 := NewSimpleFieldValidator[string](
	)
	
	validateNestedObject1 := NewObjectsNestedRequestNestedObject1Validator()
	
	validateNestedObject2 := NewObjectsNestedRequestNestedObject2Validator()
	
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectsNestedRequest) {
		validateSimpleRequiredField1(bindingCtx, "simpleRequiredField1", location, value.SimpleRequiredField1)
		validateNestedObject1(bindingCtx, "nestedObject1", location, &value.NestedObject1)
		validateNestedObject2(bindingCtx, "nestedObject2", location, &value.NestedObject2)
	}
}
