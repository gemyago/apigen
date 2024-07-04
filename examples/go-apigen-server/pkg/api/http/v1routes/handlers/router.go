package handlers

import "net/http"

type httpRouter interface {
	// PathValue returns a named path parameter of a given name
	PathValue(r *http.Request, paramName string) string

	// HandleRoute register a given handler function to handle given route
	HandleRoute(method, pathPattern string, h http.Handler)

	// HandleError will be called for any error produced by handlers
	HandleError(r *http.Request, w http.ResponseWriter, err error)
}
