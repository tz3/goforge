// Package web provides a set of templates for the specified web router.
package web

import _ "embed"

// FiberTemplate is a struct that provides methods to generate templates for a Fiber-based HTTP server.
type FiberTemplate struct{}

// Main returns the main template for the Fiber-based HTTP server.
func (c FiberTemplate) Main() []byte {
	return MakeFiberMain
}

// Server returns the server template for the Fiber-based HTTP server.
func (c FiberTemplate) Server() []byte {
	return MakeFiberServer
}

// Routes returns the routes template for the Fiber-based HTTP server.
func (c FiberTemplate) Routes() []byte {
	return MakeFiberRoutes
}

//go:embed static/server/fiber.go.tmpl
var MakeFiberServer []byte

//go:embed static/routes/fiber.go.tmpl
var MakeFiberRoutes []byte

//go:embed static/main/fiber.go.tmpl
var MakeFiberMain []byte
