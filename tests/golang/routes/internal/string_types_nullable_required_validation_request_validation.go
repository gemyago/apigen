package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableRequiredValidationRequestValidator() FieldValidator[*models.StringTypesNullableRequiredValidationRequest] {
	validateUnformattedStr := NewSimpleFieldValidator[*string](
		EnsureNonDefault,
		NewMinMaxLengthValidator[*string](10, true),
	)
	validateOptionalUnformattedStr := NewSimpleFieldValidator[*string](
		NewMinMaxLengthValidator[*string](10, true),
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesNullableRequiredValidationRequest) {
		validateUnformattedStr(bindingCtx, "unformattedStr", location, value.UnformattedStr)
		validateOptionalUnformattedStr(bindingCtx, "optionalUnformattedStr", location, value.OptionalUnformattedStr)
	}
}
