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
func (c FiberTemplate) Main() []byte {
	return fiberMain
}

// Server returns the server template for the Fiber-based HTTP server.
func (c FiberTemplate) Server() []byte {
	return fiberServer
}

// Routes returns the routes template for the Fiber-based HTTP server.
func (c FiberTemplate) Routes() []byte {
	return fiberRoutes
}

func (s FiberTemplate) ServerWithDB() []byte {
	return fiberDatabaseServerTemplate
}

func (s FiberTemplate) RoutesWithDB() []byte {
	return fiberDatabaseRoutesTemplate
}
