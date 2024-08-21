package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesArrayItemsRequestValidator() FieldValidator[*models.NumericTypesArrayItemsRequest] {
	validateNumberAny := NewArrayValidator[float32](
		NewSimpleFieldValidator[[]float32](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](100.01, false, true),
				NewMinMaxValueValidator[float32](200.02, false, false),
			),
	)
	validateNumberFloat := NewArrayValidator[float32](
		NewSimpleFieldValidator[[]float32](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[float32](
				NewMinMaxValueValidator[float32](200.02, false, true),
				NewMinMaxValueValidator[float32](300.03, false, false),
			),
	)
	validateNumberDouble := NewArrayValidator[float64](
		NewSimpleFieldValidator[[]float64](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[float64](
				NewMinMaxValueValidator[float64](300.03, false, true),
				NewMinMaxValueValidator[float64](400.04, false, false),
			),
	)
	validateNumberInt := NewArrayValidator[int32](
		NewSimpleFieldValidator[[]int32](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](400, false, true),
				NewMinMaxValueValidator[int32](500, false, false),
			),
	)
	validateNumberInt32 := NewArrayValidator[int32](
		NewSimpleFieldValidator[[]int32](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[int32](
				NewMinMaxValueValidator[int32](500, false, true),
				NewMinMaxValueValidator[int32](600, false, false),
			),
	)
	validateNumberInt64 := NewArrayValidator[int64](
		NewSimpleFieldValidator[[]int64](
			EnsureArrayFieldRequired,
		),
		NewSimpleFieldValidator[int64](
				NewMinMaxValueValidator[int64](600, false, true),
				NewMinMaxValueValidator[int64](700, false, false),
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
