package handlers

import (
	"encoding/json"
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
		ListPets:   c.HandleListPets.httpHandlerFactory,
		CreatePet:  c.HandleCreatePet.httpHandlerFactory,
		GetPetById: c.HandleGetPetById.httpHandlerFactory,
	}
}

func BuildPetsController() *PetsControllerBuilder {
	controllerBuilder := &PetsControllerBuilder{}
	controllerBuilder.HandleListPets = actionBuilder[*PetsControllerBuilder, *PetsListPetsRequest, *models.PetsResponse]{
		controllerBuilder: controllerBuilder,
		defaultStatusCode: 200,
		paramsParser: func(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsListPetsRequest, error) {
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
			defaultStatusCode: 201,
			paramsParser: func(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsCreatePetRequest, error) {
				var payload models.Pet
				if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
					return nil, fmt.Errorf("TODO: handle parsing errors (field=payload): %w", err)
				}
				return &PetsCreatePetRequest{Payload: &payload}, nil
			},
		},
	}
	controllerBuilder.HandleGetPetById = actionBuilder[*PetsControllerBuilder, *PetsGetPetById, *models.PetResponse]{
		controllerBuilder: controllerBuilder,
		defaultStatusCode: 200,
		paramsParser: func(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsGetPetById, error) {
			var err error
			var petId int64
			petId, err = strconv.ParseInt(router.PathValue(req, "petId"), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("TODO: handle parsing errors (field=petId): %w", err)
			}

			return &PetsGetPetById{
				PetId: petId,
			}, nil
		},
	}
	return controllerBuilder
}

type PetsController struct {
	ListPets   httpHandlerFactory
	CreatePet  httpHandlerFactory
	GetPetById httpHandlerFactory
}

func MountPetsRoutes(controller *PetsController, r httpRouter) {
	r.Handle("GET", "/pets", controller.ListPets(r))
	r.Handle("POST", "/pets", controller.CreatePet(r))
	r.Handle("GET", "/pets/{petId}", controller.GetPetById(r))
}
