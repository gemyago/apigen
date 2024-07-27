package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesArraysParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesArraysParsingRequest] {
	validateUnformattedStr := NewArrayValidator[string](
		NewstringValidator(ModelValidatorParams{Location: params.Location + ".UnformattedStr"}),
	)
	validateCustomFormatStr := NewArrayValidator[string](
		NewstringValidator(ModelValidatorParams{Location: params.Location + ".CustomFormatStr"}),
	)
	validateDateStr := NewArrayValidator[time.Time](
		Newtime.TimeValidator(ModelValidatorParams{Location: params.Location + ".DateStr"}),
	)
	validateDateTimeStr := NewArrayValidator[time.Time](
		Newtime.TimeValidator(ModelValidatorParams{Location: params.Location + ".DateTimeStr"}),
	)
	validateByteStr := NewArrayValidator[string](
		NewstringValidator(ModelValidatorParams{Location: params.Location + ".ByteStr"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesArraysParsingRequest) {
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateCustomFormatStr(bindingCtx, value.CustomFormatStr)
		validateDateStr(bindingCtx, value.DateStr)
		validateDateTimeStr(bindingCtx, value.DateTimeStr)
		validateByteStr(bindingCtx, value.ByteStr)
	}
}
