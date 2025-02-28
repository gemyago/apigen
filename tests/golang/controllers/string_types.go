package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type stringTypesControllerTestActions struct {
	stringTypesParsing                    mockAction[*models.StringTypesParsingParams]
	stringTypesArraysParsing              mockAction[*models.StringTypesArraysParsingParams]
	stringTypesNullableParsing            mockAction[*models.StringTypesNullableParsingParams]
	stringTypesNullableArrayItems         mockAction[*models.StringTypesNullableArrayItemsParams]
	stringTypesRangeValidation            mockAction[*models.StringTypesRangeValidationParams]
	stringTypesArrayItemsRangeValidation  mockAction[*models.StringTypesArrayItemsRangeValidationParams]
	stringTypesPatternValidation          mockAction[*models.StringTypesPatternValidationParams]
	stringTypesRequiredValidation         mockAction[*models.StringTypesRequiredValidationParams]
	stringTypesNullableRequiredValidation mockAction[*models.StringTypesNullableRequiredValidationParams]
	stringTypesEnums                      mockAction[*models.StringTypesEnumsParams]
}

type stringTypesController struct {
	testActions *stringTypesControllerTestActions
}

func (c *stringTypesController) StringTypesArrayItemsRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesArrayItemsRangeValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.stringTypesArrayItemsRangeValidation.action,
	)
}

func (c *stringTypesController) StringTypesArraysParsing(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesArraysParsingParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesArraysParsing.action)
}

func (c *stringTypesController) StringTypesEnums(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesEnumsParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesEnums.action)
}

func (c *stringTypesController) StringTypesNullableArrayItems(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesNullableArrayItemsParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesNullableArrayItems.action)
}

func (c *stringTypesController) StringTypesNullableParsing(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesNullableParsingParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesNullableParsing.action)
}

func (c *stringTypesController) StringTypesNullableRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesNullableRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(
		c.testActions.stringTypesNullableRequiredValidation.action,
	)
}

func (c *stringTypesController) StringTypesParsing(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesParsingParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesParsing.action)
}

func (c *stringTypesController) StringTypesPatternValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesPatternValidationParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesPatternValidation.action)
}

func (c *stringTypesController) StringTypesRangeValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesRangeValidationParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesRangeValidation.action)
}

func (c *stringTypesController) StringTypesRequiredValidation(
	builder handlers.NoResponseHandlerBuilder[*models.StringTypesRequiredValidationParams],
) http.Handler {
	return builder.HandleWith(c.testActions.stringTypesRequiredValidation.action)
}

var _ handlers.StringTypesController = &stringTypesController{}
