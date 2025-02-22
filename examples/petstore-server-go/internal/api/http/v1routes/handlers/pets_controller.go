package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/v1routes/models"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint

type _ func() Error

// PetsCreatePetRequest represents params for createPet operation
//
// Request: POST /pets.
type PetsCreatePetRequest struct {
	// Payload is parsed from request body and declared as payload.
	Payload *Pet
}

// PetsGetPetByIdRequest represents params for getPetById operation
//
// Request: GET /pets/{petId}.
type PetsGetPetByIdRequest struct {
	// PetId is parsed from request path and declared as petId.
	PetId int64
}

// PetsListPetsRequest represents params for listPets operation
//
// Request: GET /pets.
type PetsListPetsRequest struct {
	// Limit is parsed from request query and declared as limit.
	Limit int64
	// Offset is parsed from request query and declared as offset.
	Offset int64
}

type petsControllerBuilder struct {
	// POST /pets
	//
	// Request type: PetsCreatePetRequest,
	//
	// Response type: none
	CreatePet genericHandlerBuilder[
		*PetsCreatePetRequest,
		void,
		func(context.Context, *PetsCreatePetRequest) error,
		func(http.ResponseWriter, *http.Request, *PetsCreatePetRequest) error,
	]

	// GET /pets/{petId}
	//
	// Request type: PetsGetPetByIdRequest,
	//
	// Response type: PetResponse
	GetPetById genericHandlerBuilder[
		*PetsGetPetByIdRequest,
		*PetResponse,
		func(context.Context, *PetsGetPetByIdRequest) (*PetResponse, error),
		func(http.ResponseWriter, *http.Request, *PetsGetPetByIdRequest) (*PetResponse, error),
	]

	// GET /pets
	//
	// Request type: PetsListPetsRequest,
	//
	// Response type: PetsResponse
	ListPets genericHandlerBuilder[
		*PetsListPetsRequest,
		*PetsResponse,
		func(context.Context, *PetsListPetsRequest) (*PetsResponse, error),
		func(http.ResponseWriter, *http.Request, *PetsListPetsRequest) (*PetsResponse, error),
	]
}

func newPetsControllerBuilder(app *RootHandler) *petsControllerBuilder {
	return &petsControllerBuilder{
		// POST /pets
		CreatePet: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*PetsCreatePetRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*PetsCreatePetRequest,
				void,
			](),
			makeActionBuilderParams[
				*PetsCreatePetRequest,
				void,
			]{
				defaultStatus: 201,
				voidResult:    true,
				paramsParser:  newParamsParserPetsCreatePet(app),
			},
		),

		// GET /pets/{petId}
		GetPetById: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*PetsGetPetByIdRequest,
				*PetResponse,
			](),
			newHTTPHandlerAdapter[
				*PetsGetPetByIdRequest,
				*PetResponse,
			](),
			makeActionBuilderParams[
				*PetsGetPetByIdRequest,
				*PetResponse,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserPetsGetPetById(app),
			},
		),

		// GET /pets
		ListPets: newGenericHandlerBuilder(
			app,
			newHandlerAdapter[
				*PetsListPetsRequest,
				*PetsResponse,
			](),
			newHTTPHandlerAdapter[
				*PetsListPetsRequest,
				*PetsResponse,
			](),
			makeActionBuilderParams[
				*PetsListPetsRequest,
				*PetsResponse,
			]{
				defaultStatus: 200,
				paramsParser:  newParamsParserPetsListPets(app),
			},
		),
	}
}

type PetsController interface {
	// POST /pets
	//
	// Request type: PetsCreatePetRequest,
	//
	// Response type: none
	CreatePet(NoResponseHandlerBuilder[*PetsCreatePetRequest]) http.Handler

	// GET /pets/{petId}
	//
	// Request type: PetsGetPetByIdRequest,
	//
	// Response type: PetResponse
	GetPetById(HandlerBuilder[
		*PetsGetPetByIdRequest,
		*PetResponse,
	]) http.Handler

	// GET /pets
	//
	// Request type: PetsListPetsRequest,
	//
	// Response type: PetsResponse
	ListPets(HandlerBuilder[
		*PetsListPetsRequest,
		*PetsResponse,
	]) http.Handler
}

func RegisterPetsRoutes(rootHandler *RootHandler, controller PetsController) {
	builder := newPetsControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/pets", controller.CreatePet(builder.CreatePet))
	rootHandler.router.HandleRoute("GET", "/pets/{petId}", controller.GetPetById(builder.GetPetById))
	rootHandler.router.HandleRoute("GET", "/pets", controller.ListPets(builder.ListPets))
}
