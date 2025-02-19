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

type BehaviorWithParamsAndResponseResponseBody struct { 
	Field1 string `json:"field1,omitempty"`
	Field2 *BehaviorWithParamsAndResponseResponseBody `json:"field2,omitempty"`
}
