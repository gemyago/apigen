package router

import (
	"errors"
	"log"
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/app"
)

// httpRouter is a ServerMux adapter to use with generated routes.
type httpRouter struct {
	*http.ServeMux
}

func (httpRouter) PathValue(r *http.Request, paramName string) string {
	return r.PathValue(paramName)
}

func (a httpRouter) HandleRoute(method, pathPattern string, h http.Handler) {
	a.ServeMux.Handle(method+" "+pathPattern, h)
}

// handleActionError is a custom error handler to process action errors.
// It's a good place to map errors returned by controller actions to HTTP status codes.
func handleActionError(w http.ResponseWriter, _ *http.Request, err error) {
	code := 500
	switch {
	case errors.Is(err, app.ErrNotFound):
		code = 404
	case errors.Is(err, app.ErrConflict):
		code = 409
	}
	w.WriteHeader(code)
	log.Printf("Failed to process request: %v\n", err)
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *responseWriterWrapper) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// HandlerDeps holds dependencies of the generated routes
// usually controller implementations at least.
type HandlerDeps struct {
	PetsController handlers.PetsController
}

// NewHandler creates an minimal example implementation of the router handler
// based on the standard http.ServeMux.
func NewHandler(deps HandlerDeps) http.Handler {
	mux := http.NewServeMux()

	// Real world instance of the router handler will likely to have a more advanced setup
	// that may include various middleware to authenticate, log, trace e.t.c...
	// Below is a simple access logs on requests processing
	muxHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET %v %v %v\n", r.URL.String(), r.Proto, r.UserAgent())
		defer func() {
			if r := recover(); r != nil {
				log.Println("Request panic", r)
			}
		}()
		wrapper := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		mux.ServeHTTP(wrapper, r)
		log.Printf("%d %v %v\n", wrapper.statusCode, r.Method, r.URL.String())
	})

	// The httpApp provides a configuration layer of the generated routes
	// and also serves as an adapter that allows using different router implementations
	httpApp := handlers.NewHTTPApp(
		httpRouter{ServeMux: mux},
		handlers.WithActionErrorHandler(handleActionError),
	)

	// Register generated Pets routes. There can be multiple different
	// routes registered into the same httpApp instance.
	handlers.RegisterPetsRoutes(deps.PetsController, httpApp)

	return muxHandler
}
