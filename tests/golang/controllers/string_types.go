package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type stringTypesControllerTestActions struct {
	stringTypesParsing                    mockAction[*handlers.StringTypesStringTypesParsingRequest]
	stringTypesArraysParsing              mockAction[*handlers.StringTypesStringTypesArraysParsingRequest]
	stringTypesNullableParsing            mockAction[*handlers.StringTypesStringTypesNullableParsingRequest]
	stringTypesNullableArrayItems         mockAction[*handlers.StringTypesStringTypesNullableArrayItemsRequest]
	stringTypesRangeValidation            mockAction[*handlers.StringTypesStringTypesRangeValidationRequest]
	stringTypesArrayItemsRangeValidation  mockAction[*handlers.StringTypesStringTypesArrayItemsRangeValidationRequest]
	stringTypesPatternValidation          mockAction[*handlers.StringTypesStringTypesPatternValidationRequest]
	stringTypesRequiredValidation         mockAction[*handlers.StringTypesStringTypesRequiredValidationRequest]
	stringTypesNullableRequiredValidation mockAction[*handlers.StringTypesStringTypesNullableRequiredValidationRequest]
	stringTypesEnums                      mockAction[*handlers.StringTypesStringTypesEnumsRequest]
}

type stringTypesController struct {
	testActions *stringTypesControllerTestActions
}

func (c *stringTypesController) StringTypesArrayItemsRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesArrayItemsRangeValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.stringTypesArrayItemsRangeValidation.action,
	)
}

func (c *stringTypesController) StringTypesArraysParsing(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesArraysParsingRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesArraysParsing.action)
}

func (c *stringTypesController) StringTypesEnums(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesEnumsRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesEnums.action)
}

func (c *stringTypesController) StringTypesNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesNullableArrayItemsRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesNullableArrayItems.action)
}

func (c *stringTypesController) StringTypesNullableParsing(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesNullableParsingRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesNullableParsing.action)
}

func (c *stringTypesController) StringTypesNullableRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesNullableRequiredValidationRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.stringTypesNullableRequiredValidation.action,
	)
}

func (c *stringTypesController) StringTypesParsing(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesParsingRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesParsing.action)
}

func (c *stringTypesController) StringTypesPatternValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesPatternValidationRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesPatternValidation.action)
}

func (c *stringTypesController) StringTypesRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesRangeValidationRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesRangeValidation.action)
}

func (c *stringTypesController) StringTypesRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*handlers.StringTypesStringTypesRequiredValidationRequest],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesRequiredValidation.action)
}

var _ handlers.StringTypesControllerV3 = &stringTypesController{}
