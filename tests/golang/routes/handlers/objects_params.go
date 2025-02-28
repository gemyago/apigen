package handlers

import (
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
	. "github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports if that happens.
var _ = BindingContext{}
var _ = http.MethodGet
var _ = time.Time{}
type _ func() ObjectArraysSimpleObject

type paramsParserObjectsObjectsArrayBodyDirect struct {
	bindPayload requestParamBinder[*http.Request, []*ObjectArraysSimpleObject]
}

func (p *paramsParserObjectsObjectsArrayBodyDirect) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayBodyDirectParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsArrayBodyDirectParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyDirect(rootHandler *RootHandler) paramsParser[*ObjectsObjectsArrayBodyDirectParams] {
	return &paramsParserObjectsObjectsArrayBodyDirect{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, []*ObjectArraysSimpleObject]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[[]*ObjectArraysSimpleObject],
			),
			validateValue: NewArrayValidator(
				NewSimpleFieldValidator[[]*ObjectArraysSimpleObject](
				),
				NewObjectArraysSimpleObjectValidator(),
			),
		}),
	}
}

type paramsParserObjectsObjectsArrayBodyNested struct {
	bindPayload requestParamBinder[*http.Request, *ObjectsArrayBodyNestedRequest]
}

func (p *paramsParserObjectsObjectsArrayBodyNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayBodyNestedParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsArrayBodyNestedParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyNested(rootHandler *RootHandler) paramsParser[*ObjectsObjectsArrayBodyNestedParams] {
	return &paramsParserObjectsObjectsArrayBodyNested{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *ObjectsArrayBodyNestedRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*ObjectsArrayBodyNestedRequest],
			),
			validateValue: NewObjectsArrayBodyNestedRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsDeeplyNested struct {
	bindPayload requestParamBinder[*http.Request, *ObjectsDeeplyNestedRequest]
}

func (p *paramsParserObjectsObjectsDeeplyNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsDeeplyNestedParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsDeeplyNestedParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsDeeplyNested(rootHandler *RootHandler) paramsParser[*ObjectsObjectsDeeplyNestedParams] {
	return &paramsParserObjectsObjectsDeeplyNested{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *ObjectsDeeplyNestedRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*ObjectsDeeplyNestedRequest],
			),
			validateValue: NewObjectsDeeplyNestedRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsNullableOptionalBody struct {
	bindPayload requestParamBinder[*http.Request, *SimpleNullableObject]
}

func (p *paramsParserObjectsObjectsNullableOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableOptionalBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsNullableOptionalBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableOptionalBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsNullableOptionalBodyParams] {
	return &paramsParserObjectsObjectsNullableOptionalBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *SimpleNullableObject]{
			required: false,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*SimpleNullableObject],
			),
			validateValue: SkipNullFieldValidator(NewSimpleNullableObjectValidator()),
		}),
	}
}

type paramsParserObjectsObjectsNullableRequiredBody struct {
	bindPayload requestParamBinder[*http.Request, *SimpleNullableObject]
}

func (p *paramsParserObjectsObjectsNullableRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableRequiredBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsNullableRequiredBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableRequiredBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsNullableRequiredBodyParams] {
	return &paramsParserObjectsObjectsNullableRequiredBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *SimpleNullableObject]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*SimpleNullableObject],
			),
			validateValue: SkipNullFieldValidator(NewSimpleNullableObjectValidator()),
		}),
	}
}

type paramsParserObjectsObjectsOptionalBody struct {
	bindPayload requestParamBinder[*http.Request, *SimpleObject]
}

func (p *paramsParserObjectsObjectsOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsOptionalBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsOptionalBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsOptionalBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsOptionalBodyParams] {
	return &paramsParserObjectsObjectsOptionalBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *SimpleObject]{
			required: false,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*SimpleObject],
			),
			validateValue: NewSimpleObjectValidator(),
		}),
	}
}

type paramsParserObjectsObjectsRequiredBody struct {
	bindPayload requestParamBinder[*http.Request, *SimpleObject]
}

func (p *paramsParserObjectsObjectsRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsRequiredBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsRequiredBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsRequiredBodyParams] {
	return &paramsParserObjectsObjectsRequiredBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *SimpleObject]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*SimpleObject],
			),
			validateValue: NewSimpleObjectValidator(),
		}),
	}
}

type paramsParserObjectsObjectsRequiredNestedObjects struct {
	bindPayload requestParamBinder[*http.Request, *SimpleObjectsContainer]
}

