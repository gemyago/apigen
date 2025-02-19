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
	builder handlers.NoParamsNoResponseHandlerBuilder,
) http.Handler {
	return builder.HandleWith(
		func(ctx context.Context) error {
			return c.testActions.noParamsNoResponse.action(ctx, struct{}{})
		},
	)
}

var _ handlers.BehaviorNoParamsNoResponseIsolatedControllerV3 = &behaviorControllerNoParamsNoResponse{}
