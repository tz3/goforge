// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
	"github.com/tz3/goforge/internal/templates/advanced"
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

// Routes returns the DB server template for the Echo-based HTTP server.
func (e EchoTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

// Routes returns the DB routes template for the Echo-based HTTP server.
func (e EchoTemplate) RoutesWithDB() []byte {
	return echoDatabaseRoutesTemplate
}

func (e EchoTemplate) HtmxTemplateImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (e EchoTemplate) HtmxTemplateRoutes() []byte {
	return advanced.EchoHtmxTemplRoutesTemplate()
}