func (p *paramsParserObjectsObjectsRequiredNestedObjects) parse(router httpRouter, req *http.Request) (*ObjectsObjectsRequiredNestedObjectsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsRequiredNestedObjectsParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredNestedObjects(rootHandler *RootHandler) paramsParser[*ObjectsObjectsRequiredNestedObjectsParams] {
	return &paramsParserObjectsObjectsRequiredNestedObjects{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *SimpleObjectsContainer]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*SimpleObjectsContainer],
			),
			validateValue: NewSimpleObjectsContainerValidator(),
		}),
	}
}

type objectsControllerBuilder struct {
	// POST /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyDirectParams,
	//
	// Response type: none
	ObjectsArrayBodyDirect genericHandlerBuilder[
		*ObjectsObjectsArrayBodyDirectParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsArrayBodyDirectParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsArrayBodyDirectParams, void],
	]

	// PUT /objects/arrays
	//
	// Request type: ObjectsObjectsArrayBodyNestedParams,
	//
	// Response type: none
	ObjectsArrayBodyNested genericHandlerBuilder[
		*ObjectsObjectsArrayBodyNestedParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsArrayBodyNestedParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsArrayBodyNestedParams, void],
	]

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsObjectsDeeplyNestedParams,
	//
	// Response type: none
	ObjectsDeeplyNested genericHandlerBuilder[
		*ObjectsObjectsDeeplyNestedParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsDeeplyNestedParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsDeeplyNestedParams, void],
	]

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableOptionalBodyParams,
	//
	// Response type: none
	ObjectsNullableOptionalBody genericHandlerBuilder[
		*ObjectsObjectsNullableOptionalBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsNullableOptionalBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsNullableOptionalBodyParams, void],
	]

	// POST /objects/nullable-body
	//
	// Request type: ObjectsObjectsNullableRequiredBodyParams,
	//
	// Response type: none
	ObjectsNullableRequiredBody genericHandlerBuilder[
		*ObjectsObjectsNullableRequiredBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsNullableRequiredBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsNullableRequiredBodyParams, void],
	]

	// PUT /objects/required-body
	//
	// Request type: ObjectsObjectsOptionalBodyParams,
	//
	// Response type: none
	ObjectsOptionalBody genericHandlerBuilder[
		*ObjectsObjectsOptionalBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsOptionalBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsOptionalBodyParams, void],
	]

	// POST /objects/required-body
	//
	// Request type: ObjectsObjectsRequiredBodyParams,
	//
	// Response type: none
	ObjectsRequiredBody genericHandlerBuilder[
		*ObjectsObjectsRequiredBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsRequiredBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsRequiredBodyParams, void],
	]

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsObjectsRequiredNestedObjectsParams,
	//
	// Response type: none
	ObjectsRequiredNestedObjects genericHandlerBuilder[
		*ObjectsObjectsRequiredNestedObjectsParams,
		void,
		handlerActionFuncNoResponse[*ObjectsObjectsRequiredNestedObjectsParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsObjectsRequiredNestedObjectsParams, void],
	]
}

func newObjectsControllerBuilder(app *RootHandler) *objectsControllerBuilder {
	return &objectsControllerBuilder{
		// POST /objects/arrays
		ObjectsArrayBodyDirect: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyDirectParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyDirectParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsArrayBodyDirectParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsArrayBodyDirect(app),
			},
		),

		// PUT /objects/arrays
		ObjectsArrayBodyNested: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyNestedParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsArrayBodyNestedParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsArrayBodyNestedParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsArrayBodyNested(app),
			},
		),

		// POST /objects/deeply-nested
		ObjectsDeeplyNested: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsDeeplyNestedParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsDeeplyNestedParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsDeeplyNestedParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsDeeplyNested(app),
			},
		),

		// PUT /objects/nullable-body
		ObjectsNullableOptionalBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsNullableOptionalBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsNullableOptionalBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsNullableOptionalBodyParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsNullableOptionalBody(app),
			},
		),

		// POST /objects/nullable-body
		ObjectsNullableRequiredBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsNullableRequiredBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsNullableRequiredBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsNullableRequiredBodyParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsNullableRequiredBody(app),
			},
		),

		// PUT /objects/required-body
		ObjectsOptionalBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsOptionalBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsOptionalBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsOptionalBodyParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsOptionalBody(app),
			},
		),

		// POST /objects/required-body
		ObjectsRequiredBody: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsRequiredBodyParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsRequiredBody(app),
			},
		),

		// POST /objects/required-nested-objects
		ObjectsRequiredNestedObjects: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredNestedObjectsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsObjectsRequiredNestedObjectsParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsObjectsRequiredNestedObjectsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsRequiredNestedObjects(app),
			},
		),
	}
}
