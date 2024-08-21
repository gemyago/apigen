package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleObjectsContainerValidator() FieldValidator[*models.SimpleObjectsContainer] {
	validateSimpleField1 := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateSimpleField2 := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateSimpleObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewSimpleObjectValidator(),
	)
	validateSimpleObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: false},
		NewSimpleObjectValidator(),
	)
	validateSimpleNullableObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: true},
		NewSimpleNullableObjectValidator(),
	)
	validateSimpleNullableObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: true, Nullable: true},
		NewSimpleNullableObjectValidator(),
	)
	validateOptionalSimpleObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewSimpleObjectValidator(),
	)
	validateOptionalSimpleObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewSimpleObjectValidator(),
	)
	validateOptionalNullableSimpleObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: true},
		NewSimpleNullableObjectValidator(),
	)
	validateOptionalNullableSimpleObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: true},
		NewSimpleNullableObjectValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *models.SimpleObjectsContainer) {
		validateSimpleField1(bindingCtx.Fork("simpleField1"), value.SimpleField1)
		validateSimpleField2(bindingCtx.Fork("simpleField2"), value.SimpleField2)
		validateSimpleObject1(bindingCtx.Fork("simpleObject1"), value.SimpleObject1)
		validateSimpleObject2(bindingCtx.Fork("simpleObject2"), value.SimpleObject2)
		validateSimpleNullableObject1(bindingCtx.Fork("simpleNullableObject1"), value.SimpleNullableObject1)
		validateSimpleNullableObject2(bindingCtx.Fork("simpleNullableObject2"), value.SimpleNullableObject2)
		validateOptionalSimpleObject1(bindingCtx.Fork("optionalSimpleObject1"), value.OptionalSimpleObject1)
		validateOptionalSimpleObject2(bindingCtx.Fork("optionalSimpleObject2"), value.OptionalSimpleObject2)
		validateOptionalNullableSimpleObject1(bindingCtx.Fork("optionalNullableSimpleObject1"), value.OptionalNullableSimpleObject1)
		validateOptionalNullableSimpleObject2(bindingCtx.Fork("optionalNullableSimpleObject2"), value.OptionalNullableSimpleObject2)
	}
}
