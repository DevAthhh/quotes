package http

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func NewServer(port string, routes http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         ":" + port,
			Handler:      routes,
			ReadTimeout:  4 * time.Second,
			WriteTimeout: 4 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}
}
