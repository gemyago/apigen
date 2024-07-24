package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/internal"
	"github.com/gemyago/apigen/tests/golang/routes/models"
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
			field:      "payload",
			location:   "body",
			required:   true,
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
			field:         "payload",
			location:      "body",
			required:      true,
			parseValue:    parseJSONPayload[*models.ObjectsArrayParsingBodyNestedRequest],
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
			field:         "payload",
			location:      "body",
			required:      true,
			parseValue:    parseJSONPayload[*models.ObjectsNestedRequest],
			validateValue: internal.NewObjectsNestedRequestValidator(),
		}),
	}
}

type paramsParserObjectsObjectsNullable struct {
	bindPayload requestParamBinder[*http.Request, *models.ObjectNullableSimpleObjectsContainer]
}

func (p *paramsParserObjectsObjectsNullable) parse(router httpRouter, req *http.Request) (*ObjectsObjectsNullableRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ObjectsObjectsNullableRequest{}
	// body params
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullable(app *HTTPApp) paramsParser[*ObjectsObjectsNullableRequest] {
	return &paramsParserObjectsObjectsNullable{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectNullableSimpleObjectsContainer]{
			field:         "payload",
			location:      "body",
			required:      true,
			parseValue:    parseJSONPayload[*models.ObjectNullableSimpleObjectsContainer],
			validateValue: internal.NewObjectNullableSimpleObjectsContainerValidator(),
		}),
	}
}
