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

// ArraysArraysNullableRequiredValidationParams represents params for arraysNullableRequiredValidation operation
//
// Request: POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}.
type ArraysArraysNullableRequiredValidationParams struct {
	// SimpleItems1 is parsed from request path and declared as simpleItems1.
	SimpleItems1 []string
	// SimpleItems2 is parsed from request path and declared as simpleItems2.
	SimpleItems2 []string
	// SimpleItems1InQuery is parsed from request query and declared as simpleItems1InQuery.
	SimpleItems1InQuery []string
	// SimpleItems2InQuery is parsed from request query and declared as simpleItems2InQuery.
	SimpleItems2InQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *ArraysNullableRequiredValidationRequest
	// OptionalSimpleItems1InQuery is parsed from request query and declared as optionalSimpleItems1InQuery.
	OptionalSimpleItems1InQuery []string
	// OptionalSimpleItems2InQuery is parsed from request query and declared as optionalSimpleItems2InQuery.
	OptionalSimpleItems2InQuery []string
}

// ArraysArraysRangeValidationParams represents params for arraysRangeValidation operation
//
// Request: POST /arrays/range-validation/{simpleItems1}/{simpleItems2}.
type ArraysArraysRangeValidationParams struct {
	// SimpleItems1 is parsed from request path and declared as simpleItems1.
	SimpleItems1 []string
	// SimpleItems2 is parsed from request path and declared as simpleItems2.
	SimpleItems2 []string
	// SimpleItems1InQuery is parsed from request query and declared as simpleItems1InQuery.
	SimpleItems1InQuery []string
	// SimpleItems2InQuery is parsed from request query and declared as simpleItems2InQuery.
	SimpleItems2InQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *ArraysRangeValidationRequest
	// OptionalSimpleItems1InQuery is parsed from request query and declared as optionalSimpleItems1InQuery.
	OptionalSimpleItems1InQuery []string
	// OptionalSimpleItems2InQuery is parsed from request query and declared as optionalSimpleItems2InQuery.
	OptionalSimpleItems2InQuery []string
}

// ArraysArraysRequiredValidationParams represents params for arraysRequiredValidation operation
//
// Request: POST /arrays/required-validation/{simpleItems1}/{simpleItems2}.
type ArraysArraysRequiredValidationParams struct {
	// SimpleItems1 is parsed from request path and declared as simpleItems1.
	SimpleItems1 []string
	// SimpleItems2 is parsed from request path and declared as simpleItems2.
	SimpleItems2 []string
	// SimpleItems1InQuery is parsed from request query and declared as simpleItems1InQuery.
	SimpleItems1InQuery []string
	// SimpleItems2InQuery is parsed from request query and declared as simpleItems2InQuery.
	SimpleItems2InQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *ArraysRequiredValidationRequest
	// OptionalSimpleItems1InQuery is parsed from request query and declared as optionalSimpleItems1InQuery.
	OptionalSimpleItems1InQuery []string
	// OptionalSimpleItems2InQuery is parsed from request query and declared as optionalSimpleItems2InQuery.
	OptionalSimpleItems2InQuery []string
}
