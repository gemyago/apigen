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

// ObjectsRequiredNestedObjectsParams - Parameters for the objectsRequiredNestedObjects operation.
type ObjectsRequiredNestedObjectsParams struct { 
	Payload *SimpleObjectsContainer `json:"payload,omitempty"`
}
