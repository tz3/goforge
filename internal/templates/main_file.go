// Package template provides a set of templates for the main function, HTTP server, README, and Makefile.
package template

// MainTemplate returns a byte slice containing a Go source code template for the main function.
// The main function creates a new server and starts it.
func MainTemplate() []byte {
	return []byte(`package main

import (
    "{{.ProjectName}}/internal/server"
)

func main() {

    server := server.NewServer()

    err := server.ListenAndServe()
    if err != nil {
        panic("cannot start server")
    }
}
`)
}
