package config

import (
	"sync"
)

var config model
var once sync.Once

type Config interface {
	Postgres() *postgres
}

func Get() Config {
	once.Do(
		func() {
			config.postgres = *new(postgres).fromEnv()
		},
	)
	return config
}

type model struct {
	postgres
}

func (m model) Postgres() *postgres {
	return &m.postgres
}
