// Package web provides a set of functions common routers.
package web

// MakeHTTPServer returns a byte slice containing a Go source code template for an HTTP server.
// The server is configured with a port number and timeouts, and it uses the routes registered by the Server struct.
func MakeHTTPServer() []byte {
	return []byte(`package server

import (
    "fmt"
    "net/http"
    "time"
)

var port = 8080

type Server struct {
    port int
}

func NewServer() *http.Server {

    NewServer := &Server{
        port: port,
    }

    // Declare Server config
    server := &http.Server{
        Addr:         fmt.Sprintf(":%d", NewServer.port),
        Handler:      NewServer.RegisterRoutes(),
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    return server
}
`)
}
