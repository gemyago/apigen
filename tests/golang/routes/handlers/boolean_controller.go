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

type BooleanControllerBuilder struct {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsRequest,
	//
	// Response type: none
	BooleanArrayItems ActionBuilder[
		*BooleanBooleanArrayItemsRequest,
		void,
		func(context.Context, *BooleanBooleanArrayItemsRequest) (error),
		func(http.ResponseWriter, *http.Request, *BooleanBooleanArrayItemsRequest) (error),
	]

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableRequest,
	//
	// Response type: none
	BooleanNullable ActionBuilder[
		*BooleanBooleanNullableRequest,
		void,
		func(context.Context, *BooleanBooleanNullableRequest) (error),
		func(http.ResponseWriter, *http.Request, *BooleanBooleanNullableRequest) (error),
	]

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsRequest,
	//
	// Response type: none
	BooleanNullableArrayItems ActionBuilder[
		*BooleanBooleanNullableArrayItemsRequest,
		void,
		func(context.Context, *BooleanBooleanNullableArrayItemsRequest) (error),
		func(http.ResponseWriter, *http.Request, *BooleanBooleanNullableArrayItemsRequest) (error),
	]

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	BooleanParsing ActionBuilder[
		*BooleanBooleanParsingRequest,
		void,
		func(context.Context, *BooleanBooleanParsingRequest) (error),
		func(http.ResponseWriter, *http.Request, *BooleanBooleanParsingRequest) (error),
	]

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	BooleanRequiredValidation ActionBuilder[
		*BooleanBooleanRequiredValidationRequest,
		void,
		func(context.Context, *BooleanBooleanRequiredValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *BooleanBooleanRequiredValidationRequest) (error),
	]
}

func newBooleanControllerBuilder(app *HTTPApp) *BooleanControllerBuilder {
	return &BooleanControllerBuilder{
		// POST /boolean/array-items/{boolParam1}/{boolParam2}
		BooleanArrayItems: makeActionBuilder(
			app,
			newHandlerAdapterNoResponse[
				*BooleanBooleanArrayItemsRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanArrayItemsRequest,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanArrayItemsRequest,
				void,
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
				*BooleanBooleanNullableRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanNullableRequest,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanNullableRequest,
				void,
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
				*BooleanBooleanNullableArrayItemsRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanNullableArrayItemsRequest,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanNullableArrayItemsRequest,
				void,
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
				*BooleanBooleanParsingRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanParsingRequest,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanParsingRequest,
				void,
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
				*BooleanBooleanRequiredValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*BooleanBooleanRequiredValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*BooleanBooleanRequiredValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserBooleanBooleanRequiredValidation(app),
			},
		),
	}
}

type BooleanController interface {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsRequest,
	//
	// Response type: none
	BooleanArrayItems(NoResponseHandlerBuilder[
		*BooleanBooleanArrayItemsRequest,
	]) http.Handler

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableRequest,
	//
	// Response type: none
	BooleanNullable(NoResponseHandlerBuilder[
		*BooleanBooleanNullableRequest,
	]) http.Handler

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsRequest,
	//
	// Response type: none
	BooleanNullableArrayItems(NoResponseHandlerBuilder[
		*BooleanBooleanNullableArrayItemsRequest,
	]) http.Handler

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	BooleanParsing(NoResponseHandlerBuilder[
		*BooleanBooleanParsingRequest,
	]) http.Handler

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	BooleanRequiredValidation(NoResponseHandlerBuilder[
		*BooleanBooleanRequiredValidationRequest,
	]) http.Handler
}

func RegisterBooleanRoutesV3(controller BooleanController, app *HTTPApp) {
	builder := newBooleanControllerBuilder(app)
	app.router.HandleRoute("POST", "/boolean/array-items/{boolParam1}/{boolParam2}", controller.BooleanArrayItems(builder.BooleanArrayItems))
	app.router.HandleRoute("POST", "/boolean/nullable/{boolParam1}/{boolParam2}", controller.BooleanNullable(builder.BooleanNullable))
	app.router.HandleRoute("POST", "/boolean/nullable-array-items/{boolParam1}/{boolParam2}", controller.BooleanNullableArrayItems(builder.BooleanNullableArrayItems))
	app.router.HandleRoute("POST", "/boolean/parsing/{boolParam1}/{boolParam2}", controller.BooleanParsing(builder.BooleanParsing))
	app.router.HandleRoute("POST", "/boolean/required-validation", controller.BooleanRequiredValidation(builder.BooleanRequiredValidation))
}