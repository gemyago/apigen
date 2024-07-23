package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type StringTypesRequiredValidationRequest struct { 
	UnformattedStr string `json:"unformattedStr"`
	CustomFormatStr string `json:"customFormatStr"`
	DateStr time.Time `json:"dateStr"`
	DateTimeStr time.Time `json:"dateTimeStr"`
	ByteStr string `json:"byteStr"`
	OptionalUnformattedStr string `json:"optionalUnformattedStr,omitempty"`
	OptionalCustomFormatStr string `json:"optionalCustomFormatStr,omitempty"`
	OptionalDateStr time.Time `json:"optionalDateStr,omitempty"`
	OptionalDateTimeStr time.Time `json:"optionalDateTimeStr,omitempty"`
	OptionalByteStr string `json:"optionalByteStr,omitempty"`
	OptionalValidatedUnformattedStr string `json:"optionalValidatedUnformattedStr,omitempty"`
	OptionalValidatedCustomFormatStr string `json:"optionalValidatedCustomFormatStr,omitempty"`
	OptionalValidatedByteStr string `json:"optionalValidatedByteStr,omitempty"`
}
