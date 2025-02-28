package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

type StringTypesController interface {
	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArrayItemsRangeValidationParams,
	//
	// Response type: none
	StringTypesArrayItemsRangeValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesArrayItemsRangeValidationParams,
	]) http.Handler

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArraysParsingParams,
	//
	// Response type: none
	StringTypesArraysParsing(NoResponseHandlerBuilder[
		*StringTypesStringTypesArraysParsingParams,
	]) http.Handler

	// POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
	//
	// Request type: StringTypesStringTypesEnumsParams,
	//
	// Response type: none
	StringTypesEnums(NoResponseHandlerBuilder[
		*StringTypesStringTypesEnumsParams,
	]) http.Handler

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableArrayItemsParams,
	//
	// Response type: none
	StringTypesNullableArrayItems(NoResponseHandlerBuilder[
		*StringTypesStringTypesNullableArrayItemsParams,
	]) http.Handler

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingParams,
	//
	// Response type: none
	StringTypesNullableParsing(NoResponseHandlerBuilder[
		*StringTypesStringTypesNullableParsingParams,
	]) http.Handler

	// POST /string-types/nullable-required-validation
	//
	// Request type: StringTypesStringTypesNullableRequiredValidationParams,
	//
	// Response type: none
	StringTypesNullableRequiredValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesNullableRequiredValidationParams,
	]) http.Handler

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingParams,
	//
	// Response type: none
	StringTypesParsing(NoResponseHandlerBuilder[
		*StringTypesStringTypesParsingParams,
	]) http.Handler

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationParams,
	//
	// Response type: none
	StringTypesPatternValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesPatternValidationParams,
	]) http.Handler

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationParams,
	//
	// Response type: none
	StringTypesRangeValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesRangeValidationParams,
	]) http.Handler

	// POST /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationParams,
	//
	// Response type: none
	StringTypesRequiredValidation(NoResponseHandlerBuilder[
		*StringTypesStringTypesRequiredValidationParams,
	]) http.Handler
}

// RegisterStringTypesRoutes will attach the following routes to the root handler:
// 
// - POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
// 
// - POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/nullable-required-validation
// 
// - POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
// 
// - POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
// 
// - POST /string-types/required-validation
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterStringTypesRoutes(controller StringTypesController) *RootHandler {
	builder := newStringTypesControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesArrayItemsRangeValidation(builder.StringTypesArrayItemsRangeValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesArraysParsing(builder.StringTypesArraysParsing))
	rootHandler.router.HandleRoute("POST", "/string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}", controller.StringTypesEnums(builder.StringTypesEnums))
	rootHandler.router.HandleRoute("POST", "/string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesNullableArrayItems(builder.StringTypesNullableArrayItems))
	rootHandler.router.HandleRoute("POST", "/string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesNullableParsing(builder.StringTypesNullableParsing))
	rootHandler.router.HandleRoute("POST", "/string-types/nullable-required-validation", controller.StringTypesNullableRequiredValidation(builder.StringTypesNullableRequiredValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesParsing(builder.StringTypesParsing))
	rootHandler.router.HandleRoute("POST", "/string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}", controller.StringTypesPatternValidation(builder.StringTypesPatternValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesRangeValidation(builder.StringTypesRangeValidation))
	rootHandler.router.HandleRoute("POST", "/string-types/required-validation", controller.StringTypesRequiredValidation(builder.StringTypesRequiredValidation))
	return rootHandler
}