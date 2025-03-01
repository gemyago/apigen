package internal

import (
	"time"
	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewNumericTypesNullableArrayItemsParamsValidator() FieldValidator[*NumericTypesNullableArrayItemsParams] {
	validateNumberAny := NewArrayValidator[*float32](
		NewSimpleFieldValidator[[]*float32](
		),
		NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
			),
	)
	validateNumberFloat := NewArrayValidator[*float32](
		NewSimpleFieldValidator[[]*float32](
		),
		NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
			),
	)
	validateNumberDouble := NewArrayValidator[*float64](
		NewSimpleFieldValidator[[]*float64](
		),
		NewSimpleFieldValidator[*float64](
				SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
			),
	)
	validateNumberInt := NewArrayValidator[*int32](
		NewSimpleFieldValidator[[]*int32](
		),
		NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
			),
	)
	validateNumberInt32 := NewArrayValidator[*int32](
		NewSimpleFieldValidator[[]*int32](
		),
		NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
			),
	)
	validateNumberInt64 := NewArrayValidator[*int64](
		NewSimpleFieldValidator[[]*int64](
		),
		NewSimpleFieldValidator[*int64](
				SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
			),
	)
	validateNumberAnyInQuery := NewArrayValidator[*float32](
		NewSimpleFieldValidator[[]*float32](
		),
		NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](100.01, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, false)),
			),
	)
	validateNumberFloatInQuery := NewArrayValidator[*float32](
		NewSimpleFieldValidator[[]*float32](
		),
		NewSimpleFieldValidator[*float32](
				SkipNullValidator(NewMinMaxValueValidator[float32](200.02, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float32](300.03, false, false)),
			),
	)
	validateNumberDoubleInQuery := NewArrayValidator[*float64](
		NewSimpleFieldValidator[[]*float64](
		),
		NewSimpleFieldValidator[*float64](
				SkipNullValidator(NewMinMaxValueValidator[float64](300.03, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[float64](400.04, false, false)),
			),
	)
	validateNumberIntInQuery := NewArrayValidator[*int32](
		NewSimpleFieldValidator[[]*int32](
		),
		NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](400, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, false)),
			),
	)
	validateNumberInt32InQuery := NewArrayValidator[*int32](
		NewSimpleFieldValidator[[]*int32](
		),
		NewSimpleFieldValidator[*int32](
				SkipNullValidator(NewMinMaxValueValidator[int32](500, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int32](600, false, false)),
			),
	)
	validateNumberInt64InQuery := NewArrayValidator[*int64](
		NewSimpleFieldValidator[[]*int64](
		),
		NewSimpleFieldValidator[*int64](
				SkipNullValidator(NewMinMaxValueValidator[int64](600, false, true)),
				SkipNullValidator(NewMinMaxValueValidator[int64](700, false, false)),
			),
	)
	validatePayload := NewObjectFieldValidator(
		ObjectFieldValidatorParams{Required: false, Nullable: false},
		NewNumericTypesNullableArrayItemsRequestValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *NumericTypesNullableArrayItemsParams) {
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
