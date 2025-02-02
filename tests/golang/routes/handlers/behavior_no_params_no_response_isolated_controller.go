package handlers

import (
	"encoding/json"
	"fmt"
	"time"
)

// Below is to workaround unused imports.
var _ = time.Time{}
var _ = json.Unmarshal
var _ = fmt.Sprint



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
