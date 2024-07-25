package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type SimpleObject struct { 
	SimpleField1 string `json:"simpleField1"`
}
