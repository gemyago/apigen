package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ArraysRangeValidationRequest struct { 
	SimpleItems1 []string `json:"simpleItems1"`
	SimpleItems2 []string `json:"simpleItems2"`
	OptionalSimpleItems1 []string `json:"optionalSimpleItems1,omitempty"`
	OptionalSimpleItems2 []string `json:"optionalSimpleItems2,omitempty"`
}
