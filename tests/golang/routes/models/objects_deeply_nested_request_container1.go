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

type ObjectsDeeplyNestedRequestContainer1 struct { 
	Container11 *SimpleObjectsContainer `json:"container11"`
	Container12 *SimpleObjectsContainer `json:"container12"`
}
