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

type ObjectsDeeplyNestedRequest struct { 
	Container1 *ObjectsDeeplyNestedRequestContainer1 `json:"container1"`
	Container2 *ObjectsDeeplyNestedRequestContainer2 `json:"container2"`
}
