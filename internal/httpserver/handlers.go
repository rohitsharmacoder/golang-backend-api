package httpserver

import (
	"encoding/json"
	"net/http"
	"time"
)

type healthResp struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func LiveHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResp{
		Status:    "UP",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// For now readiness = UP.
// Later you can add DB ping / Redis ping / dependency checks.
func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResp{
		Status:    "UP",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
