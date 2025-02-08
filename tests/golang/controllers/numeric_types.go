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
	builder *handlers.NumericTypesControllerBuilder,
) http.Handler {
	return builder.NumericTypesParsing.HandleWith(
		c.testActions.numericTypesParsing.action,
	)
}

func (c *numericTypesController) NumericTypesArrayItems(
	builder *handlers.NumericTypesControllerBuilder,
) http.Handler {
	return builder.NumericTypesArrayItems.HandleWith(
		c.testActions.numericTypesArrayItems.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidation(
	builder *handlers.NumericTypesControllerBuilder,
) http.Handler {
	return builder.NumericTypesRangeValidation.HandleWith(
		c.testActions.numericTypesRangeValidation.action,
	)
}

func (c *numericTypesController) NumericTypesRangeValidationExclusive(
	builder *handlers.NumericTypesControllerBuilder,
) http.Handler {
	return builder.NumericTypesRangeValidationExclusive.HandleWith(
		c.testActions.numericTypesRangeValidationExclusive.action,
	)
}

func (c *numericTypesController) NumericTypesRequiredValidation(
	builder *handlers.NumericTypesControllerBuilder,
) http.Handler {
	return builder.NumericTypesRequiredValidation.HandleWith(
		c.testActions.numericTypesRequiredValidation.action,
	)
}

func (c *numericTypesController) NumericTypesNullable(
	builder *handlers.NumericTypesControllerBuilder,
) http.Handler {
	return builder.NumericTypesNullable.HandleWith(
		c.testActions.numericTypesNullable.action,
	)
}

func (c *numericTypesController) NumericTypesNullableArrayItems(
	builder *handlers.NumericTypesControllerBuilder,
) http.Handler {
	return builder.NumericTypesNullableArrayItems.HandleWith(
		c.testActions.numericTypesNullableArrayItems.action,
	)
}

var _ handlers.NumericTypesController = &numericTypesController{}
