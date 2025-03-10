// Code generated by apigen DO NOT EDIT.

package internal

import (
	"time"
	. "github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewPetsResponseValidator() FieldValidator[*PetsResponse] {
	validateData := NewArrayValidator[*Pet](
		NewSimpleFieldValidator[[]*Pet](
			EnsureArrayFieldRequired,
		),
		NewPetValidator(),
	)
	
	return func(bindingCtx *BindingContext, value *PetsResponse) {
		validateData(bindingCtx.Fork("data"), value.Data)
	}
}
