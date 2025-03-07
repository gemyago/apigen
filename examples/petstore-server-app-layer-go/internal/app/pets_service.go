package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app/models"
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

func (svc *PetsService) CreatePet(
	ctx context.Context, params *models.CreatePetParams,
) error {
	svc.logger.InfoContext(ctx, "Creating pet", slog.Int64("petID", params.Payload.ID))
	if _, ok := svc.petsByID[params.Payload.ID]; ok {
		return fmt.Errorf("pet %d already exists: %w", params.Payload.ID, ErrConflict)
	}

	svc.allPets = append(svc.allPets, params.Payload)
	svc.petsByID[params.Payload.ID] = params.Payload

	return nil
}

func (svc *PetsService) GetPetByID(
	ctx context.Context, params *models.GetPetByIDParams,
) (*models.PetResponse, error) {
	pet, ok := svc.petsByID[params.PetID]
	svc.logger.DebugContext(ctx, "Pet found", slog.Int64("petID", params.PetID))
	if !ok {
		return nil, fmt.Errorf("pet %d not found: %w", params.PetID, ErrNotFound)
	}
	return &models.PetResponse{Data: pet}, nil
}

func (svc *PetsService) ListPets(
	ctx context.Context, params *models.ListPetsParams,
) (*models.PetsResponse, error) {
	allPetsLen := int64(len(svc.allPets))
	limit := params.Limit
	offset := params.Offset
	if offset >= allPetsLen {
		return &models.PetsResponse{Data: []*models.Pet{}}, nil
	}
	if offset+limit > allPetsLen {
		limit = allPetsLen - offset
	}
	result := svc.allPets[offset:limit]

	svc.logger.DebugContext(ctx, "Found pets", slog.Int("count", len(result)))

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
