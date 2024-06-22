// Package web provides a set of templates for the specified web router.
package web

import _ "embed"

//go:embed static/server/fiber.go.tmpl
var fiberServer []byte

//go:embed static/routes/fiber.go.tmpl
var fiberRoutes []byte

//go:embed static/main/fiber.go.tmpl
var fiberMain []byte

//go:embed static/db/routes/fiber.go.tmpl
var fiberDatabaseRoutesTemplate []byte

//go:embed static/db/server/fiber.go.tmpl
var fiberDatabaseServerTemplate []byte

// FiberTemplate is a struct that provides methods to generate templates for a Fiber-based HTTP server.
type FiberTemplate struct{}

// Main returns the main template for the Fiber-based HTTP server.
func (f FiberTemplate) Main() []byte {
	return fiberMain
}

// Server returns the server template for the Fiber-based HTTP server.
func (f FiberTemplate) Server() []byte {
	return fiberServer
}

// Routes returns the routes template for the Fiber-based HTTP server.
func (f FiberTemplate) Routes() []byte {
	return fiberRoutes
}

// Routes returns the DB server template for the Fiber-based HTTP server.
func (f FiberTemplate) ServerWithDB() []byte {
	return fiberDatabaseServerTemplate
}

// Routes returns the DB routes template for the Fiber-based HTTP server.
func (f FiberTemplate) RoutesWithDB() []byte {
	return fiberDatabaseRoutesTemplate
}
