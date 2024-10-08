// Code generated by apigen DO NOT EDIT.

package internal

import (
	"time"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewPetValidator() FieldValidator[*models.Pet] {
	validateId := NewSimpleFieldValidator[int64](
		EnsureNonDefault[int64],
	)
	validateName := NewSimpleFieldValidator[string](
		EnsureNonDefault[string],
	)
	validateComments := NewSimpleFieldValidator[string](
	)
	
	return func(bindingCtx *BindingContext, value *models.Pet) {
		validateId(bindingCtx.Fork("id"), value.Id)
		validateName(bindingCtx.Fork("name"), value.Name)
		validateComments(bindingCtx.Fork("comments"), value.Comments)
	}
}
