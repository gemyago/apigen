package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableArrayItemsParamsValidator() FieldValidator[*StringTypesNullableArrayItemsParams] {
	validateUnformattedStr := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](
		),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](20, false)),
				SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
			),
	)
	validateCustomFormatStr := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](
		),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](20, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](30, false)),
				SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
			),
	)
	validateDateStr := NewArrayValidator[*time.Time](
		NewSimpleFieldValidator[[]*time.Time](
		),
		NewSimpleFieldValidator[*time.Time](
			),
	)
	validateDateTimeStr := NewArrayValidator[*time.Time](
		NewSimpleFieldValidator[[]*time.Time](
		),
		NewSimpleFieldValidator[*time.Time](
			),
	)
	validateByteStr := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](
		),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](30, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](40, false)),
			),
	)
	validateUnformattedStrInQuery := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](
		),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](20, false)),
				SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
			),
	)
	validateCustomFormatStrInQuery := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](
		),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](20, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](30, false)),
				SkipNullValidator(NewPatternValidator[string]("^[A-Za-z]*$")),
			),
	)
	validateDateStrInQuery := NewArrayValidator[*time.Time](
		NewSimpleFieldValidator[[]*time.Time](
		),
		NewSimpleFieldValidator[*time.Time](
			),
	)
	validateDateTimeStrInQuery := NewArrayValidator[*time.Time](
		NewSimpleFieldValidator[[]*time.Time](
		),
		NewSimpleFieldValidator[*time.Time](
			),
	)
	validateByteStrInQuery := NewArrayValidator[*string](
		NewSimpleFieldValidator[[]*string](
		),
		NewSimpleFieldValidator[*string](
				SkipNullValidator(NewMinMaxLengthValidator[string, string](30, true)),
				SkipNullValidator(NewMinMaxLengthValidator[string, string](40, false)),
			),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesNullableArrayItemsRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesNullableArrayItemsParams) {
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
