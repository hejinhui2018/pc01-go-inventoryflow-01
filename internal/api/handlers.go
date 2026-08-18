package api

import (
	"encoding/json"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/config"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/health"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/workflow"
	"net/http"
)

type Handler struct {
	svc *workflow.Service
	cfg config.Config
}

func NewHandler(svc *workflow.Service, cfg config.Config) *Handler {
	return &Handler{svc: svc, cfg: cfg}
}
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, health.Check("inventoryflow"))
}
func (h *Handler) Events(w http.ResponseWriter, r *http.Request) {
	var e domain.InventoryEvent
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, h.cfg.MaxBodyBytes)).Decode(&e); err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	if err := h.svc.Receive(e); err != nil {
		writeJSON(w, 422, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 201, e)
}
func (h *Handler) Reconcile(w http.ResponseWriter, r *http.Request) {
	var e domain.InventoryEvent
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, h.cfg.MaxBodyBytes)).Decode(&e); err != nil {
		writeJSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	result, err := h.svc.Reconcile(e)
	if err != nil {
		writeJSON(w, 422, map[string]any{"error": err.Error(), "result": result})
		return
	}
	writeJSON(w, 200, result)
}
func (h *Handler) Report(w http.ResponseWriter, _ *http.Request) {
	rows, summary := h.svc.Report()
	writeJSON(w, 200, map[string]any{"rows": rows, "summary": summary})
}
