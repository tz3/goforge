// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
	"github.com/tz3/goforge/internal/templates/advanced"
)

//go:embed static/routes/http_router.go.tmpl
var httpRouterRoutesTemplate []byte

//go:embed static/db/routes/http_router.go.tmpl
var httpDBRouterRoutesTemplate []byte

// HttpRouterTemplate is a struct that provides methods to generate templates for a HttpRouter-based HTTP server.
type HttpRouterTemplate struct{}

// Main returns the main template for the HttpRouter-based HTTP server.
func (h HttpRouterTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the HttpRouter-based HTTP server.
func (h HttpRouterTemplate) Server() []byte {
	return standardServerTemplate
}

// Routes returns the routes template for the HttpRouter-based HTTP server.
func (h HttpRouterTemplate) Routes() []byte {
	return httpRouterRoutesTemplate
}

// Routes returns the DB server template for the HttpRouter-based HTTP server.
func (h HttpRouterTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

// Routes returns the DB routes template for the HttpRouter-based HTTP server.
func (h HttpRouterTemplate) RoutesWithDB() []byte {
	return httpDBRouterRoutesTemplate
}

func (r HttpRouterTemplate) HtmxTemplateImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (r HttpRouterTemplate) HtmxTemplateRoutes() []byte {
	return advanced.HttpRouterHtmxTemplRoutesTemplate()
}
