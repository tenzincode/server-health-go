package main

import (
	"fmt"
	"net/http"
	"server-health-go/internal/health"
	"server-health-go/internal/routes"
	"time"
)

func main() {
	// Goroutine ticker to fetch real data
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()
		for {
			health.FetchMetrics()
			<-ticker.C
		}
	}()

	// Get HTTP router with all routes defined
	mux := routes.RegisterRoutes()

	// Start HTTP server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
