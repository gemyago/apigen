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

// NumericTypesNullableParams - Parameters for the numericTypesNullable operation.
type NumericTypesNullableParams struct { 
	NumberAny *float32 `json:"numberAny"`
	NumberFloat *float32 `json:"numberFloat"`
	NumberDouble *float64 `json:"numberDouble"`
	NumberInt *int32 `json:"numberInt"`
	NumberInt32 *int32 `json:"numberInt32"`
	NumberInt64 *int64 `json:"numberInt64"`
	NumberAnyInQuery *float32 `json:"numberAnyInQuery"`
	OptionalNumberAnyInQuery *float32 `json:"optionalNumberAnyInQuery,omitempty"`
	NumberFloatInQuery *float32 `json:"numberFloatInQuery"`
	NumberDoubleInQuery *float64 `json:"numberDoubleInQuery"`
	NumberIntInQuery *int32 `json:"numberIntInQuery"`
	NumberInt32InQuery *int32 `json:"numberInt32InQuery"`
	NumberInt64InQuery *int64 `json:"numberInt64InQuery"`
	Payload *NumericTypesNullableRequest `json:"payload"`
}
