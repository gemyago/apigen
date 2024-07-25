package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.NumericTypesParsingRequest] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		SimpleFieldValidatorParams{Field: "numberAny", Location: params.Location},
		EnsureNonDefault[float32],
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		SimpleFieldValidatorParams{Field: "numberFloat", Location: params.Location},
		EnsureNonDefault[float32],
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		SimpleFieldValidatorParams{Field: "numberDouble", Location: params.Location},
		EnsureNonDefault[float64],
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		SimpleFieldValidatorParams{Field: "numberInt", Location: params.Location},
		EnsureNonDefault[int32],
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		SimpleFieldValidatorParams{Field: "numberInt32", Location: params.Location},
		EnsureNonDefault[int32],
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		SimpleFieldValidatorParams{Field: "numberInt64", Location: params.Location},
		EnsureNonDefault[int64],
	)
	
	return func(bindingCtx *BindingContext, value *models.NumericTypesParsingRequest) {
		validateNumberAny(bindingCtx, value.NumberAny)
		validateNumberFloat(bindingCtx, value.NumberFloat)
		validateNumberDouble(bindingCtx, value.NumberDouble)
		validateNumberInt(bindingCtx, value.NumberInt)
		validateNumberInt32(bindingCtx, value.NumberInt32)
		validateNumberInt64(bindingCtx, value.NumberInt64)
	}
}
