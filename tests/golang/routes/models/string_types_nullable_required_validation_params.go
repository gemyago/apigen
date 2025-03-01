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

// StringTypesNullableRequiredValidationParams - Parameters for the stringTypesNullableRequiredValidation operation
type StringTypesNullableRequiredValidationParams struct { 
	UnformattedStrInQuery *string `json:"unformattedStrInQuery,omitempty"`
	OptionalUnformattedStrInQuery *string `json:"optionalUnformattedStrInQuery,omitempty"`
	Payload *StringTypesNullableRequiredValidationRequest `json:"payload,omitempty"`
}
