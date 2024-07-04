package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type voidResult *int

type actionBuilder[TControllerBuilder any, TReqParams any, TResData any] struct {
	defaultStatusCode int
	httpHandler       http.Handler
	controllerBuilder TControllerBuilder
}

func (ab *actionBuilder[TControllerBuilder, TReqParams, TResData]) With(
	handler func(context.Context, TReqParams) (TResData, error),
) TControllerBuilder {
	defaultStatus := ab.defaultStatusCode
	ab.httpHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params TReqParams
		resData, err := handler(r.Context(), params)
		if err != nil {
			// TODO: We need a error handler
			panic(err)
		}

		// TODO: Do it only if result is not void (same as sending the response)
		w.Header().Add("Content-Type", "application/json; utf-8")

		w.WriteHeader(defaultStatus)

		if err := json.NewEncoder(w).Encode(resData); err != nil {
			// TODO: We need to better handle those
			panic(err)
		}
	})
	return ab.controllerBuilder
}

type actionBuilderVoidResult[TControllerBuilder any, TReqParams any] struct {
	actionBuilder[TControllerBuilder, TReqParams, voidResult]
}

func (ab *actionBuilderVoidResult[TControllerBuilder, TReqParams]) With(
	handler func(context.Context, TReqParams) error,
) TControllerBuilder {
	return ab.actionBuilder.With(func(ctx context.Context, tp TReqParams) (voidResult, error) {
		return nil, handler(ctx, tp)
	})
}
