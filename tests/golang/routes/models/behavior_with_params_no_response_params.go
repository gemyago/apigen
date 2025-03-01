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

// BehaviorWithParamsNoResponseParams - Parameters for the behaviorWithParamsNoResponse operation
type BehaviorWithParamsNoResponseParams struct { 
	QueryParam1 string `json:"queryParam1,omitempty"`
}
