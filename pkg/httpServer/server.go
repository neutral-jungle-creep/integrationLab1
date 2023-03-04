package httpServer

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
	defaultAddr         = ":8008"
)

type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler, opts ...Option) *Server {
	s := &Server{
		server: &http.Server{
			Addr:         defaultAddr,
			Handler:      handler,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
