package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewArraysRangeValidationParamsValidator() FieldValidator[*ArraysRangeValidationParams] {
	validateSimpleItems1 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](5, true),
			NewMinMaxLengthValidator[string, []string](10, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems2 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](10, true),
			NewMinMaxLengthValidator[string, []string](15, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems1InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](5, true),
			NewMinMaxLengthValidator[string, []string](10, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems2InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](10, true),
			NewMinMaxLengthValidator[string, []string](15, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems1InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](15, true),
			NewMinMaxLengthValidator[string, []string](20, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems2InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			NewMinMaxLengthValidator[string, []string](20, true),
			NewMinMaxLengthValidator[string, []string](25, false),
		),
		NewSimpleFieldValidator[string](
			),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewArraysRangeValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ArraysRangeValidationParams) {
		validateSimpleItems1(bindingCtx.Fork("simpleItems1"), value.SimpleItems1)
		validateSimpleItems2(bindingCtx.Fork("simpleItems2"), value.SimpleItems2)
		validateSimpleItems1InQuery(bindingCtx.Fork("simpleItems1InQuery"), value.SimpleItems1InQuery)
		validateSimpleItems2InQuery(bindingCtx.Fork("simpleItems2InQuery"), value.SimpleItems2InQuery)
		validateOptionalSimpleItems1InQuery(bindingCtx.Fork("optionalSimpleItems1InQuery"), value.OptionalSimpleItems1InQuery)
		validateOptionalSimpleItems2InQuery(bindingCtx.Fork("optionalSimpleItems2InQuery"), value.OptionalSimpleItems2InQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
