package internal

import (
	"github.com/gemyago/apigen/tests/golang/routes/models"
)
func NewStringTypesRangeValidationRequestValidator() FieldValidator[*models.StringTypesRangeValidationRequest] {
	return func(bindingCtx *BindingContext, field, location string, value *models.StringTypesRangeValidationRequest) {
	}
}
