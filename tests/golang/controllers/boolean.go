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
	builder handlers.NoResponseHandlerBuilder[*handlers.BooleanBooleanParsingRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanParsing.action,
	)
}

func (c *booleanController) BooleanRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.BooleanBooleanRequiredValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanRequiredValidation.action,
	)
}

func (c *booleanController) BooleanNullable(
	builder handlers.NoResponseHandlerBuilder[*handlers.BooleanBooleanNullableRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanNullable.action,
	)
}

func (c *booleanController) BooleanArrayItems(
	builder handlers.NoResponseHandlerBuilder[*handlers.BooleanBooleanArrayItemsRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanArrayItems.action,
	)
}

func (c *booleanController) BooleanNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*handlers.BooleanBooleanNullableArrayItemsRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.nullableBooleanArrayItems.action,
	)
}

var _ handlers.BooleanController = &booleanController{}
