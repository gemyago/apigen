package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type Void string

const VoidValue Void = "void"

type voidParamsParserV2 struct{}

func (p voidParamsParserV2) parse(_ httpRouter, _ *http.Request) (*Void, error) {
	return nil, nil
}

func makeVoidParamsParserV2(_ *HTTPApp) paramsParser[*Void] {
	return voidParamsParserV2{}
}

// ActionHandlerFunc represents possible combination of action handler functions.
// Each function can be with or without parameters and with or without response.
// Additionally each function can have access to http objects for possible direct manipulation.
type ActionHandlerFunc[TReq any, TRes any] interface {
	func(context.Context, *TReq) (*TRes, error) | // with params with response
		func(context.Context) (*TRes, error) | // no params with response
		func(context.Context, *TReq) error | // with params no response
		func(context.Context) error | // no params no response

		// handlers with http context exposed
		func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) | // with params with response
		func(http.ResponseWriter, *http.Request) (*TRes, error) | // no params with response
		func(http.ResponseWriter, *http.Request, *TReq) error | // with params no response
		func(http.ResponseWriter, *http.Request) error // no params no response
}

type ActionBuilderFunc[TReq any, TRes any, THandler ActionHandlerFunc[TReq, TRes]] func(THandler) http.Handler

type ActionBuilder[
	TReq any,
	TRes any,
	TPlainHandler ActionHandlerFunc[TReq, TRes],
	THttpHandler ActionHandlerFunc[TReq, TRes],
] struct {
	HandleWith     func(TPlainHandler) http.Handler
	HandleWithHTTP func(THttpHandler) http.Handler
}

func BuildActionWithTransformers[
	TDeclaredReq any,
	TDeclaredRes any,
	TDeclaredFn ActionHandlerFunc[TDeclaredReq, TDeclaredRes],
	TDeclaredHttpFn ActionHandlerFunc[TDeclaredReq, TDeclaredRes],
	TAppReq any,
	TAppRes any,
	TAppFn ActionHandlerFunc[TAppReq, TAppRes],
](
	actionBuilder ActionBuilder[TDeclaredReq, TDeclaredRes, TDeclaredFn, TDeclaredHttpFn],
	action TAppFn,
	inputTransformer func(req *http.Request, input *TDeclaredReq) (*TAppReq, error),
	outputTransformer func(input *TAppRes) (*TDeclaredRes, error),
) http.Handler {
	panic("not implemented")
}

type makeActionBuilderParams[
	TReqParams any,
	TResData any,
] struct {
	defaultStatus int
	voidResult    bool
	paramsParser  paramsParser[*TReqParams]
}

type actionBuilderHandlerAdapter[
	TReq any,
	TRes any,
	THandler ActionHandlerFunc[TReq, TRes],
] func(THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error)

func newHandlerAdapter[
	TReq any,
	TRes any,
	THandler func(context.Context, *TReq) (*TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(_ http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			return t(r.Context(), req)
		}
	}
}

func newHandlerAdapterNoParams[
	TReq any,
	TRes any,
	THandler func(context.Context) (*TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(_ http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			return t(r.Context())
		}
	}
}

func newHandlerAdapterNoResponse[
	TReq any,
	TRes any,
	THandler func(context.Context, *TReq) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(_ http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			if err := t(r.Context(), req); err != nil {
				return nil, err
			}
			return nil, nil
		}
	}
}

func newHandlerAdapterNoParamsNoResponse[
	TReq any,
	TRes any,
	THandler func(context.Context) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			return nil, t(r.Context())
		}
	}
}

func newHTTPHandlerAdapter[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			return t(w, r, req)
		}
	}
}

func newHTTPHandlerAdapterNoParams[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request) (*TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			return t(w, r)
		}
	}
}

func newHTTPHandlerAdapterNoResponse[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request, *TReq) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			if err := t(w, r, req); err != nil {
				return nil, err
			}
			return nil, nil
		}
	}
}

func newHTTPHandlerAdapterNoParamsNoResponse[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, req *TReq) (*TRes, error) {
			return nil, t(w, r)
		}
	}
}

func makeActionBuilder[
	TReq any,
	TRes any,
	TPlainHandler ActionHandlerFunc[TReq, TRes],
	THttpHandler ActionHandlerFunc[TReq, TRes],
](
	app *HTTPApp,
	handlerAdapter actionBuilderHandlerAdapter[TReq, TRes, TPlainHandler],
	httpHandlerAdapter actionBuilderHandlerAdapter[TReq, TRes, THttpHandler],
	params makeActionBuilderParams[TReq, TRes],
) ActionBuilder[TReq, TRes, TPlainHandler, THttpHandler] {
	createHandler := func(handler func(http.ResponseWriter, *http.Request, *TReq) (*TRes, error)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqParams, err := params.paramsParser.parse(app.router, r)
			if err != nil {
				app.handleParsingErrors(r, w, err)
				return
			}

			resData, err := handler(w, r, reqParams)
			if err != nil {
				app.handleActionErrors(r, w, err)
				return
			}
			if params.voidResult {
				w.WriteHeader(params.defaultStatus)
				return
			}

			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(params.defaultStatus)
			if encodingErr := json.NewEncoder(w).Encode(resData); encodingErr != nil {
				app.handleResponseErrors(r, encodingErr)
			}
		})
	}

	return ActionBuilder[TReq, TRes, TPlainHandler, THttpHandler]{
		HandleWith: func(inputHandler TPlainHandler) http.Handler {
			return createHandler(handlerAdapter(inputHandler))
		},
		HandleWithHTTP: func(handler THttpHandler) http.Handler {
			return createHandler(httpHandlerAdapter(handler))
		},
	}
}
