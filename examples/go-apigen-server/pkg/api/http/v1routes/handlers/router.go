package handlers

import "net/http"

type router interface {
	// PathValue returns a named path parameter of a given name
	PathValue(r *http.Request, paramName string) string

	// HandleFunc register a given handler function to handle given request
	HandleFunc(method, pathPattern string, h http.HandlerFunc)
}
