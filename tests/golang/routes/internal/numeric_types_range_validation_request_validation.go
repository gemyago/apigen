package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesRangeValidationRequestValidator(params ModelValidatorParams) FieldValidator[*models.NumericTypesRangeValidationRequest] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		SimpleFieldValidatorParams{Field: "numberAny", Location: params.Location},
		EnsureNonDefault[float32],
		NewMinMaxValueValidator[float32](100.01, false, true),
		NewMinMaxValueValidator[float32](200.02, false, false),
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		SimpleFieldValidatorParams{Field: "numberFloat", Location: params.Location},
		EnsureNonDefault[float32],
		NewMinMaxValueValidator[float32](200.02, false, true),
		NewMinMaxValueValidator[float32](300.03, false, false),
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		SimpleFieldValidatorParams{Field: "numberDouble", Location: params.Location},
		EnsureNonDefault[float64],
		NewMinMaxValueValidator[float64](300.03, false, true),
		NewMinMaxValueValidator[float64](400.04, false, false),
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		SimpleFieldValidatorParams{Field: "numberInt", Location: params.Location},
		EnsureNonDefault[int32],
		NewMinMaxValueValidator[int32](400, false, true),
		NewMinMaxValueValidator[int32](500, false, false),
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		SimpleFieldValidatorParams{Field: "numberInt32", Location: params.Location},
		EnsureNonDefault[int32],
		NewMinMaxValueValidator[int32](500, false, true),
		NewMinMaxValueValidator[int32](600, false, false),
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		SimpleFieldValidatorParams{Field: "numberInt64", Location: params.Location},
		EnsureNonDefault[int64],
		NewMinMaxValueValidator[int64](600, false, true),
		NewMinMaxValueValidator[int64](700, false, false),
	)
	
	return func(bindingCtx *BindingContext, value *models.NumericTypesRangeValidationRequest) {
		validateNumberAny(bindingCtx.Fork("numberAny"), value.NumberAny)
		validateNumberFloat(bindingCtx.Fork("numberFloat"), value.NumberFloat)
		validateNumberDouble(bindingCtx.Fork("numberDouble"), value.NumberDouble)
		validateNumberInt(bindingCtx.Fork("numberInt"), value.NumberInt)
		validateNumberInt32(bindingCtx.Fork("numberInt32"), value.NumberInt32)
		validateNumberInt64(bindingCtx.Fork("numberInt64"), value.NumberInt64)
	}
}
