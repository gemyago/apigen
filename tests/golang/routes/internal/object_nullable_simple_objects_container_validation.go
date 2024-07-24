package internal

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectNullableSimpleObjectsContainerValidator() FieldValidator[*models.ObjectNullableSimpleObjectsContainer] {
	validateSimpleObject1 := NewObjectNullableSimpleObjectValidator()
	validateOptionalSimpleObject2 := NewObjectNullableSimpleObjectValidator()

	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectNullableSimpleObjectsContainer) {
		validateSimpleObject1(bindingCtx, "simpleObject1", location, value.SimpleObject1)
		validateOptionalSimpleObject2(bindingCtx, "optionalSimpleObject2", location, value.OptionalSimpleObject2)
	}
}
