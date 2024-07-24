package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type BooleanBooleanParsingRequest struct {
	BoolParam1 bool
	BoolParam2 bool
	BoolParam1InQuery bool
	BoolParam2InQuery bool
	Payload *models.BooleanParsingRequest
}

type BooleanBooleanRequiredValidationRequest struct {
	BoolParam1InQuery bool
	BoolParam2InQuery bool
	Payload *models.BooleanRequiredValidationRequest
	OptionalBoolParam1InQuery bool
	OptionalBoolParam2InQuery bool
}

type BooleanController struct {
	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	BooleanParsing httpHandlerFactory

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	BooleanRequiredValidation httpHandlerFactory
}

type BooleanControllerBuilder struct {
	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	HandleBooleanParsing actionBuilderVoidResult[*BooleanControllerBuilder, *BooleanBooleanParsingRequest]

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	HandleBooleanRequiredValidation actionBuilderVoidResult[*BooleanControllerBuilder, *BooleanBooleanRequiredValidationRequest]
}

func (c *BooleanControllerBuilder) Finalize() *BooleanController {
	return &BooleanController{
		BooleanParsing: mustInitializeAction("booleanParsing", c.HandleBooleanParsing.httpHandlerFactory),
		BooleanRequiredValidation: mustInitializeAction("booleanRequiredValidation", c.HandleBooleanRequiredValidation.httpHandlerFactory),
	}
}

func BuildBooleanController() *BooleanControllerBuilder {
	controllerBuilder := &BooleanControllerBuilder{}

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	controllerBuilder.HandleBooleanParsing.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBooleanParsing.defaultStatusCode = 204
	controllerBuilder.HandleBooleanParsing.voidResult = true
	controllerBuilder.HandleBooleanParsing.paramsParserFactory = newParamsParserBooleanBooleanParsing

	// POST /boolean/required-validation
	controllerBuilder.HandleBooleanRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBooleanRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleBooleanRequiredValidation.voidResult = true
	controllerBuilder.HandleBooleanRequiredValidation.paramsParserFactory = newParamsParserBooleanBooleanRequiredValidation

	return controllerBuilder
}

func RegisterBooleanRoutes(controller *BooleanController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/boolean/parsing/{boolParam1}/{boolParam2}", controller.BooleanParsing(app))
	app.router.HandleRoute("POST", "/boolean/required-validation", controller.BooleanRequiredValidation(app))
}
