package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = models.BehaviorNoParamsWithResponse202Response{}

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
    // isNullable: false
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


