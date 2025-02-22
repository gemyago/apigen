package v1controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/v1routes/handlers"
	"github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/v1routes/models"
)

var (
	ErrNotFound = errors.New("NOT_FOUND")
	ErrConflict = errors.New("CONFLICT")
)

type PetsController struct {
	// In practice (depending on the project structure) controllers should
	// interact with a application layer and have application layer components as dependencies.
	// For this example, we are keeping it simple and storing data in memory.
	allPets  []*models.Pet
	petsByID map[int64]*models.Pet
}

func (c *PetsController) CreatePet(
	b handlers.NoResponseHandlerBuilder[*handlers.PetsCreatePetRequest],
) http.Handler {
	return b.HandleWith(func(_ context.Context, req *handlers.PetsCreatePetRequest) error {
		if _, ok := c.petsByID[req.Payload.ID]; ok {
			return fmt.Errorf("pet %d already exists: %w", req.Payload.ID, ErrConflict)
		}

		c.allPets = append(c.allPets, req.Payload)
		c.petsByID[req.Payload.ID] = req.Payload

		return nil
	})
}

func (c *PetsController) GetPetByID(
	b handlers.HandlerBuilder[*handlers.PetsGetPetByIDRequest, *models.PetResponse],
) http.Handler {
	return b.HandleWith(func(_ context.Context, req *handlers.PetsGetPetByIDRequest) (*models.PetResponse, error) {
		pet, ok := c.petsByID[req.PetID]
		if !ok {
			return nil, fmt.Errorf("pet %d not found: %w", req.PetID, ErrNotFound)
		}
		return &models.PetResponse{Data: pet}, nil
	})
}

func (c *PetsController) ListPets(
	b handlers.HandlerBuilder[*handlers.PetsListPetsRequest, *models.PetsResponse],
) http.Handler {
	return b.HandleWith(func(_ context.Context, req *handlers.PetsListPetsRequest) (*models.PetsResponse, error) {
		allPetsLen := int64(len(c.allPets))
		limit := req.Limit
		offset := req.Offset
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
var _ handlers.PetsController = (*PetsController)(nil)

func NewPetsController() *PetsController {
	return &PetsController{
		allPets:  []*models.Pet{},
		petsByID: map[int64]*models.Pet{},
	}
}
