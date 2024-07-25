package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesNullableRequestValidator(params ModelValidatorParams) FieldValidator[*models.NumericTypesNullableRequest] {
	validateNumberAny := NewSimpleFieldValidator[*float32](
		SimpleFieldValidatorParams{Field: "numberAny", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[float32]),
		SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
	)
	validateOptionalNumberAny := NewSimpleFieldValidator[*float32](
		SimpleFieldValidatorParams{Field: "optionalNumberAny", Location: params.Location},
		SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
	)
	validateNumberFloat := NewSimpleFieldValidator[*float32](
		SimpleFieldValidatorParams{Field: "numberFloat", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[float32]),
		SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
	)
	validateNumberDouble := NewSimpleFieldValidator[*float64](
		SimpleFieldValidatorParams{Field: "numberDouble", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[float64]),
		SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
	)
	validateNumberInt := NewSimpleFieldValidator[*int32](
		SimpleFieldValidatorParams{Field: "numberInt", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[int32]),
		SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
	)
	validateNumberInt32 := NewSimpleFieldValidator[*int32](
		SimpleFieldValidatorParams{Field: "numberInt32", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[int32]),
		SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
	)
	validateNumberInt64 := NewSimpleFieldValidator[*int64](
		SimpleFieldValidatorParams{Field: "numberInt64", Location: params.Location},
		SkipNullValidator(EnsureNonDefault[int64]),
		SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
	)
	
	return func(bindingCtx *BindingContext, value *models.NumericTypesNullableRequest) {
		validateNumberAny(bindingCtx, value.NumberAny)
		validateOptionalNumberAny(bindingCtx, value.OptionalNumberAny)
		validateNumberFloat(bindingCtx, value.NumberFloat)
		validateNumberDouble(bindingCtx, value.NumberDouble)
		validateNumberInt(bindingCtx, value.NumberInt)
		validateNumberInt32(bindingCtx, value.NumberInt32)
		validateNumberInt64(bindingCtx, value.NumberInt64)
	}
}
