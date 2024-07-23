package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesParsingRequestValidator() FieldValidator[*models.NumericTypesParsingRequest] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		EnsureNonDefault,
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		EnsureNonDefault,
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		EnsureNonDefault,
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		EnsureNonDefault,
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		EnsureNonDefault,
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		EnsureNonDefault,
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.NumericTypesParsingRequest) {
		validateNumberAny(bindingCtx, "numberAny", location, value.NumberAny)
		validateNumberFloat(bindingCtx, "numberFloat", location, value.NumberFloat)
		validateNumberDouble(bindingCtx, "numberDouble", location, value.NumberDouble)
		validateNumberInt(bindingCtx, "numberInt", location, value.NumberInt)
		validateNumberInt32(bindingCtx, "numberInt32", location, value.NumberInt32)
		validateNumberInt64(bindingCtx, "numberInt64", location, value.NumberInt64)
	}
}
