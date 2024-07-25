package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRangeValidationRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesRangeValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "unformattedStr", Location: params.Location},
		EnsureNonDefault[string],
		NewMinMaxLengthValidator[string](10, true),
		NewMinMaxLengthValidator[string](20, false),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "customFormatStr", Location: params.Location},
		EnsureNonDefault[string],
		NewMinMaxLengthValidator[string](20, true),
		NewMinMaxLengthValidator[string](30, false),
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
		NewMinMaxLengthValidator[string](30, true),
		NewMinMaxLengthValidator[string](40, false),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesRangeValidationRequest) {
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, value.CustomFormatStr)
		validateDateStr(bindingCtx, value.DateStr)
		validateDateTimeStr(bindingCtx, value.DateTimeStr)
		validateByteStr(bindingCtx, value.ByteStr)
	}
}
