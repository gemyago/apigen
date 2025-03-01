package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRequiredValidationParamsValidator() FieldValidator[*StringTypesRequiredValidationParams] {
	validateUnformattedStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](10, true),
	)
	validateCustomFormatStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](20, true),
	)
	validateDateStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validateDateTimeStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validateByteStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](30, true),
	)
	validateOptionalUnformattedStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](10, true),
	)
	validateOptionalCustomFormatStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](20, true),
	)
	validateOptionalDateStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validateOptionalDateTimeStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validateOptionalByteStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](30, true),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesRequiredValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesRequiredValidationParams) {
		validateUnformattedStrInQuery(bindingCtx.Fork("unformattedStrInQuery"), value.UnformattedStrInQuery)
		validateCustomFormatStrInQuery(bindingCtx.Fork("customFormatStrInQuery"), value.CustomFormatStrInQuery)
		validateDateStrInQuery(bindingCtx.Fork("dateStrInQuery"), value.DateStrInQuery)
		validateDateTimeStrInQuery(bindingCtx.Fork("dateTimeStrInQuery"), value.DateTimeStrInQuery)
		validateByteStrInQuery(bindingCtx.Fork("byteStrInQuery"), value.ByteStrInQuery)
		validateOptionalUnformattedStrInQuery(bindingCtx.Fork("optionalUnformattedStrInQuery"), value.OptionalUnformattedStrInQuery)
		validateOptionalCustomFormatStrInQuery(bindingCtx.Fork("optionalCustomFormatStrInQuery"), value.OptionalCustomFormatStrInQuery)
		validateOptionalDateStrInQuery(bindingCtx.Fork("optionalDateStrInQuery"), value.OptionalDateStrInQuery)
		validateOptionalDateTimeStrInQuery(bindingCtx.Fork("optionalDateTimeStrInQuery"), value.OptionalDateTimeStrInQuery)
		validateOptionalByteStrInQuery(bindingCtx.Fork("optionalByteStrInQuery"), value.OptionalByteStrInQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
