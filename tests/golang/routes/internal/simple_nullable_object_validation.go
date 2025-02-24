package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleNullableObjectValidator() FieldValidator[*SimpleNullableObject] {
	validateSimpleField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, value *SimpleNullableObject) {
		validateSimpleField1(bindingCtx.Fork("simpleField1"), value.SimpleField1)
	}
}
