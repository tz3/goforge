// Package template provides a set of templates for the main function, HTTP server, and Makefile.
package template

// MainTemplate returns a byte slice containing a Go source code template for the main function.
// The main function creates a new server and starts it.
func MainTemplate() []byte {
	return []byte(`package main

import (
    "{{.ProjectName}}/internal/server"
)

func main() {

    server := server.NewServer()

    err := server.ListenAndServe()
    if err != nil {
        panic("cannot start server")
    }
}
`)
}

// MakeHTTPServer returns a byte slice containing a Go source code template for an HTTP server.
// The server is configured with a port number and timeouts, and it uses the routes registered by the Server struct.
func MakeHTTPServer() []byte {
	return []byte(`package server

import (
    "fmt"
    "net/http"
    "time"
)

var port = 8080

type Server struct {
    port int
}

func NewServer() *http.Server {

    NewServer := &Server{
        port: port,
    }

    // Declare Server config
    server := &http.Server{
        Addr:         fmt.Sprintf(":%d", NewServer.port),
        Handler:      NewServer.RegisterRoutes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    return server
}
`)
}

// MakeTemplate returns a byte slice containing a Makefile template for a Go project.
// The Makefile includes targets to build, run, test, and clean the project.
func MakeTemplate() []byte {
	return []byte(
		`# Simple Makefile for a Go project
# Build the application
all: build

build:
    @echo "Building..."
    @go build -o main cmd/api/main.go

# Run the application
run:
    @go run cmd/api/main.go 

# Test the application
test:
    @echo "Testing..."
    @go test ./...

# Clean the binary
clean:
    @echo "Cleaning..."
    @rm -f main

.PHONY: all build run test clean
        `)
}
