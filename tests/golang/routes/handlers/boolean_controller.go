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

// BooleanBooleanArrayItemsRequest represents params for booleanArrayItems operation
//
// Request: POST /boolean/array-items/{boolParam1}/{boolParam2}.
type BooleanBooleanArrayItemsRequest struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 []bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 []bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery []bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery []bool
	// Payload is parsed from request body and declared as payload.
	Payload *models.BooleanArrayItemsRequest
}

// BooleanBooleanNullableRequest represents params for booleanNullable operation
//
// Request: POST /boolean/nullable/{boolParam1}/{boolParam2}.
type BooleanBooleanNullableRequest struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 *bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 *bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery *bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery *bool
	// Payload is parsed from request body and declared as payload.
	Payload *models.BooleanNullableRequest
	// OptionalBoolParam1InQuery is parsed from request query and declared as optionalBoolParam1InQuery.
	OptionalBoolParam1InQuery *bool
}

// BooleanBooleanNullableArrayItemsRequest represents params for booleanNullableArrayItems operation
//
// Request: POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}.
type BooleanBooleanNullableArrayItemsRequest struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 []*bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 []*bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery []*bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery []*bool
	// Payload is parsed from request body and declared as payload.
	Payload *models.BooleanNullableArrayItemsRequest
}

// BooleanBooleanParsingRequest represents params for booleanParsing operation
//
// Request: POST /boolean/parsing/{boolParam1}/{boolParam2}.
type BooleanBooleanParsingRequest struct {
	// BoolParam1 is parsed from request path and declared as boolParam1.
	BoolParam1 bool
	// BoolParam2 is parsed from request path and declared as boolParam2.
	BoolParam2 bool
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery bool
	// Payload is parsed from request body and declared as payload.
	Payload *models.BooleanParsingRequest
}

// BooleanBooleanRequiredValidationRequest represents params for booleanRequiredValidation operation
//
// Request: POST /boolean/required-validation.
type BooleanBooleanRequiredValidationRequest struct {
	// BoolParam1InQuery is parsed from request query and declared as boolParam1InQuery.
	BoolParam1InQuery bool
	// BoolParam2InQuery is parsed from request query and declared as boolParam2InQuery.
	BoolParam2InQuery bool
	// Payload is parsed from request body and declared as payload.
	Payload *models.BooleanRequiredValidationRequest
	// OptionalBoolParam1InQuery is parsed from request query and declared as optionalBoolParam1InQuery.
	OptionalBoolParam1InQuery bool
	// OptionalBoolParam2InQuery is parsed from request query and declared as optionalBoolParam2InQuery.
	OptionalBoolParam2InQuery bool
}

type BooleanController struct {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsRequest,
	//
	// Response type: none
	BooleanArrayItems httpHandlerFactory

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableRequest,
	//
	// Response type: none
	BooleanNullable httpHandlerFactory

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsRequest,
	//
	// Response type: none
	BooleanNullableArrayItems httpHandlerFactory

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
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsRequest,
	//
	// Response type: none
	HandleBooleanArrayItems actionBuilderVoidResult[*BooleanControllerBuilder, *BooleanBooleanArrayItemsRequest]

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableRequest,
	//
	// Response type: none
	HandleBooleanNullable actionBuilderVoidResult[*BooleanControllerBuilder, *BooleanBooleanNullableRequest]

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsRequest,
	//
	// Response type: none
	HandleBooleanNullableArrayItems actionBuilderVoidResult[*BooleanControllerBuilder, *BooleanBooleanNullableArrayItemsRequest]

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
		BooleanArrayItems: mustInitializeAction("booleanArrayItems", c.HandleBooleanArrayItems.httpHandlerFactory),
		BooleanNullable: mustInitializeAction("booleanNullable", c.HandleBooleanNullable.httpHandlerFactory),
		BooleanNullableArrayItems: mustInitializeAction("booleanNullableArrayItems", c.HandleBooleanNullableArrayItems.httpHandlerFactory),
		BooleanParsing: mustInitializeAction("booleanParsing", c.HandleBooleanParsing.httpHandlerFactory),
		BooleanRequiredValidation: mustInitializeAction("booleanRequiredValidation", c.HandleBooleanRequiredValidation.httpHandlerFactory),
	}
}

func BuildBooleanController() *BooleanControllerBuilder {
	controllerBuilder := &BooleanControllerBuilder{}

	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	controllerBuilder.HandleBooleanArrayItems.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBooleanArrayItems.defaultStatusCode = 204
	controllerBuilder.HandleBooleanArrayItems.voidResult = true
	controllerBuilder.HandleBooleanArrayItems.paramsParserFactory = newParamsParserBooleanBooleanArrayItems

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	controllerBuilder.HandleBooleanNullable.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBooleanNullable.defaultStatusCode = 204
	controllerBuilder.HandleBooleanNullable.voidResult = true
	controllerBuilder.HandleBooleanNullable.paramsParserFactory = newParamsParserBooleanBooleanNullable

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	controllerBuilder.HandleBooleanNullableArrayItems.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBooleanNullableArrayItems.defaultStatusCode = 204
	controllerBuilder.HandleBooleanNullableArrayItems.voidResult = true
	controllerBuilder.HandleBooleanNullableArrayItems.paramsParserFactory = newParamsParserBooleanBooleanNullableArrayItems

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
	app.router.HandleRoute("POST", "/boolean/array-items/{boolParam1}/{boolParam2}", controller.BooleanArrayItems(app))
	app.router.HandleRoute("POST", "/boolean/nullable/{boolParam1}/{boolParam2}", controller.BooleanNullable(app))
	app.router.HandleRoute("POST", "/boolean/nullable-array-items/{boolParam1}/{boolParam2}", controller.BooleanNullableArrayItems(app))
	app.router.HandleRoute("POST", "/boolean/parsing/{boolParam1}/{boolParam2}", controller.BooleanParsing(app))
	app.router.HandleRoute("POST", "/boolean/required-validation", controller.BooleanRequiredValidation(app))
}
