package server

import "net/http"

type RoutesDeps struct {
	PetsController *handlers.PetsController
}

func NewRouter(deps RoutesDeps) http.Handler {
}
