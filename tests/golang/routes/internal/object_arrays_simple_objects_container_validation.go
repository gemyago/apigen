package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectArraysSimpleObjectsContainerValidator() FieldValidator[*ObjectArraysSimpleObjectsContainer] {
	validateSimpleObjects1 := NewArrayValidator[*ObjectArraysSimpleObject](
		NewSimpleFieldValidator[[]*ObjectArraysSimpleObject](
			EnsureArrayFieldRequired,
		),
		NewObjectArraysSimpleObjectValidator(),
	)
	validateSimpleObjects2 := NewArrayValidator[*ObjectArraysSimpleObject](
		NewSimpleFieldValidator[[]*ObjectArraysSimpleObject](
		),
		NewObjectArraysSimpleObjectValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectArraysSimpleObjectsContainer) {
		validateSimpleObjects1(bindingCtx.Fork("simpleObjects1"), value.SimpleObjects1)
		validateSimpleObjects2(bindingCtx.Fork("simpleObjects2"), value.SimpleObjects2)
	}
}
