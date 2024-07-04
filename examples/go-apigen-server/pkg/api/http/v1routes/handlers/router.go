package handlers

import "net/http"

type httpRouter interface {
	// PathValue returns a named path parameter of a given name
	PathValue(r *http.Request, paramName string) string

	// Handle register a given handler function to handle given request
	Handle(method, pathPattern string, h http.Handler)
}
