// Code generated by apigen DO NOT EDIT.

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

// CreatePetParams - Parameters for the createPet operation.
type CreatePetParams struct { 
	Payload *Pet `json:"payload"`
}
