package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsNestedRequestNestedObject1 struct { 
	SimpleRequiredField1 string `json:"simpleRequiredField1"`
	NestedObject11 *ObjectsNestedRequestNestedObject1NestedObject11 `json:"nestedObject11"`
	NestedObject12 *ObjectsNestedRequestNestedObject1NestedObject11 `json:"nestedObject12"`
}
