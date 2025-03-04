package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesNullableRequestValidator() FieldValidator[*NumericTypesNullableRequest] {
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
	
	return func(bindingCtx *BindingContext, value *NumericTypesNullableRequest) {
		validateNumberAny(bindingCtx.Fork("numberAny"), value.NumberAny)
		validateOptionalNumberAny(bindingCtx.Fork("optionalNumberAny"), value.OptionalNumberAny)
		validateNumberFloat(bindingCtx.Fork("numberFloat"), value.NumberFloat)
		validateNumberDouble(bindingCtx.Fork("numberDouble"), value.NumberDouble)
		validateNumberInt(bindingCtx.Fork("numberInt"), value.NumberInt)
		validateNumberInt32(bindingCtx.Fork("numberInt32"), value.NumberInt32)
		validateNumberInt64(bindingCtx.Fork("numberInt64"), value.NumberInt64)
	}
}
