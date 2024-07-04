package handlers

import "context"

type voidResult *int

type actionBuilder[TControllerBuilder any, TReqParams any, TResData any] struct {
	controllerBuilder TControllerBuilder
}

func (ab actionBuilder[TControllerBuilder, TReqParams, TResData]) With(
	handler func(context.Context, TReqParams) (TResData, error),
) TControllerBuilder {
	return ab.controllerBuilder
}

type actionBuilderVoidResult[TControllerBuilder any, TReqParams any] struct {
	actionBuilder[TControllerBuilder, TReqParams, voidResult]
}

func (ab actionBuilderVoidResult[TControllerBuilder, TReqParams]) With(
	handler func(context.Context, TReqParams) error,
) TControllerBuilder {
	return ab.controllerBuilder
}
