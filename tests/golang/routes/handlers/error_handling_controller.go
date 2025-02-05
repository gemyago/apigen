package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
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
	// GET /error-handling/action-errors
	//
	// Request type: none
	//
	// Response type: none
	ErrorHandlingActionErrors httpHandlerFactory

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
	// GET /error-handling/action-errors
	//
	// Request type: none
	//
	// Response type: none
	HandleErrorHandlingActionErrors actionBuilderNoParamsVoidResult[*ErrorHandlingControllerBuilder]

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
		ErrorHandlingActionErrors: mustInitializeAction("errorHandlingActionErrors", c.HandleErrorHandlingActionErrors.httpHandlerFactory),
		ErrorHandlingParsingErrors: mustInitializeAction("errorHandlingParsingErrors", c.HandleErrorHandlingParsingErrors.httpHandlerFactory),
		ErrorHandlingValidationErrors: mustInitializeAction("errorHandlingValidationErrors", c.HandleErrorHandlingValidationErrors.httpHandlerFactory),
	}
}

func BuildErrorHandlingController() *ErrorHandlingControllerBuilder {
	controllerBuilder := &ErrorHandlingControllerBuilder{}

	// GET /error-handling/action-errors
	controllerBuilder.HandleErrorHandlingActionErrors.controllerBuilder = controllerBuilder
	controllerBuilder.HandleErrorHandlingActionErrors.defaultStatusCode = 204
	controllerBuilder.HandleErrorHandlingActionErrors.voidResult = true
	controllerBuilder.HandleErrorHandlingActionErrors.paramsParserFactory = makeVoidParamsParser

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
	app.router.HandleRoute("GET", "/error-handling/action-errors", controller.ErrorHandlingActionErrors(app))
	app.router.HandleRoute("GET", "/error-handling/parsing-errors/{pathParam1}/{pathParam2}", controller.ErrorHandlingParsingErrors(app))
	app.router.HandleRoute("GET", "/error-handling/validation-errors", controller.ErrorHandlingValidationErrors(app))
}

type ErrorHandlingControllerBuilderV2 struct {
	// GET /error-handling/action-errors
	//
	// Request type: none
	//
	// Response type: none
	ErrorHandlingActionErrors ActionBuilder[
	  Void,
	  Void,
	  func(context.Context) (error),
	  func(http.ResponseWriter, *http.Request) (error),
	]

	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	ErrorHandlingParsingErrors ActionBuilder[
	  ErrorHandlingErrorHandlingParsingErrorsRequest,
	  Void,
	  func(context.Context, *ErrorHandlingErrorHandlingParsingErrorsRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ErrorHandlingErrorHandlingParsingErrorsRequest) (error),
	]

	// GET /error-handling/validation-errors
	//
	// Request type: ErrorHandlingErrorHandlingValidationErrorsRequest,
	//
	// Response type: none
	ErrorHandlingValidationErrors ActionBuilder[
	  ErrorHandlingErrorHandlingValidationErrorsRequest,
	  Void,
	  func(context.Context, *ErrorHandlingErrorHandlingValidationErrorsRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ErrorHandlingErrorHandlingValidationErrorsRequest) (error),
	]
}

func newErrorHandlingControllerBuilderV2(app *HTTPApp) *ErrorHandlingControllerBuilderV2 {
	return &ErrorHandlingControllerBuilderV2{
		// GET /error-handling/action-errors
		ErrorHandlingActionErrors: makeActionBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				Void,
				Void,
			](),
			makeActionBuilderParams[
				Void,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  makeVoidParamsParserV2(app),
			},
		),

		// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
		ErrorHandlingParsingErrors: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				ErrorHandlingErrorHandlingParsingErrorsRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				ErrorHandlingErrorHandlingParsingErrorsRequest,
				Void,
			](),
			makeActionBuilderParams[
				ErrorHandlingErrorHandlingParsingErrorsRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserErrorHandlingErrorHandlingParsingErrors(app),
			},
		),

		// GET /error-handling/validation-errors
		ErrorHandlingValidationErrors: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				ErrorHandlingErrorHandlingValidationErrorsRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				ErrorHandlingErrorHandlingValidationErrorsRequest,
				Void,
			](),
			makeActionBuilderParams[
				ErrorHandlingErrorHandlingValidationErrorsRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserErrorHandlingErrorHandlingValidationErrors(app),
			},
		),
	}
}

type ErrorHandlingControllerV2 interface {
	// GET /error-handling/action-errors
	//
	// Request type: none
	//
	// Response type: none
	ErrorHandlingActionErrors(b *ErrorHandlingControllerBuilderV2) http.Handler

	// GET /error-handling/parsing-errors/{pathParam1}/{pathParam2}
	//
	// Request type: ErrorHandlingErrorHandlingParsingErrorsRequest,
	//
	// Response type: none
	ErrorHandlingParsingErrors(b *ErrorHandlingControllerBuilderV2) http.Handler

	// GET /error-handling/validation-errors
	//
	// Request type: ErrorHandlingErrorHandlingValidationErrorsRequest,
	//
	// Response type: none
	ErrorHandlingValidationErrors(b *ErrorHandlingControllerBuilderV2) http.Handler
}

func RegisterErrorHandlingRoutesV2(controller ErrorHandlingControllerV2, app *HTTPApp) {
	builder := newErrorHandlingControllerBuilderV2(app)
	app.router.HandleRoute("GET", "/error-handling/action-errors", controller.ErrorHandlingActionErrors(builder))
	app.router.HandleRoute("GET", "/error-handling/parsing-errors/{pathParam1}/{pathParam2}", controller.ErrorHandlingParsingErrors(builder))
	app.router.HandleRoute("GET", "/error-handling/validation-errors", controller.ErrorHandlingValidationErrors(builder))
}