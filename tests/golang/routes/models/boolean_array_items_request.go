package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type BooleanArrayItemsRequest struct { 
	BoolParam1 []bool `json:"boolParam1"`
	BoolParam2 []bool `json:"boolParam2"`
}
