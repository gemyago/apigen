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

type ObjectArraysSimpleObject struct { 
	SimpleField1 string `json:"simpleField1"`
}
