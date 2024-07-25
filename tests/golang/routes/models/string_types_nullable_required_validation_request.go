package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type StringTypesNullableRequiredValidationRequest struct { 
	UnformattedStr *string `json:"unformattedStr"`
	OptionalUnformattedStr *string `json:"optionalUnformattedStr,omitempty"`
}
