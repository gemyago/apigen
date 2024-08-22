package handlers

import (
	"encoding/json"
	"fmt"
	"time"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint

// ErrorHandlingErrorHandlingParsingErrorsRequest represents params for errorHandlingParsingErrors operation
//
// Request: GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}.
type ErrorHandlingErrorHandlingParsingErrorsRequest struct {
	// PathParam1 is parsed from request path and declared as pathParam1.
	PathParam1 float32
	// PathParam2 is parsed from request path and declared as pathParam2.
	PathParam2 float32
	// RequiredQuery1 is parsed from request query and declared as requiredQuery1.
	RequiredQuery1 float32
	// RequiredQuery2 is parsed from request query and declared as requiredQuery2.
	RequiredQuery2 float32
}

// ErrorHandlingErrorHandlingValidationErrorsRequest represents params for errorHandlingValidationErrors operation
//
// Request: GET /error-handling/validation-errors.
type ErrorHandlingErrorHandlingValidationErrorsRequest struct {
	// RequiredQuery1 is parsed from request query and declared as requiredQuery1.
	RequiredQuery1 float32
	// RequiredQuery2 is parsed from request query and declared as requiredQuery2.
	RequiredQuery2 float32
}

type ErrorHandlingController struct {
	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	ErrorHandlingParsingErrors httpHandlerFactory

	// GET /error-handling/validation-errors
	//
	// Request type: ErrorHandlingErrorHandlingValidationErrorsRequest,
	//
	// Response type: none
	ErrorHandlingValidationErrors httpHandlerFactory
}

type ErrorHandlingControllerBuilder struct {
	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	HandleErrorHandlingParsingErrors actionBuilderVoidResult[*ErrorHandlingControllerBuilder, *ErrorHandlingErrorHandlingParsingErrorsRequest]

	// GET /error-handling/validation-errors
	//
	// Request type: ErrorHandlingErrorHandlingValidationErrorsRequest,
	//
	// Response type: none
	HandleErrorHandlingValidationErrors actionBuilderVoidResult[*ErrorHandlingControllerBuilder, *ErrorHandlingErrorHandlingValidationErrorsRequest]
}

func (c *ErrorHandlingControllerBuilder) Finalize() *ErrorHandlingController {
	return &ErrorHandlingController{
		ErrorHandlingParsingErrors: mustInitializeAction("errorHandlingParsingErrors", c.HandleErrorHandlingParsingErrors.httpHandlerFactory),
		ErrorHandlingValidationErrors: mustInitializeAction("errorHandlingValidationErrors", c.HandleErrorHandlingValidationErrors.httpHandlerFactory),
	}
}

func BuildErrorHandlingController() *ErrorHandlingControllerBuilder {
	controllerBuilder := &ErrorHandlingControllerBuilder{}

	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	controllerBuilder.HandleErrorHandlingParsingErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleErrorHandlingParsingErrors.defaultStatusCode = 204
	controllerBuilder.HandleErrorHandlingParsingErrors.voidResult = true
	controllerBuilder.HandleErrorHandlingParsingErrors.paramsParserFactory = newParamsParserErrorHandlingErrorHandlingParsingErrors

	// GET /error-handling/validation-errors
	controllerBuilder.HandleErrorHandlingValidationErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleErrorHandlingValidationErrors.defaultStatusCode = 204
	controllerBuilder.HandleErrorHandlingValidationErrors.voidResult = true
	controllerBuilder.HandleErrorHandlingValidationErrors.paramsParserFactory = newParamsParserErrorHandlingErrorHandlingValidationErrors

	return controllerBuilder
}

func RegisterErrorHandlingRoutes(controller *ErrorHandlingController, app *HTTPApp) {
	app.router.HandleRoute("GET", "/error-handling/parsing-errors/{pathParam1}/{pathParam2}", controller.ErrorHandlingParsingErrors(app))
	app.router.HandleRoute("GET", "/error-handling/validation-errors", controller.ErrorHandlingValidationErrors(app))
}
