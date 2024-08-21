package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRequiredValidationRequestValidator() FieldValidator[*models.StringTypesRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
		NewMinMaxLengthValidator[string, string](10, true),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
		NewMinMaxLengthValidator[string, string](20, true),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault[time.Time],
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault[time.Time],
	)
	validateByteStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateOptionalUnformattedStr := NewSimpleFieldValidator[string](
	)
	validateOptionalCustomFormatStr := NewSimpleFieldValidator[string](
	)
	validateOptionalDateStr := NewSimpleFieldValidator[time.Time](
	)
	validateOptionalDateTimeStr := NewSimpleFieldValidator[time.Time](
	)
	validateOptionalByteStr := NewSimpleFieldValidator[string](
	)
	validateOptionalValidatedUnformattedStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](10, true),
	)
	validateOptionalValidatedCustomFormatStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](20, true),
	)
	validateOptionalValidatedByteStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](30, true),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesRequiredValidationRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateByteStr(bindingCtx.Fork("byteStr"), value.ByteStr)
		validateOptionalUnformattedStr(bindingCtx.Fork("optionalUnformattedStr"), value.OptionalUnformattedStr)
		validateOptionalCustomFormatStr(bindingCtx.Fork("optionalCustomFormatStr"), value.OptionalCustomFormatStr)
		validateOptionalDateStr(bindingCtx.Fork("optionalDateStr"), value.OptionalDateStr)
		validateOptionalDateTimeStr(bindingCtx.Fork("optionalDateTimeStr"), value.OptionalDateTimeStr)
		validateOptionalByteStr(bindingCtx.Fork("optionalByteStr"), value.OptionalByteStr)
		validateOptionalValidatedUnformattedStr(bindingCtx.Fork("optionalValidatedUnformattedStr"), value.OptionalValidatedUnformattedStr)
		validateOptionalValidatedCustomFormatStr(bindingCtx.Fork("optionalValidatedCustomFormatStr"), value.OptionalValidatedCustomFormatStr)
		validateOptionalValidatedByteStr(bindingCtx.Fork("optionalValidatedByteStr"), value.OptionalValidatedByteStr)
	}
}
