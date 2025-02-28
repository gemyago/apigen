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

// ObjectsArrayBodyDirectParams represents params for objectsArrayBodyDirect operation
//
// Request: POST /objects/arrays.
type ObjectsArrayBodyDirectParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload []*ObjectArraysSimpleObject
}

// ObjectsArrayBodyNestedParams represents params for objectsArrayBodyNested operation
//
// Request: PUT /objects/arrays.
type ObjectsArrayBodyNestedParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *ObjectsArrayBodyNestedRequest
}

// ObjectsDeeplyNestedParams represents params for objectsDeeplyNested operation
//
// Request: POST /objects/deeply-nested.
type ObjectsDeeplyNestedParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *ObjectsDeeplyNestedRequest
}

// ObjectsNullableOptionalBodyParams represents params for objectsNullableOptionalBody operation
//
// Request: PUT /objects/nullable-body.
type ObjectsNullableOptionalBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleNullableObject
}

// ObjectsNullableRequiredBodyParams represents params for objectsNullableRequiredBody operation
//
// Request: POST /objects/nullable-body.
type ObjectsNullableRequiredBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleNullableObject
}

// ObjectsOptionalBodyParams represents params for objectsOptionalBody operation
//
// Request: PUT /objects/required-body.
type ObjectsOptionalBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObject
}

// ObjectsRequiredBodyParams represents params for objectsRequiredBody operation
//
// Request: POST /objects/required-body.
type ObjectsRequiredBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObject
}

// ObjectsRequiredNestedObjectsParams represents params for objectsRequiredNestedObjects operation
//
// Request: POST /objects/required-nested-objects.
type ObjectsRequiredNestedObjectsParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObjectsContainer
}
