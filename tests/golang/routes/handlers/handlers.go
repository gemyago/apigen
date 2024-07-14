package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
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
}

// ParsingErrorHandler will process errors during parsing and validation stages
// The default implementation will respond with 400 status code and standard
// serialization of ParsingError type.
type ParsingErrorHandler func(r *http.Request, w http.ResponseWriter, err error)

// ActionErrorHandler will process errors produced by controller actions
// The default implementation will respond with 500 and no output.
type ActionErrorHandler func(r *http.Request, w http.ResponseWriter, err error)

// ResponseErrorHandler will process errors that may occur while writing response
// At this stage either logging or panic is possible.
type ResponseErrorHandler func(r *http.Request, err error)

// SlogLogger is a fully compatible with slog and used to allow injecting the instance
type SlogLogger interface {
	Log(ctx context.Context, level slog.Level, msg string, args ...any)
	LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
}

type httpApp struct {
	router               httpRouter
	handleParsingErrors  ParsingErrorHandler
	handleActionErrors   ActionErrorHandler
	handleResponseErrors ResponseErrorHandler
	logger               SlogLogger
}

type HttpAppOpt func(app *httpApp)

func WithParsingErrorHandler(handler ParsingErrorHandler) HttpAppOpt {
	return func(app *httpApp) {
		app.handleParsingErrors = handler
	}
}

func WithActionErrorHandler(handler ActionErrorHandler) HttpAppOpt {
	return func(app *httpApp) {
		app.handleActionErrors = handler
	}
}

func WithResponseErrorHandler(handler ResponseErrorHandler) HttpAppOpt {
	return func(app *httpApp) {
		app.handleResponseErrors = handler
	}
}

func WithLogger(logger SlogLogger) HttpAppOpt {
	return func(app *httpApp) {
		app.logger = logger
	}
}

func NewHttpApp(router httpRouter, opts ...HttpAppOpt) *httpApp {

	app := &httpApp{
		router: router,
		logger: slog.Default(),
	}
	app.handleResponseErrors = func(r *http.Request, err error) {
		app.logger.LogAttrs(r.Context(), slog.LevelError, "Failed to write response", slog.Any("err", err))
	}
	app.handleParsingErrors = func(r *http.Request, w http.ResponseWriter, err error) {
		w.Header().Add("content-type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		app.logger.LogAttrs(r.Context(), slog.LevelWarn, "Failed to parse request", slog.Any("err", err))
		if httpErr, ok := err.(AggregatedBindingError); ok {
			if writeErr := json.NewEncoder(w).Encode(httpErr); writeErr != nil {
				app.handleResponseErrors(r, writeErr)
			}
			return
		}
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

type voidResult *int

type httpHandlerFactory func(app *httpApp) http.Handler
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
	return func(app *httpApp) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params, err := factoryParams.paramsParser.parse(app.router, w, r)
			if err != nil {
				app.handleParsingErrors(r, w, err)
				return
			}

			resData, err := factoryParams.handler(r.Context(), params)
			if err != nil {
				app.handleActionErrors(r, w, err)
				return
			}
			if factoryParams.voidResult {
				w.WriteHeader(factoryParams.defaultStatus)
				return
			}

			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(factoryParams.defaultStatus)
			if err := json.NewEncoder(w).Encode(resData); err != nil {
				app.handleResponseErrors(r, err)
			}
		})
	}
}

type actionBuilder[TControllerBuilder any, TReqParams any, TResData any] struct {
	defaultStatusCode  int
	voidResult         bool
	httpHandlerFactory func(app *httpApp) http.Handler
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
	float64_in_path:  newStringToNumberParser(64, strconv.ParseFloat),
	int32_in_query:   newStringSliceToNumberParser[int32](32, parseDecInt),
	int64_in_query:   newStringSliceToNumberParser[int64](64, parseDecInt),
	float32_in_query: newStringSliceToNumberParser[float32](32, parseFloat),
	float64_in_query: newStringSliceToNumberParser(64, strconv.ParseFloat),
}

