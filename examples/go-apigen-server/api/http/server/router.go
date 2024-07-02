package server

import (
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/api/http/v1routes/handlers"
	"github.com/go-chi/chi/v5"
)

type RoutesDeps struct {
	PetsController *handlers.PetsController
}

func NewRouter(deps RoutesDeps) http.Handler {
	router := chi.NewRouter()

	// TODO: Register routes

	return router
}
