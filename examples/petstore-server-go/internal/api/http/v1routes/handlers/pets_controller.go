// Code generated by apigen DO NOT EDIT.

package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/v1routes/models"
)

type PetsController interface {
	// POST /pets
	//
	// Request type: PetsCreatePetParams,
	//
	// Response type: none
	CreatePet(NoResponseHandlerBuilder[
		*PetsCreatePetParams,
	]) http.Handler

	// GET /pets/{petId}
	//
	// Request type: PetsGetPetByIDParams,
	//
	// Response type: PetResponse
	GetPetByID(HandlerBuilder[
		*PetsGetPetByIDParams,
		*PetResponse,
	]) http.Handler

	// GET /pets
	//
	// Request type: PetsListPetsParams,
	//
	// Response type: PetsResponse
	ListPets(HandlerBuilder[
		*PetsListPetsParams,
		*PetsResponse,
	]) http.Handler
}

// RegisterPetsRoutes will attach the following routes to the root handler:
// 
// - POST /pets
// 
// - GET /pets/{petId}
// 
// - GET /pets
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterPetsRoutes(controller PetsController) *RootHandler {
	builder := newPetsControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/pets", controller.CreatePet(builder.CreatePet))
	rootHandler.router.HandleRoute("GET", "/pets/{petId}", controller.GetPetByID(builder.GetPetByID))
	rootHandler.router.HandleRoute("GET", "/pets", controller.ListPets(builder.ListPets))
	return rootHandler
}