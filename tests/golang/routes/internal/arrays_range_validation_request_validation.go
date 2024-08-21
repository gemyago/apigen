package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewArraysRangeValidationRequestValidator() FieldValidator[*models.ArraysRangeValidationRequest] {
	validateSimpleItems1 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			EnsureArrayFieldRequired,
			NewMinMaxLengthValidator[string, []string](5, true),
			NewMinMaxLengthValidator[string, []string](10, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems2 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			EnsureArrayFieldRequired,
			NewMinMaxLengthValidator[string, []string](10, true),
			NewMinMaxLengthValidator[string, []string](15, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems1 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](15, true),
			NewMinMaxLengthValidator[string, []string](20, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems2 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](20, true),
			NewMinMaxLengthValidator[string, []string](25, false),
		),
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
