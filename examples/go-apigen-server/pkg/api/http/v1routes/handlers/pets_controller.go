package handlers

import (
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

type PetsCreatePetRequest struct {
	Payload models.Pet
}

type PetsGetPetByIdRequest struct {
	PetId int64
}

type PetsListPetsRequest struct {
	Limit int64
	Offset int64
}

type PetsController struct {
	// POST /pets
	//
	// Request type: PetsCreatePetRequest,
	//
	// Response type: none
	CreatePet httpHandlerFactory

	// GET /pets/{petId}
	//
	// Request type: PetsGetPetByIdRequest,
	//
	// Response type: models.PetResponse
	GetPetById httpHandlerFactory

	// GET /pets
	//
	// Request type: PetsListPetsRequest,
	//
	// Response type: models.PetsResponse
	ListPets httpHandlerFactory
}

type PetsControllerBuilder struct {
	// POST /pets
	//
	// Request type: PetsCreatePetRequest,
	//
	// Response type: none
	HandleCreatePet actionBuilderVoidResult[*PetsControllerBuilder, *PetsCreatePetRequest]

	// GET /pets/{petId}
	//
	// Request type: PetsGetPetByIdRequest,
	//
	// Response type: models.PetResponse
	HandleGetPetById actionBuilder[*PetsControllerBuilder, *PetsGetPetByIdRequest, *models.PetResponse]

	// GET /pets
	//
	// Request type: PetsListPetsRequest,
	//
	// Response type: models.PetsResponse
	HandleListPets actionBuilder[*PetsControllerBuilder, *PetsListPetsRequest, *models.PetsResponse]
}

func (c *PetsControllerBuilder) Finalize() *PetsController {
	// TODO: panic if any handler is null
	return &PetsController{
		CreatePet: c.HandleCreatePet.httpHandlerFactory,
		GetPetById: c.HandleGetPetById.httpHandlerFactory,
		ListPets: c.HandleListPets.httpHandlerFactory,
	}
}

func BuildPetsController() *PetsControllerBuilder {
	controllerBuilder := &PetsControllerBuilder{}

	// POST /pets
	controllerBuilder.HandleCreatePet.controllerBuilder = controllerBuilder
	controllerBuilder.HandleCreatePet.defaultStatusCode = 201
	controllerBuilder.HandleCreatePet.voidResult = true
	controllerBuilder.HandleCreatePet.paramsParser = newPetsCreatePetParamsParser()

	// GET /pets/{petId}
	controllerBuilder.HandleGetPetById.controllerBuilder = controllerBuilder
	controllerBuilder.HandleGetPetById.defaultStatusCode = 200
	controllerBuilder.HandleGetPetById.paramsParser = newPetsGetPetByIdParamsParser()

	// GET /pets
	controllerBuilder.HandleListPets.controllerBuilder = controllerBuilder
	controllerBuilder.HandleListPets.defaultStatusCode = 200
	controllerBuilder.HandleListPets.paramsParser = newPetsListPetsParamsParser()

	return controllerBuilder
}

func MountPetsRoutes(controller *PetsController, r httpRouter) {
	r.HandleRoute("POST", "/pets", controller.CreatePet(r))
	r.HandleRoute("GET", "/pets/{petId}", controller.GetPetById(r))
	r.HandleRoute("GET", "/pets", controller.ListPets(r))
}
