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

type BehaviorNoParamsWithResponse202Response struct { 
	Field1 string `json:"field1,omitempty"`
}
