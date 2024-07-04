package server

import (
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RoutesDeps struct {
	PetsController *handlers.PetsController
}

type chiAdapter struct {
	chi.Router
}

func (chiAdapter) PathValue(r *http.Request, paramName string) string {
	return chi.URLParam(r, paramName)
}

func (a chiAdapter) Handle(method, pathPattern string, h http.Handler) {
	a.Router.Method(method, pathPattern, h)
}

func NewRouter(deps RoutesDeps) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/v1", func(r chi.Router) {
		handlers.MountPetsRoutes(
			deps.PetsController,
			chiAdapter{Router: r},
		)
	})

	return router
}
