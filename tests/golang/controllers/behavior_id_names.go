package controllers

import (
	"context"
	"net/http"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

type behaviorIDNamesControllerTestActions struct {
	behaviorNamesWithID mockActionV2[
		*handlers.BehaviorIdNamesBehaviorNamesWithIdRequest,
		*models.BehaviorNamesWithIdData,
	]
}

type behaviorIDNamesController struct {
	testActions *behaviorIDNamesControllerTestActions
}

func (c *behaviorIDNamesController) BehaviorNamesWithId(
	builder handlers.HandlerBuilder[
		*handlers.BehaviorIdNamesBehaviorNamesWithIdRequest,
		*models.BehaviorNamesWithIdData,
	],
) http.Handler {
	return builder.HandleWith(
		func(
			ctx context.Context,
			b *handlers.BehaviorIdNamesBehaviorNamesWithIdRequest,
		) (*models.BehaviorNamesWithIdData, error) {
			return c.testActions.behaviorNamesWithID.action(ctx, b)
		},
	)
}

var _ handlers.BehaviorIdNamesController = &behaviorIDNamesController{}
