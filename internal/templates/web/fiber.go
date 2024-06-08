// Package web provides a set of templates for the specified web router.
package web

// FiberTemplate is a struct that provides methods to generate templates for a Fiber-based HTTP server.
type FiberTemplate struct{}

// Main returns the main template for the Fiber-based HTTP server.
func (c FiberTemplate) Main() []byte {
	return MakeFiberMain()
}

// Server returns the server template for the Fiber-based HTTP server.
func (c FiberTemplate) Server() []byte {
	return MakeFiberServer()
}

// Routes returns the routes template for the Fiber-based HTTP server.
func (c FiberTemplate) Routes() []byte {
	return MakeFiberRoutes()
}

// MakeFiberServer returns a byte slice containing a Go source code template for a Fiber server.
// The server is created with a new Fiber app.
func MakeFiberServer() []byte {
	return []byte(`package server

import "github.com/gofiber/fiber/v2"

type FiberServer struct {
    *fiber.App
}

func New() *FiberServer {
    server := &FiberServer{
        App: fiber.New(),
    }

    return server
}

`)
}

// MakeFiberRoutes returns a byte slice containing a Go source code template for Fiber routes.
// The template includes a function to register a hello world handler to the root path.
func MakeFiberRoutes() []byte {
	return []byte(`package server

import (
    "github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
    s.App.Get("/", s.helloWorldHandler)
}

func (s *FiberServer) helloWorldHandler(c *fiber.Ctx) error {
    resp := map[string]string{
        "message": "Hello World",
    }

    return c.JSON(resp)
}
`)
}

// MakeFiberMain returns a byte slice containing a Go source code template for the main function.
// The main function creates a new Fiber server, registers the routes, and starts the server.
func MakeFiberMain() []byte {
	return []byte(`package main

import (
    "{{.ProjectName}}/internal/server"
)

func main() {

    server := server.New()

    server.RegisterFiberRoutes()

    err := server.Listen(":8080")
    if err != nil {
        panic("cannot start server")
    }
}
`)
}
