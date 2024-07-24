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
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, value.CustomFormatStr)
		validateDateStr(bindingCtx, value.DateStr)
		validateDateTimeStr(bindingCtx, value.DateTimeStr)
		validateByteStr(bindingCtx, value.ByteStr)
		validateOptionalUnformattedStr(bindingCtx, value.OptionalUnformattedStr)
		validateOptionalCustomFormatStr(bindingCtx, value.OptionalCustomFormatStr)
		validateOptionalDateStr(bindingCtx, value.OptionalDateStr)
		validateOptionalDateTimeStr(bindingCtx, value.OptionalDateTimeStr)
		validateOptionalByteStr(bindingCtx, value.OptionalByteStr)
		validateOptionalValidatedUnformattedStr(bindingCtx, value.OptionalValidatedUnformattedStr)
		validateOptionalValidatedCustomFormatStr(bindingCtx, value.OptionalValidatedCustomFormatStr)
		validateOptionalValidatedByteStr(bindingCtx, value.OptionalValidatedByteStr)
	}
}
