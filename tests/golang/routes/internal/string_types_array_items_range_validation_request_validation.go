package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesArrayItemsRangeValidationRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesArrayItemsRangeValidationRequest] {
	validateUnformattedStr := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
				SimpleFieldValidatorParams{Field: "UnformattedStr", Location: params.Location},
				NewMinMaxLengthValidator[string](10, true),
				NewMinMaxLengthValidator[string](20, false),
			),
	)
	validateCustomFormatStr := NewArrayValidator[string](
		NewSimpleFieldValidator[string](
				SimpleFieldValidatorParams{Field: "CustomFormatStr", Location: params.Location},
				NewMinMaxLengthValidator[string](20, true),
				NewMinMaxLengthValidator[string](30, false),
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
				NewMinMaxLengthValidator[string](30, true),
				NewMinMaxLengthValidator[string](40, false),
			),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesArrayItemsRangeValidationRequest) {
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, value.CustomFormatStr)
		validateDateStr(bindingCtx, value.DateStr)
		validateDateTimeStr(bindingCtx, value.DateTimeStr)
		validateByteStr(bindingCtx, value.ByteStr)
	}
}
