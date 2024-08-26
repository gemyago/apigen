package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableRequiredValidationRequestValidator() FieldValidator[*StringTypesNullableRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[*string](
		SkipNullValidator(EnsureNonDefault[string]),
		SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string, string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)
	validateOptionalUnformattedStr := NewSimpleFieldValidator[*string](
		SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string, string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesNullableRequiredValidationRequest) {
		validateUnformattedStr(bindingCtx.Fork("unformattedStr"), value.UnformattedStr)
		validateOptionalUnformattedStr(bindingCtx.Fork("optionalUnformattedStr"), value.OptionalUnformattedStr)
	}
}
