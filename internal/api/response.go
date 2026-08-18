package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func fail(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, ErrorResponse{Code: code, Message: message})
}
func decode(w http.ResponseWriter, r *http.Request, limit int64, out any) bool {
	defer r.Body.Close()
	r.Body = http.MaxBytesReader(w, r.Body, limit)
	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
		fail(w, http.StatusBadRequest, "invalid_json", err.Error())
		return false
	}
	return true
}
func accept(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Content-Type") == ""
}
