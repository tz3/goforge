// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

//go:embed static/routes/chi.go.tmpl
var chiRoutes []byte

//go:embed static/db/routes/chi.go.tmpl
var chiDatabaseRoutesTemplate []byte

// ChiTemplate is a struct that provides methods to generate templates for a Chi-based HTTP server.
type ChiTemplate struct{}

// Main returns the main template for the Chi-based HTTP server.
func (c ChiTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the Chi-based HTTP server.
func (c ChiTemplate) Server() []byte {
	return standardServerTemplate
}

// Routes returns the routes template for the Chi-based HTTP server.
func (c ChiTemplate) Routes() []byte {
	return chiRoutes
}

// Routes returns the DB server template for the Chi-based HTTP server.
func (c ChiTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

// Routes returns the routes for DB template for the Chi-based HTTP server.
func (c ChiTemplate) RoutesWithDB() []byte {
	return chiDatabaseRoutesTemplate
}
