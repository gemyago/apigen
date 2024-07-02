package main

import (
	"log/slog"

	"github.com/gemyago/apigen/examples/go-apigen-server/api/http/server"
)

func main() {
	srv := server.NewHTTPServer(server.HTTPServerParams{
		ServerCfg: server.ServerCfg{
			Port: 8080,
		},
		Handler: server.NewRouter(server.RoutesDeps{}),
	})
	slog.Info("Starting server on port: 8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
