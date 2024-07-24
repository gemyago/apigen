package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesParsingRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "unformattedStr", Location: params.Location},
		EnsureNonDefault[string],
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "customFormatStr", Location: params.Location},
		EnsureNonDefault[string],
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
	
	return func(bindingCtx *BindingContext, value *models.StringTypesParsingRequest) {
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, value.CustomFormatStr)
		validateDateStr(bindingCtx, value.DateStr)
		validateDateTimeStr(bindingCtx, value.DateTimeStr)
		validateByteStr(bindingCtx, value.ByteStr)
	}
}
