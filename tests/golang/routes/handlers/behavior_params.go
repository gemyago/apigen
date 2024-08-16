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
	queryParamsCtx := bindingCtx.Fork("query")
	p.bindQueryParam1(queryParamsCtx.Fork("queryParam1"), readQueryValue("queryParam1", query), &reqParams.QueryParam1)
	return reqParams, bindingCtx.AggregatedError()
}

func newParamsParserBehaviorBehaviorWithParamsAndResponse(app *HTTPApp) paramsParser[*BehaviorBehaviorWithParamsAndResponseRequest] {
	return &paramsParserBehaviorBehaviorWithParamsAndResponse{
		bindQueryParam1: newRequestParamBinder(binderParams[[]string, string]{
			required: false,
			parseValue: parseMultiValueParamAsSoloValue(
				app.knownParsers.stringParser,
			),
			validateValue: internal.NewSimpleFieldValidator[string](
				internal.SimpleFieldValidatorParams{Field: "queryParam1", Location: "query"},
			),
		}),
	}
}


