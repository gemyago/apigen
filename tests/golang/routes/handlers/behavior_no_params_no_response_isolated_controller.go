package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = http.MethodGet
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint
type _ func() Error



type BehaviorNoParamsNoResponseIsolatedControllerBuilder struct {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse ActionBuilder[
		void,
		void,
		func(context.Context) (error),
		func(http.ResponseWriter, *http.Request) (error),
	]
}

func newBehaviorNoParamsNoResponseIsolatedControllerBuilder(app *HTTPApp) *BehaviorNoParamsNoResponseIsolatedControllerBuilder {
	return &BehaviorNoParamsNoResponseIsolatedControllerBuilder{
		// GET /behavior/no-params-no-response-isolated
		BehaviorNoParamsNoResponse: makeActionBuilder(
			app,
			newHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			newHTTPHandlerAdapterNoParamsNoResponse[
				void,
				void,
			](),
			makeActionBuilderParams[
				void,
				void,
			]{
				defaultStatus: 202,
				voidResult:    true,
				paramsParser:  makeVoidParamsParser(app),
			},
		),
	}
}

type BehaviorNoParamsNoResponseIsolatedController interface {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse(NoParamsNoResponseHandlerBuilder) http.Handler
}

func RegisterBehaviorNoParamsNoResponseIsolatedRoutesV3(controller BehaviorNoParamsNoResponseIsolatedController, app *HTTPApp) {
	builder := newBehaviorNoParamsNoResponseIsolatedControllerBuilder(app)
	app.router.HandleRoute("GET", "/behavior/no-params-no-response-isolated", controller.BehaviorNoParamsNoResponse(builder.BehaviorNoParamsNoResponse))
}