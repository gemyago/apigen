package server

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/api/http/v1routes/handlers"
	"github.com/gemyago/apigen/examples/go-apigen-server/pkg/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RoutesDeps struct {
	PetsController *handlers.PetsController
}

type httpRouter struct {
	chi.Router
}

func (httpRouter) PathValue(r *http.Request, paramName string) string {
	return chi.URLParam(r, paramName)
}

func (a httpRouter) Handle(method, pathPattern string, h http.Handler) {
	a.Router.Method(method, pathPattern, h)
}

func (a httpRouter) HandleError(r *http.Request, w http.ResponseWriter, err error) {
	level := slog.LevelWarn
	code := 500
	if errors.Is(err, app.ErrNotFound) {
		code = 404
	} else if errors.Is(err, app.ErrConflict) {
		code = 409
	} else {
		level = slog.LevelError
	}
	w.WriteHeader(code)
	slog.Log(r.Context(), level, "Failed to process request", slog.String("err", err.Error()))
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
			httpRouter{Router: r},
		)
	})

	return router
}
