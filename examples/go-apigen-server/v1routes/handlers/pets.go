package handlers

type PetsListPetsRequest struct {
	Limit  int32
	Offset int32
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
