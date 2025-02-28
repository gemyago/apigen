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







// BehaviorBehaviorWithParamsAndResponseParams represents params for behaviorWithParamsAndResponse operation
//
// Request: POST /behavior/with-params-and-response.
type BehaviorBehaviorWithParamsAndResponseParams struct {
	// QueryParam1 is parsed from request query and declared as queryParam1.
	QueryParam1 string
	// QueryParam2 is parsed from request query and declared as queryParam2.
	QueryParam2 int32
	// Payload is parsed from request body and declared as payload.
	Payload *BehaviorWithParamsAndResponseRequestBody
}

// BehaviorBehaviorWithParamsNoResponseParams represents params for behaviorWithParamsNoResponse operation
//
// Request: GET /behavior/with-params-no-response.
type BehaviorBehaviorWithParamsNoResponseParams struct {
	// QueryParam1 is parsed from request query and declared as queryParam1.
	QueryParam1 string
}


