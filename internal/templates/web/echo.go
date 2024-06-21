// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

//go:embed static/routes/echo.go.tmpl
var echoRoutes []byte

//go:embed static/db/routes/echo.go.tmpl
var echoDatabaseRoutesTemplate []byte

// EchoTemplate is a struct that provides methods to generate templates for an Echo-based HTTP server.
type EchoTemplate struct{}

// Main returns the main template for the Echo-based HTTP server.
func (e EchoTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the Echo-based HTTP server.
func (e EchoTemplate) Server() []byte {
	return standardServerTemplate
}

// Routes returns the routes template for the Echo-based HTTP server.
func (e EchoTemplate) Routes() []byte {
	return echoRoutes
}

func (s EchoTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

func (s EchoTemplate) RoutesWithDB() []byte {
	return echoDatabaseRoutesTemplate
}
