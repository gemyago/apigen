package handlers

type StringTypesStringTypesParsingRequest struct {
	UnformattedStr string
	CustomFormatStr string
	DateStr string
	DateTimeStr string
	ByteStr string
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	DateStrInQuery string
	DateTimeStrInQuery string
	ByteStrInQuery string
}

type StringTypesStringTypesPatternValidationRequest struct {
	UnformattedStr string
	CustomFormatStr string
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
}

type StringTypesStringTypesRangeValidationRequest struct {
	UnformattedStr string
	CustomFormatStr string
	ByteStr string
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	ByteStrInQuery string
}

type StringTypesStringTypesRequiredValidationRequest struct {
	UnformattedStrInQuery string
	CustomFormatStrInQuery string
	DateStrInQuery string
	DateTimeStrInQuery string
	ByteStrInQuery string
	OptionalUnformattedStrInQuery string
	OptionalCustomFormatStrInQuery string
	OptionalDateStrInQuery string
	OptionalDateTimeStrInQuery string
	OptionalByteStrInQuery string
}

type StringTypesController struct {
	// GET /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	StringTypesParsing httpHandlerFactory

	// GET /string-types/pattern-validation/{unformattedStr}/{customFormatStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	StringTypesPatternValidation httpHandlerFactory

	// GET /string-types/range-validation/{unformattedStr}/{customFormatStr}/{byteStr}
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
	// GET /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	//
	// Request type: StringTypesStringTypesParsingRequest,
	//
	// Response type: none
	HandleStringTypesParsing actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesParsingRequest]

	// GET /string-types/pattern-validation/{unformattedStr}/{customFormatStr}
	//
	// Request type: StringTypesStringTypesPatternValidationRequest,
	//
	// Response type: none
	HandleStringTypesPatternValidation actionBuilderVoidResult[*StringTypesControllerBuilder, *StringTypesStringTypesPatternValidationRequest]

	// GET /string-types/range-validation/{unformattedStr}/{customFormatStr}/{byteStr}
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

	// GET /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}
	controllerBuilder.HandleStringTypesParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesParsing.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesParsing.voidResult = true
	controllerBuilder.HandleStringTypesParsing.paramsParser = newParamsParserStringTypesStringTypesParsing()

	// GET /string-types/pattern-validation/{unformattedStr}/{customFormatStr}
	controllerBuilder.HandleStringTypesPatternValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesPatternValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesPatternValidation.voidResult = true
	controllerBuilder.HandleStringTypesPatternValidation.paramsParser = newParamsParserStringTypesStringTypesPatternValidation()

	// GET /string-types/range-validation/{unformattedStr}/{customFormatStr}/{byteStr}
	controllerBuilder.HandleStringTypesRangeValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesRangeValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesRangeValidation.voidResult = true
	controllerBuilder.HandleStringTypesRangeValidation.paramsParser = newParamsParserStringTypesStringTypesRangeValidation()

	// GET /string-types/required-validation
	controllerBuilder.HandleStringTypesRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleStringTypesRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleStringTypesRequiredValidation.voidResult = true
	controllerBuilder.HandleStringTypesRequiredValidation.paramsParser = newParamsParserStringTypesStringTypesRequiredValidation()

	return controllerBuilder
}

func MountStringTypesRoutes(controller *StringTypesController, app *httpApp) {
	app.router.HandleRoute("GET", "/string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}", controller.StringTypesParsing(app))
	app.router.HandleRoute("GET", "/string-types/pattern-validation/{unformattedStr}/{customFormatStr}", controller.StringTypesPatternValidation(app))
	app.router.HandleRoute("GET", "/string-types/range-validation/{unformattedStr}/{customFormatStr}/{byteStr}", controller.StringTypesRangeValidation(app))
	app.router.HandleRoute("GET", "/string-types/required-validation", controller.StringTypesRequiredValidation(app))
}
