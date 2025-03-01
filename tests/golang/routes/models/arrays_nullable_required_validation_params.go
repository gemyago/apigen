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

// ArraysNullableRequiredValidationParams - Parameters for the arraysNullableRequiredValidation operation.
type ArraysNullableRequiredValidationParams struct { 
	SimpleItems1 []string `json:"simpleItems1,omitempty"`
	SimpleItems2 []string `json:"simpleItems2,omitempty"`
	SimpleItems1InQuery []string `json:"simpleItems1InQuery,omitempty"`
	SimpleItems2InQuery []string `json:"simpleItems2InQuery,omitempty"`
	OptionalSimpleItems1InQuery []string `json:"optionalSimpleItems1InQuery,omitempty"`
	OptionalSimpleItems2InQuery []string `json:"optionalSimpleItems2InQuery,omitempty"`
	Payload *ArraysNullableRequiredValidationRequest `json:"payload,omitempty"`
}
