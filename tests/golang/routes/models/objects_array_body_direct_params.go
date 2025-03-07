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

// ObjectsArrayBodyDirectParams - Parameters for the objectsArrayBodyDirect operation.
type ObjectsArrayBodyDirectParams struct { 
	Payload []*ObjectArraysSimpleObject `json:"payload"`
}
