package controllers

import (
	"context"
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type behaviorIDNamesControllerTestActions struct {
	behaviorNamesWithID mockActionV2[
		*handlers.BehaviorIdNamesBehaviorNamesWithIDRequest,
		*models.BehaviorNamesWithIdData,
	]
}

type behaviorIDNamesController struct {
	testActions *behaviorIDNamesControllerTestActions
}

func (c *behaviorIDNamesController) BehaviorNamesWithID(
	builder handlers.HandlerBuilder[
		*handlers.BehaviorIdNamesBehaviorNamesWithIDRequest,
		*models.BehaviorNamesWithIdData,
	],
) http.Handler {
	return builder.HandleWith(
		func(
			ctx context.Context,
			b *handlers.BehaviorIdNamesBehaviorNamesWithIDRequest,
		) (*models.BehaviorNamesWithIdData, error) {
			return c.testActions.behaviorNamesWithID.action(ctx, b)
		},
	)
}

var _ handlers.BehaviorIdNamesController = &behaviorIDNamesController{}
