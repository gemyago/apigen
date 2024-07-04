// Code generated DO NOT EDIT

package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Unused imports workaround
var _ = json.Marshal
var _ = errors.New
var _ = fmt.Errorf
var _ = reflect.TypeOf
var _ = time.Time{}

type PetResponse struct { 
	Data Pet `json:"data"`
}

