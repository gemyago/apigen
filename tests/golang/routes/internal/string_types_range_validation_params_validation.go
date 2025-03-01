package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesRangeValidationParamsValidator() FieldValidator[*StringTypesRangeValidationParams] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](10, true),
		NewMinMaxLengthValidator[string, string](20, false),
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](20, true),
		NewMinMaxLengthValidator[string, string](30, false),
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
	)
	validateByteStr := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](30, true),
		NewMinMaxLengthValidator[string, string](40, false),
	)
	validateUnformattedStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](10, true),
		NewMinMaxLengthValidator[string, string](20, false),
	)
	validateCustomFormatStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](20, true),
		NewMinMaxLengthValidator[string, string](30, false),
	)
	validateDateStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validateDateTimeStrInQuery := NewSimpleFieldValidator[time.Time](
	)
	validateByteStrInQuery := NewSimpleFieldValidator[string](
		NewMinMaxLengthValidator[string, string](30, true),
		NewMinMaxLengthValidator[string, string](40, false),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesRangeValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesRangeValidationParams) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateByteStr(bindingCtx.Fork("byteStr"), value.ByteStr)
		validateUnformattedStrInQuery(bindingCtx.Fork("unformattedStrInQuery"), value.UnformattedStrInQuery)
		validateCustomFormatStrInQuery(bindingCtx.Fork("customFormatStrInQuery"), value.CustomFormatStrInQuery)
		validateDateStrInQuery(bindingCtx.Fork("dateStrInQuery"), value.DateStrInQuery)
		validateDateTimeStrInQuery(bindingCtx.Fork("dateTimeStrInQuery"), value.DateTimeStrInQuery)
		validateByteStrInQuery(bindingCtx.Fork("byteStrInQuery"), value.ByteStrInQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
