package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesArraysParsingParamsValidator() FieldValidator[*StringTypesArraysParsingParams] {
	validateUnformattedStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateCustomFormatStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateDateStr := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[[]time.Time](
		),
		NewSimpleFieldValidator[time.Time](
			),
	)
	validateDateTimeStr := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[[]time.Time](
		),
		NewSimpleFieldValidator[time.Time](
			),
	)
	validateByteStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateUnformattedStrInQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateCustomFormatStrInQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateDateStrInQuery := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[[]time.Time](
		),
		NewSimpleFieldValidator[time.Time](
			),
	)
	validateDateTimeStrInQuery := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[[]time.Time](
		),
		NewSimpleFieldValidator[time.Time](
			),
	)
	validateByteStrInQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
			),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesArraysParsingRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesArraysParsingParams) {
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
