package handlers

import (
	"net/http"

	"server-health-go/internal/health"
)

// MockMetricsHandler handles requests to update mock metrics
func MockMetricsHandler(w http.ResponseWriter, r *http.Request) {
	health.FetchMockMetrics()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mock metrics updated"))
}
