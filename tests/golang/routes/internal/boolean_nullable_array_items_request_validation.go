package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBooleanNullableArrayItemsRequestValidator() FieldValidator[*models.BooleanNullableArrayItemsRequest] {
	validateBoolParam1 := NewArrayValidator[*bool](
		NewSimpleFieldValidator[[]*bool](),
		NewSimpleFieldValidator[*bool](
			),
	)
	validateBoolParam2 := NewArrayValidator[*bool](
		NewSimpleFieldValidator[[]*bool](),
		NewSimpleFieldValidator[*bool](
			),
	)
	
	return func(bindingCtx *BindingContext, value *models.BooleanNullableArrayItemsRequest) {
		validateBoolParam1(bindingCtx.Fork("boolParam1"), value.BoolParam1)
		validateBoolParam2(bindingCtx.Fork("boolParam2"), value.BoolParam2)
	}
}
