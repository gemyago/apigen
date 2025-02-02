package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewErrorValidator() FieldValidator[*Error] {
	validateCode := NewSimpleFieldValidator[*interface{}](
		SkipNullValidator(EnsureNonDefault[interface{}]),
	)
	validateMessage := NewSimpleFieldValidator[string](
	)
	
	return func(bindingCtx *BindingContext, value *Error) {
		validateCode(bindingCtx.Fork("code"), value.Code)
		validateMessage(bindingCtx.Fork("message"), value.Message)
	}
}
