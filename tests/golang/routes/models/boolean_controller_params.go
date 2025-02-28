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

// BooleanArrayItemsParams represents params for booleanArrayItems operation
//
// Request: POST /boolean/array-items/{boolParam1}/{boolParam2}.
type BooleanArrayItemsParams struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 []bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 []bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery []bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery []bool
	// Payload is parsed from request body and declared as payload.
	Payload *BooleanArrayItemsRequest
}

// BooleanNullableParams represents params for booleanNullable operation
//
// Request: POST /boolean/nullable/{boolParam1}/{boolParam2}.
type BooleanNullableParams struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 *bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 *bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery *bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery *bool
	// Payload is parsed from request body and declared as payload.
	Payload *BooleanNullableRequest
	// OptionalBoolParam1InQuery is parsed from request query and declared as optionalBoolParam1InQuery.
	OptionalBoolParam1InQuery *bool
}

// BooleanNullableArrayItemsParams represents params for booleanNullableArrayItems operation
//
// Request: POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}.
type BooleanNullableArrayItemsParams struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 []*bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 []*bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery []*bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery []*bool
	// Payload is parsed from request body and declared as payload.
	Payload *BooleanNullableArrayItemsRequest
}

// BooleanParsingParams represents params for booleanParsing operation
//
// Request: POST /boolean/parsing/{boolParam1}/{boolParam2}.
type BooleanParsingParams struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery bool
	// Payload is parsed from request body and declared as payload.
	Payload *BooleanParsingRequest
}

// BooleanRequiredValidationParams represents params for booleanRequiredValidation operation
//
// Request: POST /boolean/required-validation.
type BooleanRequiredValidationParams struct {
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery bool
	// Payload is parsed from request body and declared as payload.
	Payload *BooleanRequiredValidationRequest
	// OptionalBoolParam1InQuery is parsed from request query and declared as optionalBoolParam1InQuery.
	OptionalBoolParam1InQuery bool
	// OptionalBoolParam2InQuery is parsed from request query and declared as optionalBoolParam2InQuery.
	OptionalBoolParam2InQuery bool
}
