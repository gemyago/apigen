package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRequiredValidationRequestValidator() FieldValidator[*models.StringTypesRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorOpts{},
		NewMinMaxLengthValidator[string](10, true),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorOpts{},
		NewMinMaxLengthValidator[string](20, true),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorOpts{},
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorOpts{},
	)
	validateByteStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorOpts{},
		NewMinMaxLengthValidator[string](30, true),
	)
	validateOptionalUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorOpts{},
		NewMinMaxLengthValidator[string](10, true),
	)
	validateOptionalCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorOpts{},
		NewMinMaxLengthValidator[string](20, true),
	)
	validateOptionalDateStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorOpts{},
	)
	validateOptionalDateTimeStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorOpts{},
	)
	validateOptionalByteStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorOpts{},
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
	}
}
