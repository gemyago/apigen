package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectArraysSimpleObjectsContainerValidator(params ModelValidatorParams) FieldValidator[*models.ObjectArraysSimpleObjectsContainer] {
	validateSimpleObjects1 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewObjectArraysSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".SimpleObjects1"}),
	)
	validateSimpleObjects2 := NewArrayValidator[*models.ObjectArraysSimpleObject](
		NewObjectArraysSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".SimpleObjects2"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.ObjectArraysSimpleObjectsContainer) {
		validateSimpleObjects1(bindingCtx, value.SimpleObjects1)
		validateSimpleObjects2(bindingCtx, value.SimpleObjects2)
	}
}
