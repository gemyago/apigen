package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectArraysSimpleObjectsContainerValidator() FieldValidator[*models.ObjectArraysSimpleObjectsContainer] {
	validateSimpleObjects1 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewSimpleFieldValidator[[]*models.ObjectArraysSimpleObject](),
		NewObjectArraysSimpleObjectValidator(),
	)
	validateSimpleObjects2 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewSimpleFieldValidator[[]*models.ObjectArraysSimpleObject](),
		NewObjectArraysSimpleObjectValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectArraysSimpleObjectsContainer) {
		validateSimpleObjects1(bindingCtx.Fork("simpleObjects1"), value.SimpleObjects1)
		validateSimpleObjects2(bindingCtx.Fork("simpleObjects2"), value.SimpleObjects2)
	}
}
