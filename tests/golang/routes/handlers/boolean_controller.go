package handlers

import (
	"time"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type BooleanBooleanParsingRequest struct {
	BoolParam1        bool
	BoolParam2        bool
	BoolParam1InQuery bool
	BoolParam2InQuery bool
}

type BooleanBooleanRequiredValidationRequest struct {
	BoolParam1InQuery         bool
	BoolParam2InQuery         bool
	OptionalBoolParam1InQuery bool
	OptionalBoolParam2InQuery bool
}

type BooleanController struct {
	// GET /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	BooleanParsing httpHandlerFactory

	// GET /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	BooleanRequiredValidation httpHandlerFactory
}

type BooleanControllerBuilder struct {
	// GET /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	HandleBooleanParsing actionBuilderVoidResult[*BooleanControllerBuilder, *BooleanBooleanParsingRequest]

	// GET /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	HandleBooleanRequiredValidation actionBuilderVoidResult[*BooleanControllerBuilder, *BooleanBooleanRequiredValidationRequest]
}

func (c *BooleanControllerBuilder) Finalize() *BooleanController {
	// TODO: panic if any handler is null
	return &BooleanController{
		BooleanParsing:            c.HandleBooleanParsing.httpHandlerFactory,
		BooleanRequiredValidation: c.HandleBooleanRequiredValidation.httpHandlerFactory,
	}
}

func BuildBooleanController() *BooleanControllerBuilder {
	controllerBuilder := &BooleanControllerBuilder{}

	// GET /boolean/parsing/{boolParam1}/{boolParam2}
	controllerBuilder.HandleBooleanParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBooleanParsing.defaultStatusCode = 204
	controllerBuilder.HandleBooleanParsing.voidResult = true
	controllerBuilder.HandleBooleanParsing.paramsParserFactory = newParamsParserBooleanBooleanParsing

	// GET /boolean/required-validation
	controllerBuilder.HandleBooleanRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBooleanRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleBooleanRequiredValidation.voidResult = true
	controllerBuilder.HandleBooleanRequiredValidation.paramsParserFactory = newParamsParserBooleanBooleanRequiredValidation

	return controllerBuilder
}

func RegisterBooleanRoutes(controller *BooleanController, app *HTTPApp) {
	app.router.HandleRoute("GET", "/boolean/parsing/{boolParam1}/{boolParam2}", controller.BooleanParsing(app))
	app.router.HandleRoute("GET", "/boolean/required-validation", controller.BooleanRequiredValidation(app))
}
