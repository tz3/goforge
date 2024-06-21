// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
)

//go:embed static/routes/http_router.go.tmpl
var httpRouterRoutesTemplate []byte

//go:embed static/db/routes/http_router.go.tmpl
var httpDBRouterRoutesTemplate []byte

// HttpRouterTemplate is a struct that provides methods to generate templates for a HttpRouter-based HTTP server.
type HttpRouterTemplate struct{}

// Main returns the main template for the HttpRouter-based HTTP server.
func (c HttpRouterTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the HttpRouter-based HTTP server.
func (c HttpRouterTemplate) Server() []byte {
	return standardServerTemplate
}

// Routes returns the routes template for the HttpRouter-based HTTP server.
func (c HttpRouterTemplate) Routes() []byte {
	return httpRouterRoutesTemplate
}

func (s HttpRouterTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

func (s HttpRouterTemplate) RoutesWithDB() []byte {
	return httpDBRouterRoutesTemplate
}
