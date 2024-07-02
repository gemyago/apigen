package app

import (
	"context"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
)

type Queries interface {
	ListPets(context.Context, *handlers.PetsListPetsRequest) error
	GetPetById(context.Context, *handlers.PetsGetPetById) error
}
