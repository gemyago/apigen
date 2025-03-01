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

// ObjectsRequiredBodyParams - Parameters for the objectsRequiredBody operation
type ObjectsRequiredBodyParams struct { 
	Payload *SimpleObject `json:"payload,omitempty"`
}
