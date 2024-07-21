package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Unused imports workaround.
var _ = json.Marshal
var _ = errors.New
var _ = fmt.Errorf
var _ = reflect.TypeOf
var _ = time.Time{}

type StringTypesParsingRequest struct { 
	UnformattedStr string `json:"unformattedStr"`
	CustomFormatStr string `json:"customFormatStr"`
	DateStr time.Time `json:"dateStr"`
	DateTimeStr time.Time `json:"dateTimeStr"`
	ByteStr string `json:"byteStr"`
}
