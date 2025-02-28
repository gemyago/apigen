package handlers

import (
	"net/http"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports if that happens.
type _ func() BasicStringEnum

type StringTypesController interface {
	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesArrayItemsRangeValidationParams,
	//
	// Response type: none
	StringTypesArrayItemsRangeValidation(NoResponseHandlerBuilder[
		*StringTypesArrayItemsRangeValidationParams,
	]) http.Handler

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesArraysParsingParams,
	//
	// Response type: none
	StringTypesArraysParsing(NoResponseHandlerBuilder[
		*StringTypesArraysParsingParams,
	]) http.Handler

	// POST /string-types/enums/{inlineEnumParam}/{nullableInlineEnumParam}/{refEnumParam}/{nullableRefEnumParam}
	//
	// Request type: StringTypesEnumsParams,
	//
	// Response type: none
	StringTypesEnums(NoResponseHandlerBuilder[
		*StringTypesEnumsParams,
	]) http.Handler

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesNullableArrayItemsParams,
	//
	// Response type: none
	StringTypesNullableArrayItems(NoResponseHandlerBuilder[
		*StringTypesNullableArrayItemsParams,
	]) http.Handler

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesNullableParsingParams,
	//
	// Response type: none
	StringTypesNullableParsing(NoResponseHandlerBuilder[
		*StringTypesNullableParsingParams,
	]) http.Handler

	// POST /string-types/nullable-required-validation
	//
	// Request type: StringTypesNullableRequiredValidationParams,
	//
	// Response type: none
	StringTypesNullableRequiredValidation(NoResponseHandlerBuilder[
		*StringTypesNullableRequiredValidationParams,
	]) http.Handler

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesParsingParams,
	//
	// Response type: none
	StringTypesParsing(NoResponseHandlerBuilder[
		*StringTypesParsingParams,
	]) http.Handler

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesPatternValidationParams,
	//
	// Response type: none
	StringTypesPatternValidation(NoResponseHandlerBuilder[
		*StringTypesPatternValidationParams,
	]) http.Handler

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesRangeValidationParams,
	//
	// Response type: none
	StringTypesRangeValidation(NoResponseHandlerBuilder[
		*StringTypesRangeValidationParams,
	]) http.Handler

	// POST /string-types/required-validation
	//
	// Request type: StringTypesRequiredValidationParams,
	//
	// Response type: none
	StringTypesRequiredValidation(NoResponseHandlerBuilder[
		*StringTypesRequiredValidationParams,
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