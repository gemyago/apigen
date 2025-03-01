package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesParsingParamsValidator() FieldValidator[*NumericTypesParsingParams] {
	validateNumberAny := NewSimpleFieldValidator[float32](
	)
	validateNumberFloat := NewSimpleFieldValidator[float32](
	)
	validateNumberDouble := NewSimpleFieldValidator[float64](
	)
	validateNumberInt := NewSimpleFieldValidator[int32](
	)
	validateNumberInt32 := NewSimpleFieldValidator[int32](
	)
	validateNumberInt64 := NewSimpleFieldValidator[int64](
	)
	validateNumberAnyInQuery := NewSimpleFieldValidator[float32](
	)
	validateNumberFloatInQuery := NewSimpleFieldValidator[float32](
	)
	validateNumberDoubleInQuery := NewSimpleFieldValidator[float64](
	)
	validateNumberIntInQuery := NewSimpleFieldValidator[int32](
	)
	validateNumberInt32InQuery := NewSimpleFieldValidator[int32](
	)
	validateNumberInt64InQuery := NewSimpleFieldValidator[int64](
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewNumericTypesParsingRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *NumericTypesParsingParams) {
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
