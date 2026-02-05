package httpserver

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Health endpoints
	mux.HandleFunc("/health/live", LiveHandler)   // liveness
	mux.HandleFunc("/health/ready", ReadyHandler) // readiness

	// Optional: root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return mux
}
