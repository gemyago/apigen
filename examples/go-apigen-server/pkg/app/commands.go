package app

import (
	"context"
	"fmt"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
)

type Commands interface {
	CreatePet(context.Context, *handlers.PetsCreatePetRequest) error
}

type CommandsDeps struct {
	Storage *Storage
}

type commandsImpl struct {
	CommandsDeps
}

func (c *commandsImpl) CreatePet(_ context.Context, req *handlers.PetsCreatePetRequest) error {
	if _, ok := c.Storage.petsByID[req.Payload.Id]; ok {
		return fmt.Errorf("pet %d already exists: %w", req.Payload.Id, ErrConflict)
	}

	c.Storage.allPets = append(c.Storage.allPets, *req.Payload)
	c.Storage.petsByID[req.Payload.Id] = *req.Payload

	return nil
}

func NewCommands(deps CommandsDeps) Commands {
	return &commandsImpl{CommandsDeps: deps}
}
