// Package template provides a set of templates for the main function, HTTP server, README, and Makefile.
package template

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

# Todo:- add watch for running air

.PHONY: all build run test clean
`)
}
