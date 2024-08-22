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

type ObjectsArrayBodyNestedRequest struct { 
	NestedArray1 []*ObjectArraysSimpleObject `json:"nestedArray1"`
	NestedArray2 []*ObjectArraysSimpleObject `json:"nestedArray2"`
	NestedArrayContainer1 []*ObjectArraysSimpleObjectsContainer `json:"nestedArrayContainer1"`
	NestedArrayContainer2 []*ObjectArraysSimpleObjectsContainer `json:"nestedArrayContainer2"`
}
