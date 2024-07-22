package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRequiredValidationRequestValidator() FieldValidator[*models.StringTypesRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string](10, true),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string](20, true),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
	)
	validateByteStr := NewSimpleFieldValidator[string](
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
		NewMinMaxLengthValidator[string](10, true),
	)
	validateOptionalValidatedCustomFormatStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string](20, true),
	)
	validateOptionalValidatedByteStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string](30, true),
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesRequiredValidationRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, "customFormatStr", location, value.CustomFormatStr)
		validateDateStr(bindingCtx, "dateStr", location, value.DateStr)
		validateDateTimeStr(bindingCtx, "dateTimeStr", location, value.DateTimeStr)
		validateByteStr(bindingCtx, "byteStr", location, value.ByteStr)
		validateOptionalUnformattedStr(bindingCtx, "optionalUnformattedStr", location, value.OptionalUnformattedStr)
		validateOptionalCustomFormatStr(bindingCtx, "optionalCustomFormatStr", location, value.OptionalCustomFormatStr)
		validateOptionalDateStr(bindingCtx, "optionalDateStr", location, value.OptionalDateStr)
		validateOptionalDateTimeStr(bindingCtx, "optionalDateTimeStr", location, value.OptionalDateTimeStr)
		validateOptionalByteStr(bindingCtx, "optionalByteStr", location, value.OptionalByteStr)
		validateOptionalValidatedUnformattedStr(bindingCtx, "optionalValidatedUnformattedStr", location, value.OptionalValidatedUnformattedStr)
		validateOptionalValidatedCustomFormatStr(bindingCtx, "optionalValidatedCustomFormatStr", location, value.OptionalValidatedCustomFormatStr)
		validateOptionalValidatedByteStr(bindingCtx, "optionalValidatedByteStr", location, value.OptionalValidatedByteStr)
	}
}
