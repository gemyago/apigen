package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesPatternValidationRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesPatternValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "unformattedStr", Location: params.Location},
		EnsureNonDefault[string],
		NewPatternValidator[string]("^\\d{10}$"),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "customFormatStr", Location: params.Location},
		EnsureNonDefault[string],
		NewPatternValidator[string]("^\\d{20}$"),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorParams{Field: "dateStr", Location: params.Location},
		EnsureNonDefault[time.Time],
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		SimpleFieldValidatorParams{Field: "dateTimeStr", Location: params.Location},
		EnsureNonDefault[time.Time],
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesPatternValidationRequest) {
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, value.CustomFormatStr)
		validateDateStr(bindingCtx, value.DateStr)
		validateDateTimeStr(bindingCtx, value.DateTimeStr)
	}
}
