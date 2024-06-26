package docker

import (
	_ "embed"
)

type MongoDockerTemplate struct{}

//go:embed static/docker-compose/mongo.tmpl
var mongoDockerTemplate []byte

func (m MongoDockerTemplate) Docker() []byte {
	return mongoDockerTemplate
}
