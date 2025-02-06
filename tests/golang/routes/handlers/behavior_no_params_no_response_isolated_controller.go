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



type BehaviorNoParamsNoResponseIsolatedControllerBuilderV2 struct {
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

func newBehaviorNoParamsNoResponseIsolatedControllerBuilderV2(app *HTTPApp) *BehaviorNoParamsNoResponseIsolatedControllerBuilderV2 {
	return &BehaviorNoParamsNoResponseIsolatedControllerBuilderV2{
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
				paramsParser:  makeVoidParamsParserV2(app),
			},
		),
	}
}

type BehaviorNoParamsNoResponseIsolatedControllerV2 interface {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse(b *BehaviorNoParamsNoResponseIsolatedControllerBuilderV2) http.Handler
}

func RegisterBehaviorNoParamsNoResponseIsolatedRoutesV2(controller BehaviorNoParamsNoResponseIsolatedControllerV2, app *HTTPApp) {
	builder := newBehaviorNoParamsNoResponseIsolatedControllerBuilderV2(app)
	app.router.HandleRoute("GET", "/behavior/no-params-no-response-isolated", controller.BehaviorNoParamsNoResponse(builder))
}