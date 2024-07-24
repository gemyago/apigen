package internal

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableRequiredValidationRequestValidator() FieldValidator[*models.StringTypesNullableRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[*string](
		SkipNullValidator(EnsureNonDefault[string]),
		SkipNullValidator(NewMinMaxLengthValidator[string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)
	validateOptionalUnformattedStr := NewSimpleFieldValidator[*string](
		SkipNullValidator(NewMinMaxLengthValidator[string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)

	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesNullableRequiredValidationRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateOptionalUnformattedStr(bindingCtx, "optionalUnformattedStr", location, value.OptionalUnformattedStr)
	}
}
