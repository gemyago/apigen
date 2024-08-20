package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type ArraysArraysNullableRequiredValidationRequest struct {
	SimpleItems1 []string
	SimpleItems2 []string
	SimpleItems1InQuery []string
	SimpleItems2InQuery []string
	Payload *models.ArraysNullableRequiredValidationRequest
	OptionalSimpleItems1InQuery []string
	OptionalSimpleItems2InQuery []string
}

type ArraysArraysRangeValidationRequest struct {
	SimpleItems1 []string
	SimpleItems2 []string
	SimpleItems1InQuery []string
	SimpleItems2InQuery []string
	Payload *models.ArraysRangeValidationRequest
	OptionalSimpleItems1InQuery []string
	OptionalSimpleItems2InQuery []string
}

type ArraysArraysRequiredValidationRequest struct {
	SimpleItems1 []string
	SimpleItems2 []string
	SimpleItems1InQuery []string
	SimpleItems2InQuery []string
	Payload *models.ArraysRequiredValidationRequest
	OptionalSimpleItems1InQuery []string
	OptionalSimpleItems2InQuery []string
}

type ArraysController struct {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	ArraysNullableRequiredValidation httpHandlerFactory

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRangeValidationRequest,
	//
	// Response type: none
	ArraysRangeValidation httpHandlerFactory

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRequiredValidationRequest,
	//
	// Response type: none
	ArraysRequiredValidation httpHandlerFactory
}

type ArraysControllerBuilder struct {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	HandleArraysNullableRequiredValidation actionBuilderVoidResult[*ArraysControllerBuilder, *ArraysArraysNullableRequiredValidationRequest]

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRangeValidationRequest,
	//
	// Response type: none
	HandleArraysRangeValidation actionBuilderVoidResult[*ArraysControllerBuilder, *ArraysArraysRangeValidationRequest]

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRequiredValidationRequest,
	//
	// Response type: none
	HandleArraysRequiredValidation actionBuilderVoidResult[*ArraysControllerBuilder, *ArraysArraysRequiredValidationRequest]
}

func (c *ArraysControllerBuilder) Finalize() *ArraysController {
	return &ArraysController{
		ArraysNullableRequiredValidation: mustInitializeAction("arraysNullableRequiredValidation", c.HandleArraysNullableRequiredValidation.httpHandlerFactory),
		ArraysRangeValidation: mustInitializeAction("arraysRangeValidation", c.HandleArraysRangeValidation.httpHandlerFactory),
		ArraysRequiredValidation: mustInitializeAction("arraysRequiredValidation", c.HandleArraysRequiredValidation.httpHandlerFactory),
	}
}

func BuildArraysController() *ArraysControllerBuilder {
	controllerBuilder := &ArraysControllerBuilder{}

	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	controllerBuilder.HandleArraysNullableRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleArraysNullableRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleArraysNullableRequiredValidation.voidResult = true
	controllerBuilder.HandleArraysNullableRequiredValidation.paramsParserFactory = newParamsParserArraysArraysNullableRequiredValidation

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	controllerBuilder.HandleArraysRangeValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleArraysRangeValidation.defaultStatusCode = 204
	controllerBuilder.HandleArraysRangeValidation.voidResult = true
	controllerBuilder.HandleArraysRangeValidation.paramsParserFactory = newParamsParserArraysArraysRangeValidation

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	controllerBuilder.HandleArraysRequiredValidation.controllerBuilder = controllerBuilder
	controllerBuilder.HandleArraysRequiredValidation.defaultStatusCode = 204
	controllerBuilder.HandleArraysRequiredValidation.voidResult = true
	controllerBuilder.HandleArraysRequiredValidation.paramsParserFactory = newParamsParserArraysArraysRequiredValidation

	return controllerBuilder
}

func RegisterArraysRoutes(controller *ArraysController, app *HTTPApp) {
	app.router.HandleRoute("POST", "/arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysNullableRequiredValidation(app))
	app.router.HandleRoute("POST", "/arrays/range-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRangeValidation(app))
	app.router.HandleRoute("POST", "/arrays/required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRequiredValidation(app))
}
