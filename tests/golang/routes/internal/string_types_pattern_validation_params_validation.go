package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesPatternValidationParamsValidator() FieldValidator[*StringTypesPatternValidationParams] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		NewPatternValidator[string]("^\\d{10}$"),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		NewPatternValidator[string]("^\\d{20}$"),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
	)
	validateUnformattedStrInQuery := NewSimpleFieldValidator[string](
		NewPatternValidator[string]("^\\d{10}$"),
	)
	validateCustomFormatStrInQuery := NewSimpleFieldValidator[string](
		NewPatternValidator[string]("^\\d{20}$"),
	)
	validateDateStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validateDateTimeStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesPatternValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesPatternValidationParams) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateUnformattedStrInQuery(bindingCtx.Fork("unformattedStrInQuery"), value.UnformattedStrInQuery)
		validateCustomFormatStrInQuery(bindingCtx.Fork("customFormatStrInQuery"), value.CustomFormatStrInQuery)
		validateDateStrInQuery(bindingCtx.Fork("dateStrInQuery"), value.DateStrInQuery)
		validateDateTimeStrInQuery(bindingCtx.Fork("dateTimeStrInQuery"), value.DateTimeStrInQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
