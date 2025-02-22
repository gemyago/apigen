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
	Id string `json:"id"`
	EndsWithId string `json:"endsWithId"`
	TheIdInTheMiddle string `json:"theIdInTheMiddle"`
}
