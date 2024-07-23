package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorNoParamsWithResponse202ResponseValidator() FieldValidator[*models.BehaviorNoParamsWithResponse202Response] {
	validateField1 := NewSimpleFieldValidator[string](
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.BehaviorNoParamsWithResponse202Response) {
		validateField1(bindingCtx, "field1", location, value.Field1)
	}
}
