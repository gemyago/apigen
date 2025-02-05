package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type booleanControllerTestActions struct {
	booleanParsing            mockAction[*handlers.BooleanBooleanParsingRequest]
	booleanRequiredValidation mockAction[*handlers.BooleanBooleanRequiredValidationRequest]
	booleanNullable           mockAction[*handlers.BooleanBooleanNullableRequest]
	booleanArrayItems         mockAction[*handlers.BooleanBooleanArrayItemsRequest]
	nullableBooleanArrayItems mockAction[*handlers.BooleanBooleanNullableArrayItemsRequest]
}

type booleanController struct {
	testActions *booleanControllerTestActions
}

func (c *booleanController) BooleanParsing(
	builder *handlers.BooleanControllerBuilderV2,
) http.Handler {
	return builder.BooleanParsing.HandleWith(
		c.testActions.booleanParsing.action,
	)
}

func (c *booleanController) BooleanRequiredValidation(
	builder *handlers.BooleanControllerBuilderV2,
) http.Handler {
	return builder.BooleanRequiredValidation.HandleWith(
		c.testActions.booleanRequiredValidation.action,
	)
}

func (c *booleanController) BooleanNullable(
	builder *handlers.BooleanControllerBuilderV2,
) http.Handler {
	return builder.BooleanNullable.HandleWith(
		c.testActions.booleanNullable.action,
	)
}

func (c *booleanController) BooleanArrayItems(
	builder *handlers.BooleanControllerBuilderV2,
) http.Handler {
	return builder.BooleanArrayItems.HandleWith(
		c.testActions.booleanArrayItems.action,
	)
}

func (c *booleanController) BooleanNullableArrayItems(
	builder *handlers.BooleanControllerBuilderV2,
) http.Handler {
	return builder.BooleanNullableArrayItems.HandleWith(
		c.testActions.nullableBooleanArrayItems.action,
	)
}

var _ handlers.BooleanControllerV2 = &booleanController{}
