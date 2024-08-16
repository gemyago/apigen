package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.BooleanParsingRequest] {
	validateBoolParam1 := NewSimpleFieldValidator[bool](
	)
	validateBoolParam2 := NewSimpleFieldValidator[bool](
	)
	
	return func(bindingCtx *BindingContext, value *models.BooleanParsingRequest) {
		validateBoolParam1(bindingCtx.Fork("boolParam1"), value.BoolParam1)
		validateBoolParam2(bindingCtx.Fork("boolParam2"), value.BoolParam2)
	}
}
