package handlers

import (
	"gopher-notes/internal/http/response"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	return response.WriteJSON(w, status, v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	response.WriteError(w, status, msg)
}
