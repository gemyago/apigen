package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type booleanControllerTestActions struct {
	booleanParsing            mockAction[*models.BooleanBooleanParsingParams]
	booleanRequiredValidation mockAction[*models.BooleanBooleanRequiredValidationParams]
	booleanNullable           mockAction[*models.BooleanBooleanNullableParams]
	booleanArrayItems         mockAction[*models.BooleanBooleanArrayItemsParams]
	nullableBooleanArrayItems mockAction[*models.BooleanBooleanNullableArrayItemsParams]
}

type booleanController struct {
	testActions *booleanControllerTestActions
}

func (c *booleanController) BooleanParsing(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanBooleanParsingParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanParsing.action,
	)
}

func (c *booleanController) BooleanRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanBooleanRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanRequiredValidation.action,
	)
}

func (c *booleanController) BooleanNullable(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanBooleanNullableParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanNullable.action,
	)
}

func (c *booleanController) BooleanArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanBooleanArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.booleanArrayItems.action,
	)
}

func (c *booleanController) BooleanNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.BooleanBooleanNullableArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.nullableBooleanArrayItems.action,
	)
}

var _ handlers.BooleanController = &booleanController{}
