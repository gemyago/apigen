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

type Error struct { 
	Code *interface{} `json:"code"`
	Message string `json:"message,omitempty"`
}
