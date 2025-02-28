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

// ObjectsObjectsArrayBodyDirectParams represents params for objectsArrayBodyDirect operation
//
// Request: POST /objects/arrays.
type ObjectsObjectsArrayBodyDirectParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload []*ObjectArraysSimpleObject
}

// ObjectsObjectsArrayBodyNestedParams represents params for objectsArrayBodyNested operation
//
// Request: PUT /objects/arrays.
type ObjectsObjectsArrayBodyNestedParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *ObjectsArrayBodyNestedRequest
}

// ObjectsObjectsDeeplyNestedParams represents params for objectsDeeplyNested operation
//
// Request: POST /objects/deeply-nested.
type ObjectsObjectsDeeplyNestedParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *ObjectsDeeplyNestedRequest
}

// ObjectsObjectsNullableOptionalBodyParams represents params for objectsNullableOptionalBody operation
//
// Request: PUT /objects/nullable-body.
type ObjectsObjectsNullableOptionalBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleNullableObject
}

// ObjectsObjectsNullableRequiredBodyParams represents params for objectsNullableRequiredBody operation
//
// Request: POST /objects/nullable-body.
type ObjectsObjectsNullableRequiredBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleNullableObject
}

// ObjectsObjectsOptionalBodyParams represents params for objectsOptionalBody operation
//
// Request: PUT /objects/required-body.
type ObjectsObjectsOptionalBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObject
}

// ObjectsObjectsRequiredBodyParams represents params for objectsRequiredBody operation
//
// Request: POST /objects/required-body.
type ObjectsObjectsRequiredBodyParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObject
}

// ObjectsObjectsRequiredNestedObjectsParams represents params for objectsRequiredNestedObjects operation
//
// Request: POST /objects/required-nested-objects.
type ObjectsObjectsRequiredNestedObjectsParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *SimpleObjectsContainer
}
