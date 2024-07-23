package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectArraysSimpleObjectsContainerValidator() FieldValidator[*models.ObjectArraysSimpleObjectsContainer] {
	validateSimpleObjects1 := NewSimpleFieldValidator[[]ObjectArraysSimpleObject](
		EnsureNonDefault,
	)
	validateSimpleObjects2 := NewSimpleFieldValidator[[]ObjectArraysSimpleObject](
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectArraysSimpleObjectsContainer) {
		validateSimpleObjects1(bindingCtx, "simpleObjects1", location, value.SimpleObjects1)
		validateSimpleObjects2(bindingCtx, "simpleObjects2", location, value.SimpleObjects2)
	}
}
