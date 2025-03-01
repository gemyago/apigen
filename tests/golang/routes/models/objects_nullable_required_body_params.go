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

// ObjectsNullableRequiredBodyParams - Parameters for the objectsNullableRequiredBody operation.
type ObjectsNullableRequiredBodyParams struct { 
	Payload *SimpleNullableObject `json:"payload,omitempty"`
}
