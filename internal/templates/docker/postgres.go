package docker

import (
	_ "embed"
)

type PostgresDockerTemplate struct{}

//go:embed static/docker-compose/postgres.tmpl
var postgresDockerTemplate []byte

func (m PostgresDockerTemplate) Docker() []byte {
	return postgresDockerTemplate
}
