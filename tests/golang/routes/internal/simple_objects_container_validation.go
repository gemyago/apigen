package internal

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewSimpleObjectsContainerValidator(params ModelValidatorParams) FieldValidator[*models.SimpleObjectsContainer] {
	validateSimpleObject1 := RequiredModelFieldValidator(
		SimpleFieldValidatorParams{Field: "simpleObject1", Location: params.Location},
		NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleObject1"}),
	)
	validateSimpleObject2 := RequiredModelFieldValidator(
		SimpleFieldValidatorParams{Field: "simpleObject2", Location: params.Location},
		NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleObject2"}),
	)
	validateSimpleNullableObject1 := SkipNullFieldValidator(NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleNullableObject1"}))
	validateSimpleNullableObject2 := SkipNullFieldValidator(NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".simpleNullableObject2"}))
	validateOptionalSimpleObject1 := SkipNullFieldValidator(NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalSimpleObject1"}))
	validateOptionalSimpleObject2 := SkipNullFieldValidator(NewSimpleObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalSimpleObject2"}))
	validateOptionalNullableSimpleObject1 := SkipNullFieldValidator(NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalNullableSimpleObject1"}))
	validateOptionalNullableSimpleObject2 := SkipNullFieldValidator(NewSimpleNullableObjectValidator(ModelValidatorParams{Location: params.Location + ".optionalNullableSimpleObject2"}))

	return func(bindingCtx *BindingContext, value *models.SimpleObjectsContainer) {
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
