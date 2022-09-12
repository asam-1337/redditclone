package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout    = 10 * time.Second
	_defaultWriteTimeout   = 10 * time.Second
	_defaultMaxHeaderBytes = 1 << 20
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: _defaultMaxHeaderBytes,
		ReadTimeout:    _defaultReadTimeout,
		WriteTimeout:   _defaultWriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Shutdown(ctx)
}
