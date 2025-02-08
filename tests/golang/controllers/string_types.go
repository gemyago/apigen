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
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesArrayItemsRangeValidation.HandleWith(
		c.testActions.stringTypesArrayItemsRangeValidation.action,
	)
}

func (c *stringTypesController) StringTypesArraysParsing(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesArraysParsing.HandleWith(c.testActions.stringTypesArraysParsing.action)
}

func (c *stringTypesController) StringTypesEnums(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesEnums.HandleWith(c.testActions.stringTypesEnums.action)
}

func (c *stringTypesController) StringTypesNullableArrayItems(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesNullableArrayItems.HandleWith(c.testActions.stringTypesNullableArrayItems.action)
}

func (c *stringTypesController) StringTypesNullableParsing(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesNullableParsing.HandleWith(c.testActions.stringTypesNullableParsing.action)
}

func (c *stringTypesController) StringTypesNullableRequiredValidation(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesNullableRequiredValidation.HandleWith(
		c.testActions.stringTypesNullableRequiredValidation.action,
	)
}

func (c *stringTypesController) StringTypesParsing(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesParsing.HandleWith(c.testActions.stringTypesParsing.action)
}

func (c *stringTypesController) StringTypesPatternValidation(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesPatternValidation.HandleWith(c.testActions.stringTypesPatternValidation.action)
}

func (c *stringTypesController) StringTypesRangeValidation(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesRangeValidation.HandleWith(c.testActions.stringTypesRangeValidation.action)
}

func (c *stringTypesController) StringTypesRequiredValidation(
	builder *handlers.StringTypesControllerBuilder,
) http.Handler {
	return builder.StringTypesRequiredValidation.HandleWith(c.testActions.stringTypesRequiredValidation.action)
}

var _ handlers.StringTypesController = &stringTypesController{}
