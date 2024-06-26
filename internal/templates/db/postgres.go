package db

import (
	_ "embed"
)

type PostgresTemplate struct{}

//go:embed static/service/postgres.go.tmpl
var postgresServiceTemplate []byte

//go:embed static/env/example/postgres.tmpl
var postgresEnvExampleTemplate []byte

//go:embed static/env/postgres.tmpl
var postgresEnvTemplate []byte

func (m PostgresTemplate) Service() []byte {
	return postgresServiceTemplate
}

func (m PostgresTemplate) Env() []byte {
	return postgresEnvTemplate
}

func (m PostgresTemplate) EnvExample() []byte {
	return postgresEnvExampleTemplate
}
