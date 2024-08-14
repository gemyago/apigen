package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesArraysParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesArraysParsingRequest] {
	validateUnformattedStr := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
				SimpleFieldValidatorParams{Field: "UnformattedStr", Location: params.Location},
			),
	)
	validateCustomFormatStr := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
				SimpleFieldValidatorParams{Field: "CustomFormatStr", Location: params.Location},
			),
	)
	validateDateStr := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[time.Time](
				SimpleFieldValidatorParams{Field: "DateStr", Location: params.Location},
			),
	)
	validateDateTimeStr := NewArrayValidator[time.Time](
		NewSimpleFieldValidator[time.Time](
				SimpleFieldValidatorParams{Field: "DateTimeStr", Location: params.Location},
			),
	)
	validateByteStr := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
				SimpleFieldValidatorParams{Field: "ByteStr", Location: params.Location},
			),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesArraysParsingRequest) {
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, value.CustomFormatStr)
		validateDateStr(bindingCtx, value.DateStr)
		validateDateTimeStr(bindingCtx, value.DateTimeStr)
		validateByteStr(bindingCtx, value.ByteStr)
	}
}