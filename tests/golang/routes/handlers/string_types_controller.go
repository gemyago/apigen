package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

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
	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingRequest,
	//
	// Response type: none
	StringTypesNullableParsing httpHandlerFactory

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
	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesNullableParsingRequest,
	//
	// Response type: none
	HandleStringTypesNullableParsing actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesNullableParsingRequest]

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
		StringTypesNullableParsing: mustInitializeAction("stringTypesNullableParsing", c.HandleStringTypesNullableParsing.httpHandlerFactory),
		StringTypesParsing: mustInitializeAction("stringTypesParsing", c.HandleStringTypesParsing.httpHandlerFactory),
		StringTypesPatternValidation: mustInitializeAction("stringTypesPatternValidation", c.HandleStringTypesPatternValidation.httpHandlerFactory),
		StringTypesRangeValidation: mustInitializeAction("stringTypesRangeValidation", c.HandleStringTypesRangeValidation.httpHandlerFactory),
		StringTypesRequiredValidation: mustInitializeAction("stringTypesRequiredValidation", c.HandleStringTypesRequiredValidation.httpHandlerFactory),
	}
}

func BuildStringTypesController() *StringTypesControllerBuilder {
	controllerBuilder := &StringTypesControllerBuilder{}

	// POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesNullableParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesNullableParsing.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesNullableParsing.voidResult = true
	controllerBuilder.HandleStringTypesNullableParsing.paramsParserFactory = newParamsParserStringTypesStringTypesNullableParsing

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
	app.router.HandleRoute("POST", "/string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesNullableParsing(app))
	app.router.HandleRoute("POST", "/string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesParsing(app))
	app.router.HandleRoute("POST", "/string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}", controller.StringTypesPatternValidation(app))
	app.router.HandleRoute("POST", "/string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesRangeValidation(app))
	app.router.HandleRoute("POST", "/string-types/required-validation", controller.StringTypesRequiredValidation(app))
}
