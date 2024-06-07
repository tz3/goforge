// Package template provides a set of templates for the Gorilla router.
package template

// GorillaTemplate is a struct that provides methods to generate templates for a Gorilla-based HTTP server.
type GorillaTemplate struct{}

// Main returns the main template for the Gorilla-based HTTP server.
func (c GorillaTemplate) Main() []byte {
	return MainTemplate()
}

// Server returns the server template for the Gorilla-based HTTP server.
func (c GorillaTemplate) Server() []byte {
	return MakeHTTPServer()
}

// Routes returns the routes template for the Gorilla-based HTTP server.
func (c GorillaTemplate) Routes() []byte {
	return MakeGorillaRoutes()
}

// MakeGorillaRoutes returns a byte slice containing a Go source code template for Gorilla routes.
// The template includes a function to register a hello world handler to the root path.
func MakeGorillaRoutes() []byte {
	return []byte(`package server

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// RegisterRoutes creates a new Gorilla router, registers a hello world handler to the root path,
// and returns the router.
func (s *Server) RegisterRoutes() http.Handler {
    r := mux.NewRouter()

    r.HandleFunc("/", s.helloWorldHandler)

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
