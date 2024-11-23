package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsArrayBodyNestedRequestValidator() FieldValidator[*ObjectsArrayBodyNestedRequest] {
	validateNestedArray1 := NewArrayValidator[*ObjectArraysSimpleObject](
		NewSimpleFieldValidator[[]*ObjectArraysSimpleObject](
			EnsureArrayFieldRequired,
		),
		NewObjectArraysSimpleObjectValidator(),
	)
	validateNestedArray2 := NewArrayValidator[*ObjectArraysSimpleObject](
		NewSimpleFieldValidator[[]*ObjectArraysSimpleObject](
			EnsureArrayFieldRequired,
		),
		NewObjectArraysSimpleObjectValidator(),
	)
	validateNestedArrayContainer1 := NewArrayValidator[*ObjectArraysSimpleObjectsContainer](
		NewSimpleFieldValidator[[]*ObjectArraysSimpleObjectsContainer](
			EnsureArrayFieldRequired,
		),
		NewObjectArraysSimpleObjectsContainerValidator(),
	)
	validateNestedArrayContainer2 := NewArrayValidator[*ObjectArraysSimpleObjectsContainer](
		NewSimpleFieldValidator[[]*ObjectArraysSimpleObjectsContainer](
			EnsureArrayFieldRequired,
		),
		NewObjectArraysSimpleObjectsContainerValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsArrayBodyNestedRequest) {
		validateNestedArray1(bindingCtx.Fork("nestedArray1"), value.NestedArray1)
		validateNestedArray2(bindingCtx.Fork("nestedArray2"), value.NestedArray2)
		validateNestedArrayContainer1(bindingCtx.Fork("nestedArrayContainer1"), value.NestedArrayContainer1)
		validateNestedArrayContainer2(bindingCtx.Fork("nestedArrayContainer2"), value.NestedArrayContainer2)
	}
}
