package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsNestedRequest struct { 
	SimpleRequiredField1 string `json:"simpleRequiredField1,omitempty"`
	NestedObject1 ObjectsNestedRequestNestedObject1 `json:"nestedObject1"`
	NestedObject2 ObjectsNestedRequestNestedObject2 `json:"nestedObject2"`
}
