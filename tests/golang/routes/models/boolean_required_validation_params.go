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

// BooleanRequiredValidationParams - Parameters for the booleanRequiredValidation operation
type BooleanRequiredValidationParams struct { 
	BoolParam1InQuery bool `json:"boolParam1InQuery,omitempty"`
	BoolParam2InQuery bool `json:"boolParam2InQuery,omitempty"`
	OptionalBoolParam1InQuery bool `json:"optionalBoolParam1InQuery,omitempty"`
	OptionalBoolParam2InQuery bool `json:"optionalBoolParam2InQuery,omitempty"`
	Payload *BooleanRequiredValidationRequest `json:"payload,omitempty"`
}
