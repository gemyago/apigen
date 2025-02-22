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

type arraysControllerBuilder struct {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	ArraysNullableRequiredValidation genericHandlerBuilder[
		*ArraysArraysNullableRequiredValidationRequest,
		void,
		func(context.Context, *ArraysArraysNullableRequiredValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *ArraysArraysNullableRequiredValidationRequest) (error),
	]

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRangeValidationRequest,
	//
	// Response type: none
	ArraysRangeValidation genericHandlerBuilder[
		*ArraysArraysRangeValidationRequest,
		void,
		func(context.Context, *ArraysArraysRangeValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *ArraysArraysRangeValidationRequest) (error),
	]

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRequiredValidationRequest,
	//
	// Response type: none
	ArraysRequiredValidation genericHandlerBuilder[
		*ArraysArraysRequiredValidationRequest,
		void,
		func(context.Context, *ArraysArraysRequiredValidationRequest) (error),
		func(http.ResponseWriter, *http.Request, *ArraysArraysRequiredValidationRequest) (error),
	]
}

func newArraysControllerBuilder(app *RootHandler) *arraysControllerBuilder {
	return &arraysControllerBuilder{
		// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
		ArraysNullableRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ArraysArraysNullableRequiredValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ArraysArraysNullableRequiredValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*ArraysArraysNullableRequiredValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysNullableRequiredValidation(app),
			},
		),

		// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
		ArraysRangeValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ArraysArraysRangeValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ArraysArraysRangeValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*ArraysArraysRangeValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysRangeValidation(app),
			},
		),

		// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
		ArraysRequiredValidation: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ArraysArraysRequiredValidationRequest,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ArraysArraysRequiredValidationRequest,
				void,
			](),
			makeActionBuilderParams[
				*ArraysArraysRequiredValidationRequest,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserArraysArraysRequiredValidation(app),
			},
		),
	}
}

type ArraysController interface {
	// POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysNullableRequiredValidationRequest,
	//
	// Response type: none
	ArraysNullableRequiredValidation(NoResponseHandlerBuilder[
		*ArraysArraysNullableRequiredValidationRequest,
	]) http.Handler

	// POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRangeValidationRequest,
	//
	// Response type: none
	ArraysRangeValidation(NoResponseHandlerBuilder[
		*ArraysArraysRangeValidationRequest,
	]) http.Handler

	// POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
	//
	// Request type: ArraysArraysRequiredValidationRequest,
	//
	// Response type: none
	ArraysRequiredValidation(NoResponseHandlerBuilder[
		*ArraysArraysRequiredValidationRequest,
	]) http.Handler
}

// RegisterArraysRoutes will attach the following routes to the root handler:
// 
// - POST /arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}
// 
// - POST /arrays/range-validation/{simpleItems1}/{simpleItems2}
// 
// - POST /arrays/required-validation/{simpleItems1}/{simpleItems2}
// 
// Routes will use provided controller to handle requests.
func RegisterArraysRoutes(rootHandler *RootHandler, controller ArraysController) {
	builder := newArraysControllerBuilder(rootHandler)
	rootHandler.router.HandleRoute("POST", "/arrays/nullable-required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysNullableRequiredValidation(builder.ArraysNullableRequiredValidation))
	rootHandler.router.HandleRoute("POST", "/arrays/range-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRangeValidation(builder.ArraysRangeValidation))
	rootHandler.router.HandleRoute("POST", "/arrays/required-validation/{simpleItems1}/{simpleItems2}", controller.ArraysRequiredValidation(builder.ArraysRequiredValidation))
}