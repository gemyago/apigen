package app

import "github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"

// In practice this should be on a service layer (e.g service package) and
// should write things to some DB like Postgres or Mongo. But keeping it simple for
// this example and storing just in memory.
type storage struct {
	allPets  []*models.Pet
	petsById map[string]*models.Pet
}

func NewStorage() *storage {
	return &storage{
		allPets:  []*models.Pet{},
		petsById: map[string]*models.Pet{},
	}
}