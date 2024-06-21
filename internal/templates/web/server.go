// Package web provides a set of templates for the specified web router.
package web

import _ "embed"

//go:embed static/server/standard.go.tmpl
var standardServerTemplate []byte
