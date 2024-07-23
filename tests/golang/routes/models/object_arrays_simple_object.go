package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectArraysSimpleObject struct { 
	SimpleField1 string `json:"simpleField1"`
}
