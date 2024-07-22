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
	)
	validateName := NewSimpleFieldValidator[string](
	)
	validateComments := NewSimpleFieldValidator[string](
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.Pet) {
		validateId(bindingCtx, "id", location, value.Id)
		validateName(bindingCtx, "name", location, value.Name)
		validateComments(bindingCtx, "comments", location, value.Comments)
	}
}