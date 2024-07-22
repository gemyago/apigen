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
	HandleBehaviorNoParamsNoResponse actionBuilderVoidResult[*BehaviorControllerBuilder, *BehaviorBehaviorNoParamsNoResponseRequest]

	// GET /behavior/no-params-with-response
	//
	// Request type: none
	//
	// Response type: models.BehaviorNoParamsWithResponse202Response
	HandleBehaviorNoParamsWithResponse actionBuilder[*BehaviorControllerBuilder, *BehaviorBehaviorNoParamsWithResponseRequest, *models.BehaviorNoParamsWithResponse202Response]

	// GET /behavior/no-status-defined
	//
	// Request type: none
	//
	// Response type: none
	HandleBehaviorNoStatusDefined actionBuilderVoidResult[*BehaviorControllerBuilder, *BehaviorBehaviorNoStatusDefinedRequest]

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
	HandleBehaviorWithStatusDefined actionBuilderVoidResult[*BehaviorControllerBuilder, *BehaviorBehaviorWithStatusDefinedRequest]
}

func (c *BehaviorControllerBuilder) Finalize() *BehaviorController {
	// TODO: panic if any handler is null
	return &BehaviorController{
		BehaviorNoParamsNoResponse: c.HandleBehaviorNoParamsNoResponse.httpHandlerFactory,
		BehaviorNoParamsWithResponse: c.HandleBehaviorNoParamsWithResponse.httpHandlerFactory,
		BehaviorNoStatusDefined: c.HandleBehaviorNoStatusDefined.httpHandlerFactory,
		BehaviorWithParamsAndResponse: c.HandleBehaviorWithParamsAndResponse.httpHandlerFactory,
		BehaviorWithStatusDefined: c.HandleBehaviorWithStatusDefined.httpHandlerFactory,
	}
}

func BuildBehaviorController() *BehaviorControllerBuilder {
	controllerBuilder := &BehaviorControllerBuilder{}

	// GET /behavior/no-params-no-response
	controllerBuilder.HandleBehaviorNoParamsNoResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoParamsNoResponse.defaultStatusCode = 202200
	controllerBuilder.HandleBehaviorNoParamsNoResponse.voidResult = true
	controllerBuilder.HandleBehaviorNoParamsNoResponse.paramsParserFactory = newParamsParserBehaviorBehaviorNoParamsNoResponse

	// GET /behavior/no-params-with-response
	controllerBuilder.HandleBehaviorNoParamsWithResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoParamsWithResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorNoParamsWithResponse.paramsParserFactory = newParamsParserBehaviorBehaviorNoParamsWithResponse

	// GET /behavior/no-status-defined
	controllerBuilder.HandleBehaviorNoStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorNoStatusDefined.defaultStatusCode = 200
	controllerBuilder.HandleBehaviorNoStatusDefined.voidResult = true
	controllerBuilder.HandleBehaviorNoStatusDefined.paramsParserFactory = newParamsParserBehaviorBehaviorNoStatusDefined

	// GET /behavior/with-params-and-response
	controllerBuilder.HandleBehaviorWithParamsAndResponse.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorWithParamsAndResponse.defaultStatusCode = 202
	controllerBuilder.HandleBehaviorWithParamsAndResponse.paramsParserFactory = newParamsParserBehaviorBehaviorWithParamsAndResponse

	// POST /behavior/with-status-defined
	controllerBuilder.HandleBehaviorWithStatusDefined.controllerBuilder = controllerBuilder
	controllerBuilder.HandleBehaviorWithStatusDefined.defaultStatusCode = 202200
	controllerBuilder.HandleBehaviorWithStatusDefined.voidResult = true
	controllerBuilder.HandleBehaviorWithStatusDefined.paramsParserFactory = newParamsParserBehaviorBehaviorWithStatusDefined

	return controllerBuilder
}

func RegisterBehaviorRoutes(controller *BehaviorController, app *HTTPApp) {
	app.router.HandleRoute("GET", "/behavior/no-params-no-response", controller.BehaviorNoParamsNoResponse(app))
	app.router.HandleRoute("GET", "/behavior/no-params-with-response", controller.BehaviorNoParamsWithResponse(app))
	app.router.HandleRoute("GET", "/behavior/no-status-defined", controller.BehaviorNoStatusDefined(app))
	app.router.HandleRoute("GET", "/behavior/with-params-and-response", controller.BehaviorWithParamsAndResponse(app))
	app.router.HandleRoute("POST", "/behavior/with-status-defined", controller.BehaviorWithStatusDefined(app))
}