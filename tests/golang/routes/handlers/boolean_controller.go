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
type _ func() BooleanArrayItemsRequest

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
	Payload *BooleanArrayItemsRequest
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
	Payload *BooleanNullableRequest
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
	Payload *BooleanNullableArrayItemsRequest
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
	Payload *BooleanParsingRequest
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
	Payload *BooleanRequiredValidationRequest
	// OptionalBoolParam1InQuery is parsed from request query and declared as optionalBoolParam1InQuery.
	OptionalBoolParam1InQuery bool
	// OptionalBoolParam2InQuery is parsed from request query and declared as optionalBoolParam2InQuery.
	OptionalBoolParam2InQuery bool
}

type BooleanControllerBuilderV2 struct {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsRequest,
	//
	// Response type: none
	BooleanArrayItems ActionBuilder[
	  BooleanBooleanArrayItemsRequest,
	  Void,
	  func(context.Context, *BooleanBooleanArrayItemsRequest) (error),
	  func(http.ResponseWriter, *http.Request, *BooleanBooleanArrayItemsRequest) (error),
	]

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableRequest,
	//
	// Response type: none
	BooleanNullable ActionBuilder[
	  BooleanBooleanNullableRequest,
	  Void,
	  func(context.Context, *BooleanBooleanNullableRequest) (error),
	  func(http.ResponseWriter, *http.Request, *BooleanBooleanNullableRequest) (error),
	]

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsRequest,
	//
	// Response type: none
	BooleanNullableArrayItems ActionBuilder[
	  BooleanBooleanNullableArrayItemsRequest,
	  Void,
	  func(context.Context, *BooleanBooleanNullableArrayItemsRequest) (error),
	  func(http.ResponseWriter, *http.Request, *BooleanBooleanNullableArrayItemsRequest) (error),
	]

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	BooleanParsing ActionBuilder[
	  BooleanBooleanParsingRequest,
	  Void,
	  func(context.Context, *BooleanBooleanParsingRequest) (error),
	  func(http.ResponseWriter, *http.Request, *BooleanBooleanParsingRequest) (error),
	]

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	BooleanRequiredValidation ActionBuilder[
	  BooleanBooleanRequiredValidationRequest,
	  Void,
	  func(context.Context, *BooleanBooleanRequiredValidationRequest) (error),
	  func(http.ResponseWriter, *http.Request, *BooleanBooleanRequiredValidationRequest) (error),
	]
}

func newBooleanControllerBuilderV2(app *HTTPApp) *BooleanControllerBuilderV2 {
	return &BooleanControllerBuilderV2{
		// POST /boolean/array-items/{boolParam1}/{boolParam2}
		BooleanArrayItems: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				BooleanBooleanArrayItemsRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				BooleanBooleanArrayItemsRequest,
				Void,
			](),
			makeActionBuilderParams[
				BooleanBooleanArrayItemsRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanArrayItems(app),
			},
		),

		// POST /boolean/nullable/{boolParam1}/{boolParam2}
		BooleanNullable: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				BooleanBooleanNullableRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				BooleanBooleanNullableRequest,
				Void,
			](),
			makeActionBuilderParams[
				BooleanBooleanNullableRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanNullable(app),
			},
		),

		// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
		BooleanNullableArrayItems: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				BooleanBooleanNullableArrayItemsRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				BooleanBooleanNullableArrayItemsRequest,
				Void,
			](),
			makeActionBuilderParams[
				BooleanBooleanNullableArrayItemsRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanNullableArrayItems(app),
			},
		),

		// POST /boolean/parsing/{boolParam1}/{boolParam2}
		BooleanParsing: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				BooleanBooleanParsingRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				BooleanBooleanParsingRequest,
				Void,
			](),
			makeActionBuilderParams[
				BooleanBooleanParsingRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanParsing(app),
			},
		),

		// POST /boolean/required-validation
		BooleanRequiredValidation: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				BooleanBooleanRequiredValidationRequest,
				Void,
			](),
			newHTTPHandlerAdapterNoResponse[
				BooleanBooleanRequiredValidationRequest,
				Void,
			](),
			makeActionBuilderParams[
				BooleanBooleanRequiredValidationRequest,
				Void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanRequiredValidation(app),
			},
		),
	}
}

type BooleanControllerV2 interface {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsRequest,
	//
	// Response type: none
	BooleanArrayItems(b *BooleanControllerBuilderV2) http.Handler

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableRequest,
	//
	// Response type: none
	BooleanNullable(b *BooleanControllerBuilderV2) http.Handler

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsRequest,
	//
	// Response type: none
	BooleanNullableArrayItems(b *BooleanControllerBuilderV2) http.Handler

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	BooleanParsing(b *BooleanControllerBuilderV2) http.Handler

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	BooleanRequiredValidation(b *BooleanControllerBuilderV2) http.Handler
}

func RegisterBooleanRoutesV2(controller BooleanControllerV2, app *HTTPApp) {
	builder := newBooleanControllerBuilderV2(app)
	app.router.HandleRoute("POST", "/boolean/array-items/{boolParam1}/{boolParam2}", controller.BooleanArrayItems(builder))
	app.router.HandleRoute("POST", "/boolean/nullable/{boolParam1}/{boolParam2}", controller.BooleanNullable(builder))
	app.router.HandleRoute("POST", "/boolean/nullable-array-items/{boolParam1}/{boolParam2}", controller.BooleanNullableArrayItems(builder))
	app.router.HandleRoute("POST", "/boolean/parsing/{boolParam1}/{boolParam2}", controller.BooleanParsing(builder))
	app.router.HandleRoute("POST", "/boolean/required-validation", controller.BooleanRequiredValidation(builder))
}