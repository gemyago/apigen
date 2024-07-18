package v1controllers

import (
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/app"
)

type PetsControllerDeps struct {
	app.Queries
	app.Commands
}

func NewPetsController(deps PetsControllerDeps) *handlers.PetsController {
	return handlers.BuildPetsController().
		HandleListPets.With(deps.Queries.ListPets).
		HandleCreatePet.With(deps.Commands.CreatePet).
		HandleGetPetById.With(deps.Queries.GetPetByID).
		Finalize()
}
