package controllers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/api/http/routes/handlers"
	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app"
)

// httpRouter is a ServerMux adapter to use with generated routes.
type httpRouter http.ServeMux

func (*httpRouter) PathValue(r *http.Request, paramName string) string {
	return r.PathValue(paramName)
}

func (router *httpRouter) HandleRoute(method, pathPattern string, h http.Handler) {
	(*http.ServeMux)(router).Handle(method+" "+pathPattern, h)
}

func (router *httpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(*http.ServeMux)(router).ServeHTTP(w, r)
}

// handleActionError is a custom error handler to process action errors.
// It's a good place to map errors returned by controller actions to HTTP status codes.
func newActionErrorHandler(logger *slog.Logger) func(w http.ResponseWriter, _ *http.Request, err error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		ctx := req.Context()
		code := 500
		switch {
		case errors.Is(err, app.ErrNotFound):
			code = 404
		case errors.Is(err, app.ErrConflict):
			code = 409
		}
		w.WriteHeader(code)
		logger.InfoContext(ctx, "Failed to process request", slog.Any("error", err))
	}
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *responseWriterWrapper) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Minimalistic access log middleware just to have some logs when running the example.
func accessLogMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger.InfoContext(ctx, fmt.Sprintf("%v %v %v %v", r.Method, r.URL.String(), r.Proto, r.UserAgent()))
		defer func() {
			if r := recover(); r != nil {
				logger.ErrorContext(ctx, "Request panic", slog.Any("panic", r))
			}
		}()
		wrapper := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)
		logger.InfoContext(ctx, fmt.Sprintf("%d %v %v", wrapper.statusCode, r.Method, r.URL.String()))
	})
}

// RoutesDeps is a set of dependencies required to setup routes.
// Usually that would include application layer services and other components
// required to handle requests.
type RoutesDeps struct {
	RootLogger *slog.Logger

	PetsService petsService
}

// SetupRoutes is an example setup of generated routes.
func SetupRoutes(deps RoutesDeps) http.Handler {
	httpLogger := deps.RootLogger.WithGroup("http")

	// Root handler instance is a central place to register all routes
	rootHandler := handlers.NewRootHandler(
		(*httpRouter)(http.NewServeMux()),
		handlers.WithActionErrorHandler(newActionErrorHandler(httpLogger)),
	)

	// Register generated Pets routes.
	rootHandler.RegisterPetsRoutes(&petsController{
		petsService: deps.PetsService,
	})

	// Root handler is a standard http.Handler so can be used in any
	// context that expects http.Handler interface.
	return accessLogMiddleware(httpLogger, rootHandler)
}
