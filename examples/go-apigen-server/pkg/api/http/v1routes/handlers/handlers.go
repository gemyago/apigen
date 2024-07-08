package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/exp/constraints"
)

type httpRouter interface {
	// PathValue returns a named path parameter of a given name
	PathValue(r *http.Request, paramName string) string

	// HandleRoute register a given handler function to handle given route
	HandleRoute(method, pathPattern string, h http.Handler)

	// HandleError will be called for any error produced by handlers
	HandleError(r *http.Request, w http.ResponseWriter, err error)
}

type voidResult *int

type httpHandlerFactory func(r httpRouter) http.Handler
type paramsParser[TReqParams any] interface {
	parse(router httpRouter, w http.ResponseWriter, req *http.Request) (TReqParams, error)
}

type handlerFactoryParams[TReqParams any, TResData any] struct {
	defaultStatus int
	voidResult    bool
	paramsParser  paramsParser[TReqParams]
	handler       func(context.Context, TReqParams) (TResData, error)
}

func createHandlerFactory[TReqParams any, TResData any](factoryParams handlerFactoryParams[TReqParams, TResData]) httpHandlerFactory {
	return func(router httpRouter) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params, err := factoryParams.paramsParser.parse(router, w, r)
			if err != nil {
				panic(fmt.Errorf("TODO: handle params parsing errors: %w", err))
			}

			resData, err := factoryParams.handler(r.Context(), params)
			if err != nil {
				router.HandleError(r, w, err)
				return
			}
			if factoryParams.voidResult {
				return
			}

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
	voidResult         bool
	httpHandlerFactory func(r httpRouter) http.Handler
	paramsParser       paramsParser[TReqParams]
	controllerBuilder  TControllerBuilder
}

func (ab *actionBuilder[TControllerBuilder, TReqParams, TResData]) With(
	handler func(context.Context, TReqParams) (TResData, error),
) TControllerBuilder {
	ab.httpHandlerFactory = createHandlerFactory(handlerFactoryParams[TReqParams, TResData]{
		defaultStatus: ab.defaultStatusCode,
		voidResult:    ab.voidResult,
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

type optionalVal[TVal any] struct {
	value TVal
	empty bool
}

func readPathValue(key string, router httpRouter, req *http.Request) optionalVal[string] {
	return optionalVal[string]{value: router.PathValue(req, key)}
}

func readQueryValue(key string, values url.Values) optionalVal[[]string] {
	if values.Has(key) {
		return optionalVal[[]string]{value: values[key], empty: false}
	}
	return optionalVal[[]string]{empty: true}
}

type rawValueParser[TRawVal any, TTargetVal any] func(optionalVal[TRawVal], *TTargetVal) error

func parseJsonPayload[TTargetVal any](req optionalVal[*http.Request], target *TTargetVal) error {
	return json.NewDecoder(req.value.Body).Decode(target)
}

var _ rawValueParser[*http.Request, string] = parseJsonPayload[string]

func newStringToSignedIntParser[TTargetVal constraints.Signed](bitSize int) rawValueParser[string, TTargetVal] {
	return func(ov optionalVal[string], target *TTargetVal) error {
		if ov.empty {
			return nil
		}
		val, err := strconv.ParseInt(ov.value, 10, bitSize)
		if err != nil {
			return err
		}
		*target = (TTargetVal)(val)
		return nil
	}
}

func newStringSliceToSignedIntParser[TTargetVal constraints.Signed](bitSize int) rawValueParser[[]string, TTargetVal] {
	return func(ov optionalVal[[]string], target *TTargetVal) error {
		if ov.empty {
			return nil
		}
		val, err := strconv.ParseInt(ov.value[0], 10, bitSize)
		if err != nil {
			return err
		}
		*target = (TTargetVal)(val)
		return nil
	}
}

type bindingError struct {
	field    string
	location string
	err      error
}

type bindingContext struct {
	errors []bindingError
}

type requestParamBinder[TRawVal any, TTargetVal any] func(
	bindingCtx *bindingContext,
	rawVal optionalVal[TRawVal],
	receiver *TTargetVal,
)

type binderParams[TRawVal any, TTargetVal any] struct {
	field      string
	location   string
	parseValue rawValueParser[TRawVal, TTargetVal]
}

func newRequestParamBinder[TRawVal any, TTargetVal any](
	params binderParams[TRawVal, TTargetVal],
) requestParamBinder[TRawVal, TTargetVal] {
	return func(
		bindingCtx *bindingContext,
		rawVal optionalVal[TRawVal],
		receiver *TTargetVal,
	) {
		if err := params.parseValue(rawVal, receiver); err != nil {
			bindingCtx.errors = append(bindingCtx.errors, bindingError{
				field:    params.field,
				location: params.location,
				err:      err,
			})
			return
		}
	}
}
