package controllers

import (
	"context"
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type behaviorIDNamesControllerTestActions struct {
	behaviorNamesWithID mockActionV2[
		*handlers.BehaviorIDNamesBehaviorNamesWithIDRequest,
		*models.BehaviorNamesWithIDData,
	]
}

type behaviorIDNamesController struct {
	testActions *behaviorIDNamesControllerTestActions
}

func (c *behaviorIDNamesController) BehaviorNamesWithID(
	builder handlers.HandlerBuilder[
		*handlers.BehaviorIDNamesBehaviorNamesWithIDRequest,
		*models.BehaviorNamesWithIDData,
	],
) http.Handler {
	return builder.HandleWith(
		func(
			ctx context.Context,
			b *handlers.BehaviorIDNamesBehaviorNamesWithIDRequest,
		) (*models.BehaviorNamesWithIDData, error) {
			return c.testActions.behaviorNamesWithID.action(ctx, b)
		},
	)
}

var _ handlers.BehaviorIDNamesController = &behaviorIDNamesController{}
