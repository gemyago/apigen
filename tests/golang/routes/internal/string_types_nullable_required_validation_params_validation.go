package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewStringTypesNullableRequiredValidationParamsValidator() FieldValidator[*StringTypesNullableRequiredValidationParams] {
	validateUnformattedStrInQuery := NewSimpleFieldValidator[*string](
		SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string, string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)
	validateOptionalUnformattedStrInQuery := NewSimpleFieldValidator[*string](
		SkipNullValidator(NewMinMaxLengthValidator[string, string](10, true)),
		SkipNullValidator(NewMinMaxLengthValidator[string, string](100, false)),
		SkipNullValidator(NewPatternValidator[string](".*")),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewStringTypesNullableRequiredValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *StringTypesNullableRequiredValidationParams) {
		validateUnformattedStrInQuery(bindingCtx.Fork("unformattedStrInQuery"), value.UnformattedStrInQuery)
		validateOptionalUnformattedStrInQuery(bindingCtx.Fork("optionalUnformattedStrInQuery"), value.OptionalUnformattedStrInQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
