package app

import (
	"context"
	"fmt"

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
	allPetsLen := int64(len(q.Storage.allPets))
	limit := req.Limit
	offset := req.Offset
	if offset >= allPetsLen {
		return &models.PetsResponse{}, nil
	}
	if offset+limit > allPetsLen {
		limit = allPetsLen - offset
	}
	result := q.Storage.allPets[offset:limit]
	return &models.PetsResponse{Data: result}, nil
}

func (q *queriesImpl) GetPetById(ctx context.Context, req *handlers.PetsGetPetById) (*models.PetResponse, error) {
	pet, ok := q.Storage.petsById[req.PetId]
	if !ok {
		return nil, fmt.Errorf("pet %d not found: %w", req.PetId, ErrNotFound)
	}
	return &models.PetResponse{Data: pet}, nil
}

func NewQueries(deps QueriesDeps) Queries {
	return &queriesImpl{QueriesDeps: deps}
}
