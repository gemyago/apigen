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

type BooleanRequiredValidationRequest struct { 
	BoolParam1 bool `json:"boolParam1"`
	BoolParam2 bool `json:"boolParam2"`
	OptionalBoolParam1 bool `json:"optionalBoolParam1"`
	OptionalBoolParam2 bool `json:"optionalBoolParam2"`
}
