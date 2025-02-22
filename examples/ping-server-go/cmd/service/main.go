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

func main() {
	port := 8080
	readHeaderTimeoutSec := 2

	rootHandler := handlers.NewRootHandler((*httpRouter)(http.NewServeMux()))
	rootHandler.RegisterPingRoutes(&pingController{})

	srv := &http.Server{
		Addr:              fmt.Sprintf("[::]:%d", port),
		ReadHeaderTimeout: time.Duration(readHeaderTimeoutSec) * time.Second,
		Handler:           rootHandler,
	}
	log.Println("Starting server on port:", port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
