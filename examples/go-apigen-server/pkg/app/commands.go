package app

import (
	"context"

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
	panic("not implemented")
}

func NewCommands(deps CommandsDeps) Commands {
	return &commandsImpl{CommandsDeps: deps}
}
