package template

import (
	_ "embed"
)

//go:embed static/env.tmpl
var envTemplate []byte

func EnvTemplate() []byte {
	return envTemplate
}
