package app

import (
	"context"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
)

type Queries interface {
	ListPets(context.Context, *handlers.PetsListPetsRequest) error
	GetPetById(context.Context, *handlers.PetsGetPetById) error
}

type QueriesDeps struct {
	Storage *storage
}

type queriesImpl struct {
	QueriesDeps
}

func (q *queriesImpl) ListPets(ctx context.Context, req *handlers.PetsListPetsRequest) error {
	panic("not implemented")
}

func (q *queriesImpl) GetPetById(ctx context.Context, req *handlers.PetsGetPetById) error {
	panic("not implemented")
}

func NewQueries(deps QueriesDeps) Queries {
	return &queriesImpl{QueriesDeps: deps}
}
