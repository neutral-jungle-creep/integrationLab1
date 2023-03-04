package httpServer

import (
	"IntegrationLab1/pkg/logger"
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
	log    *logger.Logger
}

func NewServer(handler http.Handler, log *logger.Logger, opts ...Option) *Server {
	s := &Server{
		server: &http.Server{
			Addr:         defaultAddr,
			Handler:      handler,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		},
		log: log,
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
