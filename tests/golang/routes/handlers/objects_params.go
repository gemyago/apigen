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

func (p *paramsParserObjectsObjectsArrayBodyDirect) parse(router httpRouter, req *http.Request) (*ObjectsArrayBodyDirectParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsArrayBodyDirectParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyDirect(rootHandler *RootHandler) paramsParser[*ObjectsArrayBodyDirectParams] {
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

func (p *paramsParserObjectsObjectsArrayBodyNested) parse(router httpRouter, req *http.Request) (*ObjectsArrayBodyNestedParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsArrayBodyNestedParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyNested(rootHandler *RootHandler) paramsParser[*ObjectsArrayBodyNestedParams] {
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

func (p *paramsParserObjectsObjectsDeeplyNested) parse(router httpRouter, req *http.Request) (*ObjectsDeeplyNestedParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsDeeplyNestedParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsDeeplyNested(rootHandler *RootHandler) paramsParser[*ObjectsDeeplyNestedParams] {
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

func (p *paramsParserObjectsObjectsNullableOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsNullableOptionalBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsNullableOptionalBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableOptionalBody(rootHandler *RootHandler) paramsParser[*ObjectsNullableOptionalBodyParams] {
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

func (p *paramsParserObjectsObjectsNullableRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsNullableRequiredBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsNullableRequiredBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableRequiredBody(rootHandler *RootHandler) paramsParser[*ObjectsNullableRequiredBodyParams] {
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

func (p *paramsParserObjectsObjectsOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsOptionalBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsOptionalBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsOptionalBody(rootHandler *RootHandler) paramsParser[*ObjectsOptionalBodyParams] {
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

func (p *paramsParserObjectsObjectsRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsRequiredBodyParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsRequiredBodyParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredBody(rootHandler *RootHandler) paramsParser[*ObjectsRequiredBodyParams] {
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

func (p *paramsParserObjectsObjectsRequiredNestedObjects) parse(router httpRouter, req *http.Request) (*ObjectsRequiredNestedObjectsParams, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsRequiredNestedObjectsParams{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredNestedObjects(rootHandler *RootHandler) paramsParser[*ObjectsRequiredNestedObjectsParams] {
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
	// Request type: ObjectsArrayBodyDirectParams,
	//
	// Response type: none
	ObjectsArrayBodyDirect genericHandlerBuilder[
		*ObjectsArrayBodyDirectParams,
		void,
		handlerActionFuncNoResponse[*ObjectsArrayBodyDirectParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsArrayBodyDirectParams, void],
	]

	// PUT /objects/arrays
	//
	// Request type: ObjectsArrayBodyNestedParams,
	//
	// Response type: none
	ObjectsArrayBodyNested genericHandlerBuilder[
		*ObjectsArrayBodyNestedParams,
		void,
		handlerActionFuncNoResponse[*ObjectsArrayBodyNestedParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsArrayBodyNestedParams, void],
	]

	// POST /objects/deeply-nested
	//
	// Request type: ObjectsDeeplyNestedParams,
	//
	// Response type: none
	ObjectsDeeplyNested genericHandlerBuilder[
		*ObjectsDeeplyNestedParams,
		void,
		handlerActionFuncNoResponse[*ObjectsDeeplyNestedParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsDeeplyNestedParams, void],
	]

	// PUT /objects/nullable-body
	//
	// Request type: ObjectsNullableOptionalBodyParams,
	//
	// Response type: none
	ObjectsNullableOptionalBody genericHandlerBuilder[
		*ObjectsNullableOptionalBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsNullableOptionalBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsNullableOptionalBodyParams, void],
	]

	// POST /objects/nullable-body
	//
	// Request type: ObjectsNullableRequiredBodyParams,
	//
	// Response type: none
	ObjectsNullableRequiredBody genericHandlerBuilder[
		*ObjectsNullableRequiredBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsNullableRequiredBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsNullableRequiredBodyParams, void],
	]

	// PUT /objects/required-body
	//
	// Request type: ObjectsOptionalBodyParams,
	//
	// Response type: none
	ObjectsOptionalBody genericHandlerBuilder[
		*ObjectsOptionalBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsOptionalBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsOptionalBodyParams, void],
	]

	// POST /objects/required-body
	//
	// Request type: ObjectsRequiredBodyParams,
	//
	// Response type: none
	ObjectsRequiredBody genericHandlerBuilder[
		*ObjectsRequiredBodyParams,
		void,
		handlerActionFuncNoResponse[*ObjectsRequiredBodyParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsRequiredBodyParams, void],
	]

	// POST /objects/required-nested-objects
	//
	// Request type: ObjectsRequiredNestedObjectsParams,
	//
	// Response type: none
	ObjectsRequiredNestedObjects genericHandlerBuilder[
		*ObjectsRequiredNestedObjectsParams,
		void,
		handlerActionFuncNoResponse[*ObjectsRequiredNestedObjectsParams, void],
		httpHandlerActionFuncNoResponse[*ObjectsRequiredNestedObjectsParams, void],
	]
}

func newObjectsControllerBuilder(app *RootHandler) *objectsControllerBuilder {
	return &objectsControllerBuilder{
		// POST /objects/arrays
		ObjectsArrayBodyDirect: newGenericHandlerBuilder(
			app,
			newHandlerAdapterNoResponse[
				*ObjectsArrayBodyDirectParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsArrayBodyDirectParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsArrayBodyDirectParams,
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
				*ObjectsArrayBodyNestedParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsArrayBodyNestedParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsArrayBodyNestedParams,
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
				*ObjectsDeeplyNestedParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsDeeplyNestedParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsDeeplyNestedParams,
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
				*ObjectsNullableOptionalBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsNullableOptionalBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsNullableOptionalBodyParams,
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
				*ObjectsNullableRequiredBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsNullableRequiredBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsNullableRequiredBodyParams,
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
				*ObjectsOptionalBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsOptionalBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsOptionalBodyParams,
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
				*ObjectsRequiredBodyParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsRequiredBodyParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsRequiredBodyParams,
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
				*ObjectsRequiredNestedObjectsParams,
				void,
			](),
			newHTTPHandlerAdapterNoResponse[
				*ObjectsRequiredNestedObjectsParams,
				void,
			](),
			makeActionBuilderParams[
				*ObjectsRequiredNestedObjectsParams,
				void,
			]{
				defaultStatus: 204,
				voidResult:    true,
				paramsParser:  newParamsParserObjectsObjectsRequiredNestedObjects(app),
			},
		),
	}
}
