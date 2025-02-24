package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesArrayItemsRangeValidationRequestValidator() FieldValidator[*StringTypesArrayItemsRangeValidationRequest] {
	validateUnformattedStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](10, true),
				NewMinMaxLengthValidator[string, string](20, false),
			),
	)
	validateCustomFormatStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](20, true),
				NewMinMaxLengthValidator[string, string](30, false),
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
				NewMinMaxLengthValidator[string, string](30, true),
				NewMinMaxLengthValidator[string, string](40, false),
			),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesArrayItemsRangeValidationRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateByteStr(bindingCtx.Fork("byteStr"), value.ByteStr)
	}
}
