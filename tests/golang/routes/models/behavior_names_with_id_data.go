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

type BehaviorNamesWithIdData struct { 
	ID string `json:"id"`
	EndsWithID string `json:"endsWithId"`
	TheIDInTheMiddle string `json:"theIdInTheMiddle"`
}
