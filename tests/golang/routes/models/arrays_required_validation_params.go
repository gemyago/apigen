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

// ArraysRequiredValidationParams - Parameters for the arraysRequiredValidation operation
type ArraysRequiredValidationParams struct { 
	SimpleItems1 []string `json:"simpleItems1,omitempty"`
	SimpleItems2 []string `json:"simpleItems2,omitempty"`
	SimpleItems1InQuery []string `json:"simpleItems1InQuery,omitempty"`
	SimpleItems2InQuery []string `json:"simpleItems2InQuery,omitempty"`
	OptionalSimpleItems1InQuery []string `json:"optionalSimpleItems1InQuery,omitempty"`
	OptionalSimpleItems2InQuery []string `json:"optionalSimpleItems2InQuery,omitempty"`
	Payload *ArraysRequiredValidationRequest `json:"payload,omitempty"`
}
