package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableParsingParamsValidator() FieldValidator[*StringTypesNullableParsingParams] {
	validateUnformattedStr := NewSimpleFieldValidator[*string](
	)
	validateCustomFormatStr := NewSimpleFieldValidator[*string](
	)
	validateDateStr := NewSimpleFieldValidator[*time.Time](
	)
	validateDateTimeStr := NewSimpleFieldValidator[*time.Time](
	)
	validateByteStr := NewSimpleFieldValidator[*string](
	)
	validateUnformattedStrInQuery := NewSimpleFieldValidator[*string](
	)
	validateCustomFormatStrInQuery := NewSimpleFieldValidator[*string](
	)
	validateDateStrInQuery := NewSimpleFieldValidator[*time.Time](
	)
	validateDateTimeStrInQuery := NewSimpleFieldValidator[*time.Time](
	)
	validateByteStrInQuery := NewSimpleFieldValidator[*string](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesNullableParsingRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesNullableParsingParams) {
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
