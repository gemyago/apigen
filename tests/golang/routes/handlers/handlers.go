package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

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
		router: router,
		logger: slog.Default(),
	}
	app.handleResponseErrors = func(r *http.Request, err error) {
		app.logger.LogAttrs(r.Context(), slog.LevelError, "Failed to write response", slog.Any("err", err))
	}
	app.handleParsingErrors = func(r *http.Request, w http.ResponseWriter, err error) {
		w.Header().Add("content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		app.logger.LogAttrs(r.Context(), slog.LevelWarn, "Failed to parse request", slog.Any("err", err))
		var aggregatedErr AggregatedBindingError
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

type handlerFactoryParams[TReqParams any, TResData any] struct {
	defaultStatus int
	voidResult    bool
	paramsParser  paramsParser[TReqParams]
	handler       func(context.Context, TReqParams) (TResData, error)
}

func createHandlerFactory[TReqParams any, TResData any](
	factoryParams handlerFactoryParams[TReqParams, TResData],
) httpHandlerFactory {
	return func(app *HTTPApp) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			params, err := factoryParams.paramsParser.parse(app.router, r)
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
	defaultStatusCode  int
	voidResult         bool
	httpHandlerFactory func(app *HTTPApp) http.Handler
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

func parseJSONPayload[TTargetVal any](req optionalVal[*http.Request], target *TTargetVal) error {
	return json.NewDecoder(req.value.Body).Decode(target)
}

var _ rawValueParser[*http.Request, string] = parseJSONPayload

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

func newStringToDateTimeParser(isDateOnly bool) rawValueParser[string, time.Time] {
	return func(ov optionalVal[string], t *time.Time) error {
		if !ov.assigned {
			return nil
		}
		format := time.RFC3339Nano
		if isDateOnly {
			format = time.DateOnly
		}
		val, err := time.Parse(format, ov.value)
		if err != nil {
			return err
		}
		*t = val
		return nil
	}
}

func newStringSliceToDateTimeParser(isDateOnly bool) rawValueParser[[]string, time.Time] {
	return func(ov optionalVal[[]string], t *time.Time) error {
		if !ov.assigned {
			return nil
		}
		format := time.RFC3339Nano
		if isDateOnly {
			format = time.DateOnly
		}
		val, err := time.Parse(format, ov.value[0])
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

func parseStringInPath(ov optionalVal[string], s *string) error {
	if !ov.assigned {
		return nil
	}
	*s = ov.value
	return nil
}

func parseStringInQuery(ov optionalVal[[]string], s *string) error {
	if !ov.assigned {
		return nil
	}
	*s = ov.value[0]
	return nil
}

func parseBoolInPath(ov optionalVal[string], s *bool) error {
	if !ov.assigned {
		return nil
	}
	switch ov.value {
	case "true":
		*s = true
	case "false":
		*s = false
	default:
		return fmt.Errorf("unexpected boolean format %v", ov.value)
	}
	return nil
}

func parseBoolInQuery(ov optionalVal[[]string], s *bool) error {
	if !ov.assigned {
		return nil
	}
	return parseBoolInPath(optionalVal[string]{value: ov.value[0], assigned: true}, s)
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

var knownParsers = knownParsersDef{
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

type BindingError string

const (
	// ErrBadValueFormat error means data provided can not be parsed to a target type.
	ErrBadValueFormat BindingError = "BAD_FORMAT"

	// ErrValueRequired error code indicates that the required value has not been provided.
	ErrValueRequired BindingError = "INVALID_REQUIRED"

	// ErrInvalidValueOutOfRange error code indicates that the value is out of range of allowable values
	// this is usually when number is out of min/max range, or string is outside of limits.
	ErrInvalidValueOutOfRange BindingError = "INVALID_OUT_OF_RANGE"

	// ErrInvalidValue error code a generic validation error.
	ErrInvalidValue BindingError = "INVALID"
)

func (c BindingError) Error() string {
	return string(c)
}

// FieldBindingError occurs at parsing/validation stage and holds
// context on field that the error is related to.
type FieldBindingError struct {
	Field    string       `json:"field"`
	Location string       `json:"location"`
	Err      error        `json:"-"`
	Code     BindingError `json:"code"`
}

func (be FieldBindingError) Error() string {
	return fmt.Sprintf("field %s (in %s) code=%s, error: %v", be.Field, be.Location, be.Code, be.Err)
}

type AggregatedBindingError struct {
	Errors []FieldBindingError `json:"errors"`
}

func (c AggregatedBindingError) Error() string {
	errs := make([]error, len(c.Errors))
	for i, err := range c.Errors {
		errs[i] = err
	}
	return errors.Join(errs...).Error()
}

type bindingContext struct {
	errors []FieldBindingError
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

func newMinMaxLengthValidator[TRawVal any, TTargetVal string](
	threshold int,
	isMin bool,
) valueValidator[TRawVal, TTargetVal] {
	return func(ov optionalVal[TRawVal], tv TTargetVal) error {
		if !ov.assigned {
			return nil
		}

		targetLen := len(tv)
		if isMin && targetLen < threshold {
			return fmt.Errorf(
				"value %v has length (%d) less than minimum %v: %w",
				tv, targetLen, threshold, ErrInvalidValueOutOfRange,
			)
		}
		if !isMin && targetLen > threshold {
			return fmt.Errorf(
				"value %v has length (%d) more than maximum %v: %w", tv,
				targetLen, threshold, ErrInvalidValueOutOfRange,
			)
		}

		return nil
	}
}

func newPatternValidator[TRawVal any, TTargetValue string](patternStr string) valueValidator[TRawVal, string] {
	pattern := regexp.MustCompile(patternStr)
	return func(ov optionalVal[TRawVal], tv string) error {
		if !ov.assigned {
			return nil
		}
		if !pattern.MatchString(tv) {
			return fmt.Errorf("value %v does not match pattern %v: %w", tv, patternStr, ErrInvalidValue)
		}
		return nil
	}
}

func newCompositeValidator[
	TRawVal any, TTargetVal any,
](validators ...valueValidator[TRawVal, TTargetVal]) valueValidator[TRawVal, TTargetVal] {
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
			bindingCtx.errors = append(bindingCtx.errors, FieldBindingError{
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
			bindingCtx.errors = append(bindingCtx.errors, FieldBindingError{
				Field:    params.field,
				Location: params.location,
				Code:     errCode,
				Err:      err,
			})
		}
	}
}
