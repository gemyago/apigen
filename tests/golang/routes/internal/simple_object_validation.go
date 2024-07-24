package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleObjectValidator() FieldValidator[*models.SimpleObject] {
	validateSimpleField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.SimpleObject) {
		validateSimpleField1(bindingCtx, "simpleField1", location, value.SimpleField1)
	}
}
