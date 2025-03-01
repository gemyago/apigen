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

// StringTypesPatternValidationParams - Parameters for the stringTypesPatternValidation operation
type StringTypesPatternValidationParams struct { 
	UnformattedStr string `json:"unformattedStr,omitempty"`
	CustomFormatStr string `json:"customFormatStr,omitempty"`
	DateStr time.Time `json:"dateStr,omitempty"`
	DateTimeStr time.Time `json:"dateTimeStr,omitempty"`
	UnformattedStrInQuery string `json:"unformattedStrInQuery,omitempty"`
	CustomFormatStrInQuery string `json:"customFormatStrInQuery,omitempty"`
	DateStrInQuery time.Time `json:"dateStrInQuery,omitempty"`
	DateTimeStrInQuery time.Time `json:"dateTimeStrInQuery,omitempty"`
	Payload *StringTypesPatternValidationRequest `json:"payload,omitempty"`
}
