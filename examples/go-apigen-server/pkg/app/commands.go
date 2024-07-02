package app

import (
	"context"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
)

type Commands interface {
	CreatePet(context.Context, *handlers.PetsCreatePetRequest) error
}
