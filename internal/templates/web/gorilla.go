// Package web provides a set of templates for the specified web router.
package web

import (
	_ "embed"

	template "github.com/tz3/goforge/internal/templates"
	"github.com/tz3/goforge/internal/templates/advanced"
)

//go:embed static/routes/gorilla.go.tmpl
var gorillaRoutes []byte

//go:embed static/db/routes/gorilla.go.tmpl
var gorillaDatabaseRoutesTemplate []byte

// GorillaTemplate is a struct that provides methods to generate templates for a Gorilla-based HTTP server.
type GorillaTemplate struct{}

// Main returns the main template for the Gorilla-based HTTP server.
func (g GorillaTemplate) Main() []byte {
	return template.MainTemplate
}

// Server returns the server template for the Gorilla-based HTTP server.
func (g GorillaTemplate) Server() []byte {
	return standardServerTemplate
}

// Routes returns the routes template for the Gorilla-based HTTP server.
func (g GorillaTemplate) Routes() []byte {
	return gorillaRoutes
}

// Routes returns the DB server template for the Gorilla-based HTTP server.
func (g GorillaTemplate) ServerWithDB() []byte {
	return standardDatabaseServerTemplate
}

// Routes returns the DB routes template for the Gorilla-based HTTP server.
func (g GorillaTemplate) RoutesWithDB() []byte {
	return gorillaDatabaseRoutesTemplate
}

func (g GorillaTemplate) HtmxTemplImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (g GorillaTemplate) HtmxTemplRoutes() []byte {
	return advanced.GorillaHtmxTemplRoutesTemplate()
}

func (g GorillaTemplate) HtmxTemplateImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (g GorillaTemplate) HtmxTemplateRoutes() []byte {
	return advanced.GorillaHtmxTemplRoutesTemplate()
}
