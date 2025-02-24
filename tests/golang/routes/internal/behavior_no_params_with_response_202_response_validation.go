package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorNoParamsWithResponse202ResponseValidator() FieldValidator[*BehaviorNoParamsWithResponse202Response] {
	validateField1 := NewSimpleFieldValidator[string](
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorNoParamsWithResponse202Response) {
		validateField1(bindingCtx.Fork("field1"), value.Field1)
	}
}
