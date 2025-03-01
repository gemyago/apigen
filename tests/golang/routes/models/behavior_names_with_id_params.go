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

// BehaviorNamesWithIDParams - Parameters for the behaviorNamesWithId operation.
type BehaviorNamesWithIDParams struct { 
	ID string `json:"id,omitempty"`
	EndsWithID string `json:"endsWithId,omitempty"`
	TheIDInTheMiddle string `json:"theIdInTheMiddle,omitempty"`
	QueryEndsWithID string `json:"queryEndsWithId,omitempty"`
	QueryTheIDInTheMiddle string `json:"queryTheIdInTheMiddle,omitempty"`
	Payload *BehaviorNamesWithIDData `json:"payload,omitempty"`
}
