package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/controllers"
	"github.com/gemyago/apigen/examples/petstore-server-go/internal/api/http/router"
)

// Start it using command below:
// go run ./examples/petstore-server-go/cmd/service/
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

	// Generated routes need a controller implementation to process requests
	petsController := controllers.NewPetsController()

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
