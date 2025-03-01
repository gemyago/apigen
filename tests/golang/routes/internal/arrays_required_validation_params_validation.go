package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewArraysRequiredValidationParamsValidator() FieldValidator[*ArraysRequiredValidationParams] {
	validateSimpleItems1 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems2 := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems1InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateSimpleItems2InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems1InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateOptionalSimpleItems2InQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewArraysRequiredValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *ArraysRequiredValidationParams) {
		validateSimpleItems1(bindingCtx.Fork("simpleItems1"), value.SimpleItems1)
		validateSimpleItems2(bindingCtx.Fork("simpleItems2"), value.SimpleItems2)
		validateSimpleItems1InQuery(bindingCtx.Fork("simpleItems1InQuery"), value.SimpleItems1InQuery)
		validateSimpleItems2InQuery(bindingCtx.Fork("simpleItems2InQuery"), value.SimpleItems2InQuery)
		validateOptionalSimpleItems1InQuery(bindingCtx.Fork("optionalSimpleItems1InQuery"), value.OptionalSimpleItems1InQuery)
		validateOptionalSimpleItems2InQuery(bindingCtx.Fork("optionalSimpleItems2InQuery"), value.OptionalSimpleItems2InQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
