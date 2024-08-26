package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleObjectValidator() FieldValidator[*SimpleObject] {
	validateSimpleField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
		NewMinMaxLengthValidator[string, string](2, true),
	)
	
	return func(bindingCtx *BindingContext, value *SimpleObject) {
		validateSimpleField1(bindingCtx.Fork("simpleField1"), value.SimpleField1)
	}
}
