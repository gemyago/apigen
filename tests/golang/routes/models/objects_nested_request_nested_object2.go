package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsNestedRequestNestedObject2 struct { 
	SimpleRequiredField1 string `json:"simpleRequiredField1"`
	NestedObject21 ObjectsNestedRequestNestedObject1NestedObject11 `json:"nestedObject21"`
	NestedObject22 ObjectsNestedRequestNestedObject1NestedObject11 `json:"nestedObject22"`
}
