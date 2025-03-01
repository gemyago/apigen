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

// BehaviorWithParamsAndResponseParams - Parameters for the behaviorWithParamsAndResponse operation.
type BehaviorWithParamsAndResponseParams struct { 
	QueryParam1 string `json:"queryParam1,omitempty"`
	QueryParam2 int32 `json:"queryParam2,omitempty"`
	Payload *BehaviorWithParamsAndResponseRequestBody `json:"payload,omitempty"`
}
