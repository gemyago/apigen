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
			petData := bytes.NewBufferString(`{"id":1,"name":"Bingo"}`)
			req := httptest.NewRequest(http.MethodPost, "/pets", petData)
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

	t.Run("GET /pets/{id}", func(t *testing.T) {
		t.Run("process get pet by id request", func(t *testing.T) {
			deps := newDeps()
			handler := SetupRoutes(deps)

			wantNextPet := &models.PetResponse{Data: &models.Pet{ID: 1, Name: "Bingo"}}

			service, _ := deps.PetsService.(*mockPetsService)
			service.nextGetPetByID = wantNextPet

			req := httptest.NewRequest(http.MethodGet, "/pets/1", nil)
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, req)
			assert.Equal(t, 200, res.Code)

			assert.JSONEq(t, `{"data":{"id":1,"name":"Bingo"}}`, res.Body.String())
		})
	})

	t.Run("GET /pets", func(t *testing.T) {
		t.Run("process list pets request", func(t *testing.T) {
			deps := newDeps()
			handler := SetupRoutes(deps)

			wantNextPets := &models.PetsResponse{
				Data: []*models.Pet{
					{ID: 1, Name: "Bingo"},
					{ID: 2, Name: "Bongo"},
				},
			}

			service, _ := deps.PetsService.(*mockPetsService)
			service.nextListPets = wantNextPets

			req := httptest.NewRequest(http.MethodGet, "/pets?limit=10", nil)
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, req)
			assert.Equal(t, 200, res.Code)

			assert.JSONEq(t, `{"data":[{"id":1,"name":"Bingo"},{"id":2,"name":"Bongo"}]}`, res.Body.String())
		})
	})
}
