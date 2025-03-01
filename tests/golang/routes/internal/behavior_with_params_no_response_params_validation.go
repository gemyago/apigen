package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorWithParamsNoResponseParamsValidator() FieldValidator[*BehaviorWithParamsNoResponseParams] {
	validateQueryParam1 := NewSimpleFieldValidator[string](
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorWithParamsNoResponseParams) {
		validateQueryParam1(bindingCtx.Fork("queryParam1"), value.QueryParam1)
	}
}
