package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesArraysParsingRequestValidator() FieldValidator[*models.StringTypesArraysParsingRequest] {
	validateUnformattedStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateCustomFormatStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[string](
			),
	)
	validateDateStr := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[[]time.Time](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[time.Time](
			),
	)
	validateDateTimeStr := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[[]time.Time](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[time.Time](
			),
	)
	validateByteStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[string](
			),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesArraysParsingRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateByteStr(bindingCtx.Fork("byteStr"), value.ByteStr)
	}
}
