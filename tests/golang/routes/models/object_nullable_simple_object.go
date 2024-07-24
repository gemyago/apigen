package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type ObjectNullableSimpleObject struct { 
	SimpleField1 string `json:"simpleField1"`
}
