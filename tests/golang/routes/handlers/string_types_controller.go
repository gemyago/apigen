package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type StringTypesArraysNullableRequiredValidationRequest struct {
	SimpleItems1 []string
	SimpleItems2 []string
	SimpleItems1InQuery []string
	SimpleItems2InQuery []string
	Payload *models.ArraysNullableRequiredValidationRequest
	OptionalSimpleItems1InQuery []string
	OptionalSimpleItems2InQuery []string
}

type StringTypesArraysRequiredValidationRequest struct {
	SimpleItems1 []string
	SimpleItems2 []string
	SimpleItems1InQuery []string
	SimpleItems2InQuery []string
	Payload *models.ArraysRequiredValidationRequest
	OptionalSimpleItems1InQuery []string
	OptionalSimpleItems2InQuery []string
}

type StringTypesStringTypesArrayItemsRangeValidationRequest struct {
	UnformattedStr []string
	CustomFormatStr []string
	DateStr []time.Time
	DateTimeStr []time.Time
	ByteStr []string
	UnformattedStrInQuery []string
	CustomFormatStrInQuery []string
	DateStrInQuery []time.Time
	DateTimeStrInQuery []time.Time
	ByteStrInQuery []string
	Payload *models.StringTypesArrayItemsRangeValidationRequest
}

type StringTypesStringTypesArraysParsingRequest struct {
	UnformattedStr []string
	CustomFormatStr []string
	DateStr []time.Time
	DateTimeStr []time.Time
	ByteStr []string
	UnformattedStrInQuery []string
	CustomFormatStrInQuery []string
	DateStrInQuery []time.Time
	DateTimeStrInQuery []time.Time
	ByteStrInQuery []string
	Payload *models.StringTypesArraysParsingRequest
}

type StringTypesStringTypesNullableArrayItemsRequest struct {
	UnformattedStr []*string
	CustomFormatStr []*string
	DateStr []*time.Time
	DateTimeStr []*time.Time
	ByteStr []*string
	UnformattedStrInQuery []*string
	CustomFormatStrInQuery []*string
	DateStrInQuery []*time.Time
	DateTimeStrInQuery []*time.Time
	ByteStrInQuery []*string
	Payload *models.StringTypesNullableArrayItemsRequest
}

type StringTypesStringTypesNullableParsingRequest struct {
	UnformattedStr *string
	CustomFormatStr *string
	DateStr *time.Time
	DateTimeStr *time.Time
	ByteStr *string
	UnformattedStrInQuery *string
	CustomFormatStrInQuery *string
	DateStrInQuery *time.Time
	DateTimeStrInQuery *time.Time
	ByteStrInQuery *string
	Payload *models.StringTypesNullableParsingRequest
}

type StringTypesStringTypesNullableRequiredValidationRequest struct {
	UnformattedStrInQuery *string
	Payload *models.StringTypesNullableRequiredValidationRequest
	OptionalUnformattedStrInQuery *string
}

type StringTypesStringTypesParsingRequest struct {
	UnformattedStr string
	CustomFormatStr string
	DateStr time.Time
	DateTimeStr time.Time
	ByteStr string
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	DateStrInQuery time.Time
	DateTimeStrInQuery time.Time
	ByteStrInQuery string
	Payload *models.StringTypesParsingRequest
}

type StringTypesStringTypesPatternValidationRequest struct {
	UnformattedStr string
	CustomFormatStr string
	DateStr time.Time
	DateTimeStr time.Time
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	DateStrInQuery time.Time
	DateTimeStrInQuery time.Time
	Payload *models.StringTypesPatternValidationRequest
}

type StringTypesStringTypesRangeValidationRequest struct {
	UnformattedStr string
	CustomFormatStr string
	DateStr time.Time
	DateTimeStr time.Time
	ByteStr string
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	DateStrInQuery time.Time
	DateTimeStrInQuery time.Time
	ByteStrInQuery string
	Payload *models.StringTypesRangeValidationRequest
}

type StringTypesStringTypesRequiredValidationRequest struct {
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	DateStrInQuery time.Time
	DateTimeStrInQuery time.Time
	ByteStrInQuery string
	Payload *models.StringTypesRequiredValidationRequest
	OptionalUnformattedStrInQuery string
	OptionalCustomFormatStrInQuery string
	OptionalDateStrInQuery time.Time
	OptionalDateTimeStrInQuery time.Time
	OptionalByteStrInQuery string
}

