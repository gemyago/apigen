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

type ObjectArraysSimpleObjectsContainer struct { 
	SimpleObjects1 []*ObjectArraysSimpleObject `json:"simpleObjects1"`
	SimpleObjects2 []*ObjectArraysSimpleObject `json:"simpleObjects2,omitempty"`
}
