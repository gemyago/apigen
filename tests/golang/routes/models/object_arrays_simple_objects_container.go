package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectArraysSimpleObjectsContainer struct {
	SimpleObjects1 []*ObjectArraysSimpleObject `json:"simpleObjects1"`
	SimpleObjects2 []*ObjectArraysSimpleObject `json:"simpleObjects2,omitempty"`
}
