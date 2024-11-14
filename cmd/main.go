package main

import (
	"fmt"
	"net/http"

	"server-health-go/internal/routes"
	"server-health-go/internal/server"
)

func main() {
	// Initialize server with routes
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	// Start the HTTP server
	srv := server.NewServer(mux)
	fmt.Println("Starting server on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
