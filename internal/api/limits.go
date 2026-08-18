package api

import (
	"net/http"
	"strconv"
)

func limitFromHeader(r *http.Request, fallback int) int {
	raw := r.Header.Get("X-Page-Limit")
	if raw == "" {
		return fallback
	}
	v, err := strconv.Atoi(raw)
	if err != nil || v < 1 {
		return fallback
	}
	if v > 500 {
		return 500
	}
	return v
}
func noStore(w http.ResponseWriter) { w.Header().Set("Cache-Control", "no-store") }
func jsonOnly(w http.ResponseWriter, r *http.Request) bool {
	if !accept(r) {
		fail(w, http.StatusUnsupportedMediaType, "content_type", "application/json required")
		return false
	}
	return true
}
