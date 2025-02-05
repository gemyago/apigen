package controllers

import (
	"context"
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type behaviorControllerTestActions struct {
	noParamsNoResponse mockAction[struct{}]
}

type behaviorController struct {
	testActions *behaviorControllerTestActions
}

func (c *behaviorController) BehaviorNoParamsNoResponse(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorNoParamsNoResponse.HandleWith(
		func(ctx context.Context) error {
			return c.testActions.noParamsNoResponse.action(ctx, struct{}{})
		},
	)
}

func (c *behaviorController) BehaviorNoParamsWithResponse(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorNoParamsWithResponse.HandleWith(
		func(context.Context) (*models.BehaviorNoParamsWithResponse202Response, error) {
			panic("not implemented")
		},
	)
}

func (c *behaviorController) BehaviorWithParamsAndResponse(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorWithParamsAndResponse.HandleWith(
		func(
			context.Context,
			*handlers.BehaviorBehaviorWithParamsAndResponseRequest,
		) (*models.BehaviorNoParamsWithResponse202Response, error) {
			panic("not implemented")
		},
	)
}

func (c *behaviorController) BehaviorNoStatusDefined(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorNoStatusDefined.HandleWith(
		func(context.Context) error {
			panic("not implemented")
		},
	)
}

func (c *behaviorController) BehaviorWithStatusDefined(
	builder *handlers.BehaviorControllerBuilderV2,
) http.Handler {
	return builder.BehaviorWithStatusDefined.HandleWith(
		func(context.Context) error {
			panic("not implemented")
		},
	)
}

var _ handlers.BehaviorControllerV2 = &behaviorController{}
