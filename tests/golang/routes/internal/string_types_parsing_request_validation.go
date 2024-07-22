package internal

import (
	"github.com/gemyago/apigen/tests/golang/routes/models"
)
func NewStringTypesParsingRequestValidator() FieldValidator[*models.StringTypesParsingRequest] {
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesParsingRequest) {
	}
}
