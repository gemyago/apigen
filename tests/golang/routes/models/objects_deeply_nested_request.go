package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectsDeeplyNestedRequest struct { 
	Container1 *ObjectsDeeplyNestedRequestContainer1 `json:"container1"`
	Container2 *ObjectsDeeplyNestedRequestContainer2 `json:"container2"`
}
