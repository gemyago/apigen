package server

import (
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/api/http/v1routes/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RoutesDeps struct {
	PetsController *handlers.PetsController
}

func NewRouter(deps RoutesDeps) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/pets", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
		})

		r.Post("/pets", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
		})

		r.Get("/pets/{petId}", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
		})
	})

	return router
}
