package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootPath(t *testing.T) {
	r := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}
	if body := w.Body.String(); body != "OK" {
		t.Fatalf("expected body %q, got %q", "OK", body)
	}
}

func TestUnknownPathReturns404(t *testing.T) {
	r := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestRootMethodNotAllowed(t *testing.T) {
	r := NewRouter()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}
}

func TestHealthEndpoints(t *testing.T) {
	tests := []string{"/health/live", "/health/ready"}

	for _, path := range tests {
		t.Run(path, func(t *testing.T) {
			r := NewRouter()
			req := httptest.NewRequest(http.MethodGet, path, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
			}
			if ct := w.Header().Get("Content-Type"); ct != "application/json" {
				t.Fatalf("expected content-type application/json, got %s", ct)
			}

			var resp map[string]string
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("expected valid json response: %v", err)
			}
			if resp["status"] != "UP" {
				t.Fatalf("expected status UP, got %s", resp["status"])
			}
			if resp["timestamp"] == "" {
				t.Fatal("expected timestamp to be present")
			}
		})
	}
}
