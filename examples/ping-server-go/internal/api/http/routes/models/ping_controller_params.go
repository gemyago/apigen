// Code generated by apigen DO NOT EDIT.

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

// PingParams represents params for ping operation
//
// Request: GET /ping.
type PingParams struct {
	// Message is parsed from request query and declared as message.
	Message string
}
