package main

import (
	"log/slog"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/server"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1controllers"
)

func main() {
	srv := server.NewHTTPServer(server.HTTPServerParams{
		ServerCfg: server.ServerCfg{Port: 8080},
		Handler: server.NewRouter(server.RoutesDeps{
			PetsController: v1controllers.NewPetsController(
				v1controllers.PetsControllerDeps{},
			),
		}),
	})
	slog.Info("Starting server on port: 8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
