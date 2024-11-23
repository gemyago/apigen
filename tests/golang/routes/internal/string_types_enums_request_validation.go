package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesEnumsRequestValidator() FieldValidator[*StringTypesEnumsRequest] {
	validateInlineEnumProp := NewSimpleFieldValidator[StringTypesEnumsRequestInlineEnumProp](
		EnsureNonDefault[StringTypesEnumsRequestInlineEnumProp],
	)
	validateOptionalInlineEnumProp := NewSimpleFieldValidator[StringTypesEnumsRequestOptionalInlineEnumProp](
	)
	validateNullableInlineEnumProp := NewSimpleFieldValidator[*StringTypesEnumsRequestNullableInlineEnumProp](
		SkipNullValidator(EnsureNonDefault[StringTypesEnumsRequestNullableInlineEnumProp]),
	)
	validateOptionalNullableInlineEnumProp := NewSimpleFieldValidator[*StringTypesEnumsRequestOptionalNullableInlineEnumProp](
	)
	validateRefEnumProp := NewSimpleFieldValidator[BasicStringEnum](
		EnsureNonDefault[BasicStringEnum],
	)
	validateOptionalRefEnumProp := NewSimpleFieldValidator[BasicStringEnum](
	)
	validateNullableRefEnumProp := NewSimpleFieldValidator[*NullableStringEnum](
		SkipNullValidator(EnsureNonDefault[NullableStringEnum]),
	)
	validateOptionalNullableRefEnumProp := NewSimpleFieldValidator[*NullableStringEnum](
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesEnumsRequest) {
		validateInlineEnumProp(bindingCtx.Fork("inlineEnumProp"), value.InlineEnumProp)
		validateOptionalInlineEnumProp(bindingCtx.Fork("optionalInlineEnumProp"), value.OptionalInlineEnumProp)
		validateNullableInlineEnumProp(bindingCtx.Fork("nullableInlineEnumProp"), value.NullableInlineEnumProp)
		validateOptionalNullableInlineEnumProp(bindingCtx.Fork("optionalNullableInlineEnumProp"), value.OptionalNullableInlineEnumProp)
		validateRefEnumProp(bindingCtx.Fork("refEnumProp"), value.RefEnumProp)
		validateOptionalRefEnumProp(bindingCtx.Fork("optionalRefEnumProp"), value.OptionalRefEnumProp)
		validateNullableRefEnumProp(bindingCtx.Fork("nullableRefEnumProp"), value.NullableRefEnumProp)
		validateOptionalNullableRefEnumProp(bindingCtx.Fork("optionalNullableRefEnumProp"), value.OptionalNullableRefEnumProp)
	}
}
