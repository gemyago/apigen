package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsOptionalNestedObjectsRequest struct { 
	SimpleObject1 SimpleObject `json:"simpleObject1,omitempty"`
	SimpleObject2 SimpleObject `json:"simpleObject2,omitempty"`
}
