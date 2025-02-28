package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type stringTypesControllerTestActions struct {
	stringTypesParsing                    mockAction[*models.StringTypesStringTypesParsingParams]
	stringTypesArraysParsing              mockAction[*models.StringTypesStringTypesArraysParsingParams]
	stringTypesNullableParsing            mockAction[*models.StringTypesStringTypesNullableParsingParams]
	stringTypesNullableArrayItems         mockAction[*models.StringTypesStringTypesNullableArrayItemsParams]
	stringTypesRangeValidation            mockAction[*models.StringTypesStringTypesRangeValidationParams]
	stringTypesArrayItemsRangeValidation  mockAction[*models.StringTypesStringTypesArrayItemsRangeValidationParams]
	stringTypesPatternValidation          mockAction[*models.StringTypesStringTypesPatternValidationParams]
	stringTypesRequiredValidation         mockAction[*models.StringTypesStringTypesRequiredValidationParams]
	stringTypesNullableRequiredValidation mockAction[*models.StringTypesStringTypesNullableRequiredValidationParams]
	stringTypesEnums                      mockAction[*models.StringTypesStringTypesEnumsParams]
}

type stringTypesController struct {
	testActions *stringTypesControllerTestActions
}

func (c *stringTypesController) StringTypesArrayItemsRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesArrayItemsRangeValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.stringTypesArrayItemsRangeValidation.action,
	)
}

func (c *stringTypesController) StringTypesArraysParsing(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesArraysParsingParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesArraysParsing.action)
}

func (c *stringTypesController) StringTypesEnums(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesEnumsParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesEnums.action)
}

func (c *stringTypesController) StringTypesNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesNullableArrayItemsParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesNullableArrayItems.action)
}

func (c *stringTypesController) StringTypesNullableParsing(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesNullableParsingParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesNullableParsing.action)
}

func (c *stringTypesController) StringTypesNullableRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesNullableRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.stringTypesNullableRequiredValidation.action,
	)
}

func (c *stringTypesController) StringTypesParsing(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesParsingParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesParsing.action)
}

func (c *stringTypesController) StringTypesPatternValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesPatternValidationParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesPatternValidation.action)
}

func (c *stringTypesController) StringTypesRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesRangeValidationParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesRangeValidation.action)
}

func (c *stringTypesController) StringTypesRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesStringTypesRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesRequiredValidation.action)
}

var _ handlers.StringTypesController = &stringTypesController{}
