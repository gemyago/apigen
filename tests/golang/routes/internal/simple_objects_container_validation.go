package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleObjectsContainerValidator(params ModelValidatorParams) FieldValidator[*models.SimpleObjectsContainer] {
	validateSimpleField1 := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "simpleField1", Location: params.Location},
		EnsureNonDefault[string],
	)
	validateSimpleField2 := NewSimpleFieldValidator[string](
		SimpleFieldValidatorParams{Field: "simpleField2", Location: params.Location},
		EnsureNonDefault[string],
	)
	validateSimpleObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "simpleObject1", Location: params.Location, Required: true, Nullable: false},
		NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleObject1"}),
	)
	validateSimpleObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "simpleObject2", Location: params.Location, Required: true, Nullable: false},
		NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleObject2"}),
	)
	validateSimpleNullableObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "simpleNullableObject1", Location: params.Location, Required: true, Nullable: true},
		NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleNullableObject1"}),
	)
	validateSimpleNullableObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "simpleNullableObject2", Location: params.Location, Required: true, Nullable: true},
		NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleNullableObject2"}),
	)
	validateOptionalSimpleObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "optionalSimpleObject1", Location: params.Location, Required: false, Nullable: false},
		NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalSimpleObject1"}),
	)
	validateOptionalSimpleObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "optionalSimpleObject2", Location: params.Location, Required: false, Nullable: false},
		NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalSimpleObject2"}),
	)
	validateOptionalNullableSimpleObject1 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "optionalNullableSimpleObject1", Location: params.Location, Required: false, Nullable: true},
		NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalNullableSimpleObject1"}),
	)
	validateOptionalNullableSimpleObject2 := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Field: "optionalNullableSimpleObject2", Location: params.Location, Required: false, Nullable: true},
		NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalNullableSimpleObject2"}),
	)
	
	return func(bindingCtx *BindingContext, value *models.SimpleObjectsContainer) {
		validateSimpleField1(bindingCtx, value.SimpleField1)
		validateSimpleField2(bindingCtx, value.SimpleField2)
		validateSimpleObject1(bindingCtx, value.SimpleObject1)
		validateSimpleObject2(bindingCtx, value.SimpleObject2)
		validateSimpleNullableObject1(bindingCtx, value.SimpleNullableObject1)
		validateSimpleNullableObject2(bindingCtx, value.SimpleNullableObject2)
		validateOptionalSimpleObject1(bindingCtx, value.OptionalSimpleObject1)
		validateOptionalSimpleObject2(bindingCtx, value.OptionalSimpleObject2)
		validateOptionalNullableSimpleObject1(bindingCtx, value.OptionalNullableSimpleObject1)
		validateOptionalNullableSimpleObject2(bindingCtx, value.OptionalNullableSimpleObject2)
	}
}
