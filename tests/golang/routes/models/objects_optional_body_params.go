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

// ObjectsOptionalBodyParams - Parameters for the objectsOptionalBody operation.
type ObjectsOptionalBodyParams struct { 
	Payload *SimpleObject `json:"payload"`
}
