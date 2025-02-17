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
	"strings"
	"time"

	"github.com/gemyago/apigen/tests/golang/routes/internal"
	"golang.org/x/exp/constraints"
)

// TrackedResponseWriter allows checking if headers or body were written.
type TrackedResponseWriter interface {
	// HeaderWritten returns true if headers were written.
	HeaderWritten() bool

	// BodyWritten returns true if body was written.
	BodyWritten() bool
}

type httpRouter interface {
	// PathValue returns a named path parameter of a given name
	PathValue(r *http.Request, paramName string) string

	// HandleRoute register a given handler function to handle given route
	HandleRoute(method, pathPattern string, h http.Handler)
}

// ParsingErrorHandler will process errors during parsing and validation stages
// The default implementation will respond with 400 status code and standard
// serialization of ParsingError type.
// The writer may implement TrackedResponseWriter that can be used to check if
// headers or body were written.
type ParsingErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

// ActionErrorHandler will process errors produced by controller actions
// The default implementation will respond with 500 and no output.
// The writer may implement TrackedResponseWriter that can be used to check if
// headers or body were written.
type ActionErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

// ResponseErrorHandler will process errors that may occur while writing response.
// The writer may implement TrackedResponseWriter that can be used to check if
// headers or body were written. If headers are written it then usually reasonable
// to either log error or panic.
type ResponseErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

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
	app.handleActionErrors = func(w http.ResponseWriter, r *http.Request, err error) {
		app.logger.LogAttrs(r.Context(), slog.LevelError, "Failed to process request", slog.Any("error", err))
		if tw, ok := w.(TrackedResponseWriter); ok && !tw.HeaderWritten() {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	app.handleResponseErrors = func(w http.ResponseWriter, r *http.Request, err error) {
		app.logger.LogAttrs(r.Context(), slog.LevelError, "Failed to write response", slog.Any("err", err))
		if tw, ok := w.(TrackedResponseWriter); ok && !tw.HeaderWritten() {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	app.handleParsingErrors = func(w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		app.logger.LogAttrs(r.Context(), slog.LevelWarn, "Failed to parse request", slog.Any("err", err))
		var aggregatedErr internal.AggregatedBindingError
		if ok := errors.As(err, &aggregatedErr); ok {
			if writeErr := json.NewEncoder(w).Encode(aggregatedErr); writeErr != nil {
				app.handleResponseErrors(w, r, writeErr)
			}
			return
		}
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

type paramsParser[TReqParams any] interface {
	parse(router httpRouter, req *http.Request) (TReqParams, error)
}

type void *int

type voidParamsParser struct{}

func (p voidParamsParser) parse(_ httpRouter, _ *http.Request) (void, error) {
	return void(nil), nil
}

func makeVoidParamsParser(_ *HTTPApp) paramsParser[void] {
	return voidParamsParser{}
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
	// We may need a different method to check if the body is empty
	// if content length approach will be causing issues. For this case
	// best would be to read the body to buffer and check its length. It will be fully consumed anyway.
	if req.ContentLength > 0 {
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

func parseDecInt[TInt constraints.Integer](str string, bitSize int) (TInt, error) {
	res, err := strconv.ParseInt(str, 10, bitSize)
	return (TInt)(res), err
}

func parseFloat[TFloat constraints.Float](str string, bitSize int) (TFloat, error) {
	res, err := strconv.ParseFloat(str, bitSize)
	return (TFloat)(res), err
}

func parseMultiValueParamAsSoloValue[TTargetVal any](
	targetParser rawValueParser[string, TTargetVal],
) rawValueParser[[]string, TTargetVal] {
	return func(s []string, tv *TTargetVal) error {
		return targetParser(s[0], tv)
	}
}

// parseSoloValueParamAsSoloValue is noop and used just to simplify templates.
func parseSoloValueParamAsSoloValue[TRawVal any, TTargetVal any](
	targetParser rawValueParser[TRawVal, TTargetVal],
) rawValueParser[TRawVal, TTargetVal] {
	return targetParser
}

// parseSoloValueParamAsSlice will always parse the input as csv string.
func parseSoloValueParamAsSlice[TTargetVal any](
	targetParser rawValueParser[string, TTargetVal],
) rawValueParser[string, []TTargetVal] {
	return func(s string, tv *[]TTargetVal) error {
		items := strings.Split(s, ",")
		resultingSlice := make([]TTargetVal, 0, len(items))
		for i, item := range items {
			var val TTargetVal
			if err := targetParser(item, &val); err != nil {
				return fmt.Errorf("failed to parse value at index %d: %w", i, err)
			}
			resultingSlice = append(resultingSlice, val)
		}
		*tv = resultingSlice
		return nil
	}
}

// parseMultiValueParamAsSlice will parse each value in the input slice separately.
func parseMultiValueParamAsSlice[TTargetVal any](
	targetParser rawValueParser[string, TTargetVal],
) rawValueParser[[]string, []TTargetVal] {
	return func(s []string, tv *[]TTargetVal) error {
		resultingSlice := make([]TTargetVal, 0, len(s))
		for i, item := range s {
			var val TTargetVal
			if err := targetParser(item, &val); err != nil {
				return fmt.Errorf("failed to parse value at index %d: %w", i, err)
			}
			resultingSlice = append(resultingSlice, val)
		}
		*tv = resultingSlice
		return nil
	}
}

func parseNullableParam[TTargetVal any](
	targetParser rawValueParser[string, TTargetVal],
) rawValueParser[string, *TTargetVal] {
	return func(s string, tv **TTargetVal) error {
		if s == "" || s == "null" {
			return nil
		}
		*tv = new(TTargetVal)
		return targetParser(s, *tv)
	}
}

type knownParsersDef struct {
	stringParser  rawValueParser[string, string]
	dateParser    rawValueParser[string, time.Time]
	timeParser    rawValueParser[string, time.Time]
	int32Parser   rawValueParser[string, int32]
	int64Parser   rawValueParser[string, int64]
	float32Parser rawValueParser[string, float32]
	float64Parser rawValueParser[string, float64]
	boolParser    rawValueParser[string, bool]
}

const bitSize32 = 32
const bitSize64 = 64

func newKnownParsers() *knownParsersDef {
	return &knownParsersDef{
		stringParser: func(ov string, s *string) error {
			*s = ov
			return nil
		},
		dateParser:    newStringToDateTimeParser(true),
		timeParser:    newStringToDateTimeParser(false),
		int32Parser:   newStringToNumberParser[int32](bitSize32, parseDecInt),
		int64Parser:   newStringToNumberParser[int64](bitSize64, parseDecInt),
		float32Parser: newStringToNumberParser[float32](bitSize32, parseFloat),
		float64Parser: newStringToNumberParser(bitSize64, strconv.ParseFloat),
		boolParser: func(ov string, s *bool) error {
			switch ov {
			case "true":
				*s = true
			case "false":
				*s = false
			default:
				return fmt.Errorf("unexpected boolean format %v", ov)
			}
			return nil
		},
	}
}

type requestParamBinder[TRawVal any, TTargetVal any] func(
	bindingCtx *internal.BindingContext,
	rawVal internal.OptionalVal[TRawVal],
	receiver *TTargetVal,
)

type binderParams[TRawVal any, TTargetVal any] struct {
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
					Location: bindingCtx.BuildPath(),
					Code:     internal.ErrValueRequired.Error(),
				})
			}
			return
		}
		if err := params.parseValue(rawVal.Value, receiver); err != nil {
			bindingCtx.AppendFieldError(internal.FieldBindingError{
				Location: bindingCtx.BuildPath(),
				Code:     internal.ErrBadValueFormat.Error(),
				Err:      err,
			})
			return
		}
		params.validateValue(bindingCtx, *receiver)
	}
}

// ActionHandlerFunc represents possible combination of action handler functions.
// Each function can be with or without parameters and with or without response.
// Additionally each function can have access to http objects for possible direct manipulation.
type ActionHandlerFunc[TReq any, TRes any] interface {
	func(context.Context, TReq) (TRes, error) | // with params with response
		func(context.Context) (TRes, error) | // no params with response
		func(context.Context, TReq) error | // with params no response
		func(context.Context) error | // no params no response

		// handlers with http context exposed
		func(http.ResponseWriter, *http.Request, TReq) (TRes, error) | // with params with response
		func(http.ResponseWriter, *http.Request) (TRes, error) | // no params with response
		func(http.ResponseWriter, *http.Request, TReq) error | // with params no response
		func(http.ResponseWriter, *http.Request) error // no params no response
}

type ActionBuilderFunc[TReq any, TRes any, THandler ActionHandlerFunc[TReq, TRes]] func(THandler) http.Handler

type ActionBuilder[
	TReq any,
	TRes any,
	TPlainHandler ActionHandlerFunc[TReq, TRes],
	THttpHandler ActionHandlerFunc[TReq, TRes],
] struct {
	// HandleWith creates a new http.Handler from a given action handler.
	//
	// The action handler is not supposed to have access to http objects and
	// in most scenarios can just delegate the work to the application logic layer.
	// If you need access to http objects use HandleWithHTTP
	HandleWith func(TPlainHandler) http.Handler

	// HandleWithHTTP creates a new http.Handler from a given action handler.
	//
	// The action handler allows direct access to http.ResponseWriter and *http.Request.
	// It also provides parsed request parameters and allows sending structured response.
	// If you need fully customized behavior, feel free not to use the builder and
	// return the handler directly.
	HandleWithHTTP func(THttpHandler) http.Handler
}

type genericHandlerBuilder[
	TReq any,
	TRes any,
	TPlainHandler ActionHandlerFunc[TReq, TRes],
	THttpHandler ActionHandlerFunc[TReq, TRes],
] interface {
	// HandleWith creates a new http.Handler from a given func.
	//
	// The action handler is not supposed to have access to http objects and
	// in most scenarios can just delegate the work to the application logic layer.
	// If you need access to http objects use HandleWithHTTP
	HandleWith(TPlainHandler) http.Handler

	// HandleWithHTTP creates a new http.Handler from a given func.
	//
	// The action handler allows direct access to http.ResponseWriter and *http.Request.
	// It also provides parsed request parameters and allows sending structured response.
	// If you need fully customized behavior, feel free not to use the builder and
	// return the handler directly.
	HandleWithHTTP(THttpHandler) http.Handler
}

type HandlerBuilder[TReq any, TRes any] genericHandlerBuilder[
	TReq,
	TRes,
	func(context.Context, TReq) (TRes, error),
	func(http.ResponseWriter, *http.Request, TReq) (TRes, error),
]

type NoParamsHandlerBuilder[TRes any] genericHandlerBuilder[
	struct{},
	TRes,
	func(context.Context) (TRes, error),
	func(http.ResponseWriter, *http.Request) (TRes, error),
]

type NoResponseHandlerBuilder[TReq any] genericHandlerBuilder[
	TReq,
	struct{},
	func(context.Context, TReq) (struct{}, error),
	func(http.ResponseWriter, *http.Request, TReq) (struct{}, error),
]

type NoParamsNoResponseHandlerBuilder genericHandlerBuilder[
	struct{},
	struct{},
	func(context.Context) error,
	func(http.ResponseWriter, *http.Request) error,
]

/* TODO: implement this function
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
*/

type makeActionBuilderParams[
	TReqParams any,
	TResData any,
] struct {
	defaultStatus int
	voidResult    bool
	paramsParser  paramsParser[TReqParams]
}

type actionBuilderHandlerAdapter[
	TReq any,
	TRes any,
	THandler ActionHandlerFunc[TReq, TRes],
] func(THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error)

func newHandlerAdapter[
	TReq any,
	TRes any,
	THandler func(context.Context, TReq) (TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(_ http.ResponseWriter, r *http.Request, req TReq) (TRes, error) {
			return t(r.Context(), req)
		}
	}
}

func newHandlerAdapterNoParams[
	TReq any,
	TRes any,
	THandler func(context.Context) (TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(_ http.ResponseWriter, r *http.Request, _ TReq) (TRes, error) {
			return t(r.Context())
		}
	}
}

func newHandlerAdapterNoResponse[
	TReq any,
	TRes any,
	THandler func(context.Context, TReq) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(_ http.ResponseWriter, r *http.Request, req TReq) (TRes, error) {
			var emptyRes TRes
			if err := t(r.Context(), req); err != nil {
				return emptyRes, err
			}
			return emptyRes, nil
		}
	}
}

func newHandlerAdapterNoParamsNoResponse[
	TReq any,
	TRes any,
	THandler func(context.Context) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(_ http.ResponseWriter, r *http.Request, _ TReq) (TRes, error) {
			var emptyRes TRes
			return emptyRes, t(r.Context())
		}
	}
}

