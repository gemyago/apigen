// Code generated by apigen DO NOT EDIT.

package models

import (
	"time"
)

// Unused imports workaround.
var _ = time.Time{}

type Pet struct { 
	Id int64 `json:"id"`
	Name string `json:"name"`
	Comments string `json:"comments,omitempty"`
}
