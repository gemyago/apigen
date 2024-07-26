package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.BooleanParsingRequest] {
	validateBoolParam1 := NewSimpleFieldValidator[bool](
		SimpleFieldValidatorParams{Field: "boolParam1", Location: params.Location},
	)
	validateBoolParam2 := NewSimpleFieldValidator[bool](
		SimpleFieldValidatorParams{Field: "boolParam2", Location: params.Location},
	)
	
	return func(bindingCtx *BindingContext, value *models.BooleanParsingRequest) {
		validateBoolParam1(bindingCtx, value.BoolParam1)
		validateBoolParam2(bindingCtx, value.BoolParam2)
	}
}