package internal

import (
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

func NewObjectsRequiredNestedObjectsRequestValidator() FieldValidator[*models.ObjectsRequiredNestedObjectsRequest] {
	validateSimpleObject1 := NewSimpleObjectValidator()
	validateSimpleObject2 := NewSimpleObjectValidator()
	
	return func(bindingCtx *BindingContext, field, location string, value *models.ObjectsRequiredNestedObjectsRequest) {
		validateSimpleObject1(bindingCtx, "simpleObject1", location, &value.SimpleObject1)
		validateSimpleObject2(bindingCtx, "simpleObject2", location, &value.SimpleObject2)
	}
}
