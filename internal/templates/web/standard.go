// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

// StandardLibraryTemplate is a struct that provides methods to generate templates for a standard library-based HTTP server.
type StandardLibraryTemplate struct{}

// Main returns the main template for the standard library-based HTTP server.
func (c StandardLibraryTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the standard library-based HTTP server.
func (c StandardLibraryTemplate) Server() []byte {
	return MakeHTTPServer
}

// Routes returns the routes template for the standard library-based HTTP server.
func (c StandardLibraryTemplate) Routes() []byte {
	return MakeHTTPRoutes
}

//go:embed static/routes/standard.go.tmpl
var MakeHTTPRoutes []byte
