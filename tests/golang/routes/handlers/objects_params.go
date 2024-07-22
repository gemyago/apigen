package handlers

import (
	"net/http"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
	"github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports.
var _ = time.Time{}

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
