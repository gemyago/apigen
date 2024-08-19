package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesArrayItemsRequestValidator() FieldValidator[*models.NumericTypesArrayItemsRequest] {
	validateNumberAny := NewArrayValidator[float32](
		NewSimpleFieldValidator[float32](
			),
	)
	validateNumberFloat := NewArrayValidator[float32](
		NewSimpleFieldValidator[float32](
			),
	)
	validateNumberDouble := NewArrayValidator[float64](
		NewSimpleFieldValidator[float64](
			),
	)
	validateNumberInt := NewArrayValidator[int32](
		NewSimpleFieldValidator[int32](
			),
	)
	validateNumberInt32 := NewArrayValidator[int32](
		NewSimpleFieldValidator[int32](
			),
	)
	validateNumberInt64 := NewArrayValidator[int64](
		NewSimpleFieldValidator[int64](
			),
	)
	
	return func(bindingCtx *BindingContext, value *models.NumericTypesArrayItemsRequest) {
		validateNumberAny(bindingCtx.Fork("numberAny"), value.NumberAny)
		validateNumberFloat(bindingCtx.Fork("numberFloat"), value.NumberFloat)
		validateNumberDouble(bindingCtx.Fork("numberDouble"), value.NumberDouble)
		validateNumberInt(bindingCtx.Fork("numberInt"), value.NumberInt)
		validateNumberInt32(bindingCtx.Fork("numberInt32"), value.NumberInt32)
		validateNumberInt64(bindingCtx.Fork("numberInt64"), value.NumberInt64)
	}
}
