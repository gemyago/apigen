// Code generated by apigen DO NOT EDIT.

package internal

import (
	"time"
	. "github.com/gemyago/apigen/examples/ping-server-go/internal/api/http/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewPing200ResponseValidator() FieldValidator[*Ping200Response] {
	validateMessage := NewSimpleFieldValidator[string](
	)
	
	return func(bindingCtx *BindingContext, value *Ping200Response) {
		validateMessage(bindingCtx.Fork("message"), value.Message)
	}
}
