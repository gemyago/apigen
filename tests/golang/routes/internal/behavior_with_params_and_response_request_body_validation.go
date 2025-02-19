package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorWithParamsAndResponseRequestBodyValidator() FieldValidator[*BehaviorWithParamsAndResponseRequestBody] {
	validateField1 := NewSimpleFieldValidator[string](
	)
	validateField2 := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](0, false, true),
		NewMinMaxValueValidator[int32](5000, false, false),
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorWithParamsAndResponseRequestBody) {
		validateField1(bindingCtx.Fork("field1"), value.Field1)
		validateField2(bindingCtx.Fork("field2"), value.Field2)
	}
}
