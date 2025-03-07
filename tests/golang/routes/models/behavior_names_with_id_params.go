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
	ID string `json:"id"`
	EndsWithID string `json:"endsWithId"`
	TheIDInTheMiddle string `json:"theIdInTheMiddle"`
	QueryEndsWithID string `json:"queryEndsWithId"`
	QueryTheIDInTheMiddle string `json:"queryTheIdInTheMiddle"`
	Payload *BehaviorNamesWithIDData `json:"payload"`
}
