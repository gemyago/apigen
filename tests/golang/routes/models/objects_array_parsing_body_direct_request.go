package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsArrayParsingBodyDirectRequest struct { 
	NestedArray1 []ObjectArraysSimpleObject `json:"nestedArray1"`
	NestedArray2 []ObjectArraysSimpleObject `json:"nestedArray2"`
	NestedArrayContainer1 []ObjectArraysSimpleObjectsContainer `json:"nestedArrayContainer1"`
	NestedArrayContainer2 []ObjectArraysSimpleObjectsContainer `json:"nestedArrayContainer2"`
}