func newHTTPHandlerAdapter[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request, TReq) (TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, req TReq) (TRes, error) {
			return t(w, r, req)
		}
	}
}

func newHTTPHandlerAdapterNoParams[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request) (TRes, error),
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, _ TReq) (TRes, error) {
			return t(w, r)
		}
	}
}

func newHTTPHandlerAdapterNoResponse[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request, TReq) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, req TReq) (TRes, error) {
			var emptyRes TRes
			if err := t(w, r, req); err != nil {
				return emptyRes, err
			}
			return emptyRes, nil
		}
	}
}

func newHTTPHandlerAdapterNoParamsNoResponse[
	TReq any,
	TRes any,
	THandler func(http.ResponseWriter, *http.Request) error,
]() actionBuilderHandlerAdapter[TReq, TRes, THandler] {
	return func(t THandler) func(http.ResponseWriter, *http.Request, TReq) (TRes, error) {
		return func(w http.ResponseWriter, r *http.Request, _ TReq) (TRes, error) {
			var emptyRes TRes
			return emptyRes, t(w, r)
		}
	}
}

type actionsResponseWriter struct {
	targetWriter  http.ResponseWriter
	headerWritten bool
	bodyWritten   bool
	defaultStatus int
}

func (w *actionsResponseWriter) HeaderWritten() bool {
	return w.headerWritten
}

