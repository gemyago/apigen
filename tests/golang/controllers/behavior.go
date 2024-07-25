package controllers

import (
	"context"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/gemyago/apigen/tests/golang/routes/models"
)

func newBehaviorController() *handlers.BehaviorController {
	return handlers.BuildBehaviorController().
		HandleBehaviorNoParamsNoResponse.With(
		func(context.Context) error {
			panic("not implemented")
		}).
		HandleBehaviorNoParamsWithResponse.With(
		func(context.Context) (*models.BehaviorNoParamsWithResponse202Response, error) {
			panic("not implemented")
		}).
		HandleBehaviorWithParamsAndResponse.With(
		func(
			context.Context,
			*handlers.BehaviorBehaviorWithParamsAndResponseRequest,
		) (*models.BehaviorNoParamsWithResponse202Response, error) {
			panic("not implemented")
		}).
		HandleBehaviorNoStatusDefined.With(
		func(context.Context) error {
			panic("not implemented")
		}).
		HandleBehaviorWithStatusDefined.With(
		func(context.Context) error {
			panic("not implemented")
		}).
		Finalize()
}
