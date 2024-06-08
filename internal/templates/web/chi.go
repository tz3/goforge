// Package web provides a set of templates for the specified web router.
package web

import template "github.com/tz3/goforge/internal/templates"

// ChiTemplate is a struct that provides methods to generate templates for a Chi-based HTTP server.
type ChiTemplate struct{}

// Main returns the main template for the Chi-based HTTP server.
func (c ChiTemplate) Main() []byte {
	return template.MainTemplate()
}

// Server returns the server template for the Chi-based HTTP server.
func (c ChiTemplate) Server() []byte {
	return MakeHTTPServer()
}

// Routes returns the routes template for the Chi-based HTTP server.
func (c ChiTemplate) Routes() []byte {
	return MakeChiRoutes()
}

// MakeChiRoutes returns a byte slice containing a Go source code template for a Chi router.
// The template includes a function to register routes and a hello world handler.
func MakeChiRoutes() []byte {
	return []byte(`package server

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

// RegisterRoutes creates a new Chi router, registers a hello world handler to the root path,
// and returns the router.
func (s *Server) RegisterRoutes() http.Handler {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Get("/", s.helloWorldHandler)

    return r
}

// helloWorldHandler is an HTTP handler that responds with a JSON containing a hello world message.
func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    resp := make(map[string]string)
    resp["message"] = "Hello World"

    jsonResp, err := json.Marshal(resp)
    if err != nil {
        log.Fatalf("error handling JSON marshal. Err: %v", err)
    }

    w.Write(jsonResp)
}

`)
}
