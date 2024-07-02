package handlers

import "net/http"

type PetsListPetsRequest struct {
	Limit  int32
	Offset int32
}

type PetsController struct{}

func MountPetsRoutes(controller PetsController, r router) {
	r.HandleFunc("GET", "/pets", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})

	r.HandleFunc("POST", "/pets", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})

	r.HandleFunc("GET", "/pets/{petId}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})
}

/*
# operations block
classname: Pets

# loop over each operation in the API:

# each operation has an `operationId`:
operationId: CreatePets

# and parameters:
pet: Pet


# each operation has an `operationId`:
operationId: GetPetById

# and parameters:
petId: string


# each operation has an `operationId`:
operationId: ListPets

# and parameters:
limit: int32
offset: int32


# end of operations block
*/
