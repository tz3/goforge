// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

//go:embed static/routes/gorilla.go.tmpl
var gorillaRoutes []byte

//go:embed static/db/routes/gorilla.go.tmpl
var gorillaDatabaseRoutesTemplate []byte

// GorillaTemplate is a struct that provides methods to generate templates for a Gorilla-based HTTP server.
type GorillaTemplate struct{}

// Main returns the main template for the Gorilla-based HTTP server.
func (c GorillaTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the Gorilla-based HTTP server.
func (c GorillaTemplate) Server() []byte {
	return standardServerTemplate
}

// Routes returns the routes template for the Gorilla-based HTTP server.
func (c GorillaTemplate) Routes() []byte {
	return gorillaRoutes
}

func (s GorillaTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

func (s GorillaTemplate) RoutesWithDB() []byte {
	return gorillaDatabaseRoutesTemplate
}
