package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesPatternValidationRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesPatternValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
		NewPatternValidator[string]("^\\d{10}$"),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
		NewPatternValidator[string]("^\\d{20}$"),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault[time.Time],
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault[time.Time],
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesPatternValidationRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
	}
}
