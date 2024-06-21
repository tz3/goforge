// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

//go:embed static/routes/gin.go.tmpl
var ginRoutes []byte

//go:embed static/db/routes/gin.go.tmpl
var ginDatabaseRoutesTemplate []byte

// GinTemplate is a struct that provides methods to generate templates for a Gin-based HTTP server.
type GinTemplate struct{}

// Main returns the main template for the Gin-based HTTP server.
func (c GinTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the Gin-based HTTP server.
func (c GinTemplate) Server() []byte {
	return standardServerTemplate
}

// Routes returns the routes template for the Gin-based HTTP server.
func (c GinTemplate) Routes() []byte {
	return ginRoutes
}

func (s GinTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

func (s GinTemplate) RoutesWithDB() []byte {
	return ginDatabaseRoutesTemplate
}
