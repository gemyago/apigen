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

// NumericTypesParsingParams - Parameters for the numericTypesParsing operation
type NumericTypesParsingParams struct { 
	NumberAny float32 `json:"numberAny,omitempty"`
	NumberFloat float32 `json:"numberFloat,omitempty"`
	NumberDouble float64 `json:"numberDouble,omitempty"`
	NumberInt int32 `json:"numberInt,omitempty"`
	NumberInt32 int32 `json:"numberInt32,omitempty"`
	NumberInt64 int64 `json:"numberInt64,omitempty"`
	NumberAnyInQuery float32 `json:"numberAnyInQuery,omitempty"`
	NumberFloatInQuery float32 `json:"numberFloatInQuery,omitempty"`
	NumberDoubleInQuery float64 `json:"numberDoubleInQuery,omitempty"`
	NumberIntInQuery int32 `json:"numberIntInQuery,omitempty"`
	NumberInt32InQuery int32 `json:"numberInt32InQuery,omitempty"`
	NumberInt64InQuery int64 `json:"numberInt64InQuery,omitempty"`
	Payload *NumericTypesParsingRequest `json:"payload,omitempty"`
}
