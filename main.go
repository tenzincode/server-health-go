package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

// PROMETHEUS

// Define Prometheus metrics
var (
	cpuUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage",
		Help: "Current CPU usage",
	})
	memoryUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage",
		Help: "Current memory usage",
	})
)

// Register Prometheus metrics
func init() {
	prometheus.MustRegister(cpuUsageGauge)
	prometheus.MustRegister(memoryUsageGauge)
}

// SERVER HEALTH
func fetchMetrics() {
	cpuUsage, _ := cpu.Percent(0, false) // Get CPU usage percentage for all CPUs over 1 second
	vmStat, _ := mem.VirtualMemory()     // Get memory usage statistics

	// Update Prometheus metrics
	cpuUsageGauge.Set(cpuUsage[0])
	memoryUsageGauge.Set(vmStat.UsedPercent)
}

// Function to update mock metrics for testing
func fetchMockMetrics() {
	mockCPU := 25.0
	mockMemory := 60.0

	// Update Prometheus metrics with mock data
	cpuUsageGauge.Set(mockCPU)
	memoryUsageGauge.Set(mockMemory)
}

// Handler for updating mock metrics periodically
func mockMetricsHandler(w http.ResponseWriter, r *http.Request) {
	fetchMockMetrics()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mock metrics updated"))
}

// SERVER

func main() {
	// Start background goroutine to update the real data every 10 seconds
	go func() {
		for {
			fetchMetrics() // Update Prometheus metrics with real data every 10 seconds
			time.Sleep(10 * time.Second)
		}
	}()

	// Routes for real and mock metrics checks
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/mock-metrics", mockMetricsHandler)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Println("Starting server on :8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
