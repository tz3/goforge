// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

// EchoTemplate is a struct that provides methods to generate templates for an Echo-based HTTP server.
type EchoTemplate struct{}

// Main returns the main template for the Echo-based HTTP server.
func (e EchoTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the Echo-based HTTP server.
func (e EchoTemplate) Server() []byte {
	return MakeHTTPServer
}

// Routes returns the routes template for the Echo-based HTTP server.
func (e EchoTemplate) Routes() []byte {
	return MakeEchoRoutes
}

//go:embed static/routes/echo.go.tmpl
var MakeEchoRoutes []byte
