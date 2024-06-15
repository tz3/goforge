// Package template provides a set of templates for the main function, HTTP server, README, and Makefile.
package template

import _ "embed"

//go:embed static/.gitignore.tmpl
var GitIgnoreTemplate []byte
