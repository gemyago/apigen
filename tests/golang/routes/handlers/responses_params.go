package handlers

import (
	"net/http"
	"time"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}

type paramsParserResponsesResponsesNoStatusDefined struct {
}

func (p *paramsParserResponsesResponsesNoStatusDefined) parse(router httpRouter, req *http.Request) (*ResponsesResponsesNoStatusDefinedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ResponsesResponsesNoStatusDefinedRequest{}
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserResponsesResponsesNoStatusDefined(app *HTTPApp) paramsParser[*ResponsesResponsesNoStatusDefinedRequest] {
	return &paramsParserResponsesResponsesNoStatusDefined{
	}
}

type paramsParserResponsesResponsesWithStatusDefined struct {
}

func (p *paramsParserResponsesResponsesWithStatusDefined) parse(router httpRouter, req *http.Request) (*ResponsesResponsesWithStatusDefinedRequest, error) {
	bindingCtx := internal.BindingContext{}
	reqParams := &ResponsesResponsesWithStatusDefinedRequest{}
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserResponsesResponsesWithStatusDefined(app *HTTPApp) paramsParser[*ResponsesResponsesWithStatusDefinedRequest] {
	return &paramsParserResponsesResponsesWithStatusDefined{
	}
}
