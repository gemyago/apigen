package handlers

import (
	"net/http"
	"time"

	. "github.com/gemyago/apigen/tests/golang/routes/models"
	. "github.com/gemyago/apigen/tests/golang/routes/internal"
)

// Below is to workaround unused imports if that happens.
var _ = BindingContext{}
var _ = http.MethodGet
var _ = time.Time{}
type _ func() Error



type behaviorNoParamsNoResponseIsolatedControllerBuilder struct {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse genericHandlerBuilder[
		void,
		void,
		handlerActionFuncNoParamsNoResponse[void, void],
		httpHandlerActionFuncNoParamsNoResponse[void, void],
	]
}

func newBehaviorNoParamsNoResponseIsolatedControllerBuilder(app *RootHandler) *behaviorNoParamsNoResponseIsolatedControllerBuilder {
	return &behaviorNoParamsNoResponseIsolatedControllerBuilder{
		// GET /behavior/no-params-no-response-isolated
		BehaviorNoParamsNoResponse: newGenericHandlerBuilder(
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
