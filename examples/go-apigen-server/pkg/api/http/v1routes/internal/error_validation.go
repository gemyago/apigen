// Code generated by apigen DO NOT EDIT.

package internal

import (
	"time"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewErrorValidator() FieldValidator[*models.Error] {
	validateCode := NewSimpleFieldValidator[interface{}](
	)
	validateMessage := NewSimpleFieldValidator[string](
	)
	return func(bindingCtx *BindingContext, field, location string, value *models.Error) {
		validateCode(bindingCtx, "code", location, value.Code)
		validateMessage(bindingCtx, "message", location, value.Message)
	}
}
