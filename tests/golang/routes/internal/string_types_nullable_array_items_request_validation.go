package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableArrayItemsRequestValidator() FieldValidator[*models.StringTypesNullableArrayItemsRequest] {
	validateUnformattedStr := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string](10, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string](20, false)),
				SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
			),
	)
	validateCustomFormatStr := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string](20, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string](30, false)),
				SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
			),
	)
	validateDateStr := NewArrayValidator[*time.Time](
		NewSimpleFieldValidator[[]*time.Time](),
		NewSimpleFieldValidator[*time.Time](
			),
	)
	validateDateTimeStr := NewArrayValidator[*time.Time](
		NewSimpleFieldValidator[[]*time.Time](),
		NewSimpleFieldValidator[*time.Time](
			),
	)
	validateByteStr := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string](30, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string](40, false)),
			),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesNullableArrayItemsRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateCustomFormatStr(bindingCtx.Fork("customFormatStr"), value.CustomFormatStr)
		validateDateStr(bindingCtx.Fork("dateStr"), value.DateStr)
		validateDateTimeStr(bindingCtx.Fork("dateTimeStr"), value.DateTimeStr)
		validateByteStr(bindingCtx.Fork("byteStr"), value.ByteStr)
	}
}
