package router

import (
	"errors"
	"log"
	"net/http"

	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/api/http/controllers"
	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/api/http/routes/handlers"
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
func handleActionError(w http.ResponseWriter, _ *http.Request, err error) {
	code := 500
	switch {
	case errors.Is(err, controllers.ErrNotFound):
		code = 404
	case errors.Is(err, controllers.ErrConflict):
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

// Minimalistic access log middleware just to have some logs when running the example.
func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET %v %v %v\n", r.URL.String(), r.Proto, r.UserAgent())
		defer func() {
			if r := recover(); r != nil {
				log.Println("Request panic", r)
			}
		}()
		wrapper := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)
		log.Printf("%d %v %v\n", wrapper.statusCode, r.Method, r.URL.String())
	})
}

// HandlerDeps holds dependencies of the generated routes
// usually controller implementations at least.
type HandlerDeps struct {
	PetsController *controllers.PetsController
}

// NewHandler creates an minimal example implementation of the router handler
// based on the standard http.ServeMux.
func NewHandler(deps HandlerDeps) http.Handler {
	// Root handler instance is a central place to register all routes
	rootHandler := handlers.NewRootHandler(
		(*httpRouter)(http.NewServeMux()),
		handlers.WithActionErrorHandler(handleActionError),
	)

	// Register generated Pets routes.
	rootHandler.RegisterPetsRoutes(deps.PetsController)

	// Root handler is a standard http.Handler so can be used in any
	// context that expects http.Handler interface.
	return accessLogMiddleware(rootHandler)
}
