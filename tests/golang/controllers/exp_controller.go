package controllers

// TODO: This is stub and should be removed

import (
	"context"
	"net/http"
)

type Pet struct {
	ID   int64
	Name string
}

type GetPetRequest struct {
	PetID int64
}

type GetPetResponse struct {
	Data Pet
}

type GetPetsResponse struct {
	Data []Pet
}

type HTTPApp struct{}

type Void string

const VoidValue Void = "void"

// ActionHandlerFunc represents possible combination of action handler functions.
// Each function can be with or without parameters and with or without response.
// Additionally each function can have access to http objects for possible direct manipulation.
type ActionHandlerFunc[TReq any, TRes any] interface {
	func(context.Context, *TReq) (*TRes, error) | // with params with response
		func(context.Context) (*TRes, error) | // no params with response
		func(context.Context, *TReq) error | // with params no response
		func(context.Context) error | // no params no response

		// handlers with http context exposed
		func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) | // with params with response
		func(http.ResponseWriter, *http.Request) (*TRes, error) | // no params with response
		func(http.ResponseWriter, *http.Request, *TReq) error | // with params no response
		func(http.ResponseWriter, *http.Request) error // no params no response
}

type ActionBuilderFunc[TReq any, TRes any, THandler ActionHandlerFunc[TReq, TRes]] func(THandler) http.Handler

type ActionBuilder[
	TReq any,
	TRes any,
	TPlainHandler ActionHandlerFunc[TReq, TRes],
	THttpHandler ActionHandlerFunc[TReq, TRes],
] struct {
	HandleWith     func(TPlainHandler) http.Handler
	HandleWithHTTP func(THttpHandler) http.Handler
}

type PetsControllerActionsBuilder struct {
	NotImplemented func() http.Handler

	// GET /ping
	Ping ActionBuilder[
		Void,
		Void,
		func(context.Context) error,
		func(http.ResponseWriter, *http.Request) error,
	]
	GetPet ActionBuilder[
		GetPetRequest,
		GetPetResponse,
		func(context.Context, *GetPetRequest) (*GetPetResponse, error),
		func(http.ResponseWriter, *http.Request, *GetPetRequest) (*GetPetResponse, error),
	]
	GetPets ActionBuilder[
		Void,
		GetPetsResponse,
		func(context.Context) (*GetPetsResponse, error),
		func(http.ResponseWriter, *http.Request) (*GetPetsResponse, error),
	]
}

type PetsController interface {
	// GetPet creates a handler for the GetPet endpoint.
	// Action builders:
	// - BuildPetsGetPetAction
	Ping(b *PetsControllerActionsBuilder) http.Handler
	GetPet(b *PetsControllerActionsBuilder) http.Handler
	GetPets(b *PetsControllerActionsBuilder) http.Handler
}

func RegisterRoutes(_ PetsController) {
	panic("not implemented")
}
