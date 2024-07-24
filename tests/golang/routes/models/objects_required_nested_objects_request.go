package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsRequiredNestedObjectsRequest struct { 
	SimpleObject1 SimpleObject `json:"simpleObject1"`
	SimpleObject2 SimpleObject `json:"simpleObject2"`
}
