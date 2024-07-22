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
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/internal"
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

// SlogLogger is a fully compatible with slog and used to allow injecting the instance.
type SlogLogger interface {
	Log(ctx context.Context, level slog.Level, msg string, args ...any)
	LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
}

type HTTPApp struct {
	router               httpRouter
	handleParsingErrors  ParsingErrorHandler
	handleActionErrors   ActionErrorHandler
	handleResponseErrors ResponseErrorHandler
	knownParsers         *knownParsersDef
	logger               SlogLogger
}

type HTTPAppOpt func(app *HTTPApp)

func WithParsingErrorHandler(handler ParsingErrorHandler) HTTPAppOpt {
	return func(app *HTTPApp) {
		app.handleParsingErrors = handler
	}
}

func WithActionErrorHandler(handler ActionErrorHandler) HTTPAppOpt {
	return func(app *HTTPApp) {
		app.handleActionErrors = handler
	}
}

func WithResponseErrorHandler(handler ResponseErrorHandler) HTTPAppOpt {
	return func(app *HTTPApp) {
		app.handleResponseErrors = handler
	}
}

func WithLogger(logger SlogLogger) HTTPAppOpt {
	return func(app *HTTPApp) {
		app.logger = logger
	}
}

func NewHTTPApp(router httpRouter, opts ...HTTPAppOpt) *HTTPApp {
	app := &HTTPApp{
		router:       router,
		logger:       slog.Default(),
		knownParsers: newKnownParsers(),
	}
	app.handleResponseErrors = func(r *http.Request, err error) {
		app.logger.LogAttrs(r.Context(), slog.LevelError, "Failed to write response", slog.Any("err", err))
	}
	app.handleParsingErrors = func(r *http.Request, w http.ResponseWriter, err error) {
		w.Header().Add("content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		app.logger.LogAttrs(r.Context(), slog.LevelWarn, "Failed to parse request", slog.Any("err", err))
		var aggregatedErr internal.AggregatedBindingError
		if ok := errors.As(err, &aggregatedErr); ok {
			if writeErr := json.NewEncoder(w).Encode(aggregatedErr); writeErr != nil {
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

type httpHandlerFactory func(app *HTTPApp) http.Handler
type paramsParser[TReqParams any] interface {
	parse(router httpRouter, req *http.Request) (TReqParams, error)
}
type paramsParserFactory[TReqParams any] func(app *HTTPApp) paramsParser[TReqParams]

type handlerFactoryParams[TReqParams any, TResData any] struct {
	defaultStatus       int
	voidResult          bool
	paramsParserFactory func(app *HTTPApp) paramsParser[TReqParams]
	handler             func(context.Context, TReqParams) (TResData, error)
}

func createHandlerFactory[TReqParams any, TResData any](
	factoryParams handlerFactoryParams[TReqParams, TResData],
) httpHandlerFactory {
	return func(app *HTTPApp) http.Handler {
		paramsParser := factoryParams.paramsParserFactory(app)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params, err := paramsParser.parse(app.router, r)
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
			if encodingErr := json.NewEncoder(w).Encode(resData); encodingErr != nil {
				app.handleResponseErrors(r, encodingErr)
			}
		})
	}
}

type actionBuilder[TControllerBuilder any, TReqParams any, TResData any] struct {
	defaultStatusCode   int
	voidResult          bool
	httpHandlerFactory  func(app *HTTPApp) http.Handler
	paramsParserFactory paramsParserFactory[TReqParams]
	controllerBuilder   TControllerBuilder
}

func (ab *actionBuilder[TControllerBuilder, TReqParams, TResData]) With(
	handler func(context.Context, TReqParams) (TResData, error),
) TControllerBuilder {
	ab.httpHandlerFactory = createHandlerFactory(handlerFactoryParams[TReqParams, TResData]{
		defaultStatus:       ab.defaultStatusCode,
		voidResult:          ab.voidResult,
		paramsParserFactory: ab.paramsParserFactory,
		handler:             handler,
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

func readPathValue(key string, router httpRouter, req *http.Request) internal.OptionalVal[string] {
	return internal.OptionalVal[string]{Value: router.PathValue(req, key), Assigned: true}
}

func readQueryValue(key string, values url.Values) internal.OptionalVal[[]string] {
	if values.Has(key) {
		return internal.OptionalVal[[]string]{Value: values[key], Assigned: true}
	}
	return internal.OptionalVal[[]string]{}
}

func readRequestBodyValue(req *http.Request) internal.OptionalVal[*http.Request] {
	if req.ContentLength != 0 {
		return internal.OptionalVal[*http.Request]{Value: req, Assigned: true}
	}
	return internal.OptionalVal[*http.Request]{}
}

type rawValueParser[TRawVal any, TTargetVal any] func(TRawVal, *TTargetVal) error

func parseJSONPayload[TTargetVal any](req *http.Request, target *TTargetVal) error {
	return json.NewDecoder(req.Body).Decode(target)
}

var _ rawValueParser[*http.Request, string] = parseJSONPayload

func newStringToNumberParser[TTargetVal constraints.Integer | constraints.Float](
	bitSize int, parseFn func(string, int) (TTargetVal, error),
) rawValueParser[string, TTargetVal] {
	return func(ov string, target *TTargetVal) error {
		val, err := parseFn(ov, bitSize)
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
	return func(ov []string, target *TTargetVal) error {
		val, err := parseFn(ov[0], bitSize)
		if err != nil {
			return err
		}
		*target = val
		return nil
	}
}

func newStringToDateTimeParser(isDateOnly bool) rawValueParser[string, time.Time] {
	return func(ov string, t *time.Time) error {
		format := time.RFC3339Nano
		if isDateOnly {
			format = time.DateOnly
		}
		val, err := time.Parse(format, ov)
		if err != nil {
			return err
		}
		*t = val
		return nil
	}
}

func newStringSliceToDateTimeParser(isDateOnly bool) rawValueParser[[]string, time.Time] {
	return func(ov []string, t *time.Time) error {
		format := time.RFC3339Nano
		if isDateOnly {
			format = time.DateOnly
		}
		val, err := time.Parse(format, ov[0])
		if err != nil {
			return err
		}
		*t = val
		return nil
	}
}

func parseDecInt[TInt constraints.Integer](str string, bitSize int) (TInt, error) {
	res, err := strconv.ParseInt(str, 10, bitSize)
	return (TInt)(res), err
}

func parseFloat[TFloat constraints.Float](str string, bitSize int) (TFloat, error) {
	res, err := strconv.ParseFloat(str, bitSize)
	return (TFloat)(res), err
}

func parseStringInPath(ov string, s *string) error {
	*s = ov
	return nil
}

func parseStringInQuery(ov []string, s *string) error {
	*s = ov[0]
	return nil
}

func parseBoolInPath(ov string, s *bool) error {
	switch ov {
	case "true":
		*s = true
	case "false":
		*s = false
	default:
		return fmt.Errorf("unexpected boolean format %v", ov)
	}
	return nil
}

func parseBoolInQuery(ov []string, s *bool) error {
	return parseBoolInPath(ov[0], s)
}

type knownParsersDef struct {
	// path
	stringInPath  rawValueParser[string, string]
	dateInPath    rawValueParser[string, time.Time]
	timeInPath    rawValueParser[string, time.Time]
	int32InPath   rawValueParser[string, int32]
	int64InPath   rawValueParser[string, int64]
	float32InPath rawValueParser[string, float32]
	float64InPath rawValueParser[string, float64]
	boolInPath    rawValueParser[string, bool]

	// query
	stringInQuery  rawValueParser[[]string, string]
	dateInQuery    rawValueParser[[]string, time.Time]
	timeInQuery    rawValueParser[[]string, time.Time]
	int32InQuery   rawValueParser[[]string, int32]
	int64InQuery   rawValueParser[[]string, int64]
	float32InQuery rawValueParser[[]string, float32]
	float64InQuery rawValueParser[[]string, float64]
	boolInQuery    rawValueParser[[]string, bool]
}

const bitSize32 = 32
const bitSize64 = 64

func newKnownParsers() *knownParsersDef {
	return &knownParsersDef{
		// path
		stringInPath:  parseStringInPath,
		dateInPath:    newStringToDateTimeParser(true),
		timeInPath:    newStringToDateTimeParser(false),
		int32InPath:   newStringToNumberParser[int32](bitSize32, parseDecInt),
		int64InPath:   newStringToNumberParser[int64](bitSize64, parseDecInt),
		float32InPath: newStringToNumberParser[float32](bitSize32, parseFloat),
		float64InPath: newStringToNumberParser(bitSize64, strconv.ParseFloat),
		boolInPath:    parseBoolInPath,

		// query
		stringInQuery:  parseStringInQuery,
		dateInQuery:    newStringSliceToDateTimeParser(true),
		timeInQuery:    newStringSliceToDateTimeParser(false),
		int32InQuery:   newStringSliceToNumberParser[int32](bitSize32, parseDecInt),
		int64InQuery:   newStringSliceToNumberParser[int64](bitSize64, parseDecInt),
		float32InQuery: newStringSliceToNumberParser[float32](bitSize32, parseFloat),
		float64InQuery: newStringSliceToNumberParser(bitSize64, strconv.ParseFloat),
		boolInQuery:    parseBoolInQuery,
	}
}

type requestParamBinder[TRawVal any, TTargetVal any] func(
	bindingCtx *internal.BindingContext,
	rawVal internal.OptionalVal[TRawVal],
	receiver *TTargetVal,
)

type binderParams[TRawVal any, TTargetVal any] struct {
	field         string
	location      string
	required      bool
	parseValue    rawValueParser[TRawVal, TTargetVal]
	validateValue internal.FieldValidator[TTargetVal]
}

func newRequestParamBinder[TRawVal any, TTargetVal any](
	params binderParams[TRawVal, TTargetVal],
) requestParamBinder[TRawVal, TTargetVal] {
	return func(
		bindingCtx *internal.BindingContext,
		rawVal internal.OptionalVal[TRawVal],
		receiver *TTargetVal,
	) {
		if !rawVal.Assigned {
			if params.required {
				bindingCtx.AppendFieldError(internal.FieldBindingError{
					Field:    params.field,
					Location: params.location,
					Code:     internal.ErrValueRequired.Error(),
				})
			}
			return
		}
		if err := params.parseValue(rawVal.Value, receiver); err != nil {
			bindingCtx.AppendFieldError(internal.FieldBindingError{
				Field:    params.field,
				Location: params.location,
				Code:     internal.ErrBadValueFormat.Error(),
				Err:      err,
			})
			return
		}
		params.validateValue(bindingCtx, params.field, params.location, *receiver)
	}
}
