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

// StringTypesPatternValidationParams - Parameters for the stringTypesPatternValidation operation.
type StringTypesPatternValidationParams struct { 
	UnformattedStr string `json:"unformattedStr"`
	CustomFormatStr string `json:"customFormatStr"`
	DateStr time.Time `json:"dateStr"`
	DateTimeStr time.Time `json:"dateTimeStr"`
	UnformattedStrInQuery string `json:"unformattedStrInQuery"`
	CustomFormatStrInQuery string `json:"customFormatStrInQuery"`
	DateStrInQuery time.Time `json:"dateStrInQuery"`
	DateTimeStrInQuery time.Time `json:"dateTimeStrInQuery"`
	Payload *StringTypesPatternValidationRequest `json:"payload"`
}
