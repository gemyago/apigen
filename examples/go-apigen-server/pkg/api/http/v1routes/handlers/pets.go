package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

type PetsListPetsRequest struct {
	Limit  int64
	Offset int64
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
		paramsParser: func(w http.ResponseWriter, req *http.Request) (*PetsListPetsRequest, error) {
			query := req.URL.Query()
			var err error
			var limit int64
			limit, err = strconv.ParseInt(query.Get("limit"), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("TODO: handle parsing errors (field=limit): %w", err)
			}

			var offset int64
			if query.Has("offset") {
				offset, err = strconv.ParseInt(query.Get("offset"), 10, 64)
				if err != nil {
					return nil, fmt.Errorf("TODO: handle parsing errors (field=offset): %w", err)
				}
			}

			return &PetsListPetsRequest{
				Limit:  limit,
				Offset: offset,
			}, nil
		},
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
