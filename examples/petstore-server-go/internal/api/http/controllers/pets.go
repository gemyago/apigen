package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/routes/handlers"
	"github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/routes/models"
)

var (
	ErrNotFound = errors.New("NOT_FOUND")
	ErrConflict = errors.New("CONFLICT")
)

type petsController struct {
	// In practice (depending on the project structure) controllers should
	// interact with a application layer and have application layer components as dependencies.
	// For this example, we are keeping it simple and storing data in memory.
	allPets  []*models.Pet
	petsByID map[int64]*models.Pet
}

func (c *petsController) CreatePet(
	b handlers.NoResponseHandlerBuilder[*models.CreatePetParams],
) http.Handler {
	return b.HandleWith(func(_ context.Context, params *models.CreatePetParams) error {
		if _, ok := c.petsByID[params.Payload.ID]; ok {
			return fmt.Errorf("pet %d already exists: %w", params.Payload.ID, ErrConflict)
		}

		c.allPets = append(c.allPets, params.Payload)
		c.petsByID[params.Payload.ID] = params.Payload

		return nil
	})
}

func (c *petsController) GetPetByID(
	b handlers.HandlerBuilder[*models.GetPetByIDParams, *models.PetResponse],
) http.Handler {
	return b.HandleWith(func(_ context.Context, params *models.GetPetByIDParams) (*models.PetResponse, error) {
		pet, ok := c.petsByID[params.PetID]
		if !ok {
			return nil, fmt.Errorf("pet %d not found: %w", params.PetID, ErrNotFound)
		}
		return &models.PetResponse{Data: pet}, nil
	})
}

func (c *petsController) ListPets(
	b handlers.HandlerBuilder[*models.ListPetsParams, *models.PetsResponse],
) http.Handler {
	return b.HandleWith(func(_ context.Context, params *models.ListPetsParams) (*models.PetsResponse, error) {
		allPetsLen := int64(len(c.allPets))
		limit := params.Limit
		offset := params.Offset
		if offset >= allPetsLen {
			return &models.PetsResponse{Data: []*models.Pet{}}, nil
		}
		if offset+limit > allPetsLen {
			limit = allPetsLen - offset
		}
		result := c.allPets[offset:limit]

		return &models.PetsResponse{Data: result}, nil
	})
}

// Make sure it implements the interface properly.
var _ handlers.PetsController = (*petsController)(nil)

func newPetsController() *petsController {
	return &petsController{
		allPets:  []*models.Pet{},
		petsByID: map[int64]*models.Pet{},
	}
}
