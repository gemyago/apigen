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

// ListPetsParams - Parameters for the listPets operation
type ListPetsParams struct { 
	Limit int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}
