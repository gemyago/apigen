package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesParsingRequestValidator() FieldValidator[*models.StringTypesParsingRequest] {
	
	validateUnformattedStr := NewSimpleFieldValidator[string](
		EnsureNonDefault,
	)
	
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		EnsureNonDefault,
	)
	
	validateDateStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault,
	)
	
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault,
	)
	
	validateByteStr := NewSimpleFieldValidator[string](
		EnsureNonDefault,
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesParsingRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, "customFormatStr", location, value.CustomFormatStr)
		validateDateStr(bindingCtx, "dateStr", location, value.DateStr)
		validateDateTimeStr(bindingCtx, "dateTimeStr", location, value.DateTimeStr)
		validateByteStr(bindingCtx, "byteStr", location, value.ByteStr)
	}
}