func (w *actionsResponseWriter) BodyWritten() bool {
	return w.bodyWritten
}

func (w *actionsResponseWriter) Header() http.Header {
	return w.targetWriter.Header()
}

func (w *actionsResponseWriter) WriteHeader(statusCode int) {
	w.headerWritten = true
	w.targetWriter.WriteHeader(statusCode)
}

func (w *actionsResponseWriter) Write(data []byte) (int, error) {
	if !w.headerWritten {
		w.WriteHeader(w.defaultStatus)
	}
	w.bodyWritten = true
	return w.targetWriter.Write(data)
}

var _ TrackedResponseWriter = &actionsResponseWriter{}
var _ http.ResponseWriter = &actionsResponseWriter{}

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
	createHandler := func(handler func(http.ResponseWriter, *http.Request, TReq) (TRes, error)) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			aw := &actionsResponseWriter{
				targetWriter:  w,
				defaultStatus: params.defaultStatus,
			}

			reqParams, err := params.paramsParser.parse(app.router, r)
			if err != nil {
				app.handleParsingErrors(aw, r, err)
				return
			}

			resData, err := handler(aw, r, reqParams)
			if err != nil {
				app.handleActionErrors(aw, r, err)
				return
			}

			if params.voidResult {
				// Do not write header twice
				if !aw.HeaderWritten() {
					aw.WriteHeader(params.defaultStatus)
				}
				return
			}

			// This means the action handler has written the response itself.
			if aw.BodyWritten() {
				return
			}

			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			// Not sending the status here. The action writer will send it in case of
			// success. If error has happened while encoding, then the error handler will have
			// a chance to set the status.
			if encodingErr := json.NewEncoder(aw).Encode(resData); encodingErr != nil {
				app.handleResponseErrors(aw, r, encodingErr)
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
