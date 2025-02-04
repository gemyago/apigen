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

type ActionHandlerFunc[TRq any, TRs any] interface {
	func(context.Context, *TRq) (*TRs, error) | // with params with response
		func(context.Context) (*TRs, error) | // no params with response
		func(context.Context, *TRq) error | // with params no response
		func(context.Context) error // no params no response
}

type ActionBuilder[TRq any, TRs any, Tb ActionHandlerFunc[TRq, TRs]] struct {
	HandleWith func(Tb) http.Handler
}

type PetsControllerActionsBuilder struct {
	NotImplemented func() http.Handler

	// GET /ping
	Ping ActionBuilder[
		Void,
		Void,
		func(context.Context) error,
	]
	GetPet ActionBuilder[
		GetPetRequest,
		GetPetResponse,
		func(context.Context, *GetPetRequest) (*GetPetResponse, error),
	]
	GetPets ActionBuilder[
		Void,
		GetPetsResponse,
		func(context.Context) (*GetPetsResponse, error),
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
