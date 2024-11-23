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

type StringTypesArrayItemsRangeValidationRequest struct { 
	UnformattedStr []string `json:"unformattedStr"`
	CustomFormatStr []string `json:"customFormatStr"`
	DateStr []time.Time `json:"dateStr"`
	DateTimeStr []time.Time `json:"dateTimeStr"`
	ByteStr []string `json:"byteStr"`
}
