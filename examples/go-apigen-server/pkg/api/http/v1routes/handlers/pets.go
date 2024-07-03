package handlers

import (
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

type PetsListPetsRequest struct {
	Limit  int32
	Offset int32
}

type PetsCreatePetRequest struct {
	Payload *models.Pet
}

type PetsGetPetById struct {
	PetId string
}

type PetsControllerBuilder struct {
	HandleListPets   actionBuilder[*PetsListPetsRequest, *models.PetsResponse]
	HandleCreatePet  actionBuilderVoidResult[*PetsCreatePetRequest]
	HandleGetPetById actionBuilder[*PetsGetPetById, *models.PetResponse]
}

func (c *PetsControllerBuilder) Finalize() PetsController {
	return nil
}

func BuildPetsController() *PetsControllerBuilder {
	return nil
}

type PetsController interface {
	ListPets(w http.ResponseWriter, r *http.Request)
	CreatePet(w http.ResponseWriter, r *http.Request)
	GetPetById(w http.ResponseWriter, r *http.Request)
}

func MountPetsRoutes(controller PetsController, r router) {
	r.HandleFunc("GET", "/pets", controller.ListPets)
	r.HandleFunc("POST", "/pets", controller.CreatePet)
	r.HandleFunc("GET", "/pets/{petId}", controller.GetPetById)
}

/*
# operations block
classname: Pets

# loop over each operation in the API:

# each operation has an `operationId`:
operationId: CreatePets

# and parameters:
pet: Pet


# each operation has an `operationId`:
operationId: GetPetById

# and parameters:
petId: string


# each operation has an `operationId`:
operationId: ListPets

# and parameters:
limit: int32
offset: int32


# end of operations block
*/
