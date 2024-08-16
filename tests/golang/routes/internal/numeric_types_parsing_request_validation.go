package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesParsingRequestValidator(params ModelValidatorParams) FieldValidator[*models.NumericTypesParsingRequest] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		EnsureNonDefault[float32],
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		EnsureNonDefault[float32],
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		EnsureNonDefault[float64],
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		EnsureNonDefault[int32],
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		EnsureNonDefault[int32],
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		EnsureNonDefault[int64],
	)
	
	return func(bindingCtx *BindingContext, value *models.NumericTypesParsingRequest) {
		validateNumberAny(bindingCtx.Fork("numberAny"), value.NumberAny)
		validateNumberFloat(bindingCtx.Fork("numberFloat"), value.NumberFloat)
		validateNumberDouble(bindingCtx.Fork("numberDouble"), value.NumberDouble)
		validateNumberInt(bindingCtx.Fork("numberInt"), value.NumberInt)
		validateNumberInt32(bindingCtx.Fork("numberInt32"), value.NumberInt32)
		validateNumberInt64(bindingCtx.Fork("numberInt64"), value.NumberInt64)
	}
}
