package handlers

import (
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/models"
)

// Below is to workaround unused imports.
var _ = time.Time{}







type BehaviorBehaviorWithParamsAndResponseRequest struct {
	QueryParam1 string
}



type BehaviorController struct {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoParamsNoResponse httpHandlerFactory

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: models.BehaviorNoParamsWithResponse202Response
	BehaviorNoParamsWithResponse httpHandlerFactory

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorNoStatusDefined httpHandlerFactory

	// GET /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: models.BehaviorNoParamsWithResponse202Response
	BehaviorWithParamsAndResponse httpHandlerFactory

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	BehaviorWithStatusDefined httpHandlerFactory
}

type BehaviorControllerBuilder struct {
	// GET /behavior/no-params-no-response
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorNoParamsNoResponse actionBuilderNoParamsVoidResult[*BehaviorControllerBuilder]

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: models.BehaviorNoParamsWithResponse202Response
	HandleBehaviorNoParamsWithResponse actionBuilderNoParams[*BehaviorControllerBuilder, *models.BehaviorNoParamsWithResponse202Response]

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorNoStatusDefined actionBuilderNoParamsVoidResult[*BehaviorControllerBuilder]

	// GET /behavior/with-params-and-response
	//
	// Request type: BehaviorBehaviorWithParamsAndResponseRequest,
	//
	// Response type: models.BehaviorNoParamsWithResponse202Response
	HandleBehaviorWithParamsAndResponse actionBuilder[*BehaviorControllerBuilder, *BehaviorBehaviorWithParamsAndResponseRequest, *models.BehaviorNoParamsWithResponse202Response]

	// POST /behavior/with-status-defined
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorWithStatusDefined actionBuilderNoParamsVoidResult[*BehaviorControllerBuilder]
}

func (c *BehaviorControllerBuilder) Finalize() *BehaviorController {
	return &BehaviorController{
		BehaviorNoParamsNoResponse: mustInitializeAction("behaviorNoParamsNoResponse", c.HandleBehaviorNoParamsNoResponse.httpHandlerFactory),
		BehaviorNoParamsWithResponse: mustInitializeAction("behaviorNoParamsWithResponse", c.HandleBehaviorNoParamsWithResponse.httpHandlerFactory),
		BehaviorNoStatusDefined: mustInitializeAction("behaviorNoStatusDefined", c.HandleBehaviorNoStatusDefined.httpHandlerFactory),
		BehaviorWithParamsAndResponse: mustInitializeAction("behaviorWithParamsAndResponse", c.HandleBehaviorWithParamsAndResponse.httpHandlerFactory),
		BehaviorWithStatusDefined: mustInitializeAction("behaviorWithStatusDefined", c.HandleBehaviorWithStatusDefined.httpHandlerFactory),
	}
}

func BuildBehaviorController() *BehaviorControllerBuilder {
	controllerBuilder := &BehaviorControllerBuilder{}

	// GET /behavior/no-params-no-response
	controllerBuilder.HandleBehaviorNoParamsNoResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoParamsNoResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorNoParamsNoResponse.voidResult = true

	// GET /behavior/no-params-with-response
	controllerBuilder.HandleBehaviorNoParamsWithResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoParamsWithResponse.defaultStatusCode = 202

	// GET /behavior/no-status-defined
	controllerBuilder.HandleBehaviorNoStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoStatusDefined.defaultStatusCode = 200
	controllerBuilder.HandleBehaviorNoStatusDefined.voidResult = true

	// GET /behavior/with-params-and-response
	controllerBuilder.HandleBehaviorWithParamsAndResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorWithParamsAndResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorWithParamsAndResponse.paramsParserFactory = newParamsParserBehaviorBehaviorWithParamsAndResponse

	// POST /behavior/with-status-defined
	controllerBuilder.HandleBehaviorWithStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorWithStatusDefined.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorWithStatusDefined.voidResult = true

	return controllerBuilder
}

func RegisterBehaviorRoutes(controller *BehaviorController, app *HTTPApp) {
	app.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(app))
	app.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(app))
	app.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(app))
	app.router.HandleRoute("GET", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(app))
	app.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(app))
}
