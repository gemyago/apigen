package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorWithParamsAndResponseResponseBodyValidator() FieldValidator[*BehaviorWithParamsAndResponseResponseBody] {
	validateField1 := NewSimpleFieldValidator[string](
	)
	validateField2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewBehaviorWithParamsAndResponseResponseBodyValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorWithParamsAndResponseResponseBody) {
		validateField1(bindingCtx.Fork("field1"), value.Field1)
		validateField2(bindingCtx.Fork("field2"), value.Field2)
	}
}
