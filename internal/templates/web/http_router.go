// Package web provides a set of templates for the Chi router.
package web

import template "github.com/tz3/goforge/internal/templates"

// HttpRouterTemplate is a struct that provides methods to generate templates for a HttpRouter-based HTTP server.
type HttpRouterTemplate struct{}

// Main returns the main template for the HttpRouter-based HTTP server.
func (c HttpRouterTemplate) Main() []byte {
	return template.MainTemplate()
}

// Server returns the server template for the HttpRouter-based HTTP server.
func (c HttpRouterTemplate) Server() []byte {
	return MakeHTTPServer()
}

// Routes returns the routes template for the HttpRouter-based HTTP server.
func (c HttpRouterTemplate) Routes() []byte {
	return MakeRouterRoutes()
}

// MakeRouterRoutes returns a byte slice containing a Go source code template for HttpRouter routes.
// The template includes a function to register a hello world handler to the root path.
func MakeRouterRoutes() []byte {
	return []byte(`package server

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

// RegisterRoutes creates a new HttpRouter, registers a hello world handler to the root path,
// and returns the router.
func (s *Server) RegisterRoutes() http.Handler {
    r := httprouter.New()
    r.HandlerFunc(http.MethodGet, "/", s.helloWorldHandler)

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
