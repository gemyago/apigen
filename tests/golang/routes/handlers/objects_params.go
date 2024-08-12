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
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[[]*models.ObjectArraysSimpleObject]),
			validateValue: internal.NewArrayValidator(
				internal.NewObjectArraysSimpleObjectValidator(internal.ModelValidatorParams{Location: "body"}),
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
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.ObjectsArrayParsingBodyNestedRequest]),
			validateValue: internal.NewObjectsArrayParsingBodyNestedRequestValidator(internal.ModelValidatorParams{Location: "body"}),
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
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsDeeplyNested(app *HTTPApp) paramsParser[*ObjectsObjectsDeeplyNestedRequest] {
	return &paramsParserObjectsObjectsDeeplyNested{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.ObjectsDeeplyNestedRequest]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.ObjectsDeeplyNestedRequest]),
			validateValue: internal.NewObjectsDeeplyNestedRequestValidator(internal.ModelValidatorParams{Location: "body"}),
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
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableOptionalBody(app *HTTPApp) paramsParser[*ObjectsObjectsNullableOptionalBodyRequest] {
	return &paramsParserObjectsObjectsNullableOptionalBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleNullableObject]{
			field: "payload",
			location: "body",
			required: false,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.SimpleNullableObject]),
			validateValue: internal.SkipNullFieldValidator(internal.NewSimpleNullableObjectValidator(internal.ModelValidatorParams{Location: "body"})),
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
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsNullableRequiredBody(app *HTTPApp) paramsParser[*ObjectsObjectsNullableRequiredBodyRequest] {
	return &paramsParserObjectsObjectsNullableRequiredBody{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleNullableObject]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.SimpleNullableObject]),
			validateValue: internal.SkipNullFieldValidator(internal.NewSimpleNullableObjectValidator(internal.ModelValidatorParams{Location: "body"})),
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
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.SimpleObject]),
			validateValue: internal.NewSimpleObjectValidator(internal.ModelValidatorParams{Location: "body"}),
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
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.SimpleObject]),
			validateValue: internal.NewSimpleObjectValidator(internal.ModelValidatorParams{Location: "body"}),
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
	p.bindPayload(&bindingCtx, readRequestBodyValue(req), &reqParams.Payload)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserObjectsObjectsRequiredNestedObjects(app *HTTPApp) paramsParser[*ObjectsObjectsRequiredNestedObjectsRequest] {
	return &paramsParserObjectsObjectsRequiredNestedObjects{
		bindPayload: newRequestParamBinder(binderParams[*http.Request, *models.SimpleObjectsContainer]{
			field: "payload",
			location: "body",
			required: true,
			parseValue: parseSingleValueParamAsSingleValue(parseJSONPayload[*models.SimpleObjectsContainer]),
			validateValue: internal.NewSimpleObjectsContainerValidator(internal.ModelValidatorParams{Location: "body"}),
		}),
	}
}
