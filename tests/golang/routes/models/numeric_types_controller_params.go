package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Unused imports workaround.
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint

// NumericTypesArrayItemsParams represents params for numericTypesArrayItems operation
//
// Request: POST /numeric-types/array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesArrayItemsParams struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny []float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat []float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble []float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt []int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 []int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 []int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery []float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery []float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery []float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery []int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery []int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery []int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesArrayItemsRequest
}

// NumericTypesNullableParams represents params for numericTypesNullable operation
//
// Request: POST /numeric-types/nullable/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNullableParams struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny *float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat *float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble *float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt *int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 *int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 *int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery *float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery *float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery *float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery *int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery *int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery *int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesNullableRequest
	// OptionalNumberAnyInQuery is parsed from request query and declared as optionalNumberAnyInQuery.
	OptionalNumberAnyInQuery *float32
}

// NumericTypesNullableArrayItemsParams represents params for numericTypesNullableArrayItems operation
//
// Request: POST /numeric-types/nullable-array-items/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesNullableArrayItemsParams struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny []*float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat []*float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble []*float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt []*int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 []*int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 []*int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery []*float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery []*float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery []*float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery []*int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery []*int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery []*int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesNullableArrayItemsRequest
}

// NumericTypesParsingParams represents params for numericTypesParsing operation
//
// Request: POST /numeric-types/parsing/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesParsingParams struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesParsingRequest
}

// NumericTypesRangeValidationParams represents params for numericTypesRangeValidation operation
//
// Request: POST /numeric-types/range-validation/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesRangeValidationParams struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesRangeValidationRequest
}

// NumericTypesRangeValidationExclusiveParams represents params for numericTypesRangeValidationExclusive operation
//
// Request: POST /numeric-types/range-validation-exclusive/{numberAny}/{numberFloat}/{numberDouble}/{numberInt}/{numberInt32}/{numberInt64}.
type NumericTypesRangeValidationExclusiveParams struct {
	// NumberAny is parsed from request path and declared as numberAny.
	NumberAny float32
	// NumberFloat is parsed from request path and declared as numberFloat.
	NumberFloat float32
	// NumberDouble is parsed from request path and declared as numberDouble.
	NumberDouble float64
	// NumberInt is parsed from request path and declared as numberInt.
	NumberInt int32
	// NumberInt32 is parsed from request path and declared as numberInt32.
	NumberInt32 int32
	// NumberInt64 is parsed from request path and declared as numberInt64.
	NumberInt64 int64
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// Payload is parsed from request body and declared as payload.
	Payload *NumericTypesRangeValidationExclusiveRequest
}

// NumericTypesRequiredValidationParams represents params for numericTypesRequiredValidation operation
//
// Request: GET /numeric-types/required-validation.
type NumericTypesRequiredValidationParams struct {
	// NumberAnyInQuery is parsed from request query and declared as numberAnyInQuery.
	NumberAnyInQuery float32
	// NumberFloatInQuery is parsed from request query and declared as numberFloatInQuery.
	NumberFloatInQuery float32
	// NumberDoubleInQuery is parsed from request query and declared as numberDoubleInQuery.
	NumberDoubleInQuery float64
	// NumberIntInQuery is parsed from request query and declared as numberIntInQuery.
	NumberIntInQuery int32
	// NumberInt32InQuery is parsed from request query and declared as numberInt32InQuery.
	NumberInt32InQuery int32
	// NumberInt64InQuery is parsed from request query and declared as numberInt64InQuery.
	NumberInt64InQuery int64
	// OptionalNumberAnyInQuery is parsed from request query and declared as optionalNumberAnyInQuery.
	OptionalNumberAnyInQuery float32
	// OptionalNumberFloatInQuery is parsed from request query and declared as optionalNumberFloatInQuery.
	OptionalNumberFloatInQuery float32
	// OptionalNumberDoubleInQuery is parsed from request query and declared as optionalNumberDoubleInQuery.
	OptionalNumberDoubleInQuery float64
	// OptionalNumberIntInQuery is parsed from request query and declared as optionalNumberIntInQuery.
	OptionalNumberIntInQuery int32
	// OptionalNumberInt32InQuery is parsed from request query and declared as optionalNumberInt32InQuery.
	OptionalNumberInt32InQuery int32
	// OptionalNumberInt64InQuery is parsed from request query and declared as optionalNumberInt64InQuery.
	OptionalNumberInt64InQuery int64
}
