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
	Storage *storage
}

type commandsImpl struct {
	CommandsDeps
}

func (c *commandsImpl) CreatePet(ctx context.Context, req *handlers.PetsCreatePetRequest) error {
	if _, ok := c.Storage.petsById[req.Payload.Id]; ok {
		return fmt.Errorf("pet %d already exists: %w", req.Payload.Id, ErrConflict)
	}

	c.Storage.allPets = append(c.Storage.allPets, req.Payload)
	c.Storage.petsById[req.Payload.Id] = req.Payload

	return nil
}

func NewCommands(deps CommandsDeps) Commands {
	return &commandsImpl{CommandsDeps: deps}
}
