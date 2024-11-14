package routes

import (
	"net/http"
	"server-health-go/internal/handlers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/mock-metrics", handlers.MockMetricsHandler)
	return mux
}
