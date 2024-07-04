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
	HandleListPets   actionBuilder[*PetsControllerBuilder, *PetsListPetsRequest, *models.PetsResponse]
	HandleCreatePet  actionBuilderVoidResult[*PetsControllerBuilder, *PetsCreatePetRequest]
	HandleGetPetById actionBuilder[*PetsControllerBuilder, *PetsGetPetById, *models.PetResponse]
}

func (c *PetsControllerBuilder) Finalize() *PetsController {
	return nil
}

func BuildPetsController() *PetsControllerBuilder {
	return nil
}

type PetsController struct {
	ListPets   http.Handler
	CreatePet  http.Handler
	GetPetById http.Handler
}

func MountPetsRoutes(controller *PetsController, r router) {
	r.Handle("GET", "/pets", controller.ListPets)
	r.Handle("POST", "/pets", controller.CreatePet)
	r.Handle("GET", "/pets/{petId}", controller.GetPetById)
}
