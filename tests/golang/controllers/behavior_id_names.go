package controllers

import (
	"context"
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type behaviorIDNamesControllerTestActions struct {
	behaviorNamesWithID mockActionV2[
		*models.BehaviorNamesWithIDParams,
		*models.BehaviorNamesWithIDData,
	]
}

type behaviorIDNamesController struct {
	testActions *behaviorIDNamesControllerTestActions
}

func (c *behaviorIDNamesController) BehaviorNamesWithID(
	builder handlers.HandlerBuilder[
		*models.BehaviorNamesWithIDParams,
		*models.BehaviorNamesWithIDData,
	],
) http.Handler {
	return builder.HandleWith(
		func(
			ctx context.Context,
			b *models.BehaviorNamesWithIDParams,
		) (*models.BehaviorNamesWithIDData, error) {
			return c.testActions.behaviorNamesWithID.action(ctx, b)
		},
	)
}

var _ handlers.BehaviorIDNamesController = &behaviorIDNamesController{}
