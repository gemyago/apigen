package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesRangeValidationExclusiveRequestValidator() FieldValidator[*models.NumericTypesRangeValidationExclusiveRequest] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		EnsureNonDefault,
		NewMinMaxValueValidator[float32](100.01, true, true),
		NewMinMaxValueValidator[float32](200.02, true, false),
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		EnsureNonDefault,
		NewMinMaxValueValidator[float32](200.02, true, true),
		NewMinMaxValueValidator[float32](300.03, true, false),
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		EnsureNonDefault,
		NewMinMaxValueValidator[float64](300.03, true, true),
		NewMinMaxValueValidator[float64](400.04, true, false),
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		EnsureNonDefault,
		NewMinMaxValueValidator[int32](400, true, true),
		NewMinMaxValueValidator[int32](500, true, false),
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		EnsureNonDefault,
		NewMinMaxValueValidator[int32](500, true, true),
		NewMinMaxValueValidator[int32](600, true, false),
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		EnsureNonDefault,
		NewMinMaxValueValidator[int64](600, true, true),
		NewMinMaxValueValidator[int64](700, true, false),
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.NumericTypesRangeValidationExclusiveRequest) {
		validateNumberAny(bindingCtx, "numberAny", location, value.NumberAny)
		validateNumberFloat(bindingCtx, "numberFloat", location, value.NumberFloat)
		validateNumberDouble(bindingCtx, "numberDouble", location, value.NumberDouble)
		validateNumberInt(bindingCtx, "numberInt", location, value.NumberInt)
		validateNumberInt32(bindingCtx, "numberInt32", location, value.NumberInt32)
		validateNumberInt64(bindingCtx, "numberInt64", location, value.NumberInt64)
	}
}
