package controllers

import (
	"context"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
)

type behaviorControllerNoParamsNoResponseTestActions struct {
	noParamsNoResponse mockAction[struct{}]
}

func newBehaviorControllerNoParamsNoResponse(
	testActions *behaviorControllerNoParamsNoResponseTestActions,
) *handlers.BehaviorNoParamsNoResponseIsolatedController {
	return handlers.BuildBehaviorNoParamsNoResponseIsolatedController().
		HandleBehaviorNoParamsNoResponse.With(
		func(ctx context.Context) error {
			return testActions.noParamsNoResponse.action(ctx, struct{}{})
		}).
		Finalize()
}
