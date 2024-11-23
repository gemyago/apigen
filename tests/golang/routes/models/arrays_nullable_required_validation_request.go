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

type ArraysNullableRequiredValidationRequest struct { 
	SimpleItems1 []string `json:"simpleItems1"`
	SimpleItems2 []string `json:"simpleItems2"`
	OptionalSimpleItems1 []string `json:"optionalSimpleItems1,omitempty"`
	OptionalSimpleItems2 []string `json:"optionalSimpleItems2,omitempty"`
}
