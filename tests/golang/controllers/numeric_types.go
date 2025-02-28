package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type numericTypesControllerTestActions struct {
	numericTypesParsing                  mockAction[*models.NumericTypesParsingParams]
	numericTypesArrayItems               mockAction[*models.NumericTypesArrayItemsParams]
	numericTypesRangeValidation          mockAction[*models.NumericTypesRangeValidationParams]
	numericTypesRangeValidationExclusive mockAction[*models.NumericTypesRangeValidationExclusiveParams]
	numericTypesRequiredValidation       mockAction[*models.NumericTypesRequiredValidationParams]
	numericTypesNullable                 mockAction[*models.NumericTypesNullableParams]
	numericTypesNullableArrayItems       mockAction[*models.NumericTypesNullableArrayItemsParams]
}

type numericTypesController struct {
	testActions *numericTypesControllerTestActions
}

func (c *numericTypesController) NumericTypesParsing(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesParsingParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesParsing.action,
	)
}

func (c *numericTypesController) NumericTypesArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesArrayItems.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesRangeValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRangeValidation.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidationExclusive(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesRangeValidationExclusiveParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRangeValidationExclusive.action,
	)
}

func (c *numericTypesController) NumericTypesRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRequiredValidation.action,
	)
}

func (c *numericTypesController) NumericTypesNullable(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNullableParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesNullable.action,
	)
}

func (c *numericTypesController) NumericTypesNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNullableArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesNullableArrayItems.action,
	)
}

var _ handlers.NumericTypesController = &numericTypesController{}
