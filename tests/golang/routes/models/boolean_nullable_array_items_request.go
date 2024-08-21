package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type BooleanNullableArrayItemsRequest struct { 
	BoolParam1 []*bool `json:"boolParam1"`
	BoolParam2 []*bool `json:"boolParam2"`
}
