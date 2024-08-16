package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesNullableParsingRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[*string](
		SimpleFieldValidatorParams{Field: "unformattedStr", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[string]),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[*string](
		SimpleFieldValidatorParams{Field: "customFormatStr", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[string]),
	)
	validateDateStr := NewSimpleFieldValidator[*time.Time](
		SimpleFieldValidatorParams{Field: "dateStr", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[time.Time]),
	)
	validateDateTimeStr := NewSimpleFieldValidator[*time.Time](
		SimpleFieldValidatorParams{Field: "dateTimeStr", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[time.Time]),
	)
	validateByteStr := NewSimpleFieldValidator[*string](
		SimpleFieldValidatorParams{Field: "byteStr", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[string]),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesNullableParsingRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateByteStr(bindingCtx.Fork("byteStr"), value.ByteStr)
	}
}
