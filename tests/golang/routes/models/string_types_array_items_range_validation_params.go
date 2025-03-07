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

// StringTypesArrayItemsRangeValidationParams - Parameters for the stringTypesArrayItemsRangeValidation operation.
type StringTypesArrayItemsRangeValidationParams struct { 
	UnformattedStr []string `json:"unformattedStr"`
	CustomFormatStr []string `json:"customFormatStr"`
	DateStr []time.Time `json:"dateStr"`
	DateTimeStr []time.Time `json:"dateTimeStr"`
	ByteStr []string `json:"byteStr"`
	UnformattedStrInQuery []string `json:"unformattedStrInQuery"`
	CustomFormatStrInQuery []string `json:"customFormatStrInQuery"`
	DateStrInQuery []time.Time `json:"dateStrInQuery"`
	DateTimeStrInQuery []time.Time `json:"dateTimeStrInQuery"`
	ByteStrInQuery []string `json:"byteStrInQuery"`
	Payload *StringTypesArrayItemsRangeValidationRequest `json:"payload"`
}
