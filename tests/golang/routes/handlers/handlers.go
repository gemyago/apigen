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

			if !factoryParams.voidResult {
				w.Header().Add("Content-Type", "application/json; utf-8")
			}
			w.WriteHeader(factoryParams.defaultStatus)

			if factoryParams.voidResult {
				return
			}

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

func newStringToNumberParser[TTargetVal constraints.Integer | constraints.Float](
	bitSize int, parseFn func(string, int) (TTargetVal, error),
) rawValueParser[string, TTargetVal] {
	return func(ov optionalVal[string], target *TTargetVal) error {
		if !ov.assigned {
			return nil
		}
		val, err := parseFn(ov.value, bitSize)
		if err != nil {
			return err
		}
		*target = val
		return nil
	}
}

func newStringSliceToNumberParser[TTargetVal constraints.Integer | constraints.Float](
	bitSize int, parseFn func(string, int) (TTargetVal, error),
) rawValueParser[[]string, TTargetVal] {
	return func(ov optionalVal[[]string], target *TTargetVal) error {
		if !ov.assigned {
			return nil
		}
		val, err := parseFn(ov.value[0], bitSize)
		if err != nil {
			return err
		}
		*target = val
		return nil
	}
}

type knownParsersDef struct {
	int32_in_path    rawValueParser[string, int32]
	int64_in_path    rawValueParser[string, int64]
	float32_in_path  rawValueParser[string, float32]
	float64_in_path  rawValueParser[string, float64]
	int32_in_query   rawValueParser[[]string, int32]
	int64_in_query   rawValueParser[[]string, int64]
	float32_in_query rawValueParser[[]string, float32]
	float64_in_query rawValueParser[[]string, float64]
}

func parseDecInt[TInt constraints.Integer](str string, bitSize int) (TInt, error) {
	res, err := strconv.ParseInt(str, 10, bitSize)
	return (TInt)(res), err
}

func parseFloat[TFloat constraints.Float](str string, bitSize int) (TFloat, error) {
	res, err := strconv.ParseFloat(str, bitSize)
	return (TFloat)(res), err
}

var knownParsers = knownParsersDef{
	int32_in_path:    newStringToNumberParser[int32](32, parseDecInt),
	int64_in_path:    newStringToNumberParser[int64](64, parseDecInt),
	float32_in_path:  newStringToNumberParser[float32](32, parseFloat),
	float64_in_path:  newStringToNumberParser[float64](64, strconv.ParseFloat),
	int32_in_query:   newStringSliceToNumberParser[int32](32, parseDecInt),
	int64_in_query:   newStringSliceToNumberParser[int64](64, parseDecInt),
	float32_in_query: newStringSliceToNumberParser[float32](32, parseFloat),
	float64_in_query: newStringSliceToNumberParser[float64](64, strconv.ParseFloat),
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

func newMinMaxValueValidator[TRawVal any, TTargetVal constraints.Ordered](
	threshold TTargetVal,
	inclusive bool,
	isMin bool,
) valueValidator[TRawVal, TTargetVal] {
	return func(ov optionalVal[TRawVal], tv TTargetVal) error {
		if !ov.assigned {
			return nil
		}

		if isMin && ((inclusive && tv <= threshold) || (!inclusive && tv < threshold)) {
			return fmt.Errorf("value %v is less than minimum %v", tv, threshold)
		}
		if !isMin && ((inclusive && tv >= threshold) || (!inclusive && tv > threshold)) {
			return fmt.Errorf("value %v is greater than maximum %v", tv, threshold)
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
