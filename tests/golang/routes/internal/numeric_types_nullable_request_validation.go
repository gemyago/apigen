package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesNullableRequestValidator() FieldValidator[*models.NumericTypesNullableRequest] {
	validateNumberAny := NewSimpleFieldValidator[*float32](
		SkipNullValidator(EnsureNonDefault[float32]),
		SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
	)
	validateOptionalNumberAny := NewSimpleFieldValidator[*float32](
		SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
	)
	validateNumberFloat := NewSimpleFieldValidator[*float32](
		SkipNullValidator(EnsureNonDefault[float32]),
		SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
	)
	validateNumberDouble := NewSimpleFieldValidator[*float64](
		SkipNullValidator(EnsureNonDefault[float64]),
		SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
	)
	validateNumberInt := NewSimpleFieldValidator[*int32](
		SkipNullValidator(EnsureNonDefault[int32]),
		SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
	)
	validateNumberInt32 := NewSimpleFieldValidator[*int32](
		SkipNullValidator(EnsureNonDefault[int32]),
		SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
	)
	validateNumberInt64 := NewSimpleFieldValidator[*int64](
		SkipNullValidator(EnsureNonDefault[int64]),
		SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
		SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
	)
	
	return func(bindingCtx *BindingContext, field, location string, value *models.NumericTypesNullableRequest) {
		validateNumberAny(bindingCtx, "numberAny", location, value.NumberAny)
		validateOptionalNumberAny(bindingCtx, "optionalNumberAny", location, value.OptionalNumberAny)
		validateNumberFloat(bindingCtx, "numberFloat", location, value.NumberFloat)
		validateNumberDouble(bindingCtx, "numberDouble", location, value.NumberDouble)
		validateNumberInt(bindingCtx, "numberInt", location, value.NumberInt)
		validateNumberInt32(bindingCtx, "numberInt32", location, value.NumberInt32)
		validateNumberInt64(bindingCtx, "numberInt64", location, value.NumberInt64)
	}
}
