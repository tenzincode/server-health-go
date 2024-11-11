package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type HealthData struct {
	ServiceName string  `json:"service_name"`
	Status      string  `json:"status"`
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
}

func fetchRealData() HealthData {
	// Get CPU usage percentage for all CPUs over 1 second
	cpuUsage, _ := cpu.Percent(time.Second, false)

	// Get memory usage statistics
	vmStat, _ := mem.VirtualMemory()

	return HealthData{
		ServiceName: "My Go Service",
		Status:      "Healthy",
		CPUUsage:    cpuUsage[0],        // first element is total CPU usage
		MemoryUsage: vmStat.UsedPercent, // Memory usage as a percentage
	}
}

func fetchMockData() HealthData {
	return HealthData{
		ServiceName: "Mock Service",
		Status:      "Mock Healthy",
		CPUUsage:    25.0,
		MemoryUsage: 60.0,
	}
}

func healthHandler(mock bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data HealthData

		if mock {
			data = fetchMockData()
		} else {
			data = fetchRealData()
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func main() {
	// Setup routes for real and mock health checks
	http.HandleFunc("/health", healthHandler(false))
	http.HandleFunc("/mock-health", healthHandler(true))

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Println("Starting server on :8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
