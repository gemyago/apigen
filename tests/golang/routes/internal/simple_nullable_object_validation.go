package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleNullableObjectValidator(params ModelValidatorParams) FieldValidator[*models.SimpleNullableObject] {
	validateSimpleField1 := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "simpleField1", Location: params.Location},
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, value *models.SimpleNullableObject) {
		validateSimpleField1(bindingCtx, value.SimpleField1)
	}
}
