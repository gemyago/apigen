package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorNamesWithIdDataValidator() FieldValidator[*BehaviorNamesWithIdData] {
	validateId := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateEndsWithId := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateTheIdInTheMiddle := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorNamesWithIdData) {
		validateId(bindingCtx.Fork("id"), value.Id)
		validateEndsWithId(bindingCtx.Fork("endsWithId"), value.EndsWithId)
		validateTheIdInTheMiddle(bindingCtx.Fork("theIdInTheMiddle"), value.TheIdInTheMiddle)
	}
}
