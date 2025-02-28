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

// CreatePetParams represents params for createPet operation
//
// Request: POST /pets.
type CreatePetParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *Pet
}

// GetPetByIDParams represents params for getPetById operation
//
// Request: GET /pets/{petId}.
type GetPetByIDParams struct {
	// PetID is parsed from request path and declared as petId.
	PetID int64
}

// ListPetsParams represents params for listPets operation
//
// Request: GET /pets.
type ListPetsParams struct {
	// Limit is parsed from request query and declared as limit.
	Limit int64
	// Offset is parsed from request query and declared as offset.
	Offset int64
}
