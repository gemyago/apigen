package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type paramsParserBehaviorBehaviorNoParamsNoResponse struct {
}

func (p *paramsParserBehaviorBehaviorNoParamsNoResponse) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorNoParamsNoResponseRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &BehaviorBehaviorNoParamsNoResponseRequest{}
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorNoParamsNoResponse(app *HTTPApp) paramsParser[*BehaviorBehaviorNoParamsNoResponseRequest] {
	return &paramsParserBehaviorBehaviorNoParamsNoResponse{
	}
}

type paramsParserBehaviorBehaviorNoParamsWithResponse struct {
}

func (p *paramsParserBehaviorBehaviorNoParamsWithResponse) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorNoParamsWithResponseRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &BehaviorBehaviorNoParamsWithResponseRequest{}
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorNoParamsWithResponse(app *HTTPApp) paramsParser[*BehaviorBehaviorNoParamsWithResponseRequest] {
	return &paramsParserBehaviorBehaviorNoParamsWithResponse{
	}
}

type paramsParserBehaviorBehaviorNoStatusDefined struct {
}

func (p *paramsParserBehaviorBehaviorNoStatusDefined) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorNoStatusDefinedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &BehaviorBehaviorNoStatusDefinedRequest{}
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorNoStatusDefined(app *HTTPApp) paramsParser[*BehaviorBehaviorNoStatusDefinedRequest] {
	return &paramsParserBehaviorBehaviorNoStatusDefined{
	}
}

type paramsParserBehaviorBehaviorWithParamsAndResponse struct {
	bindQueryParam1 requestParamBinder[[]string, string]
}

func (p *paramsParserBehaviorBehaviorWithParamsAndResponse) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorWithParamsAndResponseRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &BehaviorBehaviorWithParamsAndResponseRequest{}
	// query params
	query := req.URL.Query()
	p.bindQueryParam1(&bindingCtx, readQueryValue("queryParam1", query), &reqParams.QueryParam1)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorWithParamsAndResponse(app *HTTPApp) paramsParser[*BehaviorBehaviorWithParamsAndResponseRequest] {
	return &paramsParserBehaviorBehaviorWithParamsAndResponse{
		bindQueryParam1: newRequestParamBinder(binderParams[[]string, string]{
			field: "queryParam1",
			location: "query",
			required: false,
			parseValue: app.knownParsers.stringInQuery,
			validateValue: internal.NewSimpleFieldValidator[string](
			),
		}),
	}
}

type paramsParserBehaviorBehaviorWithStatusDefined struct {
}

func (p *paramsParserBehaviorBehaviorWithStatusDefined) parse(router httpRouter, req *http.Request) (*BehaviorBehaviorWithStatusDefinedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &BehaviorBehaviorWithStatusDefinedRequest{}
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorWithStatusDefined(app *HTTPApp) paramsParser[*BehaviorBehaviorWithStatusDefinedRequest] {
	return &paramsParserBehaviorBehaviorWithStatusDefined{
	}
}
