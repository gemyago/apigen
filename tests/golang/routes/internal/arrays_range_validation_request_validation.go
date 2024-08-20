package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewArraysRangeValidationRequestValidator() FieldValidator[*models.ArraysRangeValidationRequest] {
	validateSimpleItems1 := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems2 := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems1 := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems2 := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
			),
	)
	
	return func(bindingCtx *BindingContext, value *models.ArraysRangeValidationRequest) {
		validateSimpleItems1(bindingCtx.Fork("simpleItems1"), value.SimpleItems1)
		validateSimpleItems2(bindingCtx.Fork("simpleItems2"), value.SimpleItems2)
		validateOptionalSimpleItems1(bindingCtx.Fork("optionalSimpleItems1"), value.OptionalSimpleItems1)
		validateOptionalSimpleItems2(bindingCtx.Fork("optionalSimpleItems2"), value.OptionalSimpleItems2)
	}
}
