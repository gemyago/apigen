package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleObjectsContainerValidator() FieldValidator[*models.SimpleObjectsContainer] {
	validateSimpleObject1 := NewSimpleObjectValidator()
	validateSimpleObject2 := NewSimpleObjectValidator()
	validateSimpleNullableObject1 := NewSimpleNullableObjectValidator()
	validateSimpleNullableObject2 := NewSimpleNullableObjectValidator()
	validateOptionalSimpleObject1 := NewSimpleObjectValidator()
	validateOptionalSimpleObject2 := NewSimpleObjectValidator()
	validateOptionalNullableSimpleObject1 := NewSimpleNullableObjectValidator()
	validateOptionalNullableSimpleObject2 := NewSimpleNullableObjectValidator()
	
	return func(bindingCtx *BindingContext, field, location string, value *models.SimpleObjectsContainer) {
		validateSimpleObject1(bindingCtx, "simpleObject1", location, &value.SimpleObject1)
		validateSimpleObject2(bindingCtx, "simpleObject2", location, &value.SimpleObject2)
		validateSimpleNullableObject1(bindingCtx, "simpleNullableObject1", location, value.SimpleNullableObject1)
		validateSimpleNullableObject2(bindingCtx, "simpleNullableObject2", location, value.SimpleNullableObject2)
		validateOptionalSimpleObject1(bindingCtx, "optionalSimpleObject1", location, &value.OptionalSimpleObject1)
		validateOptionalSimpleObject2(bindingCtx, "optionalSimpleObject2", location, &value.OptionalSimpleObject2)
		validateOptionalNullableSimpleObject1(bindingCtx, "optionalNullableSimpleObject1", location, value.OptionalNullableSimpleObject1)
		validateOptionalNullableSimpleObject2(bindingCtx, "optionalNullableSimpleObject2", location, value.OptionalNullableSimpleObject2)
	}
}
