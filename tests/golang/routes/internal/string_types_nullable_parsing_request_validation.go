package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableParsingRequestValidator() FieldValidator[*models.StringTypesNullableParsingRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[*string](
		SkipNullValidator(EnsureNonDefault[string]),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[*string](
		SkipNullValidator(EnsureNonDefault[string]),
	)
	validateDateStr := NewSimpleFieldValidator[*time.Time](
		SkipNullValidator(EnsureNonDefault[time.Time]),
	)
	validateDateTimeStr := NewSimpleFieldValidator[*time.Time](
		SkipNullValidator(EnsureNonDefault[time.Time]),
	)
	validateByteStr := NewSimpleFieldValidator[*string](
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
