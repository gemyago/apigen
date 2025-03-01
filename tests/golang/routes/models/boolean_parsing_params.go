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

// BooleanParsingParams - Parameters for the booleanParsing operation.
type BooleanParsingParams struct { 
	BoolParam1 bool `json:"boolParam1,omitempty"`
	BoolParam2 bool `json:"boolParam2,omitempty"`
	BoolParam1InQuery bool `json:"boolParam1InQuery,omitempty"`
	BoolParam2InQuery bool `json:"boolParam2InQuery,omitempty"`
	Payload *BooleanParsingRequest `json:"payload,omitempty"`
}
