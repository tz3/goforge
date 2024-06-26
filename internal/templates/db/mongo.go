package db

import (
	_ "embed"
)

type MongoTemplate struct{}

//go:embed static/service/mongo.go.tmpl
var mongoServiceTemplate []byte

//go:embed static/env/example/mongo.tmpl
var mongoEnvExampleTemplate []byte

//go:embed static/env/mongo.tmpl
var mongoEnvTemplate []byte

func (m MongoTemplate) Service() []byte {
	return mongoServiceTemplate
}

func (m MongoTemplate) Env() []byte {
	return mongoEnvTemplate
}

func (m MongoTemplate) EnvExample() []byte {
	return mongoEnvExampleTemplate
}
