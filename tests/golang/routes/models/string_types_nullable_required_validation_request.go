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

type StringTypesNullableRequiredValidationRequest struct { 
	UnformattedStr *string `json:"unformattedStr"`
	OptionalUnformattedStr *string `json:"optionalUnformattedStr,omitempty"`
}
