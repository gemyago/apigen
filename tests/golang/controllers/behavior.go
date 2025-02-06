package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type behaviorControllerTestActions struct {
	noParamsNoResponse   mockActionV2[mockVoid, mockVoid]
	noParamsWithResponse mockActionV2[
		mockVoid,
		*models.BehaviorNoParamsWithResponse202Response,
	]
	withParamsNoResponse mockActionV2[
		*handlers.BehaviorBehaviorWithParamsNoResponseRequest,
		mockVoid,
	]
	withParamsAndResponse mockActionV2[
		*handlers.BehaviorBehaviorWithParamsAndResponseRequest,
		*models.BehaviorNoParamsWithResponse202Response,
	]
	noStatusDefined   mockActionV2[mockVoid, mockVoid]
	withStatusDefined mockActionV2[mockVoid, mockVoid]
}

type behaviorController struct {
	testActions *behaviorControllerTestActions
}

func (c *behaviorController) BehaviorNoParamsNoResponse(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorNoParamsNoResponse.HandleWith(
		c.testActions.noParamsNoResponse.actionNoParamsNoResponse,
	)
}

func (c *behaviorController) BehaviorNoParamsWithResponse(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorNoParamsWithResponse.HandleWith(
		c.testActions.noParamsWithResponse.actionNoParamsWithResponse,
	)
}

func (c *behaviorController) BehaviorWithParamsNoResponse(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorWithParamsNoResponse.HandleWith(
		c.testActions.withParamsNoResponse.actionWithParamsNoResponse,
	)
}

func (c *behaviorController) BehaviorWithParamsAndResponse(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorWithParamsAndResponse.HandleWith(
		c.testActions.withParamsAndResponse.action,
	)
}

func (c *behaviorController) BehaviorNoStatusDefined(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorNoStatusDefined.HandleWith(
		c.testActions.noStatusDefined.actionNoParamsNoResponse,
	)
}

func (c *behaviorController) BehaviorWithStatusDefined(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorWithStatusDefined.HandleWith(
		c.testActions.withStatusDefined.actionNoParamsNoResponse,
	)
}

var _ handlers.BehaviorControllerV2 = &behaviorController{}
