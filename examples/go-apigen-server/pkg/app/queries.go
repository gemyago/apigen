package app

import (
	"context"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

type Queries interface {
	ListPets(context.Context, *handlers.PetsListPetsRequest) (*models.PetsResponse, error)
	GetPetById(context.Context, *handlers.PetsGetPetById) (*models.PetResponse, error)
}

type QueriesDeps struct {
	Storage *storage
}

type queriesImpl struct {
	QueriesDeps
}

func (q *queriesImpl) ListPets(ctx context.Context, req *handlers.PetsListPetsRequest) (*models.PetsResponse, error) {
	panic("not implemented")
}

func (q *queriesImpl) GetPetById(ctx context.Context, req *handlers.PetsGetPetById) (*models.PetResponse, error) {
	panic("not implemented")
}

func NewQueries(deps QueriesDeps) Queries {
	return &queriesImpl{QueriesDeps: deps}
}
