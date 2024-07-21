package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

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
	Payload models.StringTypesParsingRequest
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
}

type StringTypesStringTypesRequiredValidationRequest struct {
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	DateStrInQuery time.Time
	DateTimeStrInQuery time.Time
	ByteStrInQuery string
	OptionalUnformattedStrInQuery string
	OptionalCustomFormatStrInQuery string
	OptionalDateStrInQuery time.Time
	OptionalDateTimeStrInQuery time.Time
	OptionalByteStrInQuery string
}

type StringTypesController struct {
	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	StringTypesParsing httpHandlerFactory

	// GET /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	StringTypesPatternValidation httpHandlerFactory

	// GET /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationRequest,
	//
	// Response type: none
	StringTypesRangeValidation httpHandlerFactory

	// GET /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationRequest,
	//
	// Response type: none
	StringTypesRequiredValidation httpHandlerFactory
}

type StringTypesControllerBuilder struct {
	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	HandleStringTypesParsing actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesParsingRequest]

	// GET /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	HandleStringTypesPatternValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesPatternValidationRequest]

	// GET /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesRangeValidationRequest,
	//
	// Response type: none
	HandleStringTypesRangeValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesRangeValidationRequest]

	// GET /string-types/required-validation
	//
	// Request type: StringTypesStringTypesRequiredValidationRequest,
	//
	// Response type: none
	HandleStringTypesRequiredValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesRequiredValidationRequest]
}

func (c *StringTypesControllerBuilder) Finalize() *StringTypesController {
	// TODO: panic if any handler is null
	return &StringTypesController{
		StringTypesParsing: c.HandleStringTypesParsing.httpHandlerFactory,
		StringTypesPatternValidation: c.HandleStringTypesPatternValidation.httpHandlerFactory,
		StringTypesRangeValidation: c.HandleStringTypesRangeValidation.httpHandlerFactory,
		StringTypesRequiredValidation: c.HandleStringTypesRequiredValidation.httpHandlerFactory,
	}
}

func BuildStringTypesController() *StringTypesControllerBuilder {
	controllerBuilder := &StringTypesControllerBuilder{}

	// POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesParsing.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesParsing.voidResult = true
	controllerBuilder.HandleStringTypesParsing.paramsParserFactory = newParamsParserStringTypesStringTypesParsing

	// GET /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}
	controllerBuilder.HandleStringTypesPatternValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesPatternValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesPatternValidation.voidResult = true
	controllerBuilder.HandleStringTypesPatternValidation.paramsParserFactory = newParamsParserStringTypesStringTypesPatternValidation

	// GET /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesRangeValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesRangeValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesRangeValidation.voidResult = true
	controllerBuilder.HandleStringTypesRangeValidation.paramsParserFactory = newParamsParserStringTypesStringTypesRangeValidation

	// GET /string-types/required-validation
	controllerBuilder.HandleStringTypesRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesRequiredValidation.voidResult = true
	controllerBuilder.HandleStringTypesRequiredValidation.paramsParserFactory = newParamsParserStringTypesStringTypesRequiredValidation

	return controllerBuilder
}

func RegisterStringTypesRoutes(controller *StringTypesController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesParsing(app))
	app.router.HandleRoute("GET", "/string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}", controller.StringTypesPatternValidation(app))
	app.router.HandleRoute("GET", "/string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesRangeValidation(app))
	app.router.HandleRoute("GET", "/string-types/required-validation", controller.StringTypesRequiredValidation(app))
}
