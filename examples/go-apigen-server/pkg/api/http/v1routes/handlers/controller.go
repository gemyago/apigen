package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type voidResult *int

type httpHandlerFactory func(r httpRouter) http.Handler
type paramsParserFunc[TReqParams any] func(router httpRouter, w http.ResponseWriter, req *http.Request) (TReqParams, error)

type handlerFactoryParams[TReqParams any, TResData any] struct {
	defaultStatus int
	paramsParser  paramsParserFunc[TReqParams]
	handler       func(context.Context, TReqParams) (TResData, error)
}

func createHandlerFactory[TReqParams any, TResData any](factoryParams handlerFactoryParams[TReqParams, TResData]) httpHandlerFactory {
	return func(router httpRouter) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params, err := factoryParams.paramsParser(router, w, r)
			if err != nil {
				panic(fmt.Errorf("TODO: handle params parsing errors: %w", err))
			}

			resData, err := factoryParams.handler(r.Context(), params)
			if err != nil {
				router.HandleError(r, w, err)
				return
			}

			// TODO: Do it only if result is not void (same as sending the response)
			w.Header().Add("Content-Type", "application/json; utf-8")

			w.WriteHeader(factoryParams.defaultStatus)

			if err := json.NewEncoder(w).Encode(resData); err != nil {
				// TODO: We need to better handle those
				panic(err)
			}
		})
	}
}

type actionBuilder[TControllerBuilder any, TReqParams any, TResData any] struct {
	defaultStatusCode  int
	httpHandlerFactory func(r httpRouter) http.Handler
	paramsParser       paramsParserFunc[TReqParams]
	controllerBuilder  TControllerBuilder
}

func (ab *actionBuilder[TControllerBuilder, TReqParams, TResData]) With(
	handler func(context.Context, TReqParams) (TResData, error),
) TControllerBuilder {
	ab.httpHandlerFactory = createHandlerFactory(handlerFactoryParams[TReqParams, TResData]{
		defaultStatus: ab.defaultStatusCode,
		paramsParser:  ab.paramsParser,
		handler:       handler,
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
