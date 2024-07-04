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
	PetId int64
}

type PetsControllerBuilder struct {
	HandleListPets   actionBuilder[*PetsControllerBuilder, *PetsListPetsRequest, *models.PetsResponse]
	HandleCreatePet  actionBuilderVoidResult[*PetsControllerBuilder, *PetsCreatePetRequest]
	HandleGetPetById actionBuilder[*PetsControllerBuilder, *PetsGetPetById, *models.PetResponse]
}

func (c *PetsControllerBuilder) Finalize() *PetsController {
	// TODO: panic if any handler is null
	return &PetsController{
		ListPets:   c.HandleListPets.httpHandler,
		CreatePet:  c.HandleCreatePet.httpHandler,
		GetPetById: c.HandleGetPetById.httpHandler,
	}
}

func BuildPetsController() *PetsControllerBuilder {
	controllerBuilder := &PetsControllerBuilder{}
	controllerBuilder.HandleListPets = actionBuilder[*PetsControllerBuilder, *PetsListPetsRequest, *models.PetsResponse]{
		controllerBuilder: controllerBuilder,
	}
	controllerBuilder.HandleCreatePet = actionBuilderVoidResult[*PetsControllerBuilder, *PetsCreatePetRequest]{
		actionBuilder: actionBuilder[*PetsControllerBuilder, *PetsCreatePetRequest, voidResult]{
			controllerBuilder: controllerBuilder,
		},
	}
	controllerBuilder.HandleGetPetById = actionBuilder[*PetsControllerBuilder, *PetsGetPetById, *models.PetResponse]{
		controllerBuilder: controllerBuilder,
	}
	return controllerBuilder
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
