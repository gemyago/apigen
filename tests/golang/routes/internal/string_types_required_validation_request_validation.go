package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRequiredValidationRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "unformattedStr", Location: params.Location},
		EnsureNonDefault[string],
		NewMinMaxLengthValidator[string](10, true),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "customFormatStr", Location: params.Location},
		EnsureNonDefault[string],
		NewMinMaxLengthValidator[string](20, true),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorParams{Field: "dateStr", Location: params.Location},
		EnsureNonDefault[time.Time],
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorParams{Field: "dateTimeStr", Location: params.Location},
		EnsureNonDefault[time.Time],
	)
	validateByteStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "byteStr", Location: params.Location},
		EnsureNonDefault[string],
	)
	validateOptionalUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "optionalUnformattedStr", Location: params.Location},
	)
	validateOptionalCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "optionalCustomFormatStr", Location: params.Location},
	)
	validateOptionalDateStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorParams{Field: "optionalDateStr", Location: params.Location},
	)
	validateOptionalDateTimeStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorParams{Field: "optionalDateTimeStr", Location: params.Location},
	)
	validateOptionalByteStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "optionalByteStr", Location: params.Location},
	)
	validateOptionalValidatedUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "optionalValidatedUnformattedStr", Location: params.Location},
		NewMinMaxLengthValidator[string](10, true),
	)
	validateOptionalValidatedCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "optionalValidatedCustomFormatStr", Location: params.Location},
		NewMinMaxLengthValidator[string](20, true),
	)
	validateOptionalValidatedByteStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "optionalValidatedByteStr", Location: params.Location},
		NewMinMaxLengthValidator[string](30, true),
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
