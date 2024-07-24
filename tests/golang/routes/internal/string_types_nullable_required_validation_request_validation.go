package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableRequiredValidationRequestValidator(params ModelValidatorParams) FieldValidator[*models.StringTypesNullableRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[*string](
		SimpleFieldValidatorParams{Field: "unformattedStr", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[string]),
		SkipNullValidator(NewMinMaxLengthValidator[string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)
	validateOptionalUnformattedStr := NewSimpleFieldValidator[*string](
		SimpleFieldValidatorParams{Field: "optionalUnformattedStr", Location: params.Location},
		SkipNullValidator(NewMinMaxLengthValidator[string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)
	
	return func(bindingCtx *BindingContext, value *models.StringTypesNullableRequiredValidationRequest) {
		validateUnformattedStr(bindingCtx, value.UnformattedStr)
		validateOptionalUnformattedStr(bindingCtx, value.OptionalUnformattedStr)
	}
}
