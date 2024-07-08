package handlers

import (
	"context"
	"encoding/json"
	"errors"
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
				// TODO: This needs a different error handling
				w.WriteHeader(400)
				if _, err := w.Write([]byte(err.Error())); err != nil {
					panic(err)
				}
				return
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
	value    TVal
	assigned bool
}

func makeOptionalVal[TVal any](val TVal) optionalVal[TVal] {
	return optionalVal[TVal]{value: val, assigned: true}
}

func readPathValue(key string, router httpRouter, req *http.Request) optionalVal[string] {
	return optionalVal[string]{value: router.PathValue(req, key), assigned: true}
}

func readQueryValue(key string, values url.Values) optionalVal[[]string] {
	if values.Has(key) {
		return optionalVal[[]string]{value: values[key], assigned: true}
	}
	return optionalVal[[]string]{}
}

type rawValueParser[TRawVal any, TTargetVal any] func(optionalVal[TRawVal], *TTargetVal) error

func parseJsonPayload[TTargetVal any](req optionalVal[*http.Request], target *TTargetVal) error {
	return json.NewDecoder(req.value.Body).Decode(target)
}

var _ rawValueParser[*http.Request, string] = parseJsonPayload

func newStringToSignedIntParser[TTargetVal constraints.Signed](bitSize int) rawValueParser[string, TTargetVal] {
	return func(ov optionalVal[string], target *TTargetVal) error {
		if !ov.assigned {
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
		if !ov.assigned {
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

func (be bindingError) Error() string {
	return fmt.Sprintf("field %s (in %s) error: %v", be.field, be.location, be.err)
}

type bindingContext struct {
	errors []error
}

func (c bindingContext) Error() error {
	if len(c.errors) == 0 {
		return nil
	}
	return errors.Join(c.errors...)
}

type requestParamBinder[TRawVal any, TTargetVal any] func(
	bindingCtx *bindingContext,
	rawVal optionalVal[TRawVal],
	receiver *TTargetVal,
)

type valueValidator[TRawVal any, TTargetVal any] func(optionalVal[TRawVal], TTargetVal) error

var errValueRequired = errors.New("value is required")

func validateNonEmpty[TRawVal any, TTargetVal any](rawVal optionalVal[TRawVal], _ TTargetVal) error {
	if !rawVal.assigned {
		return errValueRequired
	}
	return nil
}

var _ valueValidator[string, string] = validateNonEmpty

type orderedValuesValidatorParams[TTargetVal constraints.Ordered] struct {
	minimum optionalVal[TTargetVal]
	maximum optionalVal[TTargetVal]
}

func newOrderedValuesValidator[TRawVal any, TTargetVal constraints.Ordered](
	params orderedValuesValidatorParams[TTargetVal],
) valueValidator[TRawVal, TTargetVal] {
	return func(ov optionalVal[TRawVal], tv TTargetVal) error {
		if !ov.assigned {
			return nil
		}

		if params.minimum.assigned && tv < params.minimum.value {
			return fmt.Errorf("value %v is less than minimum %v", tv, params.minimum.value)
		}

		if params.maximum.assigned && tv > params.maximum.value {
			return fmt.Errorf("value %v is greater than maximum %v", tv, params.maximum.value)
		}

		return nil
	}
}

func newCompositeValidator[TRawVal any, TTargetVal any](validators ...valueValidator[TRawVal, TTargetVal]) valueValidator[TRawVal, TTargetVal] {
	return func(ov optionalVal[TRawVal], tv TTargetVal) error {
		for _, v := range validators {
			if err := v(ov, tv); err != nil {
				return err
			}
		}
		return nil
	}
}

type binderParams[TRawVal any, TTargetVal any] struct {
	field         string
	location      string
	parseValue    rawValueParser[TRawVal, TTargetVal]
	validateValue valueValidator[TRawVal, TTargetVal]
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
		if err := params.validateValue(rawVal, *receiver); err != nil {
			bindingCtx.errors = append(bindingCtx.errors, bindingError{
				field:    params.field,
				location: params.location,
				err:      err,
			})
		}
	}
}
