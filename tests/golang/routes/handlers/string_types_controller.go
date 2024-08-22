package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint

// StringTypesStringTypesArrayItemsRangeValidationRequest represents params for stringTypesArrayItemsRangeValidation operation
//
// Request: POST /string-types/array-items-range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesArrayItemsRangeValidationRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr []string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr []string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr []time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr []time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr []string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery []string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery []string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery []time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery []time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesArrayItemsRangeValidationRequest
}

// StringTypesStringTypesArraysParsingRequest represents params for stringTypesArraysParsing operation
//
// Request: POST /string-types/arrays-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesArraysParsingRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr []string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr []string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr []time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr []time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr []string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery []string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery []string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery []time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery []time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesArraysParsingRequest
}

// StringTypesStringTypesNullableArrayItemsRequest represents params for stringTypesNullableArrayItems operation
//
// Request: POST /string-types/nullable-array-items/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesNullableArrayItemsRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr []*string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr []*string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr []*time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr []*time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr []*string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery []*string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery []*string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery []*time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery []*time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery []*string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesNullableArrayItemsRequest
}

// StringTypesStringTypesNullableParsingRequest represents params for stringTypesNullableParsing operation
//
// Request: POST /string-types/nullable-parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesNullableParsingRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr *string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr *string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr *time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr *time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr *string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery *string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery *string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery *time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery *time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery *string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesNullableParsingRequest
}

// StringTypesStringTypesNullableRequiredValidationRequest represents params for stringTypesNullableRequiredValidation operation
//
// Request: POST /string-types/nullable-required-validation.
type StringTypesStringTypesNullableRequiredValidationRequest struct {
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery *string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesNullableRequiredValidationRequest
	// OptionalUnformattedStrInQuery is parsed from request query and declared as optionalUnformattedStrInQuery.
	OptionalUnformattedStrInQuery *string
}

// StringTypesStringTypesParsingRequest represents params for stringTypesParsing operation
//
// Request: POST /string-types/parsing/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesParsingRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesParsingRequest
}

// StringTypesStringTypesPatternValidationRequest represents params for stringTypesPatternValidation operation
//
// Request: POST /string-types/pattern-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}.
type StringTypesStringTypesPatternValidationRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr time.Time
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesPatternValidationRequest
}

// StringTypesStringTypesRangeValidationRequest represents params for stringTypesRangeValidation operation
//
// Request: POST /string-types/range-validation/{unformattedStr}/{customFormatStr}/{dateStr}/{dateTimeStr}/{byteStr}.
type StringTypesStringTypesRangeValidationRequest struct {
	// UnformattedStr is parsed from request path and declared as unformattedStr.
	UnformattedStr string
	// CustomFormatStr is parsed from request path and declared as customFormatStr.
	CustomFormatStr string
	// DateStr is parsed from request path and declared as dateStr.
	DateStr time.Time
	// DateTimeStr is parsed from request path and declared as dateTimeStr.
	DateTimeStr time.Time
	// ByteStr is parsed from request path and declared as byteStr.
	ByteStr string
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesRangeValidationRequest
}

// StringTypesStringTypesRequiredValidationRequest represents params for stringTypesRequiredValidation operation
//
// Request: POST /string-types/required-validation.
type StringTypesStringTypesRequiredValidationRequest struct {
	// UnformattedStrInQuery is parsed from request query and declared as unformattedStrInQuery.
	UnformattedStrInQuery string
	// CustomFormatStrInQuery is parsed from request query and declared as customFormatStrInQuery.
	CustomFormatStrInQuery string
	// DateStrInQuery is parsed from request query and declared as dateStrInQuery.
	DateStrInQuery time.Time
	// DateTimeStrInQuery is parsed from request query and declared as dateTimeStrInQuery.
	DateTimeStrInQuery time.Time
	// ByteStrInQuery is parsed from request query and declared as byteStrInQuery.
	ByteStrInQuery string
	// Payload is parsed from request body and declared as payload.
	Payload *models.StringTypesRequiredValidationRequest
	// OptionalUnformattedStrInQuery is parsed from request query and declared as optionalUnformattedStrInQuery.
	OptionalUnformattedStrInQuery string
	// OptionalCustomFormatStrInQuery is parsed from request query and declared as optionalCustomFormatStrInQuery.
	OptionalCustomFormatStrInQuery string
	// OptionalDateStrInQuery is parsed from request query and declared as optionalDateStrInQuery.
	OptionalDateStrInQuery time.Time
	// OptionalDateTimeStrInQuery is parsed from request query and declared as optionalDateTimeStrInQuery.
	OptionalDateTimeStrInQuery time.Time
	// OptionalByteStrInQuery is parsed from request query and declared as optionalByteStrInQuery.
	OptionalByteStrInQuery string
}

type StringTypesController struct {
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
