# # Start from the official Golang image
FROM golang:1.23.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o server-health-go ./cmd

# Start a new stage from a minimal Alpine image
FROM alpine:latest

# Set the working directory for the runtime container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/server-health-go .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./server-health-go"]
