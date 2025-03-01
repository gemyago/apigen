package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorWithParamsAndResponseParamsValidator() FieldValidator[*BehaviorWithParamsAndResponseParams] {
	validateQueryParam1 := NewSimpleFieldValidator[string](
	)
	validateQueryParam2 := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](0, false, true),
		NewMinMaxValueValidator[int32](5000, false, false),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewBehaviorWithParamsAndResponseRequestBodyValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorWithParamsAndResponseParams) {
		validateQueryParam1(bindingCtx.Fork("queryParam1"), value.QueryParam1)
		validateQueryParam2(bindingCtx.Fork("queryParam2"), value.QueryParam2)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
