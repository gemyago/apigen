package server

import (
	"fmt"
	"net/http"
	"time"
)

type Cfg struct {
	Port              int
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
}

type HTTPServerParams struct {
	Cfg
	Handler http.Handler
}

// NewHTTPServer constructor factory for general use *http.Server.
func NewHTTPServer(params HTTPServerParams) *http.Server {
	address := fmt.Sprintf("[::]:%d", params.Port)
	return &http.Server{
		Addr:              address,
		IdleTimeout:       params.IdleTimeout,
		ReadHeaderTimeout: params.ReadHeaderTimeout,
		ReadTimeout:       params.ReadTimeout,
		WriteTimeout:      params.WriteTimeout,
		Handler:           params.Handler,
	}
}
