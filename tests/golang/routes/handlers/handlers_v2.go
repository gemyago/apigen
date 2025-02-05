package handlers

import (
	"context"
	"net/http"
)

type Void string

const VoidValue Void = "void"

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