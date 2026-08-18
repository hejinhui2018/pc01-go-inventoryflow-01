package api

import (
	"context"
	"fmt"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/config"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/workflow"
	"net/http"
	"time"
)

type Server struct {
	svc  *workflow.Service
	cfg  config.Config
	http *http.Server
}

func NewServer(svc *workflow.Service, cfg config.Config) *Server {
	mux := NewRoutes(svc, cfg)
	return &Server{svc: svc, cfg: cfg, http: &http.Server{Addr: fmt.Sprintf(":%d", cfg.Port), Handler: mux, ReadTimeout: cfg.RequestTimeout, WriteTimeout: cfg.RequestTimeout}}
}
func (s *Server) ListenAndServe() error              { return s.http.ListenAndServe() }
func (s *Server) Shutdown(ctx context.Context) error { return s.http.Shutdown(ctx) }
func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}
