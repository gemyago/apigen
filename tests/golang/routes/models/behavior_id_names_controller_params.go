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

// BehaviorIDNamesBehaviorNamesWithIDParams represents params for behaviorNamesWithId operation
//
// Request: POST /behavior/id-names/{id}/{endsWithId}/{theIdInTheMiddle}.
type BehaviorIDNamesBehaviorNamesWithIDParams struct {
	// ID is parsed from request path and declared as id.
	ID string
	// EndsWithID is parsed from request path and declared as endsWithId.
	EndsWithID string
	// TheIDInTheMiddle is parsed from request path and declared as theIdInTheMiddle.
	TheIDInTheMiddle string
	// QueryEndsWithID is parsed from request query and declared as queryEndsWithId.
	QueryEndsWithID string
	// QueryTheIDInTheMiddle is parsed from request query and declared as queryTheIdInTheMiddle.
	QueryTheIDInTheMiddle string
	// Payload is parsed from request body and declared as payload.
	Payload *BehaviorNamesWithIDData
}
