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

type BehaviorNamesWithId200Response struct { 
	Id string `json:"id,omitempty"`
	EndsWithId string `json:"endsWithId,omitempty"`
	TheIdInTheMiddle string `json:"theIdInTheMiddle,omitempty"`
}
