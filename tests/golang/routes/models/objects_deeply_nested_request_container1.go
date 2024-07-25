package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsDeeplyNestedRequestContainer1 struct { 
	Container11 *SimpleObjectsContainer `json:"container11"`
	Container12 *SimpleObjectsContainer `json:"container12"`
}
