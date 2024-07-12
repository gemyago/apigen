package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gemyago/apigen/tests/golang/routes/handlers"
	"github.com/stretchr/testify/assert"
)

type routerAdapter struct {
	mux           *http.ServeMux
	handledErrors []error
}

func (r *routerAdapter) PathValue(req *http.Request, paramName string) string {
	return req.PathValue(paramName)
}

func (r *routerAdapter) HandleRoute(method, pathPattern string, h http.Handler) {
	r.mux.Handle(method+" "+pathPattern, h)
}

func (r *routerAdapter) HandleError(req *http.Request, w http.ResponseWriter, err error) {
	r.handledErrors = append(r.handledErrors, err)
	w.WriteHeader(http.StatusInternalServerError)
}

// openTestLogFile will open a log file in a project root directory
func openTestLogFile() *os.File {
	_, filename, _, _ := runtime.Caller(0) // Will be current file
	testFilePath := filepath.Join(filename, "..", "..", "..", "..", "golang-test.log")
	f, err := os.OpenFile(testFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666)
	if err != nil {
		err = fmt.Errorf("fail to open log file %q for test logging: %w", testFilePath, err)
		panic(err)
	}
	return f
}

var testOutput = openTestLogFile()

func newLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(testOutput, nil))
}

func unmarshalBindingErrors(
	t *testing.T,
	body *bytes.Buffer,
) *handlers.AggregatedBindingError {
	var gotErrors handlers.AggregatedBindingError
	if err := json.Unmarshal(body.Bytes(), &gotErrors); !assert.NoError(t, err) {
		t.FailNow()
		return nil
	}
	return &gotErrors
}

func assertFieldError(
	t *testing.T,
	err *handlers.AggregatedBindingError,
	location string,
	field string,
	code handlers.BindingErrorCode,
) bool {
	for _, fieldErr := range err.Errors {
		if fieldErr.Location == location && fieldErr.Field == field {
			return assert.Equal(t, code, fieldErr.Code, "field %s: unexpected error code", field)
		}
	}
	assert.Fail(t, fmt.Sprintf("no error found for field %s, code %s", field, code))
	return false
}
