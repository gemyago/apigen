package controllers

import (
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type transformedBehaviorNoParamsWithResponse202Response = models.BehaviorNoParamsWithResponse202Response
type transformedBehaviorBehaviorWithParamsNoResponseRequest = handlers.BehaviorBehaviorWithParamsNoResponseRequest
type transformedBehaviorBehaviorWithParamsAndResponseRequest = handlers.BehaviorBehaviorWithParamsAndResponseRequest
type transformedBehaviorWithParamsAndResponseResponseBody = models.BehaviorWithParamsAndResponseResponseBody

type behaviorControllerTransformTestActions struct {
	noParamsWithResponse mockActionV2[
		mockVoid,
		*transformedBehaviorNoParamsWithResponse202Response,
	]
	withParamsNoResponse mockActionV2[
		*transformedBehaviorBehaviorWithParamsNoResponseRequest,
		mockVoid,
	]
	withParamsAndResponse mockActionV2[
		*transformedBehaviorBehaviorWithParamsAndResponseRequest,
		*transformedBehaviorWithParamsAndResponseResponseBody,
	]
}

type behaviorControllerTransform struct {
	stdController behaviorController
	testActions   *behaviorControllerTransformTestActions
}

func (c *behaviorControllerTransform) BehaviorNoParamsWithResponse(
	builder handlers.NoParamsHandlerBuilder[*models.BehaviorNoParamsWithResponse202Response],
) http.Handler {
	return builder.HandleWith(
		c.testActions.noParamsWithResponse.actionNoParamsWithResponse,
	)
}

func (c *behaviorControllerTransform) BehaviorWithParamsAndResponse(
	builder handlers.HandlerBuilder[
		*handlers.BehaviorBehaviorWithParamsAndResponseRequest,
		*models.BehaviorWithParamsAndResponseResponseBody,
	],
) http.Handler {
	return builder.HandleWith(
		c.testActions.withParamsAndResponse.action,
	)
}

func (c *behaviorControllerTransform) BehaviorWithParamsNoResponse(
	builder handlers.NoResponseHandlerBuilder[*handlers.BehaviorBehaviorWithParamsNoResponseRequest],
) http.Handler {
	return builder.HandleWith(
		c.testActions.withParamsNoResponse.actionWithParamsNoResponse,
	)
}

func (c *behaviorControllerTransform) BehaviorNoParamsNoResponse(
	builder handlers.NoParamsNoResponseHandlerBuilder,
) http.Handler {
	return c.stdController.BehaviorNoParamsNoResponse(builder)
}

func (c *behaviorControllerTransform) BehaviorNoStatusDefined(
	builder handlers.NoParamsNoResponseHandlerBuilder,
) http.Handler {
	return c.stdController.BehaviorNoStatusDefined(builder)
}

func (c *behaviorControllerTransform) BehaviorWithStatusDefined(
	builder handlers.NoParamsNoResponseHandlerBuilder,
) http.Handler {
	return c.stdController.BehaviorWithStatusDefined(builder)
}

var _ handlers.BehaviorController = &behaviorControllerTransform{}
