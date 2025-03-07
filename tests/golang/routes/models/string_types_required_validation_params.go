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

// StringTypesRequiredValidationParams - Parameters for the stringTypesRequiredValidation operation.
type StringTypesRequiredValidationParams struct { 
	UnformattedStrInQuery string `json:"unformattedStrInQuery"`
	CustomFormatStrInQuery string `json:"customFormatStrInQuery"`
	DateStrInQuery time.Time `json:"dateStrInQuery"`
	DateTimeStrInQuery time.Time `json:"dateTimeStrInQuery"`
	ByteStrInQuery string `json:"byteStrInQuery"`
	OptionalUnformattedStrInQuery string `json:"optionalUnformattedStrInQuery,omitempty"`
	OptionalCustomFormatStrInQuery string `json:"optionalCustomFormatStrInQuery,omitempty"`
	OptionalDateStrInQuery time.Time `json:"optionalDateStrInQuery,omitempty"`
	OptionalDateTimeStrInQuery time.Time `json:"optionalDateTimeStrInQuery,omitempty"`
	OptionalByteStrInQuery string `json:"optionalByteStrInQuery,omitempty"`
	Payload *StringTypesRequiredValidationRequest `json:"payload"`
}
