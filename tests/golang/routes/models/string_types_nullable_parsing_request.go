package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type StringTypesNullableParsingRequest struct { 
	UnformattedStr *string `json:"unformattedStr"`
	CustomFormatStr *string `json:"customFormatStr"`
	DateStr *time.Time `json:"dateStr"`
	DateTimeStr *time.Time `json:"dateTimeStr"`
	ByteStr *string `json:"byteStr"`
}
