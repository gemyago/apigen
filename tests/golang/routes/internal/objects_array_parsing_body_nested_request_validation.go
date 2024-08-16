package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsArrayParsingBodyNestedRequestValidator(params ModelValidatorParams) FieldValidator[*models.ObjectsArrayParsingBodyNestedRequest] {
	validateNestedArray1 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewObjectArraysSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".NestedArray1"}),
	)
	validateNestedArray2 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewObjectArraysSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".NestedArray2"}),
	)
	validateNestedArrayContainer1 := NewArrayValidator[*models.ObjectArraysSimpleObjectsContainer](
		NewObjectArraysSimpleObjectsContainerValidator(ModelValidatorParams{Location: params.Location + ".NestedArrayContainer1"}),
	)
	validateNestedArrayContainer2 := NewArrayValidator[*models.ObjectArraysSimpleObjectsContainer](
		NewObjectArraysSimpleObjectsContainerValidator(ModelValidatorParams{Location: params.Location + ".NestedArrayContainer2"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectsArrayParsingBodyNestedRequest) {
		validateNestedArray1(bindingCtx.Fork("nestedArray1"), value.NestedArray1)
		validateNestedArray2(bindingCtx.Fork("nestedArray2"), value.NestedArray2)
		validateNestedArrayContainer1(bindingCtx.Fork("nestedArrayContainer1"), value.NestedArrayContainer1)
		validateNestedArrayContainer2(bindingCtx.Fork("nestedArrayContainer2"), value.NestedArrayContainer2)
	}
}
