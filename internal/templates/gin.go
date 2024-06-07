// Package template provides a set of templates for the Gin router.
package template

// GinTemplate is a struct that provides methods to generate templates for a Gin-based HTTP server.
type GinTemplate struct{}

// Main returns the main template for the Gin-based HTTP server.
func (c GinTemplate) Main() []byte {
	return MainTemplate()
}

// Server returns the server template for the Gin-based HTTP server.
func (c GinTemplate) Server() []byte {
	return MakeHTTPServer()
}

// Routes returns the routes template for the Gin-based HTTP server.
func (c GinTemplate) Routes() []byte {
	return MakeGinRoutes()
}

// MakeGinRoutes returns a byte slice containing a Go source code template for Gin routes.
// The template includes a function to register a hello world handler to the root path.
func MakeGinRoutes() []byte {
	return []byte(`package server

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// RegisterRoutes creates a new Gin router, registers a hello world handler to the root path,
// and returns the router.
func (s *Server) RegisterRoutes() http.Handler {
    r := gin.Default()

    r.GET("/", s.helloWorldHandler)

    return r
}

// helloWorldHandler is an HTTP handler that responds with a JSON containing a hello world message.
func (s *Server) helloWorldHandler(c *gin.Context) {
    resp := make(map[string]string)
    resp["message"] = "Hello World"

    c.JSON(http.StatusOK, resp)
}

`)
}
