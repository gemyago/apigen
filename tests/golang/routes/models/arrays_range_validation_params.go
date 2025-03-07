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

// ArraysRangeValidationParams - Parameters for the arraysRangeValidation operation.
type ArraysRangeValidationParams struct { 
	SimpleItems1 []string `json:"simpleItems1"`
	SimpleItems2 []string `json:"simpleItems2"`
	SimpleItems1InQuery []string `json:"simpleItems1InQuery"`
	SimpleItems2InQuery []string `json:"simpleItems2InQuery"`
	OptionalSimpleItems1InQuery []string `json:"optionalSimpleItems1InQuery,omitempty"`
	OptionalSimpleItems2InQuery []string `json:"optionalSimpleItems2InQuery,omitempty"`
	Payload *ArraysRangeValidationRequest `json:"payload"`
}
