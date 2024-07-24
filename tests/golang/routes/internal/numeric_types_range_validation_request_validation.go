package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesRangeValidationRequestValidator() FieldValidator[*models.NumericTypesRangeValidationRequest] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		EnsureNonDefault[float32],
		NewMinMaxValueValidator[float32](100.01, false, true),
		NewMinMaxValueValidator[float32](200.02, false, false),
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		EnsureNonDefault[float32],
		NewMinMaxValueValidator[float32](200.02, false, true),
		NewMinMaxValueValidator[float32](300.03, false, false),
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		EnsureNonDefault[float64],
		NewMinMaxValueValidator[float64](300.03, false, true),
		NewMinMaxValueValidator[float64](400.04, false, false),
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		EnsureNonDefault[int32],
		NewMinMaxValueValidator[int32](400, false, true),
		NewMinMaxValueValidator[int32](500, false, false),
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		EnsureNonDefault[int32],
		NewMinMaxValueValidator[int32](500, false, true),
		NewMinMaxValueValidator[int32](600, false, false),
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		EnsureNonDefault[int64],
		NewMinMaxValueValidator[int64](600, false, true),
		NewMinMaxValueValidator[int64](700, false, false),
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.NumericTypesRangeValidationRequest) {
		validateNumberAny(bindingCtx, "numberAny", location, value.NumberAny)
		validateNumberFloat(bindingCtx, "numberFloat", location, value.NumberFloat)
		validateNumberDouble(bindingCtx, "numberDouble", location, value.NumberDouble)
		validateNumberInt(bindingCtx, "numberInt", location, value.NumberInt)
		validateNumberInt32(bindingCtx, "numberInt32", location, value.NumberInt32)
		validateNumberInt64(bindingCtx, "numberInt64", location, value.NumberInt64)
	}
}
