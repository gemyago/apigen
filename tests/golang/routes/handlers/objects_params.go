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
	bindPayload requestParamBinder[*http.Request, *models.ObjectsArrayParsingBodyDirectRequest]
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
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsArrayParsingBodyDirectRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.ObjectsArrayParsingBodyDirectRequest],
			validateValue: internal.NewObjectsArrayParsingBodyDirectRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsArrayParsingBodyDirect_0 struct {
	bindPayload requestParamBinder[*http.Request, *models.[]ObjectArraysSimpleObject]
}

func (p *paramsParserObjectsObjectsArrayParsingBodyDirect_0) parse(router httpRouter, req *http.Request) (*ObjectsObjectsArrayParsingBodyDirect0Request, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsArrayParsingBodyDirect0Request{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsArrayParsingBodyDirect_0(app *HTTPApp) paramsParser[*ObjectsObjectsArrayParsingBodyDirect0Request] {
	return &paramsParserObjectsObjectsArrayParsingBodyDirect_0{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.[]ObjectArraysSimpleObject]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseJSONPayload[*models.[]ObjectArraysSimpleObject],
			validateValue: internal.New[]ObjectArraysSimpleObjectValidator(),
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
