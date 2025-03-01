package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesRequiredValidationParamsValidator() FieldValidator[*NumericTypesRequiredValidationParams] {
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
	validateOptionalNumberAnyInQuery := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](100.01, false, true),
	)
	validateOptionalNumberFloatInQuery := NewSimpleFieldValidator[float32](
		NewMinMaxValueValidator[float32](200.02, false, true),
	)
	validateOptionalNumberDoubleInQuery := NewSimpleFieldValidator[float64](
		NewMinMaxValueValidator[float64](300.03, false, true),
	)
	validateOptionalNumberIntInQuery := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](400, false, true),
	)
	validateOptionalNumberInt32InQuery := NewSimpleFieldValidator[int32](
		NewMinMaxValueValidator[int32](500, false, true),
	)
	validateOptionalNumberInt64InQuery := NewSimpleFieldValidator[int64](
		NewMinMaxValueValidator[int64](600, false, true),
	)
	
	return func(bindingCtx *BindingContext, value *NumericTypesRequiredValidationParams) {
		validateNumberAnyInQuery(bindingCtx.Fork("numberAnyInQuery"), value.NumberAnyInQuery)
		validateNumberFloatInQuery(bindingCtx.Fork("numberFloatInQuery"), value.NumberFloatInQuery)
		validateNumberDoubleInQuery(bindingCtx.Fork("numberDoubleInQuery"), value.NumberDoubleInQuery)
		validateNumberIntInQuery(bindingCtx.Fork("numberIntInQuery"), value.NumberIntInQuery)
		validateNumberInt32InQuery(bindingCtx.Fork("numberInt32InQuery"), value.NumberInt32InQuery)
		validateNumberInt64InQuery(bindingCtx.Fork("numberInt64InQuery"), value.NumberInt64InQuery)
		validateOptionalNumberAnyInQuery(bindingCtx.Fork("optionalNumberAnyInQuery"), value.OptionalNumberAnyInQuery)
		validateOptionalNumberFloatInQuery(bindingCtx.Fork("optionalNumberFloatInQuery"), value.OptionalNumberFloatInQuery)
		validateOptionalNumberDoubleInQuery(bindingCtx.Fork("optionalNumberDoubleInQuery"), value.OptionalNumberDoubleInQuery)
		validateOptionalNumberIntInQuery(bindingCtx.Fork("optionalNumberIntInQuery"), value.OptionalNumberIntInQuery)
		validateOptionalNumberInt32InQuery(bindingCtx.Fork("optionalNumberInt32InQuery"), value.OptionalNumberInt32InQuery)
		validateOptionalNumberInt64InQuery(bindingCtx.Fork("optionalNumberInt64InQuery"), value.OptionalNumberInt64InQuery)
	}
}