type BindingErrorCode string

const (
	// ErrBadValueFormat error means data provided can not be parsed to a target type
	ErrBadValueFormat BindingErrorCode = "BAD_FORMAT"

	// ErrValueRequired error code indicates that the required value has not been provided
	ErrValueRequired BindingErrorCode = "INVALID_REQUIRED"

	// ErrInvalidValueOutOfRange error code indicates that the value is out of range of allowable values
	// this is usually when number is out of min/max range, or string is outside of limits
	ErrInvalidValueOutOfRange BindingErrorCode = "INVALID_OUT_OF_RANGE"

	// ErrInvalidValue error code a generic validation error
	ErrInvalidValue BindingErrorCode = "INVALID"
)

func (c BindingErrorCode) Error() string {
	return string(c)
}

// BindingError occurs at parsing/validation stage and holds
// context on field that the error is related to
type BindingError struct {
	Field    string           `json:"field"`
	Location string           `json:"location"`
	Err      error            `json:"-"`
	Code     BindingErrorCode `json:"code"`
}

func (be BindingError) Error() string {
	return fmt.Sprintf("field %s (in %s) error: %v", be.Field, be.Location, be.Err)
}

type AggregatedBindingError struct {
	Errors []BindingError `json:"errors"`
}

func (c AggregatedBindingError) Error() string {
	errs := make([]error, len(c.Errors))
	for i, err := range c.Errors {
		errs[i] = err
	}
	return errors.Join(errs...).Error()
}

type bindingContext struct {
	errors []BindingError
}

func (c bindingContext) AggregatedError() error {
	if len(c.errors) == 0 {
		return nil
	}
	return AggregatedBindingError{Errors: c.errors}
}

type requestParamBinder[TRawVal any, TTargetVal any] func(
	bindingCtx *bindingContext,
	rawVal optionalVal[TRawVal],
	receiver *TTargetVal,
)

type valueValidator[TRawVal any, TTargetVal any] func(optionalVal[TRawVal], TTargetVal) error

func validateNonEmpty[TRawVal any, TTargetVal any](rawVal optionalVal[TRawVal], _ TTargetVal) error {
	if !rawVal.assigned {
		return ErrValueRequired
	}
	return nil
}

var _ valueValidator[string, string] = validateNonEmpty

func newMinMaxValueValidator[TRawVal any, TTargetVal constraints.Ordered](
	threshold TTargetVal,
	exclusive bool,
	isMin bool,
) valueValidator[TRawVal, TTargetVal] {
	return func(ov optionalVal[TRawVal], tv TTargetVal) error {
		if !ov.assigned {
			return nil
		}

		// From OpenAPI spec:
		// exclusiveMinimum: false or not included	value ≥ minimum
		// exclusiveMinimum: true	value > minimum
		// exclusiveMaximum: false or not included	value ≤ maximum
		// exclusiveMaximum: true	value < maximum

		if isMin && ((exclusive && tv <= threshold) || (!exclusive && tv < threshold)) {
			return fmt.Errorf("value %v is less than minimum %v: %w", tv, threshold, ErrInvalidValueOutOfRange)
		}
		if !isMin && ((exclusive && tv >= threshold) || (!exclusive && tv > threshold)) {
			return fmt.Errorf("value %v is greater than maximum %v: %w", tv, threshold, ErrInvalidValueOutOfRange)
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
			bindingCtx.errors = append(bindingCtx.errors, BindingError{
				Field:    params.field,
				Location: params.location,
				Code:     ErrBadValueFormat,
				Err:      err,
			})
			return
		}
		if err := params.validateValue(rawVal, *receiver); err != nil {
			errCode := ErrInvalidValue
			errors.As(err, &errCode)
			bindingCtx.errors = append(bindingCtx.errors, BindingError{
				Field:    params.field,
				Location: params.location,
				Code:     errCode,
				Err:      err,
			})
		}
	}
}
