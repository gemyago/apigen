package handlers

import "context"

type actionBuilder[TReqParams any, TResData any] struct{}

func (actionBuilder[TReqParams, TResData]) With(
	handler func(context.Context, TReqParams) (TResData, error),
) *PetsControllerBuilder {
	return nil
}

type actionBuilderVoidResult[TReqParams any] struct{}

func (actionBuilderVoidResult[TReqParams]) With(
	handler func(context.Context, TReqParams) error,
) *PetsControllerBuilder {
	return nil
}
