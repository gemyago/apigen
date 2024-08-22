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

type NumericTypesNullableArrayItemsRequest struct { 
	NumberAny []*float32 `json:"numberAny"`
	NumberFloat []*float32 `json:"numberFloat"`
	NumberDouble []*float64 `json:"numberDouble"`
	NumberInt []*int32 `json:"numberInt"`
	NumberInt32 []*int32 `json:"numberInt32"`
	NumberInt64 []*int64 `json:"numberInt64"`
}
