package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesPatternValidationRequestValidator() FieldValidator[*models.StringTypesPatternValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		EnsureNonDefault,
		NewPatternValidator[string]("^\\d{10}$"),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		EnsureNonDefault,
		NewPatternValidator[string]("^\\d{20}$"),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault,
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault,
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesPatternValidationRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, "customFormatStr", location, value.CustomFormatStr)
		validateDateStr(bindingCtx, "dateStr", location, value.DateStr)
		validateDateTimeStr(bindingCtx, "dateTimeStr", location, value.DateTimeStr)
	}
}
