package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/api/http/routes/models"
)

var (
	ErrNotFound = errors.New("NOT_FOUND")
	ErrConflict = errors.New("CONFLICT")
)

type PetsService struct {
	logger   *slog.Logger
	allPets  []*models.Pet
	petsByID map[int64]*models.Pet
}

func (c *PetsService) CreatePet(
	ctx context.Context, params *models.CreatePetParams,
) error {
	c.logger.InfoContext(ctx, "Creating pet", slog.Int64("petID", params.Payload.ID))
	if _, ok := c.petsByID[params.Payload.ID]; ok {
		return fmt.Errorf("pet %d already exists: %w", params.Payload.ID, ErrConflict)
	}

	c.allPets = append(c.allPets, params.Payload)
	c.petsByID[params.Payload.ID] = params.Payload

	return nil
}

func (c *PetsService) GetPetByID(
	ctx context.Context, params *models.GetPetByIDParams,
) (*models.PetResponse, error) {
	pet, ok := c.petsByID[params.PetID]
	c.logger.DebugContext(ctx, "Pet found", slog.Int64("petID", params.PetID))
	if !ok {
		return nil, fmt.Errorf("pet %d not found: %w", params.PetID, ErrNotFound)
	}
	return &models.PetResponse{Data: pet}, nil
}

func (c *PetsService) ListPets(
	ctx context.Context, params *models.ListPetsParams,
) (*models.PetsResponse, error) {
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

	c.logger.DebugContext(ctx, "Found pets", slog.Int("count", len(result)))

	return &models.PetsResponse{Data: result}, nil
}

type PetsServiceDeps struct {
	RootLogger *slog.Logger
}

func NewPetsService(deps PetsServiceDeps) *PetsService {
	return &PetsService{
		logger:   deps.RootLogger.WithGroup("app.pets-service"),
		allPets:  []*models.Pet{},
		petsByID: map[int64]*models.Pet{},
	}
}
