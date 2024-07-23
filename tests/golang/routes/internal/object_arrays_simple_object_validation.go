package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectArraysSimpleObjectValidator() FieldValidator[*models.ObjectArraysSimpleObject] {
	
	validateSimpleField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault,
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectArraysSimpleObject) {
		validateSimpleField1(bindingCtx, "simpleField1", location, value.SimpleField1)
	}
}
