package internal

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRangeValidationRequestValidator() FieldValidator[*models.StringTypesRangeValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string](10, true),
		NewMinMaxLengthValidator[string](20, false),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string](20, true),
		NewMinMaxLengthValidator[string](30, false),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time]()
	validateDateTimeStr := NewSimpleFieldValidator[time.Time]()
	validateByteStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string](30, true),
		NewMinMaxLengthValidator[string](40, false),
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesRangeValidationRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, "customFormatStr", location, value.CustomFormatStr)
		validateDateStr(bindingCtx, "dateStr", location, value.DateStr)
		validateDateTimeStr(bindingCtx, "dateTimeStr", location, value.DateTimeStr)
		validateByteStr(bindingCtx, "byteStr", location, value.ByteStr)
	}
}
