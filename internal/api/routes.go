package api

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/config"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/workflow"
	"net/http"
)

func NewRoutes(svc *workflow.Service, cfg config.Config) http.Handler {
	mux := http.NewServeMux()
	h := NewHandler(svc, cfg)
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/v1/events", h.Events)
	mux.HandleFunc("/v1/reconcile", h.Reconcile)
	mux.HandleFunc("/v1/report", h.Report)
	mux.Handle("/", http.FileServer(http.Dir("web/dist")))
	return withRequestID(mux)
}
func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", r.Header.Get("X-Request-ID"))
		next.ServeHTTP(w, r)
	})
}
