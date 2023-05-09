package server

import (
	"context"
	"fmt"
	"net/http"
	"skillbox-diploma/internal/config"
	"time"
)

type Handler http.Handler

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(_ context.Context) error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func New(cfg config.Address, h Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			Handler:      h,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}
