package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesParsingRequestValidator() FieldValidator[*StringTypesParsingRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateCustomFormatStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateDateStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault[time.Time],
	)
	validateDateTimeStr := NewSimpleFieldValidator[time.Time](
		EnsureNonDefault[time.Time],
	)
	validateByteStr := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesParsingRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateByteStr(bindingCtx.Fork("byteStr"), value.ByteStr)
	}
}
