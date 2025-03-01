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

// NumericTypesRequiredValidationParams - Parameters for the numericTypesRequiredValidation operation.
type NumericTypesRequiredValidationParams struct { 
	NumberAnyInQuery float32 `json:"numberAnyInQuery"`
	NumberFloatInQuery float32 `json:"numberFloatInQuery"`
	NumberDoubleInQuery float64 `json:"numberDoubleInQuery"`
	NumberIntInQuery int32 `json:"numberIntInQuery"`
	NumberInt32InQuery int32 `json:"numberInt32InQuery"`
	NumberInt64InQuery int64 `json:"numberInt64InQuery"`
	OptionalNumberAnyInQuery float32 `json:"optionalNumberAnyInQuery,omitempty"`
	OptionalNumberFloatInQuery float32 `json:"optionalNumberFloatInQuery,omitempty"`
	OptionalNumberDoubleInQuery float64 `json:"optionalNumberDoubleInQuery,omitempty"`
	OptionalNumberIntInQuery int32 `json:"optionalNumberIntInQuery,omitempty"`
	OptionalNumberInt32InQuery int32 `json:"optionalNumberInt32InQuery,omitempty"`
	OptionalNumberInt64InQuery int64 `json:"optionalNumberInt64InQuery,omitempty"`
}
