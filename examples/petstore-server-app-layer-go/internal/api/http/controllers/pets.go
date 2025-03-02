package controllers

import (
	"context"
	"net/http"

	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/api/http/routes/handlers"
	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app/models"
)

// You may use the app.PetsService struct directly in the controller, or you may abstract it away
// with an interface which may be useful for testing if app layer is more complex and requires
// mocking.
type petsService interface {
	CreatePet(ctx context.Context, params *models.CreatePetParams) error
	GetPetByID(ctx context.Context, params *models.GetPetByIDParams) (*models.PetResponse, error)
	ListPets(ctx context.Context, params *models.ListPetsParams) (*models.PetsResponse, error)
}

type PetsController struct {
	petsService
}

func (c *PetsController) CreatePet(
	b handlers.NoResponseHandlerBuilder[*models.CreatePetParams],
) http.Handler {
	return b.HandleWith(c.petsService.CreatePet)
}

func (c *PetsController) GetPetByID(
	b handlers.HandlerBuilder[*models.GetPetByIDParams, *models.PetResponse],
) http.Handler {
	return b.HandleWith(c.petsService.GetPetByID)
}

func (c *PetsController) ListPets(
	b handlers.HandlerBuilder[*models.ListPetsParams, *models.PetsResponse],
) http.Handler {
	return b.HandleWith(c.petsService.ListPets)
}

// Make sure it implements the interface properly.
var _ handlers.PetsController = (*PetsController)(nil)

type PetsControllerDeps struct {
	PetsService petsService
}

func NewPetsController(deps PetsControllerDeps) *PetsController {
	return &PetsController{
		petsService: deps.PetsService,
	}
}
