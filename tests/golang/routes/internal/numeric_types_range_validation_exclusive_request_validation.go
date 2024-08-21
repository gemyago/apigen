package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesRangeValidationExclusiveRequestValidator() FieldValidator[*models.NumericTypesRangeValidationExclusiveRequest] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		EnsureNonDefault[float32],
		NewMinMaxValueValidator[float32](100.01, true, true),
		NewMinMaxValueValidator[float32](200.02, true, false),
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		EnsureNonDefault[float32],
		NewMinMaxValueValidator[float32](200.02, true, true),
		NewMinMaxValueValidator[float32](300.03, true, false),
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		EnsureNonDefault[float64],
		NewMinMaxValueValidator[float64](300.03, true, true),
		NewMinMaxValueValidator[float64](400.04, true, false),
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		EnsureNonDefault[int32],
		NewMinMaxValueValidator[int32](400, true, true),
		NewMinMaxValueValidator[int32](500, true, false),
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		EnsureNonDefault[int32],
		NewMinMaxValueValidator[int32](500, true, true),
		NewMinMaxValueValidator[int32](600, true, false),
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		EnsureNonDefault[int64],
		NewMinMaxValueValidator[int64](600, true, true),
		NewMinMaxValueValidator[int64](700, true, false),
	)
	
	return func(bindingCtx *BindingContext, value *models.NumericTypesRangeValidationExclusiveRequest) {
		validateNumberAny(bindingCtx.Fork("numberAny"), value.NumberAny)
		validateNumberFloat(bindingCtx.Fork("numberFloat"), value.NumberFloat)
		validateNumberDouble(bindingCtx.Fork("numberDouble"), value.NumberDouble)
		validateNumberInt(bindingCtx.Fork("numberInt"), value.NumberInt)
		validateNumberInt32(bindingCtx.Fork("numberInt32"), value.NumberInt32)
		validateNumberInt64(bindingCtx.Fork("numberInt64"), value.NumberInt64)
	}
}
