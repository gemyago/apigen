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

// BooleanNullableArrayItemsParams - Parameters for the booleanNullableArrayItems operation.
type BooleanNullableArrayItemsParams struct { 
	BoolParam1 []*bool `json:"boolParam1,omitempty"`
	BoolParam2 []*bool `json:"boolParam2,omitempty"`
	BoolParam1InQuery []*bool `json:"boolParam1InQuery,omitempty"`
	BoolParam2InQuery []*bool `json:"boolParam2InQuery,omitempty"`
	Payload *BooleanNullableArrayItemsRequest `json:"payload,omitempty"`
}
