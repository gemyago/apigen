package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorNamesWithIDDataValidator() FieldValidator[*BehaviorNamesWithIDData] {
	validateID := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateEndsWithID := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateTheIDInTheMiddle := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorNamesWithIDData) {
		validateID(bindingCtx.Fork("id"), value.ID)
		validateEndsWithID(bindingCtx.Fork("endsWithId"), value.EndsWithID)
		validateTheIDInTheMiddle(bindingCtx.Fork("theIdInTheMiddle"), value.TheIDInTheMiddle)
	}
}
