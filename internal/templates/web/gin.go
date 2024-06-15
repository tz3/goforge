// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

// GinTemplate is a struct that provides methods to generate templates for a Gin-based HTTP server.
type GinTemplate struct{}

// Main returns the main template for the Gin-based HTTP server.
func (c GinTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the Gin-based HTTP server.
func (c GinTemplate) Server() []byte {
	return MakeHTTPServer
}

// Routes returns the routes template for the Gin-based HTTP server.
func (c GinTemplate) Routes() []byte {
	return MakeGinRoutes
}

//go:embed static/routes/gin.go.tmpl
var MakeGinRoutes []byte
