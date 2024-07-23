package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type BehaviorNoParamsWithResponse202Response struct { 
	Field1 string `json:"field1,omitempty"`
}
