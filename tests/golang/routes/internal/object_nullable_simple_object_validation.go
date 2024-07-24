package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectNullableSimpleObjectValidator() FieldValidator[*models.ObjectNullableSimpleObject] {
	validateSimpleField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectNullableSimpleObject) {
		validateSimpleField1(bindingCtx, "simpleField1", location, value.SimpleField1)
	}
}
