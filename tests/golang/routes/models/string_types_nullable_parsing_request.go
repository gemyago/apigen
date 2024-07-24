package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type StringTypesNullableParsingRequest struct { 
	UnformattedStr *string `json:"unformattedStr,omitempty"`
	CustomFormatStr *string `json:"customFormatStr,omitempty"`
	DateStr *time.Time `json:"dateStr,omitempty"`
	DateTimeStr *time.Time `json:"dateTimeStr,omitempty"`
	ByteStr *string `json:"byteStr,omitempty"`
}
