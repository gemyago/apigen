package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.ObjectArraysSimpleObject{}

type paramsParserObjectsObjectsArrayBodyDirect struct {
	bindPayload requestParamBinder[*http.Request, []*models.ObjectArraysSimpleObject]
}

func (p *paramsParserObjectsObjectsArrayBodyDirect) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayBodyDirectRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsArrayBodyDirectRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyDirect(app *HTTPApp) paramsParser[*ObjectsObjectsArrayBodyDirectRequest] {
	return &paramsParserObjectsObjectsArrayBodyDirect{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, []*models.ObjectArraysSimpleObject]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[[]*models.ObjectArraysSimpleObject],
			),
			validateValue: internal.NewArrayValidator(
				internal.NewObjectArraysSimpleObjectValidator(),
			),
		}),
	}
}

type paramsParserObjectsObjectsArrayBodyNested struct {
	bindPayload requestParamBinder[*http.Request, *models.ObjectsArrayBodyNestedRequest]
}

func (p *paramsParserObjectsObjectsArrayBodyNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayBodyNestedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsArrayBodyNestedRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayBodyNested(app *HTTPApp) paramsParser[*ObjectsObjectsArrayBodyNestedRequest] {
	return &paramsParserObjectsObjectsArrayBodyNested{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsArrayBodyNestedRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.ObjectsArrayBodyNestedRequest],
			),
			validateValue: internal.NewObjectsArrayBodyNestedRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsDeeplyNested struct {
	bindPayload requestParamBinder[*http.Request, *models.ObjectsDeeplyNestedRequest]
}

func (p *paramsParserObjectsObjectsDeeplyNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsDeeplyNestedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsDeeplyNestedRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsDeeplyNested(app *HTTPApp) paramsParser[*ObjectsObjectsDeeplyNestedRequest] {
	return &paramsParserObjectsObjectsDeeplyNested{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsDeeplyNestedRequest]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.ObjectsDeeplyNestedRequest],
			),
			validateValue: internal.NewObjectsDeeplyNestedRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsNullableOptionalBody struct {
	bindPayload requestParamBinder[*http.Request, *models.SimpleNullableObject]
}

func (p *paramsParserObjectsObjectsNullableOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableOptionalBodyRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsNullableOptionalBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableOptionalBody(app *HTTPApp) paramsParser[*ObjectsObjectsNullableOptionalBodyRequest] {
	return &paramsParserObjectsObjectsNullableOptionalBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleNullableObject]{
			required: false,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.SimpleNullableObject],
			),
			validateValue: internal.SkipNullFieldValidator(internal.NewSimpleNullableObjectValidator()),
		}),
	}
}

type paramsParserObjectsObjectsNullableRequiredBody struct {
	bindPayload requestParamBinder[*http.Request, *models.SimpleNullableObject]
}

func (p *paramsParserObjectsObjectsNullableRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableRequiredBodyRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsNullableRequiredBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableRequiredBody(app *HTTPApp) paramsParser[*ObjectsObjectsNullableRequiredBodyRequest] {
	return &paramsParserObjectsObjectsNullableRequiredBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleNullableObject]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.SimpleNullableObject],
			),
			validateValue: internal.SkipNullFieldValidator(internal.NewSimpleNullableObjectValidator()),
		}),
	}
}

type paramsParserObjectsObjectsOptionalBody struct {
	bindPayload requestParamBinder[*http.Request, *models.SimpleObject]
}

func (p *paramsParserObjectsObjectsOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsOptionalBodyRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsOptionalBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsOptionalBody(app *HTTPApp) paramsParser[*ObjectsObjectsOptionalBodyRequest] {
	return &paramsParserObjectsObjectsOptionalBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObject]{
			required: false,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.SimpleObject],
			),
			validateValue: internal.NewSimpleObjectValidator(),
		}),
	}
}

type paramsParserObjectsObjectsRequiredBody struct {
	bindPayload requestParamBinder[*http.Request, *models.SimpleObject]
}

func (p *paramsParserObjectsObjectsRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsRequiredBodyRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsRequiredBodyRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredBody(app *HTTPApp) paramsParser[*ObjectsObjectsRequiredBodyRequest] {
	return &paramsParserObjectsObjectsRequiredBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObject]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.SimpleObject],
			),
			validateValue: internal.NewSimpleObjectValidator(),
		}),
	}
}

type paramsParserObjectsObjectsRequiredNestedObjects struct {
	bindPayload requestParamBinder[*http.Request, *models.SimpleObjectsContainer]
}

func (p *paramsParserObjectsObjectsRequiredNestedObjects) parse(router httpRouter, req *http.Request) (*ObjectsObjectsRequiredNestedObjectsRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsRequiredNestedObjectsRequest{}
	// body params
	p.bindPayload(bindingCtx.Fork("body"), readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredNestedObjects(app *HTTPApp) paramsParser[*ObjectsObjectsRequiredNestedObjectsRequest] {
	return &paramsParserObjectsObjectsRequiredNestedObjects{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObjectsContainer]{
			required: true,
			parseValue: parseSoloValueParamAsSoloValue(
				parseJSONPayload[*models.SimpleObjectsContainer],
			),
			validateValue: internal.NewSimpleObjectsContainerValidator(),
		}),
	}
}
