package handlers

import (
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

type PetsCreatePetRequest struct {
	Payload models.Pet
}
type PetsGetPetByIdRequest struct {
	PetId int64
}
type PetsListPetsRequest struct {
	Limit  int64
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
		CreatePet:  c.HandleCreatePet.httpHandlerFactory,
		GetPetById: c.HandleGetPetById.httpHandlerFactory,
		ListPets:   c.HandleListPets.httpHandlerFactory,
	}
}

type petsCreatePetParamsParser struct {
	bindPayload requestParamBinder[*http.Request, models.Pet]
}

func (p *petsCreatePetParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsCreatePetRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsCreatePetRequest{}
	p.bindPayload(&bindingCtx, optionalVal[*http.Request]{value: req}, &reqParams.Payload)
	return reqParams, nil
}

type petsGetPetByIdParamsParser struct {
	bindPetId requestParamBinder[string, int64]
}

func (p *petsGetPetByIdParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsGetPetByIdRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsGetPetByIdRequest{}
	p.bindPetId(&bindingCtx, readPathValue("petId", router, req), &reqParams.PetId)
	return reqParams, nil
}

type petsListPetsParamsParser struct {
	bindLimit  requestParamBinder[[]string, int64]
	bindOffset requestParamBinder[[]string, int64]
}

func (p *petsListPetsParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsListPetsRequest, error) {
	query := req.URL.Query()
	bindingCtx := bindingContext{}
	reqParams := &PetsListPetsRequest{}
	p.bindLimit(&bindingCtx, readQueryValue("limit", query), &reqParams.Limit)
	p.bindLimit(&bindingCtx, readQueryValue("offset", query), &reqParams.Offset)
	return reqParams, nil
}

func BuildPetsController() *PetsControllerBuilder {
	controllerBuilder := &PetsControllerBuilder{}

	controllerBuilder.HandleCreatePet.controllerBuilder = controllerBuilder
	controllerBuilder.HandleCreatePet.voidResult = true
	controllerBuilder.HandleCreatePet.defaultStatusCode = 201
	controllerBuilder.HandleCreatePet.paramsParser = &petsCreatePetParamsParser{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, models.Pet]{
			field:      "payload",
			location:   "body",
			parseValue: parseJsonPayload[models.Pet],
		}),
	}

	controllerBuilder.HandleGetPetById.controllerBuilder = controllerBuilder
	controllerBuilder.HandleGetPetById.defaultStatusCode = 200
	controllerBuilder.HandleGetPetById.paramsParser = &petsGetPetByIdParamsParser{
		bindPetId: newRequestParamBinder(binderParams[string, int64]{
			field:      "petId",
			location:   "path",
			parseValue: newStringToSignedIntParser[int64](64),
		}),
	}

	controllerBuilder.HandleListPets.controllerBuilder = controllerBuilder
	controllerBuilder.HandleListPets.defaultStatusCode = 200
	controllerBuilder.HandleListPets.paramsParser = &petsListPetsParamsParser{
		bindLimit: newRequestParamBinder(binderParams[[]string, int64]{
			field:      "limit",
			location:   "query",
			parseValue: newStringSliceToSignedIntParser[int64](64),
		}),
		bindOffset: newRequestParamBinder(binderParams[[]string, int64]{
			field:      "offset",
			location:   "query",
			parseValue: newStringSliceToSignedIntParser[int64](64),
		}),
	}

	return controllerBuilder
}

func MountPetsRoutes(controller *PetsController, r httpRouter) {
	r.HandleRoute("POST", "/pets", controller.CreatePet(r))
	r.HandleRoute("GET", "/pets/{petId}", controller.GetPetById(r))
	r.HandleRoute("GET", "/pets", controller.ListPets(r))
}
