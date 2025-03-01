package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesRangeValidationExclusiveParamsValidator() FieldValidator[*NumericTypesRangeValidationExclusiveParams] {
	validateNumberAny := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](100.01, true, true),
		NewMinMaxValueValidator[float32](200.02, true, false),
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](200.02, true, true),
		NewMinMaxValueValidator[float32](300.03, true, false),
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
		NewMinMaxValueValidator[float64](300.03, true, true),
		NewMinMaxValueValidator[float64](400.04, true, false),
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](400, true, true),
		NewMinMaxValueValidator[int32](500, true, false),
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](500, true, true),
		NewMinMaxValueValidator[int32](600, true, false),
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
		NewMinMaxValueValidator[int64](600, true, true),
		NewMinMaxValueValidator[int64](700, true, false),
	)
	validateNumberAnyInQuery := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](100.01, true, true),
		NewMinMaxValueValidator[float32](200.02, true, false),
	)
	validateNumberFloatInQuery := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](200.02, true, true),
		NewMinMaxValueValidator[float32](300.03, true, false),
	)
	validateNumberDoubleInQuery := NewSimpleFieldValidator[float64](
		NewMinMaxValueValidator[float64](300.03, true, true),
		NewMinMaxValueValidator[float64](400.04, true, false),
	)
	validateNumberIntInQuery := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](400, true, true),
		NewMinMaxValueValidator[int32](500, true, false),
	)
	validateNumberInt32InQuery := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](500, true, true),
		NewMinMaxValueValidator[int32](600, true, false),
	)
	validateNumberInt64InQuery := NewSimpleFieldValidator[int64](
		NewMinMaxValueValidator[int64](600, true, true),
		NewMinMaxValueValidator[int64](700, true, false),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewNumericTypesRangeValidationExclusiveRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *NumericTypesRangeValidationExclusiveParams) {
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