type StringTypesController struct {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: StringTypesArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	ArraysNullableRequiredValidation httpHandlerFactory

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: StringTypesArraysRequiredValidationRequest,
	//
	// Response type: none
	ArraysRequiredValidation httpHandlerFactory

	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArrayItemsRangeValidationRequest,
	//
	// Response type: none
	StringTypesArrayItemsRangeValidation httpHandlerFactory

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArraysParsingRequest,
	//
	// Response type: none
	StringTypesArraysParsing httpHandlerFactory

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableArrayItemsRequest,
	//
	// Response type: none
	StringTypesNullableArrayItems httpHandlerFactory

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingRequest,
	//
	// Response type: none
	StringTypesNullableParsing httpHandlerFactory

	// POST /string-types/nullable-required-validation
	//
	// Request type: StringTypesStringTypesNullableRequiredValidationRequest,
	//
	// Response type: none
	StringTypesNullableRequiredValidation httpHandlerFactory

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	StringTypesParsing httpHandlerFactory

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	StringTypesPatternValidation httpHandlerFactory

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationRequest,
	//
	// Response type: none
	StringTypesRangeValidation httpHandlerFactory

	// POST /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationRequest,
	//
	// Response type: none
	StringTypesRequiredValidation httpHandlerFactory
}

type StringTypesControllerBuilder struct {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: StringTypesArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	HandleArraysNullableRequiredValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesArraysNullableRequiredValidationRequest]

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: StringTypesArraysRequiredValidationRequest,
	//
	// Response type: none
	HandleArraysRequiredValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesArraysRequiredValidationRequest]

	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArrayItemsRangeValidationRequest,
	//
	// Response type: none
	HandleStringTypesArrayItemsRangeValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesArrayItemsRangeValidationRequest]

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesArraysParsingRequest,
	//
	// Response type: none
	HandleStringTypesArraysParsing actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesArraysParsingRequest]

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableArrayItemsRequest,
	//
	// Response type: none
	HandleStringTypesNullableArrayItems actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesNullableArrayItemsRequest]

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingRequest,
	//
	// Response type: none
	HandleStringTypesNullableParsing actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesNullableParsingRequest]

	// POST /string-types/nullable-required-validation
	//
	// Request type: StringTypesStringTypesNullableRequiredValidationRequest,
	//
	// Response type: none
	HandleStringTypesNullableRequiredValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesNullableRequiredValidationRequest]

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	HandleStringTypesParsing actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesParsingRequest]

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	HandleStringTypesPatternValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesPatternValidationRequest]

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationRequest,
	//
	// Response type: none
	HandleStringTypesRangeValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesRangeValidationRequest]

	// POST /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationRequest,
	//
	// Response type: none
	HandleStringTypesRequiredValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesRequiredValidationRequest]
}

func (c *StringTypesControllerBuilder) Finalize() *StringTypesController {
	return &StringTypesController{
		ArraysNullableRequiredValidation: mustInitializeAction("arraysNullableRequiredValidation", c.HandleArraysNullableRequiredValidation.httpHandlerFactory),
		ArraysRequiredValidation: mustInitializeAction("arraysRequiredValidation", c.HandleArraysRequiredValidation.httpHandlerFactory),
		StringTypesArrayItemsRangeValidation: mustInitializeAction("stringTypesArrayItemsRangeValidation", c.HandleStringTypesArrayItemsRangeValidation.httpHandlerFactory),
		StringTypesArraysParsing: mustInitializeAction("stringTypesArraysParsing", c.HandleStringTypesArraysParsing.httpHandlerFactory),
		StringTypesNullableArrayItems: mustInitializeAction("stringTypesNullableArrayItems", c.HandleStringTypesNullableArrayItems.httpHandlerFactory),
		StringTypesNullableParsing: mustInitializeAction("stringTypesNullableParsing", c.HandleStringTypesNullableParsing.httpHandlerFactory),
		StringTypesNullableRequiredValidation: mustInitializeAction("stringTypesNullableRequiredValidation", c.HandleStringTypesNullableRequiredValidation.httpHandlerFactory),
		StringTypesParsing: mustInitializeAction("stringTypesParsing", c.HandleStringTypesParsing.httpHandlerFactory),
		StringTypesPatternValidation: mustInitializeAction("stringTypesPatternValidation", c.HandleStringTypesPatternValidation.httpHandlerFactory),
		StringTypesRangeValidation: mustInitializeAction("stringTypesRangeValidation", c.HandleStringTypesRangeValidation.httpHandlerFactory),
		StringTypesRequiredValidation: mustInitializeAction("stringTypesRequiredValidation", c.HandleStringTypesRequiredValidation.httpHandlerFactory),
	}
}

