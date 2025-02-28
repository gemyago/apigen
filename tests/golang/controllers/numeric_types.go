package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type numericTypesControllerTestActions struct {
	numericTypesParsing                  mockAction[*models.NumericTypesNumericTypesParsingParams]
	numericTypesArrayItems               mockAction[*models.NumericTypesNumericTypesArrayItemsParams]
	numericTypesRangeValidation          mockAction[*models.NumericTypesNumericTypesRangeValidationParams]
	numericTypesRangeValidationExclusive mockAction[*models.NumericTypesNumericTypesRangeValidationExclusiveParams]
	numericTypesRequiredValidation       mockAction[*models.NumericTypesNumericTypesRequiredValidationParams]
	numericTypesNullable                 mockAction[*models.NumericTypesNumericTypesNullableParams]
	numericTypesNullableArrayItems       mockAction[*models.NumericTypesNumericTypesNullableArrayItemsParams]
}

type numericTypesController struct {
	testActions *numericTypesControllerTestActions
}

func (c *numericTypesController) NumericTypesParsing(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNumericTypesParsingParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesParsing.action,
	)
}

func (c *numericTypesController) NumericTypesArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNumericTypesArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesArrayItems.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNumericTypesRangeValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRangeValidation.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidationExclusive(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNumericTypesRangeValidationExclusiveParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRangeValidationExclusive.action,
	)
}

func (c *numericTypesController) NumericTypesRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNumericTypesRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRequiredValidation.action,
	)
}

func (c *numericTypesController) NumericTypesNullable(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNumericTypesNullableParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesNullable.action,
	)
}

func (c *numericTypesController) NumericTypesNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.NumericTypesNumericTypesNullableArrayItemsParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesNullableArrayItems.action,
	)
}

var _ handlers.NumericTypesController = &numericTypesController{}
