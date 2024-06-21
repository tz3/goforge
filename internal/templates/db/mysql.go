package db

import (
	_ "embed"
)

type MysqlTemplate struct{}

//go:embed static/service/mysql.go.tmpl
var mysqlServiceTemplate []byte

//go:embed static/env/mysql.tmpl
var mysqlEnvTemplate []byte

func (m MysqlTemplate) Service() []byte {
	return mysqlServiceTemplate
}

func (m MysqlTemplate) Env() []byte {
	return mysqlEnvTemplate
}
