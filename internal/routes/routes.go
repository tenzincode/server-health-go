package routes

import (
	"encoding/json"
	"net/http"

	"server-health-go/internal/health"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func healthHandler(mock bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data
		if mock {
			data = health.FetchMockMetrics()
		} else {
			data = health.FetchMetrics()
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", healthHandler(false))
	mux.HandleFunc("/mock-health", healthHandler(true))
	mux.Handle("/metrics", promhttp.Handler())
}
