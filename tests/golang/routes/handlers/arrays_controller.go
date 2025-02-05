package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint
type _ func() ArraysNullableRequiredValidationRequest

// ArraysArraysNullableRequiredValidationRequest represents params for arraysNullableRequiredValidation operation
//
// Request: POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}.
type ArraysArraysNullableRequiredValidationRequest struct {
	// SimpleItems1 is parsed from request path and declared as simpleItems1.
	SimpleItems1 []string
	// SimpleItems2 is parsed from request path and declared as simpleItems2.
	SimpleItems2 []string
	// SimpleItems1InQuery is parsed from request query and declared as simpleItems1InQuery.
	SimpleItems1InQuery []string
	// SimpleItems2InQuery is parsed from request query and declared as simpleItems2InQuery.
	SimpleItems2InQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *ArraysNullableRequiredValidationRequest
	// OptionalSimpleItems1InQuery is parsed from request query and declared as optionalSimpleItems1InQuery.
	OptionalSimpleItems1InQuery []string
	// OptionalSimpleItems2InQuery is parsed from request query and declared as optionalSimpleItems2InQuery.
	OptionalSimpleItems2InQuery []string
}

// ArraysArraysRangeValidationRequest represents params for arraysRangeValidation operation
//
// Request: POST /arrays/range-validation/{simpleItems1}/{simpleItems2}.
type ArraysArraysRangeValidationRequest struct {
	// SimpleItems1 is parsed from request path and declared as simpleItems1.
	SimpleItems1 []string
	// SimpleItems2 is parsed from request path and declared as simpleItems2.
	SimpleItems2 []string
	// SimpleItems1InQuery is parsed from request query and declared as simpleItems1InQuery.
	SimpleItems1InQuery []string
	// SimpleItems2InQuery is parsed from request query and declared as simpleItems2InQuery.
	SimpleItems2InQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *ArraysRangeValidationRequest
	// OptionalSimpleItems1InQuery is parsed from request query and declared as optionalSimpleItems1InQuery.
	OptionalSimpleItems1InQuery []string
	// OptionalSimpleItems2InQuery is parsed from request query and declared as optionalSimpleItems2InQuery.
	OptionalSimpleItems2InQuery []string
}

// ArraysArraysRequiredValidationRequest represents params for arraysRequiredValidation operation
//
// Request: POST /arrays/required-validation/{simpleItems1}/{simpleItems2}.
type ArraysArraysRequiredValidationRequest struct {
	// SimpleItems1 is parsed from request path and declared as simpleItems1.
	SimpleItems1 []string
	// SimpleItems2 is parsed from request path and declared as simpleItems2.
	SimpleItems2 []string
	// SimpleItems1InQuery is parsed from request query and declared as simpleItems1InQuery.
	SimpleItems1InQuery []string
	// SimpleItems2InQuery is parsed from request query and declared as simpleItems2InQuery.
	SimpleItems2InQuery []string
	// Payload is parsed from request body and declared as payload.
	Payload *ArraysRequiredValidationRequest
	// OptionalSimpleItems1InQuery is parsed from request query and declared as optionalSimpleItems1InQuery.
	OptionalSimpleItems1InQuery []string
	// OptionalSimpleItems2InQuery is parsed from request query and declared as optionalSimpleItems2InQuery.
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

type ArraysControllerBuilderV2 struct {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	ArraysNullableRequiredValidation ActionBuilder[
	  ArraysArraysNullableRequiredValidationRequest,
	  Void,
	  func(context.Context, *ArraysArraysNullableRequiredValidationRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ArraysArraysNullableRequiredValidationRequest) (error),
	]

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRangeValidationRequest,
	//
	// Response type: none
	ArraysRangeValidation ActionBuilder[
	  ArraysArraysRangeValidationRequest,
	  Void,
	  func(context.Context, *ArraysArraysRangeValidationRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ArraysArraysRangeValidationRequest) (error),
	]

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRequiredValidationRequest,
	//
	// Response type: none
	ArraysRequiredValidation ActionBuilder[
	  ArraysArraysRequiredValidationRequest,
	  Void,
	  func(context.Context, *ArraysArraysRequiredValidationRequest) (error),
	  func(http.ResponseWriter, *http.Request, *ArraysArraysRequiredValidationRequest) (error),
	]
}

func newArraysControllerBuilderV2(app *HTTPApp) *ArraysControllerBuilderV2 {
	return &ArraysControllerBuilderV2{
		// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
		ArraysNullableRequiredValidation: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				ArraysArraysNullableRequiredValidationRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				ArraysArraysNullableRequiredValidationRequest,
				Void,
			](),
			makeActionBuilderParams[
				ArraysArraysNullableRequiredValidationRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysNullableRequiredValidation(app),
			},
		),

		// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
		ArraysRangeValidation: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				ArraysArraysRangeValidationRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				ArraysArraysRangeValidationRequest,
				Void,
			](),
			makeActionBuilderParams[
				ArraysArraysRangeValidationRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysRangeValidation(app),
			},
		),

		// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
		ArraysRequiredValidation: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				ArraysArraysRequiredValidationRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				ArraysArraysRequiredValidationRequest,
				Void,
			](),
			makeActionBuilderParams[
				ArraysArraysRequiredValidationRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysRequiredValidation(app),
			},
		),
	}
}

type ArraysControllerV2 interface {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	ArraysNullableRequiredValidation(b *ArraysControllerBuilderV2) http.Handler

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRangeValidationRequest,
	//
	// Response type: none
	ArraysRangeValidation(b *ArraysControllerBuilderV2) http.Handler

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRequiredValidationRequest,
	//
	// Response type: none
	ArraysRequiredValidation(b *ArraysControllerBuilderV2) http.Handler
}

func RegisterArraysRoutesV2(controller ArraysControllerV2, app *HTTPApp) {
	builder := newArraysControllerBuilderV2(app)
	app.router.HandleRoute("POST", "/arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysNullableRequiredValidation(builder))
	app.router.HandleRoute("POST", "/arrays/range-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRangeValidation(builder))
	app.router.HandleRoute("POST", "/arrays/required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRequiredValidation(builder))
}