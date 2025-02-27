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

// PetsCreatePetParams represents params for createPet operation
//
// Request: POST /pets.
type PetsCreatePetParams struct {
	// Payload is parsed from request body and declared as payload.
	Payload *Pet
}

// PetsGetPetByIDParams represents params for getPetById operation
//
// Request: GET /pets/{petId}.
type PetsGetPetByIDParams struct {
	// PetID is parsed from request path and declared as petId.
	PetID int64
}

// PetsListPetsParams represents params for listPets operation
//
// Request: GET /pets.
type PetsListPetsParams struct {
	// Limit is parsed from request query and declared as limit.
	Limit int64
	// Offset is parsed from request query and declared as offset.
	Offset int64
}
