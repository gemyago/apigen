package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewBehaviorNamesWithId200ResponseValidator() FieldValidator[*BehaviorNamesWithId200Response] {
	validateId := NewSimpleFieldValidator[string](
	)
	validateEndsWithId := NewSimpleFieldValidator[string](
	)
	validateTheIdInTheMiddle := NewSimpleFieldValidator[string](
	)
	
	return func(bindingCtx *BindingContext, value *BehaviorNamesWithId200Response) {
		validateId(bindingCtx.Fork("id"), value.Id)
		validateEndsWithId(bindingCtx.Fork("endsWithId"), value.EndsWithId)
		validateTheIdInTheMiddle(bindingCtx.Fork("theIdInTheMiddle"), value.TheIdInTheMiddle)
	}
}
