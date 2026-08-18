package api

import "net/http"

func method(w http.ResponseWriter, r *http.Request, allowed string) bool {
	if r.Method != allowed {
		w.Header().Set("Allow", allowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}
