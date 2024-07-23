package internal

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsArrayParsingBodyNestedRequestValidator() FieldValidator[*models.ObjectsArrayParsingBodyNestedRequest] {
	validateNestedArray1 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewObjectArraysSimpleObjectValidator(),
	)
	validateNestedArray2 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewObjectArraysSimpleObjectValidator(),
	)
	validateNestedArrayContainer1 := NewArrayValidator[*models.ObjectArraysSimpleObjectsContainer](
		NewObjectArraysSimpleObjectsContainerValidator(),
	)
	validateNestedArrayContainer2 := NewArrayValidator[*models.ObjectArraysSimpleObjectsContainer](
		NewObjectArraysSimpleObjectsContainerValidator(),
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectsArrayParsingBodyNestedRequest) {
		validateNestedArray1(bindingCtx, "nestedArray1", location, value.NestedArray1)
		validateNestedArray2(bindingCtx, "nestedArray2", location, value.NestedArray2)
		validateNestedArrayContainer1(bindingCtx, "nestedArrayContainer1", location, value.NestedArrayContainer1)
		validateNestedArrayContainer2(bindingCtx, "nestedArrayContainer2", location, value.NestedArrayContainer2)
	}
}
