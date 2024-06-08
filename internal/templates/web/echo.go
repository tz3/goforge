// Package web provides a set of templates for the specified web router.
package web

import template "github.com/tz3/goforge/internal/templates"

// EchoTemplate is a struct that provides methods to generate templates for an Echo-based HTTP server.
type EchoTemplate struct{}

// Main returns the main template for the Echo-based HTTP server.
func (e EchoTemplate) Main() []byte {
	return template.MainTemplate()
}

// Server returns the server template for the Echo-based HTTP server.
func (e EchoTemplate) Server() []byte {
	return MakeHTTPServer()
}

// Routes returns the routes template for the Echo-based HTTP server.
func (e EchoTemplate) Routes() []byte {
	return MakeEchoRoutes()
}

// MakeEchoRoutes returns a byte slice containing a Go source code template for an Echo router.
// The template includes a function to register routes and a hello world handler.
func MakeEchoRoutes() []byte {
	return []byte(`package server
    import (
        "net/http"
    
        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
    )
// RegisterRoutes creates a new Echo router, registers a hello world handler to the root path,
// and returns the router.
func (s *Server) RegisterRoutes() http.Handler {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.GET("/", s.helloWorldHandler)
    return e
}
// helloWorldHandler is an HTTP handler that responds with a JSON containing a hello world message.
func (s *Server) helloWorldHandler(c echo.Context) error {
    resp := map[string]string{
        "message": "Hello World",
    }
    return c.JSON(http.StatusOK, resp)
}
`)
}
