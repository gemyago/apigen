package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesArrayItemsRangeValidationParamsValidator() FieldValidator[*StringTypesArrayItemsRangeValidationParams] {
	validateUnformattedStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](10, true),
				NewMinMaxLengthValidator[string, string](20, false),
			),
	)
	validateCustomFormatStr := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](20, true),
				NewMinMaxLengthValidator[string, string](30, false),
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
				NewMinMaxLengthValidator[string, string](30, true),
				NewMinMaxLengthValidator[string, string](40, false),
			),
	)
	validateUnformattedStrInQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](10, true),
				NewMinMaxLengthValidator[string, string](20, false),
			),
	)
	validateCustomFormatStrInQuery := NewArrayValidator[string](
		NewSimpleFieldValidator[[]string](
		),
		NewSimpleFieldValidator[string](
				NewMinMaxLengthValidator[string, string](20, true),
				NewMinMaxLengthValidator[string, string](30, false),
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
				NewMinMaxLengthValidator[string, string](30, true),
				NewMinMaxLengthValidator[string, string](40, false),
			),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesArrayItemsRangeValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesArrayItemsRangeValidationParams) {
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
