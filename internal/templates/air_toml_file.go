package template

import _ "embed"

//go:embed static/.air.toml.tmpl
var AirTomlTemplate []byte
