package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/api/http/controllers"
	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/api/http/router"
	"github.com/gemyago/apigen/examples/petstore-server-app-layer-go/internal/app"
)

// Start it using command below:
// go run ./examples/petstore-server-app-layer-go/cmd/service/
//
// Send requests with curl:
// Create few pets:
// - curl -X POST localhost:8080/pets -d '{"id": 1, "name": "dog"}'
// - curl -X POST localhost:8080/pets -d '{"id": 2, "name": "cat"}'
// - curl -X POST localhost:8080/pets -d '{"id": 3, "name": "fish"}'
// List pets: curl localhost:8080/pets?limit=10
// Get pet by id: curl localhost:8080/pets/1

func main() {
	// Minimal implementation of the http server startup.
	// Real world implementation will likely to be more advanced and have
	// things like configuration loading, logging setup, some form of DI e.t.c

	port := 8080
	readHeaderTimeoutSec := 2

	rootCtx := context.Background()
	rootLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// application layer wireup
	petsService := app.NewPetsService(app.PetsServiceDeps{
		RootLogger: rootLogger,
	})

	// controllers wireup
	petsController := controllers.NewPetsController(controllers.PetsControllerDeps{
		PetsService: petsService,
	})

	// root handler wireup
	rootHandler := router.NewHandler(router.HandlerDeps{
		RootLogger:     rootLogger,
		PetsController: petsController,
	})

	srv := &http.Server{
		Addr:              fmt.Sprintf("[::]:%d", port),
		ReadHeaderTimeout: time.Duration(readHeaderTimeoutSec) * time.Second,
		Handler:           rootHandler,
	}
	rootLogger.InfoContext(rootCtx, "Starting server", slog.Int("port", port))
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
