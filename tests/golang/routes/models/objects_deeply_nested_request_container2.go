package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Unused imports workaround.
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint

type ObjectsDeeplyNestedRequestContainer2 struct { 
	Container21 *SimpleObjectsContainer `json:"container21"`
	Container22 *SimpleObjectsContainer `json:"container22"`
}
