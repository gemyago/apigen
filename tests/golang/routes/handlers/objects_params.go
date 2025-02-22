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

func (p *paramsParserObjectsObjectsArrayBodyDirect) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayBodyDirectRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsArrayBodyDirectRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyDirect(rootHandler *RootHandler) paramsParser[*ObjectsObjectsArrayBodyDirectRequest] {
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

func (p *paramsParserObjectsObjectsArrayBodyNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayBodyNestedRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsArrayBodyNestedRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyNested(rootHandler *RootHandler) paramsParser[*ObjectsObjectsArrayBodyNestedRequest] {
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

func (p *paramsParserObjectsObjectsDeeplyNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsDeeplyNestedRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsDeeplyNestedRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsDeeplyNested(rootHandler *RootHandler) paramsParser[*ObjectsObjectsDeeplyNestedRequest] {
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

func (p *paramsParserObjectsObjectsNullableOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableOptionalBodyRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsNullableOptionalBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableOptionalBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsNullableOptionalBodyRequest] {
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

func (p *paramsParserObjectsObjectsNullableRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableRequiredBodyRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsNullableRequiredBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableRequiredBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsNullableRequiredBodyRequest] {
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

func (p *paramsParserObjectsObjectsOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsOptionalBodyRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsOptionalBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsOptionalBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsOptionalBodyRequest] {
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

func (p *paramsParserObjectsObjectsRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsRequiredBodyRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsRequiredBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredBody(rootHandler *RootHandler) paramsParser[*ObjectsObjectsRequiredBodyRequest] {
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

func (p *paramsParserObjectsObjectsRequiredNestedObjects) parse(router httpRouter, req *http.Request) (*ObjectsObjectsRequiredNestedObjectsRequest, error) {
	bindingCtx := BindingContext{}
	reqParams := &ObjectsObjectsRequiredNestedObjectsRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredNestedObjects(rootHandler *RootHandler) paramsParser[*ObjectsObjectsRequiredNestedObjectsRequest] {
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