func BuildStringTypesController() *StringTypesControllerBuilder {
	controllerBuilder := &StringTypesControllerBuilder{}

	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	controllerBuilder.HandleArraysNullableRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleArraysNullableRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleArraysNullableRequiredValidation.voidResult = true
	controllerBuilder.HandleArraysNullableRequiredValidation.paramsParserFactory = newParamsParserStringTypesArraysNullableRequiredValidation

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	controllerBuilder.HandleArraysRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleArraysRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleArraysRequiredValidation.voidResult = true
	controllerBuilder.HandleArraysRequiredValidation.paramsParserFactory = newParamsParserStringTypesArraysRequiredValidation

	// POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesArrayItemsRangeValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesArrayItemsRangeValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesArrayItemsRangeValidation.voidResult = true
	controllerBuilder.HandleStringTypesArrayItemsRangeValidation.paramsParserFactory = newParamsParserStringTypesStringTypesArrayItemsRangeValidation

	// POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesArraysParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesArraysParsing.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesArraysParsing.voidResult = true
	controllerBuilder.HandleStringTypesArraysParsing.paramsParserFactory = newParamsParserStringTypesStringTypesArraysParsing

	// POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesNullableArrayItems.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesNullableArrayItems.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesNullableArrayItems.voidResult = true
	controllerBuilder.HandleStringTypesNullableArrayItems.paramsParserFactory = newParamsParserStringTypesStringTypesNullableArrayItems

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesNullableParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesNullableParsing.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesNullableParsing.voidResult = true
	controllerBuilder.HandleStringTypesNullableParsing.paramsParserFactory = newParamsParserStringTypesStringTypesNullableParsing

	// POST /string-types/nullable-required-validation
	controllerBuilder.HandleStringTypesNullableRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesNullableRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesNullableRequiredValidation.voidResult = true
	controllerBuilder.HandleStringTypesNullableRequiredValidation.paramsParserFactory = newParamsParserStringTypesStringTypesNullableRequiredValidation

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesParsing.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesParsing.voidResult = true
	controllerBuilder.HandleStringTypesParsing.paramsParserFactory = newParamsParserStringTypesStringTypesParsing

	// POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	controllerBuilder.HandleStringTypesPatternValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesPatternValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesPatternValidation.voidResult = true
	controllerBuilder.HandleStringTypesPatternValidation.paramsParserFactory = newParamsParserStringTypesStringTypesPatternValidation

	// POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesRangeValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesRangeValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesRangeValidation.voidResult = true
	controllerBuilder.HandleStringTypesRangeValidation.paramsParserFactory = newParamsParserStringTypesStringTypesRangeValidation

	// POST /string-types/required-validation
	controllerBuilder.HandleStringTypesRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesRequiredValidation.voidResult = true
	controllerBuilder.HandleStringTypesRequiredValidation.paramsParserFactory = newParamsParserStringTypesStringTypesRequiredValidation

	return controllerBuilder
}

func RegisterStringTypesRoutes(controller *StringTypesController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysNullableRequiredValidation(app))
	app.router.HandleRoute("POST", "/arrays/required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRequiredValidation(app))
	app.router.HandleRoute("POST", "/string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesArrayItemsRangeValidation(app))
	app.router.HandleRoute("POST", "/string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesArraysParsing(app))
	app.router.HandleRoute("POST", "/string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesNullableArrayItems(app))
	app.router.HandleRoute("POST", "/string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesNullableParsing(app))
	app.router.HandleRoute("POST", "/string-types/nullable-required-validation", controller.StringTypesNullableRequiredValidation(app))
	app.router.HandleRoute("POST", "/string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesParsing(app))
	app.router.HandleRoute("POST", "/string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}", controller.StringTypesPatternValidation(app))
	app.router.HandleRoute("POST", "/string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesRangeValidation(app))
	app.router.HandleRoute("POST", "/string-types/required-validation", controller.StringTypesRequiredValidation(app))
}
