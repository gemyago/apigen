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

// ObjectsNullableOptionalBodyParams - Parameters for the objectsNullableOptionalBody operation.
type ObjectsNullableOptionalBodyParams struct { 
	Payload *SimpleNullableObject `json:"payload,omitempty"`
}
