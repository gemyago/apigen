package controllers

import (
	"context"
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type behaviorControllerNoParamsNoResponseTestActions struct {
	noParamsNoResponse mockAction[struct{}]
}

type behaviorControllerNoParamsNoResponse struct {
	testActions *behaviorControllerNoParamsNoResponseTestActions
}

func (c *behaviorControllerNoParamsNoResponse) BehaviorNoParamsNoResponse(
	builder *handlers.BehaviorNoParamsNoResponseIsolatedControllerBuilderV2,
) http.Handler {
	return builder.BehaviorNoParamsNoResponse.HandleWith(
		func(ctx context.Context) error {
			return c.testActions.noParamsNoResponse.action(ctx, struct{}{})
		},
	)
}

var _ handlers.BehaviorNoParamsNoResponseIsolatedControllerV2 = &behaviorControllerNoParamsNoResponse{}
