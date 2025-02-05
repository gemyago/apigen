package controllers

// TODO: This is stub and should be removed

import (
	"context"
	"net/http"
)

type petsQueriesGetPetParams struct {
	PetID int64
}

type petsQueriesGetPetResponse struct {
	Data Pet
}

type petsQueriesGetPetsResponse struct {
	Data []Pet
}

type petsQueries interface {
	Ping(ctx context.Context) error
	GetPetPassThrough(ctx context.Context, params *GetPetRequest) (*GetPetResponse, error)
	GetPetReqMatch(ctx context.Context) (*petsQueriesGetPetResponse, error)
	GetPetResMatch(ctx context.Context, params *petsQueriesGetPetParams) (*GetPetResponse, error)
	GetPet(ctx context.Context, params *petsQueriesGetPetParams) (*petsQueriesGetPetResponse, error)
	GetPets(ctx context.Context) (*petsQueriesGetPetsResponse, error)
	GetPetsPassThrough(ctx context.Context) (*GetPetsResponse, error)
}

type controllerImpl struct {
	queries petsQueries
}

func (impl *controllerImpl) Ping(b *PetsControllerActionsBuilder) http.Handler {
	return b.Ping.HandleWith(impl.queries.Ping)
}

func BuildActionWithTransformers[
	TDeclaredReq any,
	TDeclaredRes any,
	TDeclaredFn ActionHandlerFunc[TDeclaredReq, TDeclaredRes],
	TDeclaredHttpFn ActionHandlerFunc[TDeclaredReq, TDeclaredRes],
	TAppReq any,
	TAppRes any,
	TAppFn ActionHandlerFunc[TAppReq, TAppRes],
](
	actionBuilder ActionBuilder[TDeclaredReq, TDeclaredRes, TDeclaredFn, TDeclaredHttpFn],
	action TAppFn,
	inputTransformer func(req *http.Request, input *TDeclaredReq) (*TAppReq, error),
	outputTransformer func(input *TAppRes) (*TDeclaredRes, error),
) http.Handler {
	panic("not implemented")
}

func (impl *controllerImpl) GetPet(b *PetsControllerActionsBuilder) http.Handler {
	return /*handlers.*/ BuildActionWithTransformers(
		b.GetPet,
		impl.queries.GetPet,
		func(*http.Request, *GetPetRequest) (*petsQueriesGetPetParams, error) { panic("not implemented") },
		func(*petsQueriesGetPetResponse) (*GetPetResponse, error) { panic("not implemented") },
	)
}

func (impl *controllerImpl) GetPetPassThrough(b *PetsControllerActionsBuilder) http.Handler {
	return b.GetPet.HandleWith(func(ctx context.Context, gpr *GetPetRequest) (*GetPetResponse, error) {
		return impl.queries.GetPetPassThrough(ctx, gpr)
	})
}

func (impl *controllerImpl) GetPetHTTP(b *PetsControllerActionsBuilder) http.Handler {
	return b.GetPet.HandleWithHTTP(
		func(w http.ResponseWriter, r *http.Request, gpr *GetPetRequest) (*GetPetResponse, error) {
			panic("not implemented")
		},
	)
}

func (impl *controllerImpl) GetPets(b *PetsControllerActionsBuilder) http.Handler {
	return b.GetPets.HandleWith(impl.queries.GetPetsPassThrough)
}

var _ PetsController = &controllerImpl{}
