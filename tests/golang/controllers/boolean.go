package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type booleanControllerTestActions struct {
	booleanParsing            mockAction[*models.BooleanParsingParams]
	booleanRequiredValidation mockAction[*models.BooleanRequiredValidationParams]
	booleanNullable           mockAction[*models.BooleanNullableParams]
	booleanArrayItems         mockAction[*models.BooleanArrayItemsParams]
	nullableBooleanArrayItems mockAction[*models.BooleanNullableArrayItemsParams]
}

type booleanController struct {
	testActions *booleanControllerTestActions
}

func (c *booleanController) BooleanParsing(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanParsingParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanParsing.action,
	)
}

func (c *booleanController) BooleanRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanRequiredValidation.action,
	)
}

func (c *booleanController) BooleanNullable(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanNullableParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanNullable.action,
	)
}

func (c *booleanController) BooleanArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanArrayItems.action,
	)
}

func (c *booleanController) BooleanNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanNullableArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.nullableBooleanArrayItems.action,
	)
}

var _ handlers.BooleanController = &booleanController{}
