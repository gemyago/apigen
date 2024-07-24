package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectNullableSimpleObjectsContainer struct { 
	SimpleObject1 *ObjectNullableSimpleObject `json:"simpleObject1"`
	OptionalSimpleObject2 *ObjectNullableSimpleObject `json:"optionalSimpleObject2,omitempty"`
}
