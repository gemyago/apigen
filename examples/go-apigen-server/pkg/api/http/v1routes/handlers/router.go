package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/exp/constraints"
)

// TODO: Maybe move this to another file and also rename to something else.
type httpRouter interface {
	// PathValue returns a named path parameter of a given name
	PathValue(r *http.Request, paramName string) string

	// HandleRoute register a given handler function to handle given route
	HandleRoute(method, pathPattern string, h http.Handler)

	// HandleError will be called for any error produced by handlers
	HandleError(r *http.Request, w http.ResponseWriter, err error)
}

type bindingError struct {
	field    string
	location string
	err      error
}

type bindingContext struct {
	errors []bindingError
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

type requestParamBinder[TRawVal any, TTargetVal any] func(
	bindingCtx *bindingContext,
	rawVal optionalVal[TRawVal],
	receiver *TTargetVal,
)

type rawValueParser[TRawVal any, TTargetVal any] func(optionalVal[TRawVal], *TTargetVal) error

func parseJsonPayload[TTargetVal any](req optionalVal[*http.Request], target *TTargetVal) error {
	return json.NewDecoder(req.value.Body).Decode(target)
}

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
