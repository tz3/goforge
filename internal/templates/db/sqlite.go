package db

import (
	_ "embed"
)

type SqliteTemplate struct{}

//go:embed static/service/sqlite.go.tmpl
var sqliteServiceTemplate []byte

//go:embed static/env/sqlite.tmpl
var sqliteEnvTemplate []byte

func (m SqliteTemplate) Service() []byte {
	return sqliteServiceTemplate
}

func (m SqliteTemplate) Env() []byte {
	return sqliteEnvTemplate
}
