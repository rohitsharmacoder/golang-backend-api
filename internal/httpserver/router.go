package httpserver

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Health endpoints
	mux.HandleFunc("/health/live", LiveHandler)   // liveness
	mux.HandleFunc("/health/ready", ReadyHandler) // readiness

	// Root endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return mux
}
