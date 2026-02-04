package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	_ = writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
