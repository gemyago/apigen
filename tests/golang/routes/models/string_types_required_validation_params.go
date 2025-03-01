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
	UnformattedStrInQuery string `json:"unformattedStrInQuery,omitempty"`
	CustomFormatStrInQuery string `json:"customFormatStrInQuery,omitempty"`
	DateStrInQuery time.Time `json:"dateStrInQuery,omitempty"`
	DateTimeStrInQuery time.Time `json:"dateTimeStrInQuery,omitempty"`
	ByteStrInQuery string `json:"byteStrInQuery,omitempty"`
	OptionalUnformattedStrInQuery string `json:"optionalUnformattedStrInQuery,omitempty"`
	OptionalCustomFormatStrInQuery string `json:"optionalCustomFormatStrInQuery,omitempty"`
	OptionalDateStrInQuery time.Time `json:"optionalDateStrInQuery,omitempty"`
	OptionalDateTimeStrInQuery time.Time `json:"optionalDateTimeStrInQuery,omitempty"`
	OptionalByteStrInQuery string `json:"optionalByteStrInQuery,omitempty"`
	Payload *StringTypesRequiredValidationRequest `json:"payload,omitempty"`
}
