package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth_GET_OK(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "http://example.com/health", nil)
	w := httptest.NewRecorder()

	Health(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("status=%d, want %d", res.StatusCode, http.StatusOK)
	}
	ct := res.Header.Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Fatalf("content-type=%q, want application/json", ct)
	}

	body := w.Body.String()
	if !strings.Contains(body, `"status"`) || !strings.Contains(body, `"ok"`) {
		t.Fatalf("body=%q, want json with status ok", body)
	}
}

func TestHealth_MethodNotAllowed(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "http://example.com/health", nil)
	w := httptest.NewRecorder()

	Health(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("status=%d, want %d", res.StatusCode, http.StatusMethodNotAllowed)
	}
	body := w.Body.String()
	if body == "" {
		t.Fatalf("body is empty, want error message")
	}
}
