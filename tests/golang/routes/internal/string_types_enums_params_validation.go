package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesEnumsParamsValidator() FieldValidator[*StringTypesEnumsParams] {
	validateInlineEnumParam := NewSimpleFieldValidator[StringTypesEnumsParamsInlineEnumParam](
	)
	validateNullableInlineEnumParam := NewSimpleFieldValidator[*StringTypesEnumsParamsNullableInlineEnumParam](
	)
	validateRefEnumParam := NewSimpleFieldValidator[BasicStringEnum](
	)
	validateNullableRefEnumParam := NewSimpleFieldValidator[*NullableStringEnum](
	)
	validateInlineEnumParamInQuery := NewSimpleFieldValidator[StringTypesEnumsParamsInlineEnumParamInQuery](
	)
	validateOptionalInlineEnumParamInQuery := NewSimpleFieldValidator[StringTypesEnumsParamsOptionalInlineEnumParamInQuery](
	)
	validateNullableInlineEnumParamInQuery := NewSimpleFieldValidator[*StringTypesEnumsParamsNullableInlineEnumParamInQuery](
	)
	validateOptionalNullableInlineEnumParamInQuery := NewSimpleFieldValidator[*StringTypesEnumsParamsOptionalNullableInlineEnumParamInQuery](
	)
	validateRefEnumParamInQuery := NewSimpleFieldValidator[BasicStringEnum](
	)
	validateNullableRefEnumParamInQuery := NewSimpleFieldValidator[*NullableStringEnum](
	)
	validateOptionalRefEnumParamInQuery := NewSimpleFieldValidator[BasicStringEnum](
	)
	validateOptionalNullableRefEnumParamInQuery := NewSimpleFieldValidator[*NullableStringEnum](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesEnumsRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesEnumsParams) {
		validateInlineEnumParam(bindingCtx.Fork("inlineEnumParam"), value.InlineEnumParam)
		validateNullableInlineEnumParam(bindingCtx.Fork("nullableInlineEnumParam"), value.NullableInlineEnumParam)
		validateRefEnumParam(bindingCtx.Fork("refEnumParam"), value.RefEnumParam)
		validateNullableRefEnumParam(bindingCtx.Fork("nullableRefEnumParam"), value.NullableRefEnumParam)
		validateInlineEnumParamInQuery(bindingCtx.Fork("inlineEnumParamInQuery"), value.InlineEnumParamInQuery)
		validateOptionalInlineEnumParamInQuery(bindingCtx.Fork("optionalInlineEnumParamInQuery"), value.OptionalInlineEnumParamInQuery)
		validateNullableInlineEnumParamInQuery(bindingCtx.Fork("nullableInlineEnumParamInQuery"), value.NullableInlineEnumParamInQuery)
		validateOptionalNullableInlineEnumParamInQuery(bindingCtx.Fork("optionalNullableInlineEnumParamInQuery"), value.OptionalNullableInlineEnumParamInQuery)
		validateRefEnumParamInQuery(bindingCtx.Fork("refEnumParamInQuery"), value.RefEnumParamInQuery)
		validateNullableRefEnumParamInQuery(bindingCtx.Fork("nullableRefEnumParamInQuery"), value.NullableRefEnumParamInQuery)
		validateOptionalRefEnumParamInQuery(bindingCtx.Fork("optionalRefEnumParamInQuery"), value.OptionalRefEnumParamInQuery)
		validateOptionalNullableRefEnumParamInQuery(bindingCtx.Fork("optionalNullableRefEnumParamInQuery"), value.OptionalNullableRefEnumParamInQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
