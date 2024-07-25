package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsDeeplyNestedRequestContainer2 struct { 
	Container21 *SimpleObjectsContainer `json:"container21"`
	Container22 *SimpleObjectsContainer `json:"container22"`
}
