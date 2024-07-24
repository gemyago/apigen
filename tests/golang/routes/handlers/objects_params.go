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

type paramsParserObjectsObjectsArrayParsingBodyDirect struct {
	bindPayload requestParamBinder[*http.Request, []*models.ObjectArraysSimpleObject]
}

func (p *paramsParserObjectsObjectsArrayParsingBodyDirect) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayParsingBodyDirectRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsArrayParsingBodyDirectRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayParsingBodyDirect(app *HTTPApp) paramsParser[*ObjectsObjectsArrayParsingBodyDirectRequest] {
	return &paramsParserObjectsObjectsArrayParsingBodyDirect{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, []*models.ObjectArraysSimpleObject]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[[]*models.ObjectArraysSimpleObject],
			validateValue: internal.NewArrayValidator(
        internal.NewObjectArraysSimpleObjectValidator(),
      ),
		}),
	}
}

type paramsParserObjectsObjectsArrayParsingBodyNested struct {
	bindPayload requestParamBinder[*http.Request, *models.ObjectsArrayParsingBodyNestedRequest]
}

func (p *paramsParserObjectsObjectsArrayParsingBodyNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayParsingBodyNestedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsArrayParsingBodyNestedRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayParsingBodyNested(app *HTTPApp) paramsParser[*ObjectsObjectsArrayParsingBodyNestedRequest] {
	return &paramsParserObjectsObjectsArrayParsingBodyNested{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsArrayParsingBodyNestedRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.ObjectsArrayParsingBodyNestedRequest],
			validateValue: internal.NewObjectsArrayParsingBodyNestedRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsNested struct {
	bindPayload requestParamBinder[*http.Request, *models.ObjectsNestedRequest]
}

func (p *paramsParserObjectsObjectsNested) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNestedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsNestedRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNested(app *HTTPApp) paramsParser[*ObjectsObjectsNestedRequest] {
	return &paramsParserObjectsObjectsNested{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsNestedRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.ObjectsNestedRequest],
			validateValue: internal.NewObjectsNestedRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsNullableOptionalBody struct {
	bindPayload requestParamBinder[*http.Request, *models.SimpleObject]
}

func (p *paramsParserObjectsObjectsNullableOptionalBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableOptionalBodyRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsNullableOptionalBodyRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableOptionalBody(app *HTTPApp) paramsParser[*ObjectsObjectsNullableOptionalBodyRequest] {
	return &paramsParserObjectsObjectsNullableOptionalBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObject]{
			field: "payload",
			location: "body",
			required: false,
			parseValue: parseJSONPayload[*models.SimpleObject],
			validateValue: internal.NewSimpleObjectValidator(),
		}),
	}
}

type paramsParserObjectsObjectsNullableRequiredBody struct {
	bindPayload requestParamBinder[*http.Request, *models.SimpleObject]
}

func (p *paramsParserObjectsObjectsNullableRequiredBody) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableRequiredBodyRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsNullableRequiredBodyRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableRequiredBody(app *HTTPApp) paramsParser[*ObjectsObjectsNullableRequiredBodyRequest] {
	return &paramsParserObjectsObjectsNullableRequiredBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObject]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.SimpleObject],
			validateValue: internal.NewSimpleObjectValidator(),
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
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsOptionalBody(app *HTTPApp) paramsParser[*ObjectsObjectsOptionalBodyRequest] {
	return &paramsParserObjectsObjectsOptionalBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObject]{
			field: "payload",
			location: "body",
			required: false,
			parseValue: parseJSONPayload[*models.SimpleObject],
			validateValue: internal.NewSimpleObjectValidator(),
		}),
	}
}

type paramsParserObjectsObjectsOptionalNestedObjects struct {
	bindPayload requestParamBinder[*http.Request, *models.ObjectsOptionalNestedObjectsRequest]
}

func (p *paramsParserObjectsObjectsOptionalNestedObjects) parse(router httpRouter, req *http.Request) (*ObjectsObjectsOptionalNestedObjectsRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsOptionalNestedObjectsRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsOptionalNestedObjects(app *HTTPApp) paramsParser[*ObjectsObjectsOptionalNestedObjectsRequest] {
	return &paramsParserObjectsObjectsOptionalNestedObjects{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsOptionalNestedObjectsRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.ObjectsOptionalNestedObjectsRequest],
			validateValue: internal.NewObjectsOptionalNestedObjectsRequestValidator(),
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
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredBody(app *HTTPApp) paramsParser[*ObjectsObjectsRequiredBodyRequest] {
	return &paramsParserObjectsObjectsRequiredBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObject]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.SimpleObject],
			validateValue: internal.NewSimpleObjectValidator(),
		}),
	}
}

type paramsParserObjectsObjectsRequiredNestedObjects struct {
	bindPayload requestParamBinder[*http.Request, *models.ObjectsRequiredNestedObjectsRequest]
}

func (p *paramsParserObjectsObjectsRequiredNestedObjects) parse(router httpRouter, req *http.Request) (*ObjectsObjectsRequiredNestedObjectsRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsRequiredNestedObjectsRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredNestedObjects(app *HTTPApp) paramsParser[*ObjectsObjectsRequiredNestedObjectsRequest] {
	return &paramsParserObjectsObjectsRequiredNestedObjects{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsRequiredNestedObjectsRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.ObjectsRequiredNestedObjectsRequest],
			validateValue: internal.NewObjectsRequiredNestedObjectsRequestValidator(),
		}),
	}
}
