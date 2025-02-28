package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gemyago/apigen/examples/ping-server-go/internal/api/http/routes/handlers"
	"github.com/gemyago/apigen/examples/ping-server-go/internal/api/http/routes/models"
)

// Start it using command below:
// go run ./examples/ping-server-go/cmd/service/
//
// Send requests with curl:
// - curl localhost:8080/ping
// - curl localhost:8080/ping?message=hello

func pingHandler(_ context.Context, req *models.PingPingParams) (*models.Ping200Response, error) {
	message := req.Message
	if message == "" {
		message = "pong"
	}
	return &models.Ping200Response{Message: message}, nil
}

type pingController struct{}

func (c *pingController) Ping(b handlers.HandlerBuilder[
	*models.PingPingParams,
	*models.Ping200Response,
]) http.Handler {
	return b.HandleWith(pingHandler)
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
	rootHandler := handlers.NewRootHandler((*httpRouter)(http.NewServeMux()))
	rootHandler.RegisterPingRoutes(&pingController{})

	const readHeaderTimeout = 5 * time.Second
	srv := &http.Server{
		Addr:              "[::]:8080",
		ReadHeaderTimeout: readHeaderTimeout,
		Handler:           rootHandler,
	}
	log.Println("Starting server on port: 8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
