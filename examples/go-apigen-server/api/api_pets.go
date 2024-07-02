
# This is a sample api mustache template.  It is representing a fictitious
# language and won't be usable or compile to anything without lots of changes.
# Use it as an example.  You can access the variables in the generator object
# like such:

# use the package from the `apiPackage` variable
package: api

# operations block
classname: PetsAPI

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
