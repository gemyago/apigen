package handlers

import (
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

type booleanControllerBuilder struct {
	// POST /boolean/array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanArrayItemsRequest,
	//
	// Response type: none
	BooleanArrayItems genericHandlerBuilder[
		*BooleanBooleanArrayItemsRequest,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanArrayItemsRequest, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanArrayItemsRequest, void],
	]

	// POST /boolean/nullable/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableRequest,
	//
	// Response type: none
	BooleanNullable genericHandlerBuilder[
		*BooleanBooleanNullableRequest,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanNullableRequest, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanNullableRequest, void],
	]

	// POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanNullableArrayItemsRequest,
	//
	// Response type: none
	BooleanNullableArrayItems genericHandlerBuilder[
		*BooleanBooleanNullableArrayItemsRequest,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanNullableArrayItemsRequest, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanNullableArrayItemsRequest, void],
	]

	// POST /boolean/parsing/{boolParam1}/{boolParam2}
	//
	// Request type: BooleanBooleanParsingRequest,
	//
	// Response type: none
	BooleanParsing genericHandlerBuilder[
		*BooleanBooleanParsingRequest,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanParsingRequest, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanParsingRequest, void],
	]

	// POST /boolean/required-validation
	//
	// Request type: BooleanBooleanRequiredValidationRequest,
	//
	// Response type: none
	BooleanRequiredValidation genericHandlerBuilder[
		*BooleanBooleanRequiredValidationRequest,
		void,
		handlerActionFuncNoResponse[*BooleanBooleanRequiredValidationRequest, void],
		httpHandlerActionFuncNoResponse[*BooleanBooleanRequiredValidationRequest, void],
	]
}

func newBooleanControllerBuilder(app *RootHandler) *booleanControllerBuilder {
	return &booleanControllerBuilder{
		// POST /boolean/array-items/{boolParam1}/{boolParam2}
		BooleanArrayItems: newGenericHandlerBuilder(
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
		BooleanNullable: newGenericHandlerBuilder(
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
		BooleanNullableArrayItems: newGenericHandlerBuilder(
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
		BooleanParsing: newGenericHandlerBuilder(
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
		BooleanRequiredValidation: newGenericHandlerBuilder(
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

// RegisterBooleanRoutes will attach the following routes to the root handler:
// 
// - POST /boolean/array-items/{boolParam1}/{boolParam2}
// 
// - POST /boolean/nullable/{boolParam1}/{boolParam2}
// 
// - POST /boolean/nullable-array-items/{boolParam1}/{boolParam2}
// 
// - POST /boolean/parsing/{boolParam1}/{boolParam2}
// 
// - POST /boolean/required-validation
// 
// Routes will use provided controller to handle requests.
func(rootHandler *RootHandler) RegisterBooleanRoutes(controller BooleanController) *RootHandler {
	builder := newBooleanControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/boolean/array-items/{boolParam1}/{boolParam2}", controller.BooleanArrayItems(builder.BooleanArrayItems))
	rootHandler.router.HandleRoute("POST", "/boolean/nullable/{boolParam1}/{boolParam2}", controller.BooleanNullable(builder.BooleanNullable))
	rootHandler.router.HandleRoute("POST", "/boolean/nullable-array-items/{boolParam1}/{boolParam2}", controller.BooleanNullableArrayItems(builder.BooleanNullableArrayItems))
	rootHandler.router.HandleRoute("POST", "/boolean/parsing/{boolParam1}/{boolParam2}", controller.BooleanParsing(builder.BooleanParsing))
	rootHandler.router.HandleRoute("POST", "/boolean/required-validation", controller.BooleanRequiredValidation(builder.BooleanRequiredValidation))
	return rootHandler
}