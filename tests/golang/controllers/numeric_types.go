package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type numericTypesControllerTestActions struct {
	numericTypesParsing                  mockAction[*handlers.NumericTypesNumericTypesParsingRequest]
	numericTypesArrayItems               mockAction[*handlers.NumericTypesNumericTypesArrayItemsRequest]
	numericTypesRangeValidation          mockAction[*handlers.NumericTypesNumericTypesRangeValidationRequest]
	numericTypesRangeValidationExclusive mockAction[*handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest]
	numericTypesRequiredValidation       mockAction[*handlers.NumericTypesNumericTypesRequiredValidationRequest]
	numericTypesNullable                 mockAction[*handlers.NumericTypesNumericTypesNullableRequest]
	numericTypesNullableArrayItems       mockAction[*handlers.NumericTypesNumericTypesNullableArrayItemsRequest]
}

type numericTypesController struct {
	testActions *numericTypesControllerTestActions
}

func (c *numericTypesController) NumericTypesParsing(
	builder handlers.NoResponseHandlerBuilder[*handlers.NumericTypesNumericTypesParsingRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesParsing.action,
	)
}

func (c *numericTypesController) NumericTypesArrayItems(
	builder handlers.NoResponseHandlerBuilder[*handlers.NumericTypesNumericTypesArrayItemsRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesArrayItems.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.NumericTypesNumericTypesRangeValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRangeValidation.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidationExclusive(
	builder handlers.NoResponseHandlerBuilder[*handlers.NumericTypesNumericTypesRangeValidationExclusiveRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRangeValidationExclusive.action,
	)
}

func (c *numericTypesController) NumericTypesRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.NumericTypesNumericTypesRequiredValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesRequiredValidation.action,
	)
}

func (c *numericTypesController) NumericTypesNullable(
	builder handlers.NoResponseHandlerBuilder[*handlers.NumericTypesNumericTypesNullableRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesNullable.action,
	)
}

func (c *numericTypesController) NumericTypesNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*handlers.NumericTypesNumericTypesNullableArrayItemsRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.numericTypesNullableArrayItems.action,
	)
}

var _ handlers.NumericTypesControllerV3 = &numericTypesController{}
