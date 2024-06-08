// Package web provides a set of templates for the specified web router.
package web

import template "github.com/tz3/goforge/internal/templates"

// StandardLibraryTemplate is a struct that provides methods to generate templates for a standard library-based HTTP server.
type StandardLibraryTemplate struct{}

// Main returns the main template for the standard library-based HTTP server.
func (c StandardLibraryTemplate) Main() []byte {
	return template.MainTemplate()
}

// Server returns the server template for the standard library-based HTTP server.
func (c StandardLibraryTemplate) Server() []byte {
	return MakeHTTPServer()
}

// Routes returns the routes template for the standard library-based HTTP server.
func (c StandardLibraryTemplate) Routes() []byte {
	return MakeHTTPRoutes()
}

// MakeHTTPRoutes returns a byte slice containing a Go source code template for standard library HTTP routes.
// The template includes a function to register a hello world handler to the root path.
func MakeHTTPRoutes() []byte {
	return []byte(`package server

import (
    "net/http"
    "encoding/json"
    "log"
)

// RegisterRoutes creates a new standard library HTTP ServeMux, registers a hello world handler to the root path,
// and returns the ServeMux.
func (s *Server) RegisterRoutes() http.Handler {

    mux := http.NewServeMux()
    mux.HandleFunc("/", s.handler)

    return mux
}

// handler is an HTTP handler that responds with a JSON containing a hello world message.
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
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
