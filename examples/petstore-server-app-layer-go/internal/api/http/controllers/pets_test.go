package controllers

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app/models"
	"github.com/stretchr/testify/assert"
)

type mockPetsService struct {
	createPetCalls []*models.CreatePetParams
	nextGetPetByID *models.PetResponse
	nextListPets   *models.PetsResponse
}

func (m *mockPetsService) CreatePet(_ context.Context, params *models.CreatePetParams) error {
	m.createPetCalls = append(m.createPetCalls, params)
	return nil
}

func (m *mockPetsService) GetPetByID(_ context.Context, _ *models.GetPetByIDParams) (*models.PetResponse, error) {
	return m.nextGetPetByID, nil
}

func (m *mockPetsService) ListPets(_ context.Context, _ *models.ListPetsParams) (*models.PetsResponse, error) {
	return m.nextListPets, nil
}

func TestTestPets(t *testing.T) {
	discardLogger := slog.New(slog.NewTextHandler(io.Discard, nil))

	newDeps := func() RoutesDeps {
		return RoutesDeps{
			RootLogger:  discardLogger,
			PetsService: &mockPetsService{},
		}
	}

	t.Run("POST /pets", func(t *testing.T) {
		t.Run("process create pet request", func(t *testing.T) {
			deps := newDeps()
			handler := SetupRoutes(deps)
			req := httptest.NewRequest(http.MethodPost, "/pets", bytes.NewBufferString(`{"id":1,"name":"Bingo"}`))
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, req)
			assert.Equal(t, 201, res.Code)

			service, _ := deps.PetsService.(*mockPetsService)
			assert.Len(t, service.createPetCalls, 1)
			assert.Equal(t,
				&models.CreatePetParams{Payload: &models.Pet{ID: 1, Name: "Bingo"}},
				service.createPetCalls[0],
			)
		})
	})
}
