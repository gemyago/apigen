package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsArrayBodyDirectParamsValidator() FieldValidator[*ObjectsArrayBodyDirectParams] {
	validatePayload := NewArrayValidator[*ObjectArraysSimpleObject](
		NewSimpleFieldValidator[[]*ObjectArraysSimpleObject](
		),
		NewObjectArraysSimpleObjectValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ObjectsArrayBodyDirectParams) {
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
