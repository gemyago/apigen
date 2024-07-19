package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/router"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1controllers"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/app"
)

func main() {
	// Minimal implementation of the http server startup.
	// Real world implementation will likely to be more advanced and have
	// things like configuration loading, logging setup, some form of DI e.t.c

	port := 8080
	readHeaderTimeoutSec := 2
	storage := app.NewStorage()

	// Generated routes need a controller implementation to process requests
	petsController := v1controllers.NewPetsController(
		v1controllers.PetsControllerDeps{
			Commands: app.NewCommands(app.CommandsDeps{Storage: storage}),
			Queries:  app.NewQueries(app.QueriesDeps{Storage: storage}),
		},
	)

	srv := &http.Server{
		Addr:              fmt.Sprintf("[::]:%d", port),
		ReadHeaderTimeout: time.Duration(readHeaderTimeoutSec) * time.Second,
		Handler: router.NewHandler(router.HandlerDeps{
			PetsController: petsController,
		}),
	}
	log.Println("Starting server on port:", port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
