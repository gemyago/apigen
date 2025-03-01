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

// ObjectsArrayBodyNestedParams - Parameters for the objectsArrayBodyNested operation.
type ObjectsArrayBodyNestedParams struct { 
	Payload *ObjectsArrayBodyNestedRequest `json:"payload"`
}
