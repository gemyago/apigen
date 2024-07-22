package internal

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesParsingRequestValidator() FieldValidator[*models.StringTypesParsingRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string]()
	validateCustomFormatStr := NewSimpleFieldValidator[string]()
	validateDateStr := NewSimpleFieldValidator[time.Time]()
	validateDateTimeStr := NewSimpleFieldValidator[time.Time]()
	validateByteStr := NewSimpleFieldValidator[string]()
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesParsingRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, "customFormatStr", location, value.CustomFormatStr)
		validateDateStr(bindingCtx, "dateStr", location, value.DateStr)
		validateDateTimeStr(bindingCtx, "dateTimeStr", location, value.DateTimeStr)
		validateByteStr(bindingCtx, "byteStr", location, value.ByteStr)
	}
}
