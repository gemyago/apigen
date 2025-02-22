package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gemyago/apigen/examples/ping-server-go/internal/api/http/v1routes/handlers"
	"github.com/gemyago/apigen/examples/ping-server-go/internal/api/http/v1routes/models"
)

// Start it using command below:
// go run ./examples/ping-server-go/cmd/service/
//
// Send requests with curl:
// - curl localhost:8080/ping
// - curl localhost:8080/ping?message=hello

type pingController struct{}

func (c *pingController) Ping(b handlers.HandlerBuilder[
	*handlers.PingPingRequest,
	*models.Ping200Response,
]) http.Handler {
	return b.HandleWith(
		func(_ context.Context, req *handlers.PingPingRequest) (*models.Ping200Response, error) {
			message := req.Message
			if message == "" {
				message = "pong"
			}
			return &models.Ping200Response{Message: message}, nil
		},
	)
}

// pingController must implement handlers.PingController interface.
var _ handlers.PingController = (*pingController)(nil)

// Simple adapter to use http.ServeMux with generated routes.
type httpRouter http.ServeMux

func (*httpRouter) PathValue(r *http.Request, paramName string) string {
	return r.PathValue(paramName)
}

func (r *httpRouter) HandleRoute(method, pathPattern string, h http.Handler) {
	mux := (*http.ServeMux)(r)
	mux.Handle(method+" "+pathPattern, h)
}

func (r *httpRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux := (*http.ServeMux)(r)
	mux.ServeHTTP(w, req)
}

func main() {
	port := 8080
	readHeaderTimeoutSec := 2

	// TODO: Rename HTTPApp to something more sensible
	// Make RegisterPingRoutes be attached to the instance above
	// Make the adapter implement http.Handler interface so it could be used
	// directly as handler.

	mux := http.NewServeMux()
	httpApp := handlers.NewRootHandler((*httpRouter)(mux))
	httpApp.RegisterPingRoutes(&pingController{})

	srv := &http.Server{
		Addr:              fmt.Sprintf("[::]:%d", port),
		ReadHeaderTimeout: time.Duration(readHeaderTimeoutSec) * time.Second,
		Handler:           mux,
	}
	log.Println("Starting server on port:", port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
