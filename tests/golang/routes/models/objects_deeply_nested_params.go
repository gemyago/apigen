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

// ObjectsDeeplyNestedParams - Parameters for the objectsDeeplyNested operation.
type ObjectsDeeplyNestedParams struct { 
	Payload *ObjectsDeeplyNestedRequest `json:"payload"`
}
