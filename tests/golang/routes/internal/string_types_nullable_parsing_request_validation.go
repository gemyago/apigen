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
	
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesNullableParsingRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, "customFormatStr", location, value.CustomFormatStr)
		validateDateStr(bindingCtx, "dateStr", location, value.DateStr)
		validateDateTimeStr(bindingCtx, "dateTimeStr", location, value.DateTimeStr)
		validateByteStr(bindingCtx, "byteStr", location, value.ByteStr)
	}
}
