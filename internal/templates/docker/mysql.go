package docker

import (
	_ "embed"
)

type MysqlDockerTemplate struct{}

//go:embed static/docker-compose/mysql.tmpl
var mysqlDockerTemplate []byte

func (m MysqlDockerTemplate) Docker() []byte {
	return mysqlDockerTemplate
}
