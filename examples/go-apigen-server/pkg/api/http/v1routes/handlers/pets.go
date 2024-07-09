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

type PetsCreatePetParamsParser struct {
	bindPayload requestParamBinder[*http.Request, models.Pet]
}

func (p *PetsCreatePetParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsCreatePetRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsCreatePetRequest{}
	
	p.bindPayload(&bindingCtx, optionalVal[*http.Request]{value: req, assigned: true}, &reqParams.Payload)
	return reqParams, bindingCtx.Error()
}

func newPetsCreatePetParamsParser() *PetsCreatePetParamsParser {
	return &PetsCreatePetParamsParser{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, models.Pet]{
			field: "payload",
			location: "body",
			parseValue: parseJsonPayload[models.Pet],
			validateValue: newCompositeValidator[*http.Request, models.Pet](
				validateNonEmpty,
			),
		}),
	}
}
type PetsGetPetByIdParamsParser struct {
	bindPetId requestParamBinder[string, int64]
}

func (p *PetsGetPetByIdParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsGetPetByIdRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsGetPetByIdRequest{}
	
	p.bindPetId(&bindingCtx, readPathValue("petId", router, req), &reqParams.PetId)
	return reqParams, bindingCtx.Error()
}

func newPetsGetPetByIdParamsParser() *PetsGetPetByIdParamsParser {
	return &PetsGetPetByIdParamsParser{
		bindPetId: newRequestParamBinder(binderParams[string, int64]{
			field: "petId",
			location: "path",
			parseValue: knownParsers.int64_in_path,
			validateValue: newCompositeValidator[string, int64](
				validateNonEmpty,
			),
		}),
	}
}
type PetsListPetsParamsParser struct {
	bindLimit requestParamBinder[[]string, int64]
	bindOffset requestParamBinder[[]string, int64]
}

func (p *PetsListPetsParamsParser) parse(router httpRouter, w http.ResponseWriter, req *http.Request) (*PetsListPetsRequest, error) {
	bindingCtx := bindingContext{}
	reqParams := &PetsListPetsRequest{}
	query := req.URL.Query()
	p.bindLimit(&bindingCtx, readQueryValue("limit", query), &reqParams.Limit)
	p.bindOffset(&bindingCtx, readQueryValue("offset", query), &reqParams.Offset)
	return reqParams, bindingCtx.Error()
}

func newPetsListPetsParamsParser() *PetsListPetsParamsParser {
	return &PetsListPetsParamsParser{
		bindLimit: newRequestParamBinder(binderParams[[]string, int64]{
			field: "limit",
			location: "query",
			parseValue: knownParsers.int64_in_query,
			validateValue: newCompositeValidator[[]string, int64](
				validateNonEmpty,
				newMinMaxValueValidator[[]string, int64](1, false, true),
				newMinMaxValueValidator[[]string, int64](100, false, false),
			),
		}),
		bindOffset: newRequestParamBinder(binderParams[[]string, int64]{
			field: "offset",
			location: "query",
			parseValue: knownParsers.int64_in_query,
			validateValue: newCompositeValidator[[]string, int64](
				newMinMaxValueValidator[[]string, int64](1, false, true),
			),
		}),
	}
}

func BuildPetsController() *PetsControllerBuilder {
	controllerBuilder := &PetsControllerBuilder{}
	controllerBuilder.HandleCreatePet.controllerBuilder = controllerBuilder
	controllerBuilder.HandleCreatePet.defaultStatusCode = 201
	controllerBuilder.HandleCreatePet.voidResult = true
	controllerBuilder.HandleCreatePet.paramsParser = newPetsCreatePetParamsParser()
	controllerBuilder.HandleGetPetById.controllerBuilder = controllerBuilder
	controllerBuilder.HandleGetPetById.defaultStatusCode = 200
	controllerBuilder.HandleGetPetById.paramsParser = newPetsGetPetByIdParamsParser()
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
