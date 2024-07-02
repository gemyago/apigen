package handlers

import "context"

type actionBuilder[TReqParams any] struct{}

func (actionBuilder[TReqParams]) With(
	handler func(context.Context, TReqParams) error,
) *PetsControllerBuilder {
	return nil
}
