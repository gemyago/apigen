package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanParsingRequestValidator() FieldValidator[*models.BooleanParsingRequest] {
	validateBoolParam1 := NewSimpleFieldValidator[bool](
	)
	validateBoolParam2 := NewSimpleFieldValidator[bool](
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.BooleanParsingRequest) {
		validateBoolParam1(bindingCtx, "boolParam1", location, value.BoolParam1)
		validateBoolParam2(bindingCtx, "boolParam2", location, value.BoolParam2)
	}
}
