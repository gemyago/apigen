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



type BehaviorNoParamsNoResponseIsolatedController struct {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse httpHandlerFactory
}

type BehaviorNoParamsNoResponseIsolatedControllerBuilder struct {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorNoParamsNoResponse actionBuilderNoParamsVoidResult[*BehaviorNoParamsNoResponseIsolatedControllerBuilder]
}

func (c *BehaviorNoParamsNoResponseIsolatedControllerBuilder) Finalize() *BehaviorNoParamsNoResponseIsolatedController {
	return &BehaviorNoParamsNoResponseIsolatedController{
		BehaviorNoParamsNoResponse: mustInitializeAction("behaviorNoParamsNoResponse", c.HandleBehaviorNoParamsNoResponse.httpHandlerFactory),
	}
}

func BuildBehaviorNoParamsNoResponseIsolatedController() *BehaviorNoParamsNoResponseIsolatedControllerBuilder {
	controllerBuilder := &BehaviorNoParamsNoResponseIsolatedControllerBuilder{}

	// GET /behavior/no-params-no-response-isolated
	controllerBuilder.HandleBehaviorNoParamsNoResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoParamsNoResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorNoParamsNoResponse.voidResult = true
	controllerBuilder.HandleBehaviorNoParamsNoResponse.paramsParserFactory = makeVoidParamsParser

	return controllerBuilder
}

func RegisterBehaviorNoParamsNoResponseIsolatedRoutes(controller *BehaviorNoParamsNoResponseIsolatedController, app *HTTPApp) {
	app.router.HandleRoute("GET", "/behavior/no-params-no-response-isolated", controller.BehaviorNoParamsNoResponse(app))
}

type BehaviorNoParamsNoResponseIsolatedControllerBuilderV2 struct {
	// GET /behavior/no-params-no-response-isolated
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse ActionBuilder[
	  Void,
	  Void,
	  func(context.Context) (error),
	  func(http.ResponseWriter, *http.Request) (error),
	]
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
  builder := BehaviorNoParamsNoResponseIsolatedControllerBuilderV2{}
	app.router.HandleRoute("GET", "/behavior/no-params-no-response-isolated", controller.BehaviorNoParamsNoResponse(&builder))
}