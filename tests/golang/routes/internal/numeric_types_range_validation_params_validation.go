package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesRangeValidationParamsValidator() FieldValidator[*NumericTypesRangeValidationParams] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](100.01, false, true),
		NewMinMaxValueValidator[float32](200.02, false, false),
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](200.02, false, true),
		NewMinMaxValueValidator[float32](300.03, false, false),
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		NewMinMaxValueValidator[float64](300.03, false, true),
		NewMinMaxValueValidator[float64](400.04, false, false),
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](400, false, true),
		NewMinMaxValueValidator[int32](500, false, false),
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](500, false, true),
		NewMinMaxValueValidator[int32](600, false, false),
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		NewMinMaxValueValidator[int64](600, false, true),
		NewMinMaxValueValidator[int64](700, false, false),
	)
	validateNumberAnyInQuery := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](100.01, false, true),
		NewMinMaxValueValidator[float32](200.02, false, false),
	)
	validateNumberFloatInQuery := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](200.02, false, true),
		NewMinMaxValueValidator[float32](300.03, false, false),
	)
	validateNumberDoubleInQuery := NewSimpleFieldValidator[float64](
		NewMinMaxValueValidator[float64](300.03, false, true),
		NewMinMaxValueValidator[float64](400.04, false, false),
	)
	validateNumberIntInQuery := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](400, false, true),
		NewMinMaxValueValidator[int32](500, false, false),
	)
	validateNumberInt32InQuery := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](500, false, true),
		NewMinMaxValueValidator[int32](600, false, false),
	)
	validateNumberInt64InQuery := NewSimpleFieldValidator[int64](
		NewMinMaxValueValidator[int64](600, false, true),
		NewMinMaxValueValidator[int64](700, false, false),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewNumericTypesRangeValidationRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *NumericTypesRangeValidationParams) {
		validateNumberAny(bindingCtx.Fork("numberAny"), value.NumberAny)
		validateNumberFloat(bindingCtx.Fork("numberFloat"), value.NumberFloat)
		validateNumberDouble(bindingCtx.Fork("numberDouble"), value.NumberDouble)
		validateNumberInt(bindingCtx.Fork("numberInt"), value.NumberInt)
		validateNumberInt32(bindingCtx.Fork("numberInt32"), value.NumberInt32)
		validateNumberInt64(bindingCtx.Fork("numberInt64"), value.NumberInt64)
		validateNumberAnyInQuery(bindingCtx.Fork("numberAnyInQuery"), value.NumberAnyInQuery)
		validateNumberFloatInQuery(bindingCtx.Fork("numberFloatInQuery"), value.NumberFloatInQuery)
		validateNumberDoubleInQuery(bindingCtx.Fork("numberDoubleInQuery"), value.NumberDoubleInQuery)
		validateNumberIntInQuery(bindingCtx.Fork("numberIntInQuery"), value.NumberIntInQuery)
		validateNumberInt32InQuery(bindingCtx.Fork("numberInt32InQuery"), value.NumberInt32InQuery)
		validateNumberInt64InQuery(bindingCtx.Fork("numberInt64InQuery"), value.NumberInt64InQuery)
		validatePayload(bindingCtx.Fork("payload"), value.Payload)
	}
}
