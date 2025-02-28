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
		*models.BehaviorBehaviorWithParamsNoResponseParams,
		mockVoid,
	]
	withParamsAndResponse mockActionV2[
		*models.BehaviorBehaviorWithParamsAndResponseParams,
		*models.BehaviorWithParamsAndResponseResponseBody,
	]
	noStatusDefined   mockActionV2[mockVoid, mockVoid]
	withStatusDefined mockActionV2[mockVoid, mockVoid]
}

type behaviorController struct {
	testActions *behaviorControllerTestActions
}

func (c *behaviorController) BehaviorNoParamsNoResponse(
	builder handlers.NoParamsNoResponseHandlerBuilder,
) http.Handler {
	if c.testActions.noParamsNoResponse.isHTTPAction {
		return builder.HandleWithHTTP(
			c.testActions.noParamsNoResponse.httpActionNoParamsNoResponse,
		)
	}
	return builder.HandleWith(
		c.testActions.noParamsNoResponse.actionNoParamsNoResponse,
	)
}

func (c *behaviorController) BehaviorNoParamsWithResponse(
	builder handlers.NoParamsHandlerBuilder[*models.BehaviorNoParamsWithResponse202Response],
) http.Handler {
	if c.testActions.noParamsWithResponse.isHTTPAction {
		return builder.HandleWithHTTP(
			c.testActions.noParamsWithResponse.httpActionNoParamsWithResponse,
		)
	}
	return builder.HandleWith(
		c.testActions.noParamsWithResponse.actionNoParamsWithResponse,
	)
}

func (c *behaviorController) BehaviorWithParamsAndResponse(
	builder handlers.HandlerBuilder[
		*models.BehaviorBehaviorWithParamsAndResponseParams,
		*models.BehaviorWithParamsAndResponseResponseBody,
	],
) http.Handler {
	if c.testActions.withParamsAndResponse.isHTTPAction {
		return builder.HandleWithHTTP(
			c.testActions.withParamsAndResponse.httpAction,
		)
	}
	return builder.HandleWith(
		c.testActions.withParamsAndResponse.action,
	)
}

func (c *behaviorController) BehaviorWithParamsNoResponse(
	builder handlers.NoResponseHandlerBuilder[*models.BehaviorBehaviorWithParamsNoResponseParams],
) http.Handler {
	if c.testActions.withParamsNoResponse.isHTTPAction {
		return builder.HandleWithHTTP(
			c.testActions.withParamsNoResponse.httpActionWithParamsNoResponse,
		)
	}
	return builder.HandleWith(
		c.testActions.withParamsNoResponse.actionWithParamsNoResponse,
	)
}

func (c *behaviorController) BehaviorNoStatusDefined(
	builder handlers.NoParamsNoResponseHandlerBuilder,
) http.Handler {
	if c.testActions.noStatusDefined.isHTTPAction {
		return builder.HandleWithHTTP(
			c.testActions.noStatusDefined.httpActionNoParamsNoResponse,
		)
	}
	return builder.HandleWith(
		c.testActions.noStatusDefined.actionNoParamsNoResponse,
	)
}

func (c *behaviorController) BehaviorWithStatusDefined(
	builder handlers.NoParamsNoResponseHandlerBuilder,
) http.Handler {
	if c.testActions.withStatusDefined.isHTTPAction {
		return builder.HandleWithHTTP(
			c.testActions.withStatusDefined.httpActionNoParamsNoResponse,
		)
	}
	return builder.HandleWith(
		c.testActions.withStatusDefined.actionNoParamsNoResponse,
	)
}

var _ handlers.BehaviorController = &behaviorController{}
